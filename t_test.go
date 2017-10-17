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
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkBool(b *testing.B) {
	var result T
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(true)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt(t *testing.T) {
	var original int
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt(b *testing.B) {
	var result T
	var t int = 1
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt8(t *testing.T) {
	var original int8
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt8(b *testing.B) {
	var result T
	var t int8 = 1
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt16(t *testing.T) {
	var original int16
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt16(b *testing.B) {
	var result T
	var t int16 = 1
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt32(t *testing.T) {
	var original int32
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt32(b *testing.B) {
	var result T
	var t int32 = 1
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt64(t *testing.T) {
	var original int64
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt64(b *testing.B) {
	var result T
	var t int64 = 1
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt(t *testing.T) {
	var original uint
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt(b *testing.B) {
	var result T
	var t uint = 1
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt8(t *testing.T) {
	var original uint8
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt8(b *testing.B) {
	var result T
	var t uint8 = 1
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt16(t *testing.T) {
	var original uint16
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt16(b *testing.B) {
	var result T
	var t uint16 = 1
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt32(t *testing.T) {
	var original uint32
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt32(b *testing.B) {
	var result T
	var t uint32 = 1
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt64(t *testing.T) {
	var original uint64
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt64(b *testing.B) {
	var result T
	var t uint64 = 1
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestFloat32(t *testing.T) {
	var original float32
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestFloat64(t *testing.T) {
	var original float64
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestComplex64(t *testing.T) {
	var original complex64
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestComplex128(t *testing.T) {
	var original complex128
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestByte(t *testing.T) {
	var original byte
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestRune(t *testing.T) {
	var original rune
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	var original string
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestArray(t *testing.T) {
	var original [12]int
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestSlice(t *testing.T) {
	var original []int
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkSlice(b *testing.B) {
	var result T
	var t []int
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUintpr(t *testing.T) {
	var original uintptr
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestChan(t *testing.T) {
	var original chan struct{}
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestFunc(t *testing.T) {
	var original func(int, string, complex128)
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestReader(t *testing.T) {
	var original io.Reader
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestWriter(t *testing.T) {
	var original io.Writer
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestInterface(t *testing.T) {
	var original interface{}
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestStruct(t *testing.T) {
	var original struct{}
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestG(t *testing.T) {
	var original T
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	var original map[string]int
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkMap(b *testing.B) {
	var result T
	var t map[string]int
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestPointer(t *testing.T) {
	var original *int
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestValue(t *testing.T) {
	var original reflect.Value
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkValue(b *testing.B) {
	var result T
	var t reflect.Value
	f := func(g T) T { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestError(t *testing.T) {
	var original error
	result := func(g T) T { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}
