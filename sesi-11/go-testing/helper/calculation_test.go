package helper

import (
	"fmt"
	// "github.com/stretchr/testify/assert"
	"testing"
)

// method 1
// func TestSum(t *testing.T) {
// 	rest := Sum(10, 10)

// 	if rest != 20 {
// 		panic("Result should be 20")
// 	}
// }

// method 2
func TestSum(t *testing.T) {
	rest := Sum(10, 10)

	if rest != 20 {
		t.FailNow()
	}
	fmt.Println("Result should be 20")
}

// method fail 1
func TestFailSum(t *testing.T) {
	rest := Sum(10, 10)

	if rest != 40 {
		t.Fail()
	}

	fmt.Println("TestFailSum Eksekusi Terhenti")
}
