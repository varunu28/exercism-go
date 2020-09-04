/*
Package twofer implements ShareWith method to output the sharing string based on input
*/
package twofer

import "fmt"

// ShareWith returns a share string
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
