package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].age < s[j].age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"eko", 20},
		{"budi", 30},
		{"joko", 40},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}