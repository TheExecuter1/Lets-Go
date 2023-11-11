package integers

import (
	"fmt"
	"testing"
)



func Add(x,y int) int {
	return x+y 
}

func TestAdder(t *testing.T){
	sum := Add(2,2)
	want := 4

	if sum != want {
		t.Errorf("want '%d' but got '%d'", want , sum)
	}
}
func ExampleAdd(){
	sum := Add(1,5)
	fmt.Println(sum)
	// Output: 6
}
