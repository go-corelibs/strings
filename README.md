[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/strings)
[![codecov](https://codecov.io/gh/go-corelibs/strings/graph/badge.svg?token=ekbgYsvk8Z)](https://codecov.io/gh/go-corelibs/strings)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/strings)](https://goreportcard.com/report/github.com/go-corelibs/strings)

# strings - text and string-type utilities

strings is a package for manipulating strings of text.

# Installation

``` shell
> go get github.com/go-corelibs/strings@latest
```

# Examples

## PathToSnake

``` go
func main() {
    snakePath := PathToSnake("/this/path")
    // snakePath == "this__path"
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2023 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
