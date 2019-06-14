package secret

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Example() {
	secret := Secret("my-secret!")

	fmt.Println(secret)
	fmt.Println(string(secret))
	fmt.Println(len(secret))

	// Output:
	// ****
	// my-secret!
	// 10
}

func Example_format() {
	secret := Secret("my-secret!")
	username := "john"

	fmt.Printf("%%T: %T, %T\n", secret, username)
	fmt.Printf("%%v: %v, %v\n", secret, username)
	fmt.Printf("%%+v: %+v, %+v\n", secret, username)
	fmt.Printf("%%#v: %#v, %#v\n", secret, username)
	fmt.Printf("%%s: %s, %s\n", secret, username)
	fmt.Printf("%%q: %q, %q\n", secret, username)
	fmt.Printf("%%x: %x, %x\n", secret, username)
	fmt.Printf("%%X: %X, %X\n", secret, username)

	printMismatchFormat := func(secret, username interface{}) {
		fmt.Printf("%%d: %d, %d\n", secret, username)
	}
	printMismatchFormat(secret, username)

	// Output:
	// %T: secret.Secret, string
	// %v: ****, john
	// %+v: ****, john
	// %#v: "****", "john"
	// %s: ****, john
	// %q: "****", "john"
	// %x: 2a2a2a2a, 6a6f686e
	// %X: 2A2A2A2A, 6A6F686E
	// %d: %!d(secret.Secret=****), %!d(string=john)
}

type request struct {
	Username string `json:"username"`
	Password Secret `json:"password"`
}

func Example_json() {
	req := &request{
		Username: "john",
		Password: "my-secret!",
	}
	fmt.Printf("req: %#v\n", req)

	js, _ := json.Marshal(req)
	fmt.Printf("json: %s\n", js)

	// Output:
	// req: &secret.request{Username:"john", Password:"****"}
	// json: {"username":"john","password":"my-secret!"}
}

func Example_reflect() {
	secret := Secret("my-secret!")

	rv := reflect.ValueOf(secret)
	fmt.Printf("%#v\n", rv.Kind() == reflect.String)
	fmt.Printf("%#v\n", rv.String())
	// Output:
	// true
	// "my-secret!"
}

func Example_plus_operator() {
	secret := Secret("secret!")

	s := "my secret is " + secret

	fmt.Printf("%T\n", s)
	fmt.Println(s)
	fmt.Println(string(s))

	// Output:
	// secret.Secret
	// ****
	// my secret is secret!
}
