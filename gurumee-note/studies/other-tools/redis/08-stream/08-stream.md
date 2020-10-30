# Stream

![logo](./../logo.png)

> redis 공식 문서와, 레디스 엔터프라이즈 교육 문서를 읽고, 정리한 문서입니다.

## 개요

`Stream` 데이터 타입은 `로그` 같이 연속적으로 "추가만" 되는 데이터를 처리하기 위해서, `Redis 5.0`에서 새로 도입된 데이터 타입이다. `Stream`은 꽤 단순한 구조임에도 불구하고 `Redis` 데이터 타입 중 가장 복잡한 데이터 타입으로 공식 문서는 소개하고 있는데, 이는 개념 구현 외에 또 다른 기능이 추가적으로 구현되었기 때문이다. 

`Redis`의 `Stream`은 분산 메세지 큐 `Kafka`처럼, Producer, Consumer 개념이 도입되었다. 정확히는 같은 개념을 구현하고 있으나 용어가 다르다고 한다. 아무튼 Producer가 제공하는 스트림에 새로운 데이터가 추가되는 것을 쓰기 위해, Consumer가 기다리는 블로킹 작업 셋과, Consumer Group이라는 개념이 추가되었다. 다시 정리하면, 클라이언트 그룹이 동일한 Producer가 제공하는 메세지 스트림의 다른 부분을 함께 Consuming할 수 있다.

일단 개인적인 의견을 붙이자면, `로그`는 어느 정도 "시계열 데이터" 개념을 포함하고 있어서, 시계열 데이터를 다룰 때 써도 괜찮을 것으로 보인다.


## 관련 명령어

관련 명령어는 다음과 같다.

* 단순 명령어
  * XADD 
  * XRANGE 
  * XREAD
  * XTRIM
  * XDEL
  * XLEN
  * XINFO
* 컨슈머 그룹 명령어
  * XGROUP
  * XREADGROUP
  * XPENDING
  * XREVRANGE








각 명령어에 대한 설명과 예시는 아래 절들을 참고하길 바란다.

### XADD
### XLEN
### XRANGE
### XREVRANGE
### XREAD
### XDEL
### XTRIM
### XGROUP
### XREADGROUP
### XACK
### XPENDING
### XINFO