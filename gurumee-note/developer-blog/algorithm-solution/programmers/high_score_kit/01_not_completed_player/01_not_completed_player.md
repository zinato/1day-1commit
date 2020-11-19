# 알고리즘 문제 풀이: 완주하지 못한 선수

![logo](../../logo.png)

> 프로그래머스 > 코딩테스트 연습 > 코딩테스트 고득점 Kit > 해시 > "완주하지 못한 선수" 문제에 대한 풀이입니다.
> 
> 문제 링크는 [이 곳](https://programmers.co.kr/learn/courses/30/lessons/42576?language=java)에서 확인할 수 있습니다.


## 문제 분석

문제 조건은 다음과 같다.

* 마라톤 경기에 참여한 선수의 수는 1명 이상 100,000명 이하입니다.
* 단 한명의 선수를 제외하고 모든 선수가 마라톤을 완주 하였습니다
* completion의 길이는 participant의 길이보다 1 작습니다. 
* 참가자의 이름은 1개 이상 20개 이하의 알파벳 소문자로 이루어져 있습니다.
* 참가자 중에는 동명이인이 있을 수 있습니다.

여기서 중요한 것은 모든 입력은 `participant는 completion을 포함`하며, `participant의 길이 - completion의 길이 = 1`를 만족한다는 것이다. 또한, 선수 이름 중 `중복된 이름`이 있을 수 있다는 것이다.

그럼 2가지 경우의 수가 나온다.

1. 중복된 이름이 없는 경우
2. 중복된 이름이 있는 경우

각 경우의 수를 살펴보자.

#### 중복된 이름이 없는 경우

문제의 첫 번째 입력이 이 경우의 수에 해당한다.

문제 입력
```
participant = [leo, kiki, eden]
completion = [eden, kiki]
```

이 때, `participant`에 따르면, 등록한 선수는 3명이다. 문제 조건에서 "중복된 이름"이 있을 수 있다고 했으니까, "선수 이름 : 등록한 선수 숫자" 를 표로 나타내보자.

| 등록 선수 이름 | 선수 숫자 |
| :-- | :-- |
| leo | 1 |
| kiki | 1 |
| eden | 1 |

이제 `completion`에 따라, 완주한 선수가 있으면, 표에서 이름과 쌍을 이루는 선수 숫자 1씩 빼주자. 그럼 완주 못한 단 하나의 선수가 나올 것이다.

| 등록 선수 이름 | 선수 숫자 | 설명 | 
| :-- | :-- | :-- |
| leo | 1 | 애를 반환하면 된다. |
| kiki | 0 | 1 -> 0 (completion에 이름이 존재한다.) |
| eden | 0 | 1 -> 0 (completion에 이름이 존재한다.) |

표에 따르면 "leo"를 반환하면 된다.

#### 중복된 이름이 있는 경우

문제의 세 번째 입력이 이 경우에 속한다.

문제 입력
```
participant = [mislav, stanko, mislav, ana] # mislav 중복된다.
completion = [stanko, ana, mislav]
```

이 경우에도, `participant`에 따라 등록 선수 이름 별로 선수 숫자를 나타내는 표를 만들어보자.

| 등록 선수 이름 | 선수 숫자 |
| :-- | :-- |
| mislav | 2 |
| stanko | 1 |
| ana | 1 |

이제 `completion`에 따라서, 완주한 선수면, 표에서 이름과 쌍을 이루는 선수 숫자 1씩 빼주자. 그럼 완주 못한 단 하나의 선수가 나올 것이다.

| 등록 선수 이름 | 선수 숫자 | 설명 | 
| :-- | :-- | :-- |
| mislav | 1 | 2 -> 1애를 반환하면 된다. |
| stanko | 0 | 1 -> 0 |
| ana | 0 | 1 -> 0 |

표에 따르면, "mislav"를 반환하면 된다.

#### 어떻게 풀까?

이제 어떻게 푸는지 보이는가? 여기서 중요한 것은 **등록 선수 이름과 그 쌍으로 저장된 선수 숫자**이다. "자료구조"를 공부한 사람이라면, 이러한 데이터를 저장하는데 좋은 자료구조를 알 것이다.

잘 모르겠다면 이 문제의 파트의 이름을 기억해보자. `해시`! `해시`는 개념이다. 이를 구현하는 구현체는 여러 개 있다. 해시 테이블, 해시 맵 등등.. 각 언어에서도 사용하는 것이 여러 개 있다. 대표적으로 C++의 map, Python의 딕셔너리, Java의 HashMap이 있다. 

여기서는 이러한 자료구조체를 그냥 묶어서 `해시 테이블`이라 칭하겠다. 

> 참고!
> 
> 엄밀히 말하면 "해시"는 자료구조는 아닙니다. 임의의 길이의 데이터를 고정된 길이의 데이터로 매핑하는 개념을 나타냅니다. 위의 자료 구조들은 이러한 개념을 사용하여 데이터를 저장 및 반환하는 구현체입니다. 
> 
> 자세한 내용은 [이 곳](https://ko.wikipedia.org/wiki/%ED%95%B4%EC%8B%9C_%ED%95%A8%EC%88%98)을 참고하세요.

즉, 등록할 선수 이름과 동명 이인의 수 쌍을 갖는 `해시 테이블`을 이용하면, 이 문제를 쉽게 풀 수 있다.


## 문제 풀이

먼저 `해시 테이블`을 생성한다. 이 때, { 등록한 선수 이름 : 동명 이인의 수 } 쌍을 저장해야 한다. 즉, 문자열을 키로, 정수형을 값의 쌍을 저장해야 한다. 

이 때 키 타입을 K, 값 타입을 V라고 할때 `new hash_table[K, V]`이 K를 키 타입, V를 값 타입 쌍으로 저장하는 `hash_table`을 생성하는 함수라고 하자.

```
hash_table = new hash_table[string, int]
```

그 다음 등록한 선수 이름의 목록이 있는 `participant`를 순회한다. 여기서 `p`를 `participant`의 각 요소라고 생각하자.

```
...

for p in participant: # participant 순회
    ...
```

이제 선수 이름 `p`와 동명이인의 수를 `해시 테이블`에 저장한다. 어떻게 저장할까? 간단하다. 

먼저, 해시 테이블에 해당 선수의 이름의 키가 있는지 확인한다. 없으면, { 선수 이름 : 0 } 쌍을 갖게 만들어준다. 여기서, 그리고 `hash_table`의 키를 접근하는 것은 `[]`라고 하자.

```
for p in participant:
    if p not in hash_table: # 해시 테이블에 선수 이름이 있나 확인
        hash_table[p]=0     # 없으면, { 선수 이름 : 0 } 쌍을 저장
    ...

...
```

그리고 이제 `p`로 `hash_table`을 접근하여 숫자 1을 증가시켜준다.

```
for p in participant:
    if p not in hash_table: 
        hash_table[p]=0     
    
    count = hash_table[p]       # 해당 키로 접근하여 값을 받아온다.
    hash_table[p] = count + 1 , # 그 값을 1 증가시킨 후 키 p에 저장한다.

...
```

그럼 등록 선수 목록을 기준으로 { 선수 이름 : 동명 이인의 수 } 쌍을 갖는 `해시 테이블`이 만들어진다. 

그럼 완주한 선수 목록 `completion`을 토대로 `해시 테이블`에서 동명 이인의 수를 어떻게 뺄 수 있을까? 굉장히 쉽다. `completion` 순회하면서 각 요소를 키로 `해시 테이블`을 접근하고 값을 1씩 빼주면 된다.

```
...

for c in completion:
    count = hash_table[c] 
    count -= 1
    ...
```

이 때, 그 값이 0 이하라면, `해시 테이블에서` 키를 제거하면 된다. 여기서 키를 제거하는 동작이 `delete_key`이라고 표기하자. 값이 0보다 크다면, 다시 `해시 테이블`에 `c`를 키로, `count`를 값으로 저장해야 한다.

```
...

for c in completion:
    count = hash_table[c] 
    count -= 1
    
    if count <= 0:
        hash_table.delete_key(c)
    else:
        hash_table[c] = count
```

문제 입력에 따르면, `해시 테이블`에 남아 있는 키는 딱 1개, 그 키와 싸으로 저장된 값도 1이 된다. 그럼 그 키를 반환하면 된다. 여기서 `해시 테이블`에서 키 목록을 반환하는 함수를 `key_set`이라고 하자. 즉 답은 이렇게 된다.

```
...
answer = hash_table.key_set()[0]
return answer
```

정리해보자. 알고리즘은 단계별로 이렇게 정리할 수 있겠다.

1. { 등록한 선수 이름: 동명이인의 수 } 쌍을 저장하는 해시 테이블을 생성한다.
2. 등록 선수 목록 `participant` 기준으로 해시 테이블을 초기화한다.
   1. 맵에 선수 이름의 키가 없으면 { 선수 이름: 0 } 쌍으로 저장한다.
   2. 맵에서 선수 이름의 키를 가진 값을 가져와 1 더한다.
3. 완주 선수 목록 `completion` 기준으로 해시 테이블을 동기화한다. 
   1. 맵에서 선수 이름을 키로 등록한 동명이인의 수를 가져온다.
   2. 그 수에서 1을 뺀다.
   3. 만약 뺀 수가 0이라면, 맵에서 선수 이름의 키를 제거한다.
   4. 아니라면, 뺀 수를 선수 이름과 함께 저장한다.
4. 맵에서 저장하는 키, 즉 선수 이름 중 첫 요소를 가져와서 반환한다. (`participant의 길이-completion의 길이=1`이니까 딱 1개의 키만 남는다.)

이를 나타낸 전체 수도코드는 다음과 같다.

```
# 1. 해시 테이블 생성
hash_table = new hash_table[string, int]

# 2. 등록 선수 목록 기준으로 해시 테이블 초기화
for p in participant:
    if p not in hash_table: 
        hash_table[p]=0     
    
    hash_table[p] += 1

# 3. 완주 선수 목록 기준으로 해시 테이블 동기화
for c in completion:
    count = hash_table[c] 
    count -= 1
    
    if count <= 0:
        hash_table.delete_key(c)
    else:
        hash_table[c] = count


# 4. 완주하지 못한 선수 반환
answer = hash_table.key_set()[0]
return answer
```

이를 자바 코드로 나타내면 다음과 같다.

```java
import java.util.*;

class Solution {
    public String solution(String[] participant, String[] completion) {
        Map<String, Integer> map = new HashMap<>();

        for (String p : participant) {
            map.putIfAbsent(p, 0);
            map.put(p, map.get(p) + 1);
        }

        for (String c : completion) {
            Integer count = map.get(c);
            count -= 1;

            if (count <= 0) {
                map.remove(c);
            } else {
                map.put(c, count);
            }
        }

        String answer = map.keySet().stream().findFirst().get();
        return answer;
    }
}
```

프로그래머스 사이트에서 코드를 채점하면 통과함을 알 수 있다.