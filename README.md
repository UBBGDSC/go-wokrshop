# Go Workshop

A brief workshop overview. All materials referenced here will be updated
continuously.

## Installing the workshop tools

- Go 1.22 ([1](https://go.dev/doc/install), [2](https://go.dev/doc/manage-install))
- curl ([1]https://curl.se/download.html)
- postman (curl alternative) ([1](https://www.postman.com/downloads/))

## Useful standard library packages to explore

- `fmt` (output formatting, writing to the console)
- `slices` (basic Go slice operations, since Go 1.21)
- `maps` (basic Go maps operations, since Go 1.21)
- `strings` (basic Go string operations)
- `strconv` (efficient string conversion)
- `math` (basic math functions)
- `os` (operating system features, such as env variable access, file operations and system calls)
- `path` and `path/filepath` (working with slash separated path names)
- `flag` (working with command-line flags)
- `math/rand` and `crypto/rand` (insecure and secure, respectively, random generators, first one being much more user friendly)
- `encoding` (for encoding data in general, such as Binary, including Little & Big endian, base64, json etc.)
- `encoding/json` (JSON writing & parsing, general JSON handling)
- `io` (onput output stream interfaces and operations)
- `bufio` (buffered input output, helpers for reading streams more conveniently)
- `errors` (create and compare errors)
- `testing` (testing and benchmarking support)
- `net/http` (Go's HTTP library)
- `context` (information propagation and cancellation signals)
- `database/sql` (interacting with databases)
- `reflect` (runtime type analysis and manipulation)
- `unsafe` (if you start using this, you know enough about Go)

## Go slang

These are things you might want to look up, but don't know their name, which makes them really hard to look up.

- `interface` https://go.dev/tour/methods/9
- `named return values` https://go.dev/doc/effective_go#named-results
- `receiver argument` (from where value and pointer receivers comes from): https://go.dev/tour/methods/1
- `pointer indirection` https://go.dev/tour/methods/1
- `uninitialized slice vs empty slice` https://stackoverflow.com/questions/44305170/nil-slices-vs-non-nil-slices-vs-empty-slices-in-go-language
- `comparable types` (required for map keys): https://go101.org/article/type-system-overview.html#types-not-support-comparison

## Useful materials

- [A Tour of Go](https://go.dev/tour/list) pretty much walks through most Go language concepts (perhaps except standard library).
- [Go Blog](https://go.dev/blog/) is the official blog of the people developing Go. Includes a lot of articles about the language, including best practices and release notes of new versions.
- [Golang Cafe](https://www.youtube.com/@GolangCafe) made a lot of videos in the past about how to use the standard library effectively, as well as Go concurrency. All the videos are relevant to this day. We warmly recommend watching them, they are extremely insightful.
- [Effective Go](https://go.dev/doc/effective_go#generality) is a nice handbook about writing idiomatic Go code.
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md) contains useful information about writing Go in production (check out [functional options](https://github.com/uber-go/guide/blob/master/style.md#functional-options) -- you'll like it a lot).
- [Dave Cheney](https://dave.cheney.net/) is a GoD and writes a lot of articles about various Go topics, including common Go coding patterns, performance profiling, low-level language analysis and more.
- [Go 101](https://go101.org/) is a deep dive into Go's implementation. We recommend this once you're more familiar with the language.
- [Marrio Carrion](https://www.youtube.com/c/MarioCarrion) is a very experienced Software Engineer, making videos largely about Go. His videos require some Go and Software Engineering experience in general.
- [Golang Dojo](https://www.youtube.com/@GolangDojo/videos) is a YouTuber talking about Go (but unfortunately seems to be relatively inactive lately), which is more beginner-friendly.
- [Nic Jackson's microservices tutorial](https://www.youtube.com/watch?v=VzBGi_n65iU&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_) is an amazing one to follow along. It is a bit outdated (e.g. using Gorilla as the web framework), but most of the material is still relevant. It's mainly how Alex learned Go before for his first job.
- This [amazing presentation](https://www.youtube.com/watch?v=oV9rvDllKEg) from Rob Pike, one of Go's founders, talking about concurrency and parallelism.
- Most Gophercon videos. They are available for free on YouTube.

## Useful Go open-source projects

When looking for an open-source project, we like going to their github page and see things like the number of stars, latest commits to see if it's actively maintained, and look up online what people say about it. But, there seems to be a general consensus around these projects that they're pretty much industry standard.

Also, and this is very important: look up their dependencies. Ideally, they should not have many dependencies. Zero is ideal.

- [Kubernetes](https://github.com/kubernetes/kubernetes) well...
- [Zerolog](https://github.com/rs/zerolog) is a logger you'll probably like.
- [Chi](https://github.com/go-chi/chi) is a nice web framework with zero extra dependencies. We've used [Gin](https://github.com/gin-gonic/gin) and [Echo](https://github.com/labstack/echo) in the past, but if in doubt, Chi should do the trick.
- [Cobra](https://github.com/spf13/cobra) is nice if you develop command-line applications. They suggest [Viper](https://github.com/spf13/viper) for config, but TBH 
- [sqlx](https://github.com/jmoiron/sqlx) and to some extent [sqlc](https://github.com/sqlc-dev/sqlc) are what you will likely be using in production to interact with databases. Please, do not use [GORM](https://github.com/go-gorm/gorm), even if people still say it's cool and you need an ORM. You probably don't.

## Useful VSCODE shortcuts

Check out [this repository](https://github.com/Ozoniuss/toolconfigs). 