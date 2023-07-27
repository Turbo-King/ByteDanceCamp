package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {
	a := user{name: "wei", password: "1024"}
	b := user{"wei", "1024"}
	c := user{name: "wei"}
	c.password = "1024"

	var d user
	d.name = "wei"
	d.password = "1024"

	fmt.Println(a, b, c, d)
	fmt.Println(checkPassword1(a, "haha"))
	fmt.Println(checkPassword2(&a, "haha"))

}

func checkPassword1(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}
