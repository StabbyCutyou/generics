package generics_test

import (
	"io"
	"reflect"
	"testing"

	// Use the dot level import to really unleash the power of golangs library importation
	. "github.com/StabbyCutyou/generics"
)

func TestBool(t *testing.T) {
	var original bool
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestInt(t *testing.T) {
	var original int
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestInt8(t *testing.T) {
	var original int8
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestInt16(t *testing.T) {
	var original int16
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestInt32(t *testing.T) {
	var original int32
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestInt64(t *testing.T) {
	var original int64
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestUInt(t *testing.T) {
	var original uint
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestUInt8(t *testing.T) {
	var original uint8
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestUInt16(t *testing.T) {
	var original uint16
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestUInt32(t *testing.T) {
	var original uint32
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestUInt64(t *testing.T) {
	var original uint64
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestFloat32(t *testing.T) {
	var original float32
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestFloat64(t *testing.T) {
	var original float64
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestComplex64(t *testing.T) {
	var original complex64
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestComplex128(t *testing.T) {
	var original complex128
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	var original string
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestArray(t *testing.T) {
	var original [12]int
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestSlice(t *testing.T) {
	var original []int
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestUintpr(t *testing.T) {
	var original uintptr
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestChan(t *testing.T) {
	var original chan struct{}
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestFunc(t *testing.T) {
	var original func(int, string, complex128)
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestReader(t *testing.T) {
	var original io.Reader
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestWriter(t *testing.T) {
	var original io.Writer
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestInterface(t *testing.T) {
	var original interface{}
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestStruct(t *testing.T) {
	var original struct{}
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestG(t *testing.T) {
	var original G
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	var original map[string]int
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestPointer(t *testing.T) {
	var original *int
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestValue(t *testing.T) {
	var original reflect.Value
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestError(t *testing.T) {
	var original error
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}
