package hexdump

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	yup "github.com/yupsh/framework"
	"github.com/yupsh/framework/opt"
	localopt "github.com/yupsh/hexdump/opt"
)

// Flags represents the configuration options for the hexdump command
type Flags = localopt.Flags

// Command implementation
type command opt.Inputs[string, Flags]

// Hexdump creates a new hexdump command with the given parameters
func Hexdump(parameters ...any) yup.Command {
	cmd := command(opt.Args[string, Flags](parameters...))
	// Set default bytes per line
	if cmd.Flags.BytesPerLine == 0 {
		cmd.Flags.BytesPerLine = 16
	}
	return cmd
}

func (c command) Execute(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
	var sources []io.ReadCloser

	if len(c.Positional) == 0 {
		sources = append(sources, io.NopCloser(stdin))
	} else {
		for _, filename := range c.Positional {
			if filename == "-" {
				sources = append(sources, io.NopCloser(stdin))
			} else {
				file, err := os.Open(filename)
				if err != nil {
					fmt.Fprintf(stderr, "hexdump: %s: %v\n", filename, err)
					continue
				}
				sources = append(sources, file)
			}
		}
	}

	for _, source := range sources {
		defer source.Close()
		if err := c.processSource(source, stdout); err != nil {
			fmt.Fprintf(stderr, "hexdump: %v\n", err)
		}
	}

	return nil
}

func (c command) processSource(source io.Reader, output io.Writer) error {
	// Skip bytes if specified
	if c.Flags.SkipBytes > 0 {
		_, err := io.CopyN(io.Discard, source, int64(c.Flags.SkipBytes))
		if err != nil && err != io.EOF {
			return err
		}
	}

	// Limit bytes if specified
	var reader io.Reader = source
	if c.Flags.ReadBytes > 0 {
		reader = io.LimitReader(source, int64(c.Flags.ReadBytes))
	}

	return c.hexdump(reader, output)
}

func (c command) hexdump(reader io.Reader, output io.Writer) error {
	buffer := make([]byte, int(c.Flags.BytesPerLine))
	offset := int64(c.Flags.SkipBytes)

	for {
		n, err := reader.Read(buffer)
		if n == 0 {
			break
		}

		data := buffer[:n]

		if bool(c.Flags.Canonical) {
			c.printCanonical(output, data, offset)
		} else if bool(c.Flags.Octal) {
			c.printOctal(output, data, offset)
		} else if bool(c.Flags.Decimal) {
			c.printDecimal(output, data, offset)
		} else if bool(c.Flags.Hex) || bool(c.Flags.Uppercase) {
			c.printHex(output, data, offset)
		} else {
			// Default: canonical format
			c.printCanonical(output, data, offset)
		}

		offset += int64(n)

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	// Print final offset
	fmt.Fprintf(output, "%08x\n", offset)

	return nil
}

func (c command) printCanonical(output io.Writer, data []byte, offset int64) {
	fmt.Fprintf(output, "%08x  ", offset)

	// Print hex bytes
	for i := 0; i < 16; i++ {
		if i == 8 {
			fmt.Fprint(output, " ")
		}
		if i < len(data) {
			if bool(c.Flags.Uppercase) {
				fmt.Fprintf(output, "%02X ", data[i])
			} else {
				fmt.Fprintf(output, "%02x ", data[i])
			}
		} else {
			fmt.Fprint(output, "   ")
		}
	}

	// Print ASCII representation
	fmt.Fprint(output, " |")
	for i := 0; i < len(data); i++ {
		if data[i] >= 32 && data[i] <= 126 {
			fmt.Fprintf(output, "%c", data[i])
		} else {
			fmt.Fprint(output, ".")
		}
	}
	fmt.Fprintln(output, "|")
}

func (c command) printOctal(output io.Writer, data []byte, offset int64) {
	fmt.Fprintf(output, "%08x ", offset)
	for _, b := range data {
		fmt.Fprintf(output, " %03o", b)
	}
	fmt.Fprintln(output)
}

func (c command) printDecimal(output io.Writer, data []byte, offset int64) {
	fmt.Fprintf(output, "%08x ", offset)
	for _, b := range data {
		fmt.Fprintf(output, " %3d", b)
	}
	fmt.Fprintln(output)
}

func (c command) printHex(output io.Writer, data []byte, offset int64) {
	fmt.Fprintf(output, "%08x ", offset)
	for _, b := range data {
		if bool(c.Flags.Uppercase) {
			fmt.Fprintf(output, " %02X", b)
		} else {
			fmt.Fprintf(output, " %02x", b)
		}
	}
	fmt.Fprintln(output)
}

func (c command) printASCII(data []byte) string {
	var result strings.Builder
	for _, b := range data {
		if b >= 32 && b <= 126 {
			result.WriteByte(b)
		} else {
			result.WriteByte('.')
		}
	}
	return result.String()
}

func (c command) String() string {
	return fmt.Sprintf("hexdump %v", c.Positional)
}
