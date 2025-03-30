package pd_format

import "testing"

func TestAdd(t *testing.T) {
	r := Add(1, 2)

	if r != 3 {
		t.Error("Add(1,2) shouldbe 3")
	}
}

func Testsplit(t *testing.T) {
	s := split("Welcome,to,GeeksforGeeks")

}

// func Testaa(t *testing.T) {
// 	r := Add(1, 2)

// 	if r != 3 {
// 		t.Error("Add(1,2) shouldbe 3")
// 	}
// }
