package command

import (
	"context"
	"fmt"
	"io"

	gloo "github.com/gloo-foo/framework"
)

type command gloo.Inputs[gloo.File, flags]

func Hexdump(parameters ...any) gloo.Command {
	cmd := command(gloo.Initialize[gloo.File, flags](parameters...))
	if cmd.Flags.BytesPerLine == 0 {
		cmd.Flags.BytesPerLine = 16
	}
	return cmd
}

func (p command) Executor() gloo.CommandExecutor {
	return func(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
		bytesPerLine := int(p.Flags.BytesPerLine)
		offset := 0

		buffer := make([]byte, 4096)
		for {
			n, err := stdin.Read(buffer)
			if n > 0 {
				// Process the bytes we read
				for i := 0; i < n; i += bytesPerLine {
					end := i + bytesPerLine
					if end > n {
						end = n
					}

					// Print offset
					fmt.Fprintf(stdout, "%08x  ", offset+i)

					// Print hex values
					for j := i; j < end; j++ {
						if j > i && (j-i)%8 == 0 {
							fmt.Fprintf(stdout, " ")
						}
						fmt.Fprintf(stdout, "%02x ", buffer[j])
					}

					// Pad if needed
					for j := end; j < i+bytesPerLine; j++ {
						if j > i && (j-i)%8 == 0 {
							fmt.Fprintf(stdout, " ")
						}
						fmt.Fprintf(stdout, "   ")
					}

					// Print ASCII representation
					fmt.Fprintf(stdout, " |")
					for j := i; j < end; j++ {
						if buffer[j] >= 32 && buffer[j] <= 126 {
							fmt.Fprintf(stdout, "%c", buffer[j])
						} else {
							fmt.Fprintf(stdout, ".")
						}
					}
					fmt.Fprintf(stdout, "|\n")
				}
				offset += n
			}

			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
		}

		// Print final offset
		fmt.Fprintf(stdout, "%08x\n", offset)
		return nil
	}
}
