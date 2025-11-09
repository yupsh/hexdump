package command_test

import (
	"errors"
	"testing"

	"github.com/gloo-foo/testable/assertion"
	"github.com/gloo-foo/testable/run"
	command "github.com/yupsh/hexdump"
)

func TestHexdump_Basic(t *testing.T) {
	result := run.Command(command.Hexdump()).
		WithStdin("abc").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestHexdump_Canonical(t *testing.T) {
	result := run.Command(command.Hexdump(command.Canonical)).
		WithStdin("test").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestHexdump_BytesPerLine(t *testing.T) {
	result := run.Command(command.Hexdump(command.BytesPerLine(8))).
		WithStdin("test data").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 2)
}

func TestHexdump_EmptyInput(t *testing.T) {
	result := run.Quick(command.Hexdump())
	assertion.NoError(t, result.Err)
	assertion.Empty(t, result.Stdout)
}

func TestHexdump_InputError(t *testing.T) {
	result := run.Command(command.Hexdump()).
		WithStdinError(errors.New("read failed")).Run()
	assertion.ErrorContains(t, result.Err, "read failed")
}

