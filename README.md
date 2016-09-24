# Generics
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/StabbyCutyou/generics)

When writing golang code, there are a series of [immutable, irrefutable laws ](https://go-proverbs.github.io/) that you're required to follow. One of these laws states the following:

`interface{} says nothing`

This is why you should avoid programming against it as a replacement for the lack of generics, as it's a meaningless crutch which tells you nothing about the behavior of the types underneath it.

I couldn't agree more.

That's why I've built package generics.

# Say it with G
Package generics exposes a single, powerful interface, `G`. With `G`, a golang developer (or gopher, as they like to be called) can confidently and securely rely on the power and flexibility of the built in `interface{}`, without risking a violation of the core laws and tenets all golang developing gophers are required to adhere to.

# Getting G

First, fetch the code using one of the many different officially sanctioned options for retrieving go code. Find each of them documented below for your convenience:

## go get

`go get github.com/StabbyCutyou/generics`

# Using G

Once you've gotten the code, you can import it like so:

```golang
  import . "github.com/StabbyCutyou/generics"
```

While importing using the dot notation is optional, it is the officially supported method of using package generics, as requiring the developer to type out the word `generics` each time they wish to declare something as being generic is very un-golang. Golang is all about developer productivity, and you can't be productive if you're spending all your time typing out long type declarations.

# G is for Generics
With `G`, you can take ugly code signatures like this:

```go
func UglyUnIdiomaticQuoteGenericApproachUnquote(poorexcuse ...interface{}) []interface{}
```

And transform it into the embodiment of meaningful code for which their can be no ambiguity

```go
func Excellence(things ...G) []G
```

Clearly, the latter example is far superior, as it screams its intent loud and proud: I am `G`, and I am generic!

# Backwards Compatibility
`G` meets the standard of golang by matching its stance on backwards compatibility. Until a 2.0 release of generics, which may never happen, `G` will always be 100% Backwards compatible with it's initial 1.0 release.

# Afterword
I hope you enjoy programming with generics as much as I do. The best part about golang is the involvement of the community in shaping the future of the language. Remember:

`interface{} says nothing`

Say it with G, instead!

# License
Apache v2 - See LICENSE
