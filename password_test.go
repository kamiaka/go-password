package password

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Example() {
	password := Password("my-secret-password!")

	fmt.Println(password)
	fmt.Println(string(password))
	fmt.Println(len(password))

	// Output:
	// ****
	// my-secret-password!
	// 19
}

func Example_format() {
	password := Password("my-secret-password!")
	username := "john"

	fmt.Printf("%%T: %T, %T\n", password, username)
	fmt.Printf("%%v: %v, %v\n", password, username)
	fmt.Printf("%%+v: %+v, %+v\n", password, username)
	fmt.Printf("%%#v: %#v, %#v\n", password, username)
	fmt.Printf("%%s: %s, %s\n", password, username)
	fmt.Printf("%%q: %q, %q\n", password, username)
	fmt.Printf("%%x: %x, %x\n", password, username)
	fmt.Printf("%%X: %X, %X\n", password, username)

	printMismatchFormat := func(password, username interface{}) {
		fmt.Printf("%%d: %d, %d\n", password, username)
	}
	printMismatchFormat(password, username)

	// Output:
	// %T: password.Password, string
	// %v: ****, john
	// %+v: ****, john
	// %#v: "****", "john"
	// %s: ****, john
	// %q: "****", "john"
	// %x: 2a2a2a2a, 6a6f686e
	// %X: 2A2A2A2A, 6A6F686E
	// %d: %!d(password.Password=****), %!d(string=john)
}

type request struct {
	Username string   `json:"username"`
	Password Password `json:"password"`
}

func Example_json() {
	req := &request{
		Username: "john",
		Password: Password("my-secret-password!"),
	}
	fmt.Printf("req: %#v\n", req)

	js, _ := json.Marshal(req)
	fmt.Printf("json: %s\n", js)

	// Output:
	// req: &password.request{Username:"john", Password:"****"}
	// json: {"username":"john","password":"my-secret-password!"}
}

func Example_reflect() {
	password := Password("my-secret-password!")

	rv := reflect.ValueOf(password)
	fmt.Printf("%#v\n", rv.Kind() == reflect.String)
	fmt.Printf("%#v\n", rv.String())
	// Output:
	// true
	// "my-secret-password!"
}

func Example_plus_operator() {
	password := Password("my-secret-password!")

	str := "my password is " + password

	fmt.Println(str)
	fmt.Println(string(str))

	// Output:
	// ****
	// my password is my-secret-password!
}
