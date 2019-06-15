package sum

import "testing"

type SumData struct {
	start int
	end   int
	sum   int
}

func TestLoopSum(t *testing.T) {
	sd := getTestData()
	for _, v := range sd {
		if csum := LoopSum(v.start, v.end); csum != v.sum {
			t.Errorf("%d + ... +%d=%d,except %d", v.start, v.end, csum, v.sum)
		}
	}
}

func TestGsSum(t *testing.T) {
	sd := getTestData()
	for _, v := range sd {
		if csum := GsSum(v.start, v.end); csum != v.sum {
			t.Errorf("%d + ... +%d=%d,except %d", v.start, v.end, csum, v.sum)
		}
	}
}

func BenchmarkGsSum(b *testing.B) {
	sd := getTestData()
	for i := 1; i < b.N; i++ {
		for _, v := range sd {
			if csum := GsSum(v.start, v.end); csum != v.sum {
				b.Errorf("%d + ... +%d=%d,except %d", v.start, v.end, csum, v.sum)
			}
		}
	}
}

func BenchmarkLoopSum(b *testing.B) {
	sd := getTestData()
	for i := 1; i < b.N; i++ {
		for _, v := range sd {
			if csum := LoopSum(v.start, v.end); csum != v.sum {
				b.Errorf("%d + ... +%d=%d,except %d", v.start, v.end, csum, v.sum)
			}
		}
	}
}
func getTestData() []SumData {
	return []SumData{{1, 100, 5050}, {1, 1, 1}}
}
