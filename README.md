# go-secret
[![GoDoc](https://godoc.org/github.com/kamiaka/go-secret?status.svg)](https://godoc.org/github.com/kamiaka/go-secret)

Not easily displayed string for Go.
But almost can treat like as normal string.

## Example

```go
type request struct {
  Username string `json:"username"`
  Password Secret `json:"secret"`
}
  :
req := &request{
  Username: "john",
  Password: "my-secret!",
}
fmt.Printf("req: %#v\n", req)

js, _ := json.Marshal(req)
fmt.Printf("json: %s\n", js)
```

### Output

```
req: &secret.request{Username:"john", Password:"****"}
json: {"username":"john","password":"my-secret!"}
```

### Want literal?

Convert to string.

```go
regexp.MustCompile(`^[a-z!-]+$`).MatchString(string(secret)) // true
```

## Caution!

If you use the plus operator, you may not get the result you imagined.

```go
secret := Secret("secret")

str := "my secret is " + secret // Secret("my secret is secret")

fmt.Println(str)         // ****
fmt.Println(string(str)) // my secret is secret
```

## License

[MIT](./LICENSE)