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
	g1Got := fmt.Sprintf("%+v\n", f(1, 2, 3))
	g1Expected := "[1 2 3]\n"
	if g1Got != g1Expected {
		t.Errorf("Error: Expected %s, got %s", g1Expected, g1Got)
	}
	g2Got := fmt.Sprintf("%+v\n", f(1, "two", 3))
	g2Expected := "[1 two 3]\n"
	if g2Got != g2Expected {
		t.Errorf("Error: Expected %s, got %s", g2Expected, g2Got)
	}
	g3Got := fmt.Sprintf("%+v\n", f(1, 2.1, 3))
	g3Expected := "[1 2.1 3]\n"
	if g3Got != g3Expected {
		t.Errorf("Error: Expected %s, got %s", g3Expected, g3Got)
	}
	g4Got := fmt.Sprintf("%+v\n", f(1, []byte{}, 3))
	g4Expected := "[1 [] 3]\n"
	if g4Got != g4Expected {
		t.Errorf("Error: Expected %s, got %s", g4Expected, g4Got)
	}
	g5Got := fmt.Sprintf("%+v\n", f(1, time.Time{}, 3))
	g5Expected := "[1 0001-01-01 00:00:00 +0000 UTC 3]\n"
	if g5Got != g5Expected {
		t.Errorf("Error: Expected %s, got %s", g5Expected, g5Got)
	}
}
