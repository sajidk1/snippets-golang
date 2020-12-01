# Modules

> How todo dependency management in Go

Initialise the module:

```sh
cd hello
go mod init hello
```

This creates a `go.mod` file.

Next time you run your code, you'll get:

```sh
❯ go run hello.go  
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
```

This updates the `go.mod` file with any dependencies (i.e. anything listed under
`import` in your code that isn't a standard package). It also creates a `go.sum`
file which contain cryptographic checksums of the content of specific module
versions.

Both `go.mod` and `go.sum` should be checked into Git alongside your code

## Further reading

- [Using Go modules](https://blog.golang.org/using-go-modules)
