package golang

import (
    goformat "go/format"
)

// FormatCode formats Go source code; returns original on failure.
func FormatCode(code string) string {
    b, err := goformat.Source([]byte(code))
    if err != nil {
        return code
    }
    return string(b)
}
