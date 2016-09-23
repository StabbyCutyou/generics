# Generics
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/StabbyCutyou/generics)
In golang, it's been said that `interface{}` doesn't say anything. This is why you
should avoid programming against it as a replacement for the lack of generics, as it's
a meaningless crutch which tells you nothing about the behavior of the types underneath it.

I couldn't agree more.

That's why I've built package generics.

# Say it with G
Package generics exposes a single, powerful interface, `G`. With `G`, a golang developer
(or gopher, as they like to be called) can confidently and securely rely on the power
and flexibility of golangs built in `interface{}`, without the bad feelings and guilt
of going against the wisdom of the crowds.

# G is for Generics
With `G`, you can take ugly code signatures like this:

```go
func UglyUnIdiomaticQuoteGenericApproachUnquote(poorexcuse ...interface{}) []interface{}
```

And transform it into the embodiment of meaningful code for which their can be no ambiguity

```go
func Excellence(things ...G) []G
```

# Backwards Compatibility
`G` meets the standard of golang by matching it's stance on backwards compatibility.
Until a 2.0 release of generics, which may never happen, `G` will always be 100% Backwards
compatible.

# Afterword
I hope you enjoy programming with generics as much as I do. The best part about golang
is the involvement of the community in shaping the future of the language.

# License
Apache v2 - See LICENSE
