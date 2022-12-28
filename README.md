# werr

A package that wraps errors with a stack trace

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/iand/werr)
[![Check Status](https://github.com/iand/werr/actions/workflows/check.yml/badge.svg)](https://github.com/iand/werr/actions/workflows/check.yml)
[![Test Status](https://github.com/iand/werr/actions/workflows/test.yml/badge.svg)](https://github.com/iand/werr/actions/workflows/test.yml)

## Overview

This is a very simple package that can be used to wrap an error before returning it. 
The `Wrap` function will annotate the error with the stack frame of the function that called Wrap.
Subsequent calls to wrap will result in a nested set of stack frames forming a call chain.
When the error is formatted, the stack frame can be printed as a multiline trace by using the `%+v` verb, for example `fmt.Printf("error %+v", err)`

`Wrap` will not wrap a nil error so it is idiomatic to return `werr.Wrap(err)` without a nil check.

This package is compatible with (and uses) `golang.org/x/xerrors`. 
The `Wrap` function is essentially a missing function from the xerrors package. 

Many thanks to [@ribasushi](https://github.com/ribasushi) who gave me the initial idea and code and suggested I make it into a standalone package.

## Usage

Wrap an error:

```Go
f, err := os.Open("myfile.txt")
if err != nil {
	return werr.Wrap(err) // return the error with the stack frame embedded 
}

```

Print a wrapped error:

```Go
// Just print the wrapped error
fmt.Printf("got error %v", err)

// Print a full stack trace on multiple lines
fmt.Printf("got error %+v", err)
```

## License

This is free and unencumbered software released into the public domain. For more
information, see <http://unlicense.org/> or the accompanying [`UNLICENSE`](UNLICENSE) file.
