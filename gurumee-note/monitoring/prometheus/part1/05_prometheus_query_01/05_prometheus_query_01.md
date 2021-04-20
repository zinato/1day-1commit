# 05장. Prometheus Query (1) PromQL 기본

## 5.1 개요

`Prometheus`에 저장된 데이터를 쿼리하는 방법은 크게 다음의 2가지가 있다.

* PromQL
* HTTP API

이 장에서는 일반적으로 쿼리하는데 사용되는 `PromQL`의 "기본"이라 할 수 있는 `Scalar`, `Selector`, `Matcher`, `Instant Vector`, `Range Vector`, `Time Duration`, `Offset` 등에 대해 살펴본다.
 
이 장에서는 쉽고 빠르게 데이터를 수집하도록 `node-exporter`와 `Prometheus`를 연동할 것이다. `node-exporter`와 `Prometheus` 연동 및 모니터링 시 필요한 내용들에 대해서는 "2부 모니터링 편"에서 깊이 다룰 예정이다. 여기서는, 실행하는 방법만 살펴보도록 하자. 코드는 다음 URL에서 얻을 수 있다.

* 5장 코드 : [https://github.com/gurumee92/gurumee-prometheus-code/tree/master/part1/ch05](https://github.com/gurumee92/gurumee-prometheus-code/tree/master/part1/ch05) 

다운을 받았다면 터미널에 다음을 입력한다.

```bash
# 현재 위치 확인
$ pwd
/Users/gurumee/Workspace/gurumee-prometheus-code

# 디렉토리 이동
$ cd part1/ch05/

# prometheus, node-exporter 설치 및 구동
$ docker-compose up -d
```

다음 화면이 보인다면, 잘 구동된 것이다.

![01](./01.png)

## 5.2 String과 Scalar

`PromQL`에서는 다음의 `Expression Language Data Type`이 있다.

* String
* Scalar
* Instant Vector
* Range Vector

그 중 `Literal`로 표현될 수 있는 타입이 `String`과 `Scalar`이다. 먼저 `String` 타입을 쿼리해보자. `Prometheus UI`에서 `start`를 입력해보자.

![02](./02.png)

왼쪽에는 타입인 `String` 오른쪽엔 값인 `start`가 보인다. `Prometheus`는 `Go`로 만들어진만큼 탈출 문자 규칙도 `Go`의 규칙을 따른다. 예를 들어 `i say "love you"`를 표현하고 싶다면 다음과 같이 작성할 수 있다.

```
'i say "love you"'     (O)
"i say \"love you\""   (O)
'i say \"love you\"'   (X) '로 시작하면 \"를 하면 안된다.
```

반대도 마찬가지다. 혼용해서 사용할 수는 있지만, 혼용했을 때는 백슬래시를 쓰면 안된다. 이번엔 `Scalar` 타입을 쿼리해보자. `1.25`을 입력해본다.

![03](./03.png)

역시 왼쪽에는 타입인 `Scalar`가 오른쪽엔 값인 `1.25`가 보인다. `Scalar` 타입은 실수를 포함하는 숫자라고 보면 된다. 실수, 숫자, 지수, 16진수, Inf, -Inf, NaN이 될 수 있다.

## 5.3 Instant Vector Selector
## 5.4 Range Vector Selector
## 5.5 Time Duration과 Offset

## 5.6 PromQL 꿀팁