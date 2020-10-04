package array

import "testing"

func findMaxConsecutiveOnes(nums []int) int {
	consecutiveOnes := 0
	cnt := 0

	for _, i := range nums {
		if i == 0 {

			if consecutiveOnes < cnt {
				consecutiveOnes = cnt
			}

			cnt = 0
		} else {
			cnt++
		}
	}

	if consecutiveOnes < cnt {
		consecutiveOnes = cnt
	}

	return consecutiveOnes
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	input := []int{1, 1, 0, 1, 1, 1}
	output := findMaxConsecutiveOnes(input)

	if output == 2 {
		t.Log("test is completed")
	} else {
		t.Error("test is failed, it must be ", 3)
	}
}
