package adder

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func Test_add(t *testing.T) {
	result := add(3, 3)
	expected := 6
	if expected != result {
		t.Errorf("incorrect result: expected %v got %v", expected, result)
	}
}

func TestSecond(t *testing.T) {
	fmt.Println("it ist stub", testTime)
}

func TestMain(m *testing.M) {
	fmt.Println("Set up")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Tear Down")
	os.Exit(exitVal)
}
