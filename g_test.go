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

func TestBoolThing(t *testing.T) {
	var original bool
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkBool(b *testing.B) {
	var result G
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(true)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func BenchmarkBoolThing(b *testing.B) {
	var result G
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(true)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt(t *testing.T) {
	var original int
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt(b *testing.B) {
	var result G
	var t int = 1
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestIntThing(t *testing.T) {
	var original int
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkIntThing(b *testing.B) {
	var result Thing
	var t int = 1
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt8(t *testing.T) {
	var original int8
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt8(b *testing.B) {
	var result G
	var t int8 = 1
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt8Thing(t *testing.T) {
	var original int8
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt8Thing(b *testing.B) {
	var result Thing
	var t int8 = 1
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt16(t *testing.T) {
	var original int16
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt16(b *testing.B) {
	var result G
	var t int16 = 1
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt16Thing(t *testing.T) {
	var original int16
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt16Thing(b *testing.B) {
	var result Thing
	var t int16 = 1
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt32(t *testing.T) {
	var original int32
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt32(b *testing.B) {
	var result G
	var t int32 = 1
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt32Thing(t *testing.T) {
	var original int32
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt32Thing(b *testing.B) {
	var result Thing
	var t int32 = 1
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt64(t *testing.T) {
	var original int64
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt64(b *testing.B) {
	var result G
	var t int64 = 1
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestInt64Thing(t *testing.T) {
	var original int64
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkInt64Thing(b *testing.B) {
	var result Thing
	var t int64 = 1
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt(t *testing.T) {
	var original uint
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt(b *testing.B) {
	var result G
	var t uint = 1
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUIntThing(t *testing.T) {
	var original uint
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUIntThing(b *testing.B) {
	var result Thing
	var t uint = 1
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt8(t *testing.T) {
	var original uint8
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt8(b *testing.B) {
	var result G
	var t uint8 = 1
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt8Thing(t *testing.T) {
	var original uint8
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt8Thing(b *testing.B) {
	var result Thing
	var t uint8 = 1
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt16(t *testing.T) {
	var original uint16
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt16(b *testing.B) {
	var result G
	var t uint16 = 1
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt16Thing(t *testing.T) {
	var original uint16
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt16Thing(b *testing.B) {
	var result Thing
	var t uint16 = 1
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt32(t *testing.T) {
	var original uint32
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt32(b *testing.B) {
	var result G
	var t uint32 = 1
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt32Thing(t *testing.T) {
	var original uint32
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt32Thing(b *testing.B) {
	var result Thing
	var t uint32 = 1
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt64(t *testing.T) {
	var original uint64
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt64(b *testing.B) {
	var result G
	var t uint64 = 1
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUInt64Thing(t *testing.T) {
	var original uint64
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkUInt64Thing(b *testing.B) {
	var result Thing
	var t uint64 = 1
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
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

func TestFloat32Thing(t *testing.T) {
	var original float32
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestFloat64Thing(t *testing.T) {
	var original float64
	result := func(g Thing) Thing { return g }(original)
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

func TestComplex64Thing(t *testing.T) {
	var original complex64
	result := func(g Thing) Thing { return g }(original)
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

func TestComplex128Thing(t *testing.T) {
	var original complex128
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestByte(t *testing.T) {
	var original byte
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestByteThing(t *testing.T) {
	var original byte
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestRune(t *testing.T) {
	var original rune
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestRuneThing(t *testing.T) {
	var original rune
	result := func(g Thing) Thing { return g }(original)
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

func TestStringThing(t *testing.T) {
	var original string
	result := func(g Thing) Thing { return g }(original)
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

func TestArrayThing(t *testing.T) {
	var original [12]int
	result := func(g Thing) Thing { return g }(original)
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

func BenchmarkSlice(b *testing.B) {
	var result G
	var t []int
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestSliceThing(t *testing.T) {
	var original []int
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkSliceThing(b *testing.B) {
	var result Thing
	var t []int
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestUintpr(t *testing.T) {
	var original uintptr
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestUintprThing(t *testing.T) {
	var original uintptr
	result := func(g Thing) Thing { return g }(original)
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

func TestChanThing(t *testing.T) {
	var original chan struct{}
	result := func(g Thing) Thing { return g }(original)
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

func TestFuncThing(t *testing.T) {
	var original func(int, string, complex128)
	result := func(g Thing) Thing { return g }(original)
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

func TestReaderThing(t *testing.T) {
	var original io.Reader
	result := func(g Thing) Thing { return g }(original)
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

func TestWriterThing(t *testing.T) {
	var original io.Writer
	result := func(g Thing) Thing { return g }(original)
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

func TestInterfaceThing(t *testing.T) {
	var original interface{}
	result := func(g Thing) Thing { return g }(original)
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

func TestStructThing(t *testing.T) {
	var original struct{}
	result := func(g Thing) Thing { return g }(original)
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

func TestThing(t *testing.T) {
	var original Thing
	result := func(g Thing) Thing { return g }(original)
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

func BenchmarkMap(b *testing.B) {
	var result G
	var t map[string]int
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestMapThing(t *testing.T) {
	var original map[string]int
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkMapThing(b *testing.B) {
	var result Thing
	var t map[string]int
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestPointer(t *testing.T) {
	var original *int
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestPointerThing(t *testing.T) {
	var original *int
	result := func(g Thing) Thing { return g }(original)
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

func BenchmarkValue(b *testing.B) {
	var result G
	var t reflect.Value
	f := func(g G) G { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestValueThing(t *testing.T) {
	var original reflect.Value
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func BenchmarkValueThing(b *testing.B) {
	var result Thing
	var t reflect.Value
	f := func(g Thing) Thing { return g }
	b.ResetTimer() // Ensure the preliminary work is not factored into the benchmark
	for i := 0; i < b.N; i++ {
		result = f(t)
	}
	f(result) // Ensure the compiler does not ruin the purity of the G benchmark
}

func TestError(t *testing.T) {
	var original error
	result := func(g G) G { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}

func TestErrorThing(t *testing.T) {
	var original error
	result := func(g Thing) Thing { return g }(original)
	if !reflect.DeepEqual(original, result) {
		t.Fail()
	}
}
