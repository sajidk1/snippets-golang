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

## Using local modules

See `./greetings` as an example.

- Create a new folder in `./modules`

- Run `go mod init <folder name>`

- You can't use `package main` anywhere in that directory rather `package <folder name>`

- To import that module into another program:
  - import `"github.com/sajid-khan-js/snippets-golang/modules/<folder name>"` - This will require it to be checked into Git
  - or in your `go.mod` file you can reference a local copy:
    - replace `github.com/sajid-khan-js/snippets-golang/modules/<folder name> => ../<folder name>`

## Further reading

- [Using Go modules](https://blog.golang.org/using-go-modules)
- [Creating Go modules](https://golang.org/doc/tutorial/create-module)
