package main

import (
    "fmt"
    "sample/oop/user"
)

func main() {
	mike := user.New("mshock", "mshock@caiman-club.org")
	mike.SetEmail("librarian@caiman-club.org") // static dispatch of methods
	fmt.Printf("'%v' '%v'\n", mike.Login(), mike.Email())

	vlad := user.New("pirogov", "chairman@caiman-club.org")
	fmt.Printf("'%v' '%v' '%v'\n", vlad.Login(), vlad.Email(), vlad.Role())
	// Внедрённый тип UserRole передаёт (promote) свои поля и методы внедряющему типу User
}
