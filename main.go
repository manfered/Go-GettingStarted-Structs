package main

import "fmt"

func main() {
	// defining the struct type
	type user struct {
		Id        int
		UserName  string
		FirstName string
	}

	// explicit struct declaration
	var fredUser user
	fredUser.Id = 1
	fredUser.UserName = "Manfered"
	fredUser.FirstName = "Fred"

	fmt.Println(fredUser)

	// inline declaration
	manfredUser := user{
		Id:        2,
		UserName:  "Manfered",
		FirstName: "Fred",
	}
	fmt.Println(manfredUser)
}
