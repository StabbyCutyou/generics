package generics_test

import (
	"fmt"
	"testing"
	"time"

	// Use the dot level import to really unleash the power of golangs library importation
	. "github.com/StabbyCutyou/generics"
)

func TestG(t *testing.T) {
	f := func(a, b, c G) []G {
		return []G{a, b, c}
	}

	fmt.Printf("%+v\n", f(1, 2, 3))
	fmt.Printf("%+v\n", f(1, "two", 3))
	fmt.Printf("%+v\n", f(1, 2.1, 3))
	fmt.Printf("%+v\n", f(1, []byte{}, 3))
	fmt.Printf("%+v\n", f(1, time.Time{}, 3))
}
