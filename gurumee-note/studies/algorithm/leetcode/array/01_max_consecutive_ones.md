# 연속된 1의 개수 구하기

## 문제

다음처럼 array가 한 개 주어진다. array에서 연속되는 1의 개수 최댓값을 구하라

Example 1:
```
Input: [1,1,0,1,1,1]
Output: 3
Explanation: 1이 연속되는 데는 두군데이다. 처음에 2개와 끝에 3개, 즉 최댓값은 3이다.
```

Note:

* 입력 배열은 오로지 0, 1만 갖는다.
* 입력 배열의 길이는 양수이며, 10,000을 넘지 않는다.


## 풀이

단순하다. 배열을 순회하면서 1이면 계속 개수를 센다. 그 후 0이면, 연속된 1의 개수가 최대값인지 판단하고, 업데이트 후 다시 0부터 세면 된다. 뭐 굳이 나열하자면 이렇게 되려나

1. nums를 순회한다.
    1. i는 nums의 순회하는 원소이다.
    2. i == 0 인지 판단한다.
        1. i != 0, 즉 i가 1이면, cnt를 1 더한다.
        2. i == 0, 현재 연속된 1의 개수 cnt와 최대 연속 개수 consecutiveOnes를 비교하여, 최댓값을 consecutiveOnes에 업데이트한다. 업데이트 후 cnt는 0으로 초기화한다.
2. 한 번도 비교가 안 될 수 있기 때문에, 현재 연속된 1의 개수 cnt와 최대 연속 개수 consecutiveOnes를 비교하여, 최댓값을 consecutiveOnes에 업데이트한다.
    
2번의 경우는 예외 케이스 처리인데, 만약 예제 케이스가 다음과 같다고 해보자. 

Example 2:
```
Input: [1, 1, 1, 1]
Output: 4
```

이 때, 2번 작업이 없으면 비교를 안하기 때문에 consecutiveOnes의 값이 0이다. 따라서 마지막에도 cnt와 consecutiveOnes를 비교하는 것이 필요하다. Go 코드는 다음과 같다.

```go
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
            cnt ++
        }
    }
    
    if consecutiveOnes < cnt {
        consecutiveOnes = cnt
    }
    
    return consecutiveOnes
}
```