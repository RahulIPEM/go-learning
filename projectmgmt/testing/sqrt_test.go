package sqrt

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
	file, err := os.Open("sqrt_cases.csv")
	if err != nil {
		t.Fatalf("can't open cases file - %s", err)
	}
	defer file.Close()

	rdr := csv.NewReader(file)
	for {
		record, err := rdr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatalf("error reading cases file - %s", err)
		}

		val, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			t.Fatalf("Bad value - %s", record[0])
		}

		expected, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			t.Fatalf("Bad value - %s", record[1])
		}

		t.Run(fmt.Sprintf("%f", val), func(t *testing.T) {
			out, err := Sqrt(val)
			if err != nil {
				t.Fatal("error")
			}

			if !almostEqual(out, expected) {
				t.Fatalf("%f != %f", out, expected)
			}
		})
	}
}

// BenchmarkSqrt for sqrt func
func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Sqrt(float64(i))
		if err != nil {
			b.Fatal(err)
		}
	}
}
