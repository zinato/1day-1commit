# 04장. Prometheus Label

## 4.1 개요

`Label`은 `Prometheus`의 아주 강력한 기능 중 하나이다. `Label`은 키-값 쌍으로 이루어져 있으며, `Prometheus`가 시계열 데이터를 식별하는데 "메트릭 이름"과 더불어서 사용한다. 예를 들어보자. 모니터링 세계에서 HTTP 요청에 대한 상태 코드는 주로 다음과 같이 수집한다.

* 2xx (응답 성공)
* 3xx (응답 성공 - 리다이렉션)
* 4xx (응답 실패 - 사용자 오류)
* 5xx (응답 실패 - 서버 오류)

어떻게 메트릭 이름을 지을 것인가? 아주 간단하게 다음과 같이 지을 수 있을 것이다.

```
http_request_status_code_2xx
http_request_status_code_3xx
http_request_status_code_4xx
http_request_status_code_5xx
```

만약에, 2xx, 3xx가 아니라 각 상태 코드 별로 모아야 한다고 해보자. 현재 표준에 따르면 30개가 넘는 상태 코드가 존재한다. (실제로는 다 쓰이진 않더라도...) 이를 다 만들 것인가? `Prometheus`에서는 저런 패턴을 지양할 것을 강력하게 권고한다. 저런 패턴을 일컬어 "안티 패턴"이라고도 한다. `Prometheus`는 보통 상태 코드의 값에 대한 메트릭을 다음과 같이 수집한다.

```
http_request{ status_code="200" }
http_request{ status_code="201" }
http_request{ status_code="301" }
http_request{ status_code="404" }
http_request{ status_code="400" }
http_request{ status_code="500" }
```

여기서 `http_request`가 메트릭 이름이고, `status_code`가 `Label`이다. 위의 6개의 시계열 데이터는 각각 다른 데이터라고 보면 된다. 그럼 여기서 드는 질문이 하나 있을 것이다. 왜 `Label`이 강력한 기능일까?

`Label`을 이용해서, 메트릭에 대한 집계를 할 수 있기 때문이다. 상태코드 2xx에 대한 개수를 보고 싶으면 다음과 같이 쿼리를 만들 수 있다.

```
sum(rate(http_request{status_code="2.."}[5m]))
```

"2.."이라고 표현함으로써, `Label`의 키 `status_code`의 값이 2xx(200, 201 등의 2로 시작하는 3자리 숫자)인 모든 데이터를 집계할 수 있다.

## 4.2 Label 만들어보기

## 4.3 Label을 이용한 집계

## 4.4 Label 사용 시 Tip