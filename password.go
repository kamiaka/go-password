// Package password is not easily displayed string for Go.
// But almost can treat like as normal string.
package password

import (
	"fmt"
	"io"
)

// Password is not easily displayed string.
//
// Example:
//   password := Password("secret")
//   fmt.Print(password) // print `****`
type Password string

// Format password.
func (p Password) Format(s fmt.State, v rune) {
	switch v {
	case 'v', 's':
		if s.Flag('#') {
			io.WriteString(s, `"****"`)
		} else {
			io.WriteString(s, "****")
		}
	case 'q':
		io.WriteString(s, `"****"`)
	case 'x':
		io.WriteString(s, "2a2a2a2a")
	case 'X':
		io.WriteString(s, "2A2A2A2A")
	default:
		fmt.Fprintf(s, "%%!%v(%T=****)", string(v), p)
	}
}
