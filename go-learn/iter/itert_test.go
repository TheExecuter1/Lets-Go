package iter

import "testing"

const repeatCount = 5 
func Repeat(x string) string{
	var repeated string 
	 for i :=0; i< repeatCount ; i++{
		repeated += x 
	 }
	return repeated
}


func TestRepeat (t *testing.T){
	repeated := Repeat("f")
	want := "fffff"

	if repeated != want{
		t.Errorf("want %q and got %q",want, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i :=0; i< repeatCount ; i++{
		Repeat("f")
	 }
}