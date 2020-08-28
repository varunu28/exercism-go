/*
Package twofer implements ShareWith method to output the sharing string based on input
*/
package twofer

import "fmt"

// BaseString : Value of base share for user
const BaseString string = "One for "

// MeString : Value for share for self
const MeString string = ", one for me."

// ShareWith returns a share string
func ShareWith(name string) string {
	if name == "" {
		return fmt.Sprintf("%syou%s", BaseString, MeString)
	}
	return fmt.Sprintf("%s%s%s", BaseString, name, MeString)
}
