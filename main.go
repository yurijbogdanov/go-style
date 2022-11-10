package go_style

import (
    "io"
    "strings"
)

func Title(w io.Writer, s string) (n int, err error) {
    var output string
    output += "\n"
    output += "\u001B[33m" + strings.TrimSpace(s) + "\u001B[0m\n"
    output += "\u001B[33m" + strings.Repeat("=", len(s)-1) + "\u001B[0m\n"
    output += "\n"

    return w.Write([]byte(output))
}
