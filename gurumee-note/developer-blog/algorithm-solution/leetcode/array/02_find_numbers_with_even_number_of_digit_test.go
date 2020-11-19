package array

import (
	"strconv"
	"testing"
)

func findNumbers(nums []int) int {
    cnt := 0
    
    for _, i := range nums {
        s := strconv.Itoa(i)
        
        if len(s) % 2 == 0 {
            cnt++
        }
    }   
    
    return cnt
}

func TestFindNumbers(t *testing.T) {
	input := []int{12,345,2,6,7896}
	output := findNumbers(input)
	expected := 2

	if output == expected {
		t.Log("test is completed")
	} else {
		t.Error("test is failed, it must be", expected, "but output is", output)
	}
}