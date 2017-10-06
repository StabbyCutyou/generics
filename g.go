// Package generics offers a simple, idiomatic, and elegant way to utilize the power of generics,
// while adhering to the holy laws of golang.
package generics

// G is an interface under which all possible types apply
type G interface{}

// Thing is a helpful enhancement to this amazing interface.
type Thing G
