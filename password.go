// Package secret is not easily displayed string for Go.
// But almost can treat like as normal string.
package secret

import (
	"fmt"
	"io"
)

// Secret is not easily displayed string.
//
// Example:
//   secret := Secret("my-secret")
//   fmt.Print(secret) // print `****`
type Secret string

// Format secret.
func (p Secret) Format(s fmt.State, v rune) {
	switch v {
	case 'v', 's':
		if s.Flag('#') {
			io.WriteString(s, `"****"`)
		} else {
			io.WriteString(s, "****")
		}
	case 'q', 'x', 'X':
		fmt.Fprintf(s, fmt.Sprintf("%%%s", string(v)), "****")
	default:
		fmt.Fprintf(s, "%%!%v(%T=****)", string(v), p)
	}
}
