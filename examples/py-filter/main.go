package main

import (
	"fmt"
	"strings"

	"github.com/jabbalaci/jabbagolib/jpy"
)

type User struct {
	name       string
	occupation string
	country    string
}

func main() {
	words := []string{"war", "cup", "water", "tree", "storm"}
	selected := jpy.Filter(words, func(s string) bool {
		return strings.HasPrefix(s, "w")
	})
	fmt.Println(selected)

	fmt.Println("---")
	users := []User{
		{"John Doe", "gardener", "USA"},
		{"Roger Roe", "driver", "UK"},
		{"Ferenc Nagy", "programmer", "Hungary"},
		{"Paul Smith", "programmer", "Canada"},
		{"Lucia Mala", "teacher", "Slovakia"},
		{"Patrick Connor", "shopkeeper", "USA"},
		{"Tim Welson", "programmer", "Canada"},
		{"Tomas Smutny", "programmer", "Slovakia"},
		{"István Kiss", "programmer", "Hungary"},
	}
	huns := jpy.Filter(users, func(u User) bool {
		return u.country == "Hungary"
	})
	fmt.Printf("%q", huns)
}
