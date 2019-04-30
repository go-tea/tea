package serve

import (
	"fmt"
	"io"
	"os"
)

var (
	// Normal colors
	nBlack   = []byte{'\033', '[', '3', '0', 'm'}
	nRed     = []byte{'\033', '[', '3', '1', 'm'}
	nGreen   = []byte{'\033', '[', '3', '2', 'm'}
	nYellow  = []byte{'\033', '[', '3', '3', 'm'}
	nBlue    = []byte{'\033', '[', '3', '4', 'm'}
	nMagenta = []byte{'\033', '[', '3', '5', 'm'}
	nCyan    = []byte{'\033', '[', '3', '6', 'm'}
	nWhite   = []byte{'\033', '[', '3', '7', 'm'}
	// Bright colors
	bBlack   = []byte{'\033', '[', '3', '0', ';', '1', 'm'}
	bRed     = []byte{'\033', '[', '3', '1', ';', '1', 'm'}
	bGreen   = []byte{'\033', '[', '3', '2', ';', '1', 'm'}
	bYellow  = []byte{'\033', '[', '3', '3', ';', '1', 'm'}
	bBlue    = []byte{'\033', '[', '3', '4', ';', '1', 'm'}
	bMagenta = []byte{'\033', '[', '3', '5', ';', '1', 'm'}
	bCyan    = []byte{'\033', '[', '3', '6', ';', '1', 'm'}
	bWhite   = []byte{'\033', '[', '3', '7', ';', '1', 'm'}

	reset = []byte{'\033', '[', '0', 'm'}

	gGreen   = []byte{27, 91, 57, 55, 59, 52, 50, 109}
	gWhite   = []byte{27, 91, 57, 48, 59, 52, 55, 109}
	gYellow  = []byte{27, 91, 57, 55, 59, 52, 51, 109}
	gRed     = []byte{27, 91, 57, 55, 59, 52, 49, 109}
	gBlue    = []byte{27, 91, 57, 55, 59, 52, 52, 109}
	gMagenta = []byte{27, 91, 57, 55, 59, 52, 53, 109}
	gCyan    = []byte{27, 91, 57, 55, 59, 52, 54, 109}
	//gReset   = []byte{27, 91, 48, 109}
)

var isTTY bool

func init() {
	// This is sort of cheating: if stdout is a character device, we assume
	// that means it's a TTY. Unfortunately, there are many non-TTY
	// character devices, but fortunately stdout is rarely set to any of
	// them.
	//
	// We could solve this properly by pulling in a dependency on
	// code.google.com/p/go.crypto/ssh/terminal, for instance, but as a
	// heuristic for whether to print in color or in black-and-white, I'd
	// really rather not.
	fi, err := os.Stdout.Stat()
	if err == nil {
		m := os.ModeDevice | os.ModeCharDevice
		isTTY = fi.Mode()&m == m
	}
}

// colorWrite
func cW(w io.Writer, color []byte, s string, args ...interface{}) {
	if isTTY {
		w.Write(color)
	}
	fmt.Fprintf(w, s, args...)
	if isTTY {
		w.Write(reset)
	}
}
