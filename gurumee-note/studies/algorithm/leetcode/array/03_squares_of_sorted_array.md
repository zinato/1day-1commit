# 배열의 원소를 제곱한 후 정렬하기


## 문제 

정렬된 `A` 정수 배열이 주어진다. 각 원소를 제곱하여, 정렬된 배열로 반환하라.

Example 1:
```
Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
```

Example 2:
```
Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]
```

Note:

* 1 <= A.length <= 10000
* -10000 <= A[i] <= 10000
* A is sorted in non-decreasing order.

## 풀이

단순하다. 배열의 순회하며 각 원소를 제곱한다. 그리고 정렬하면 된다.

```go
import "sort"

func sortedSquares(A []int) []int {
    for i:=0; i<len(A); i++ {
        A[i] = A[i] * A[i]
    }
    
    sort.Ints(A)
    return A
}
```