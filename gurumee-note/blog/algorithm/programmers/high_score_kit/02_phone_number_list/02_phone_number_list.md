# 알고리즘 문제 풀이: 전화번호 목록

![logo](../../logo.png)

> 프로그래머스 > 코딩테스트 연습 > 코딩테스트 고득점 Kit > 해시 > "전화번호 목록" 문제에 대한 풀이입니다.
> 
> 문제 링크는 [이 곳](https://programmers.co.kr/learn/courses/30/lessons/42577?language=java)에서 확인할 수 있습니다.


## 문제 분석

먼저 문제의 조건을 살펴보자.

1. phone_book의 길이는 1 이상 1,000,000 이하입니다.
2. 각 전화번호의 길이는 1 이상 20 이하입니다.

그리고, "한 번호가 다른 번호의 접두어인지 확인한다"라는 지문을 보았을 때, **입력으로 중복된 번호가 들어오지 않을 것**이다. 이를 유의해서 생각해보자. 해당 문제의 경우 나올 수 있는 경우의 수는 딱 두 가지 뿐이다.

1. 입력에 다른 번호가 접두어인 번호가 있는 경우
2. 입력에 다른 번호가 접두어인 번호가 없는 경우


### 입력에 다른 번호가 접두어인 번호가 있는 경우

먼저 다른 번호가 접두어인 번호가 있는 경우를 살펴보자. 문제에서 주어진 첫 번째, 세 번째 입력이 이에 해당한다. 첫 번째 입력을 살펴보자.

문제 입력
```
phone_book = ["119", "97674223", "1195524421"]
```

먼저 "119"를 나머지 번호와 비교해보자. 

```
119 vs 97674223
```

이 경우 2 번호 모두 각각 접두어가 아닌다. 이제 그 다음 전화 번호 "1195524421"와 비교해보자.

```
119 vs 1195524421
```

이 경우 "1195524421"는 "119"의 접두어는 아니지만, "119"는 "`119`5524421"의 접두어이다. 따라서 이 경우 정답은 `false`이다. (세 번째 입력도 마찬가지이다.)


### 입력에 다른 번호가 접두어인 번호가 없는 경우

이번엔 입력에 다른 번호가 접두어인 번호가 없는 경우를 살펴보자. 문제에서 주어진 두 번째 입력이 이에 해당한다.

```
phone_book = ["123", "456", "789"]
```

첫 전화번호인 "123"부터 비교해보자. 먼저 "123"을 제외한 나머지 번호 중 "456"과 비교한다.

```
123 vs 456
```

"123", "456" 모두 서로의 접두어가 아니다. 이제 "789"와 비교한다.

```
123 vs 789
```

역시 "123", "789" 모두 서로의 접두어가 아니다. 이제, "123"과 비교할 수 있는 전화 번호는 없다. 이제 "456"을 기준으로 비교해보자. 이미 "123"과 비교했으므로 "789"와 비교하자.

```
456 vs 789
```

"456", "789"는 서로의 접두어가 아니다. 이제 "456"과 비교하지 않은 전화 번호는 없다. 또한, "789"와 비교하지 않은 전화 번호도 없다. 즉, `phone_book`의 모든 전화번호는 서로의 접두어가 아니다. 이 경우, `true`를 반환하면 된다.

### 어떻게 풀까?

프로그래머스에는 보통 정확도 테스트와 효율성 테스트가 주어진다. 어떤 문제들은 정확도 테스트만 체크하는 경우도 더러 있다. 개인적으로 알고리즘 시험을 치를 때, 가장 중요한 것은 정확도 테스트를 얼마나 빨리 풀 수 있는가이다. 효율성 테스트는 그 다음으로 생각해도 좋다. 

뭐 "해시 파트에 속해 있으니까 무조건 해시로 풀자."라고 생각할 수 있지만, 위의 경우의 수를 살펴 봐서 알 수 있듯이 해시를 쓸 필요는 없다. 명확히 풀 수 있는 방법이 있다면, 그걸 쓰는 것이 제일 좋다. 만약 그 알고리즘이 효율성 테스트에서 통과하지 못할 때, 요구하는 자료구조 및 알고리즘을 있는가를 생각하는 것이 옳다고 생각한다. 

물론 처음부터 그것을 쓸 수 있다면 써라. 이 문제의 경우, 이중 for문이면 충분하다.


## 문제 풀이

이 문제는 매우 쉽다. `phone_book`을 반복하며 순회하면서, 해당 번호와 다른 번호가 서로 접두어인지 판별하면 된다. 어떤 배열의 각 원소를 다른 원소들과 어떻게 비교할 수 있을까? 제일 쉽게는 이런 방식이 있다.

```
n = len(arr) # arr 길이

for i in 0 .. n:
    for j in 0 .. n:
        if (i == j):
            continue

        if (arr[i] vs arr[j]):
            do
```

이렇게 하면, 모든 원소를 순회하면서, 서로 다른 원소끼리 비교할 수 있다. 그러나, 같은 비교가 최소 2번은 일어나게 된다. (i vs j, j vs i) 비교횟수를 줄이고 싶다면 이런 식으로 코드를 짤 수 있다.

```
n = len(arr) # arr 길이

for i in 0 .. n-1:
    for j in i+1 .. n:
        if (arr[i] vs arr[j]):
            do
```

i, j가 같은지 비교할 필요도 없고, 다른 원소끼리의 비교는 딱 한 번 일어나게 된다. 이를 이용해서 문제를 풀어보자. 문제의 알고리즘은 다음과 같이 간단하게 정리할 수 있다.

1. `i`를 `0 ~ phone_book 길이 - 1`만큼 순회한다. 
   1. `phone_book[i]`는 `a`이다.
   2. `j`를 `i+1 ~ phone_book 길이`만큼 순회한다.
      1. `phone_book[j]`는 `b`이다.
      2. `a`가 `b`의 접두어 혹은 `b`가 `a`의 접두어라면 false를 반환한다.
2. true를 반환한다.

이를 자바 코드로 나타내면 다음과 같다.

```java
class Solution {
    public boolean solution(String[] phone_book) {
        for (int i=0; i<phone_book.length-1; i++) {
            String a = phone_book[i];
            
            for (int j=i+1; j<phone_book.length; j++) {
                String b = phone_book[j];
                
                if (a.startsWith(b) || b.startsWith(a)) {
                    return false;
                }
            }     
        }
        
        return true;
    }
}
```

코드를 제출하면, 통과하는 것을 확인할 수 있다.