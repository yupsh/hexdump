package command

type BytesPerLine int
type SkipBytes int
type ReadBytes int
type Format string

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

type flags struct {
	BytesPerLine BytesPerLine
	SkipBytes    SkipBytes
	ReadBytes    ReadBytes
	Format       Format
	Canonical    CanonicalFlag
	Octal        OctalFlag
	Decimal      DecimalFlag
	Hex          HexFlag
	Uppercase    UppercaseFlag
}

func (b BytesPerLine) Configure(flags *flags)  { flags.BytesPerLine = b }
func (s SkipBytes) Configure(flags *flags)     { flags.SkipBytes = s }
func (r ReadBytes) Configure(flags *flags)     { flags.ReadBytes = r }
func (f Format) Configure(flags *flags)        { flags.Format = f }
func (c CanonicalFlag) Configure(flags *flags) { flags.Canonical = c }
func (o OctalFlag) Configure(flags *flags)     { flags.Octal = o }
func (d DecimalFlag) Configure(flags *flags)   { flags.Decimal = d }
func (h HexFlag) Configure(flags *flags)       { flags.Hex = h }
func (u UppercaseFlag) Configure(flags *flags) { flags.Uppercase = u }
