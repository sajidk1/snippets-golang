# snippets-golang

This repo contains various snippets of Go code (`./snippets`) and serves as an example Go application repo (`main.go`, `modules`, `go.mod`, `go.sum`)

## Install Go

`brew install go` or grab a binary for your OS [here](https://golang.org/dl/)

## Read docs

[A short read (I promise), walking you through a "Hello, World!" example.](https://golang.org/doc/tutorial/getting-started)

## Run a snippet

```sh
cd snippets
go run hello_world.go
```

## Run the main program

In the root of the repo

`go run main.go`

## Build the main program i.e. create a binary

In the root of the repo

`go build`

## Test functions

```sh
cd modules/greetings
go test -v
``` 

## Format all code

`gofmt -w *.go`

## Further reading

- <https://www.youtube.com/watch?v=Q0sKAMal4WQ>
- <https://www.youtube.com/watch?v=C8LgvuEBraI&feature=youtu.be>
- <https://golang.org/doc/tutorial/>
- <https://tour.golang.org/list>
- <https://gowebexamples.com/>
- <https://golang.org/doc/effective_go#introduction>
- <https://golang.org/doc/code.html>
- <https://medium.com/appsflyer/my-journey-from-python-to-go-3859783c6b3c>
- <https://www.dotnetperls.com/>

## Useful links

- [Go standard library](https://golang.org/pkg/)

# Modules

> How todo dependency management in Go

Initialise the module (run this at the root of a repo):

```sh
go mod init github.com/sajid-khan-js/snippets-golang
```

This creates a `go.mod` file.

Next time you run your code, you'll get:

```sh
❯ go run main.go  
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
...
```

This updates the `go.mod` file with any dependencies (i.e. anything listed under
`import` in your code that isn't a standard package). It also creates a `go.sum`
file which contain cryptographic checksums of the content of specific module
versions.

Both `go.mod` and `go.sum` should be checked into Git alongside your code

## Using local packages

- Create a new folder called `greetings` in `./modules`
  - Create a file called `greetings.go` in the new directory
  - Add the line `package greetings`

- To import that package into another program:
  - import `"github.com/sajid-khan-js/snippets-golang/modules/greetings"`
    - :memo: Your module (i.e. the root `go.mod`) needs to be pushed to your Git repository for this to work
  - or in your `go.mod` file you can reference a local copy:
    - replace `github.com/sajid-khan-js/snippets-golang/modules/greetings => ./modules/greetings`

## Further reading

- [Using Go modules](https://blog.golang.org/using-go-modules)
- [Creating Go modules](https://golang.org/doc/tutorial/create-module)
