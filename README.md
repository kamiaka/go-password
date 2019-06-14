# go-password
[![GoDoc](https://godoc.org/github.com/kamiaka/go-password?status.svg)](https://godoc.org/github.com/kamiaka/go-password)

Not easily displayed string for Go.
But almost can treat like as normal string.

## Example

```go
type request struct {
  Username string   `json:"username"`
  Password Password `json:"password"`
}
  :
req := &request{
  Username: "john",
  Password: Password("my-secret-password!"),
}
fmt.Printf("req: %#v\n", req)

js, _ := json.Marshal(req)
fmt.Printf("json: %s\n", js)
```

### Output

```
req: &password.request{Username:"john", Password:"****"}
json: {"username":"john","password":"my-secret-password!"}
```

### Want literal?

Convert to string.

```go
regexp.MustCompile(`^[a-z!-]+$`).MatchString(string(password)) // true
```

## Caution!

If you use the plus operator, you may not get the result you imagined.

```go
password := Password("secret")

str := "my password is " + password // Password("my password is secret")

fmt.Println(str)         // ****
fmt.Println(string(str)) // my password is secret
```

## License

[MIT](./LICENSE)