# 짝수 자릿수 숫자 개수 구하기

## 문제

정수형 배열 `nums`가 주어진다. 짝수 자릿수를 가진 원소의 개수를 구하라.

Example 1:
```
Input: [1,1,0,1,1,1]
[12,345,2,6,7896]
Output: 2
Explanation: 
12는 2자릿 수이다. 1개
345는 3자릿 수이다.
2는 1자릿 수이다.
6은 1자릿 수이다.
7896 4자릿 수이다. 2개
결국 짝수 자릿수를 가진 원소는 총 2개이다.
```

Example 2:
```
Input: nums = [555,901,482,1771]
Output: 1 
Explanation: 
오직 1771 4자릿수로 짝수 자릿수를 가진 원소는 총 1개이다.
```

Note:

* 1 < num.length <= 500
* 1 < num[i] <= 10 ^ 5


## 풀이

단순하다. 배열을 순회하면서 원소의 자릿수가 짝수인지 판단하면 된다.

1. nums를 순회한다.
    1. i는 nums의 순회하는 원소이다.
    2. i를 문자열로 변환한다. 이를 s로 한다.
    3. s의 길이가 짝수라면, cnt를 1 올려준다.

끝이다.

```go
import "strconv"

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
```