package opt

// Custom types for parameters
type BytesPerLine int
type SkipBytes int
type ReadBytes int
type Format string

// Boolean flag types with constants
type CanonicalFlag bool
const (
	Canonical   CanonicalFlag = true
	NoCanonical CanonicalFlag = false
)

type OctalFlag bool
const (
	Octal   OctalFlag = true
	NoOctal OctalFlag = false
)

type DecimalFlag bool
const (
	Decimal   DecimalFlag = true
	NoDecimal DecimalFlag = false
)

type HexFlag bool
const (
	Hex   HexFlag = true
	NoHex HexFlag = false
)

type UppercaseFlag bool
const (
	Uppercase   UppercaseFlag = true
	NoUppercase UppercaseFlag = false
)

// Flags represents the configuration options for the hexdump command
type Flags struct {
	BytesPerLine BytesPerLine  // Bytes per output line (-l)
	SkipBytes    SkipBytes     // Skip bytes from beginning (-s)
	ReadBytes    ReadBytes     // Read only specified number of bytes (-n)
	Format       Format        // Custom format string (-f)
	Canonical    CanonicalFlag // Canonical hex+ASCII display (-C)
	Octal        OctalFlag     // One-byte octal display (-b)
	Decimal      DecimalFlag   // One-byte decimal display (-d)
	Hex          HexFlag       // One-byte hex display (-x)
	Uppercase    UppercaseFlag // Use uppercase hex digits (-X)
}

// Configure methods for the opt system
func (b BytesPerLine) Configure(flags *Flags)   { flags.BytesPerLine = b }
func (s SkipBytes) Configure(flags *Flags)      { flags.SkipBytes = s }
func (r ReadBytes) Configure(flags *Flags)      { flags.ReadBytes = r }
func (f Format) Configure(flags *Flags)         { flags.Format = f }
func (c CanonicalFlag) Configure(flags *Flags)  { flags.Canonical = c }
func (o OctalFlag) Configure(flags *Flags)      { flags.Octal = o }
func (d DecimalFlag) Configure(flags *Flags)    { flags.Decimal = d }
func (h HexFlag) Configure(flags *Flags)        { flags.Hex = h }
func (u UppercaseFlag) Configure(flags *Flags)  { flags.Uppercase = u }
