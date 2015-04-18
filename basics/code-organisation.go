// Go Bootcamp: 6.1

package models

// List of packages to import
import (
	"fmt"
)

// List of constants
const (
	ConstExample = "const before vars"
)

// List of variables
var (
	ExportedVar    = 42
	nonExportedVar = "so say we all"
)

// Main type(s) for the file.
// Try to keep the lowest amount of structs per file when possible
type User struct {
	FirstName, LastName string
	Location            *UserLocation
}

type UserLocation struct {
	City    string
	Country string
}

// List of functions
func NewUser(firstName, lastName string) *User {
	return &User{FirstName: firstName,
		LastName: lastName,
		Location: &UserLocation{
			City:    "Melbourne",
			Country: "Australia",
		},
	}
}

// List of methods
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}
