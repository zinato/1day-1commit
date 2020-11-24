# 알고리즘 문제 풀이: 위장

![logo](../../logo.png)

> 프로그래머스 > 코딩테스트 연습 > 코딩테스트 고득점 Kit > 해시 > "위장" 문제에 대한 풀이입니다.
> 
> 문제 링크는 [이 곳](https://programmers.co.kr/learn/courses/30/lessons/42578?language=java)에서 확인할 수 있습니다.


## 문제 분석

음 이 문제야말로, `해시`라는 것을 이용하기에 최적화된 문제라고 볼 수 있다. 먼저 이 문제를 풀기 위해서는 간단한 수학적 지식이 필요하다. 먼저 `A`그룹과 `B`그룹이 있다고 하자. `A` 그룹에는 `a`, `b`, `c`가 있으며 `B` 그룹에는 `d`, `e`가 있다. 그룹은 최소 1가지는 선택해야 하며, 각 그룹 당 1개씩 선택하는 경우의 수는 어떻게 구할 수 있을까?

먼저 일일이 세어보자. 먼저 한 그룹만 선택하고 한 원소만 선택하는 경우의 수이다.

```
경우의 수 : a, b, c, d, e
```

그리고 두 그룹에서 각각 1개씩 뽑아 조합하는 경우의 수이다.

```
경우의 수 : {a, d}, {a, e}, {b, d}, {b, e}, {c, d}, {c, e}
```

총 11개의 경우의 수를 구할 수 있다. 이런 경우 경우의 수를 구하는 최적의 방식이 있다. `A`그룹의 원소의 개수를 `CNTa` `B`그룹의 원소의 개수를 `CNTb`를 할 때, 다음 공식이 성립된다.

```
cnt = (CNTa + 1) x (CNTb + 1) -1
```

이를 일반화하면, 각 그룹을 Xn으로 표현할 때, 최소 한 그룹은 선택하며, 선택한 그룹에서 원소 1개씩 뽑아 만든 조합의 경우의 수는 다음과 같다.

```
cnt = (CNTx1 + 1) X * (CNTx2 + 1) X ... X (CNTxN + 1) - 1
```

이제 문제에 적용해보자. 다음은 문제의 첫번째 입력이다.

문제 입력
```
clothes = [[yellow_hat, headgear], [blue_sunglasses, eyewear], [green_turban, headgear]]
```

문제 입력에서는 2차원 배열로, 그 안에 저장된 각 배열은 또, {옷 이름, 옷 종류}가 저장되어 있다. 우리가 관심있는 것은 "옷 종류"에 따른 그 갯수이다. 입력에 따르면, 옷 종류는 `headgear`, `eyewear` 총 2가지이다. 이들을 그룹으로, 개수를 세어주면 된다. 이렇게 말이다.

```
m = {
    "headgear": 2, // yellow_hat, green_turban
    "eyewear": 1,  // blue_sunglasses
}
```

이렇게 그루핑된 데이터를 다음 공식에 제공하면 된다.

```
answer = ( CNTheadgear + 1 ) X (CNTeyewear + 1) - 1
       = (2 + 1) X (1 + 1) - 1
       = 3 X 2 - 1
       = 5
```

따라서 답은 5가 된다.


## 문제 풀이

이 문제의 핵심은 크게 2가지이다.

1. `clothes`를 `{ 옷의 종류 : 옷의 개수 }` 쌍으로 저장하는 해시 구조로 변경하는 것
2. 해시에서, 경우의 수 알고리즘을 적용하는 것.

먼저 주어진 입력 `clothes`를 해시로 변경해보자. 해시 구조를 `m`이라 하고 해시 구조 생성을 `make(map[K]V)` 라고 하자. 여기서 K, V는 키의 타입, 값의 타입이다.

```
m = make(map[string]int)
```

옷의 종류는 문자열, 이에 개수는 숫자이기 때문에, {문자열 : 정수} 쌍을 저장하는 해시를 만든다. 그 후 `clothes`를 순회한다. `clothes`는 2차원 배열인데, 그 안에는 길이가 2개인 1차원 배열이 들어있다. 첫 값은 옷의 이름, 두 번째 값은 옷의 종류이다. 우리가 관심있는 것은 옷의 종류이다.

```
...

for cloth in clothes:
    name = cloth[0]
    kind = cloth[1]
    ...

...
```

여기서, `kind`를 키로, `m`에 값을 불러온다. 이를 `cnt`라 하자. 가져오고 1을 더해준 후 다시 키 `kind`에 1을 더한 `cnt`를 값으로 `m`에 저장하면 된다. 이 때 키가 `m`에 존재하지 않을 수 있으므로 기본 값 설정을 해주어야 한다.

```
...

for cloth in clothes:
    name = cloth[0]
    kind = cloth[1]
    # 기본 값 0 설정
    cnt = 0 
    
    # 키 존재하면 해당 갯수로
    if m.isExists(kind):
        cnt = m[kind]
    
    # 1 증가
    cnt += 1

    # 해시에 키: 업데이트된 값 저장.
    m[kind] = cnt      
...
```

이제 이렇게 구조가 변경된 해시를 이용해서 경우의 수 알고리즘을 구하면 된다. 각 값 + 1 을 곱해줘야 하므로, 값들을 순회하면서, 1을 더해준 후 누적시키면 된다. 각 키에 저장된 값을 가져오는 함수를 `values`라고 하면 코드는 다음과 같이 쓸 수 있다.

```
...

answer = 1

# 각 그룹에서 1개씩 뽑아 조합하는 경우의 수를 나타낸다.
for v in m.values():
    answer = answer * (v + 1)

# 1을 뺴주어야 함을 잊지 말자
answer -= 1
```

이제 모든 코드 조각을 다 모았다. 전체적으로 보면 다음과 같다.

```
# 해시를 생성한다.
m = make(map[string]int)

# 이차원 배열 -> 해시로 구조 변경
for cloth in clothes:
    // name = cloth[0]
    kind = cloth[1]
    cnt = 0 
    
    if m.isExists(kind):
        cnt = m[kind]
    
    cnt += 1
    m[kind] = cnt  

# 경우의 수 알고리즘 적용
answer = 1

for v in m.values():
    answer = answer * (v + 1)

answer -= 1
```

이를 간략하게 정리하면 다음과 같다.

1. `{ 옷의 종류 : 옷의 개수 }` 쌍을 저장하는 해시를 만든다. 이를 `m`이라 한다.
2. `clothes`를 순회한다. 순회하는 원소는 `cloth`라고 한다.
   1. `cloth[1]`은 옷의 종류를 나타낸다. 이를 `kind`라 한다.
   2. `kind`를 키로 `m`에 저장된 값을 가져온다. 해당 키가 없다면 0을 반환한다. 이를 `cnt`라 한다.
   3. `cnt`를 1 증가시킨다.
   4. 다시 `kind`를 키로 `cnt`를 값으로 `m`에 저장한다.
3. `answer`는 1로 초기화한다.
4. `m`의 값들만 순회한다. 이 떄 순회하는 원소를 `v`라 한다.
   1. `(v + 1)`을 `answer`에 곱해준다.
5. `answer`를 1 빼주고 반환한다.

이를 자바 코드로 옮기면 다음과 같다.

```java
import java.util.*;

public class Solution {
    public int solution(String[][] clothes) {
        Map<String, Integer> m = new HashMap<>();

        for (String [] cloth : clothes) {
            String kind = cloth[1];
            int cnt = m.getOrDefault(kind, 0);
            cnt += 1;
            m.put(kind, cnt);
        }

        int answer = 1;

        for (int v : m.values()) {
            answer *= (v + 1);
        }

        answer -= 1;
        return answer;
    }
}
```