# Recording Rule

![logo](../../logo.png)

## 개요

이번 장에서는 `PromQL`의 쿼리 성능을 보다 높여줄 수 있는 `Recording Rule`에 대해 알아볼 것이다. 이 문서에서는 편의성을 위해서 `Docker` 환경에서 진행할 것이나, 실제 서버 환경에서 어떻게 작업해야 하는지까지 최대한 다루도록 하겠다. 관련 코드는 다음 링크를 참고하길 바란다.

* 이번 장 코드 : [https://github.com/gurumee92/gurumee-prometheus-code/tree/master/part4/ch02](https://github.com/gurumee92/gurumee-prometheus-code/tree/master/part4/ch02)

## Recording Rule은 무엇인가?

`Prometheus` 공식 문서에 따르면, `Recording Rule`의 정의는 다음과 같다.

> 기록 규칙은 일관된 이름 체계를 사용함으로써, 한 눈에 규칙을 쉽게 파악할 수 있다. 이것은 또한 부정확하거나 무의미한 계산을 두드러지게 함으로써 실수를 피한다.

개인적으로는 다소 뜬금 없는 표현이라고 생각되는데 쉽게 생각해서, `PromQL`을 통해서 집계한 시계열 데이터의 이름을 붙인다고 생각하면 된다. 예를 들면 지난 장 "[서비스 메트릭 모니터링하기 (1) prometheus-nginxlog-exporter]()"에서 `Nginx`의 1분간 평균 응답 시간에 대한 패널을 어떻게 구축했는가?

![grafana > response time](./01.png)

해당 패널은 다음 쿼리로 구성된다.

```
sum by (method, request_uri)(rate(nginx_http_response_time_seconds_hist_sum[1m])) 
/ 
sum by (method, request_uri)(rate(nginx_http_response_time_seconds_hist_count[1m]))
```

위의 쿼리는 단순한 수준에 속하지만 `PromQL`을 잘 모르는 경우에는 무슨 쿼리인지 도통 알아보기가 힘들다. 쿼리가 복잡해지면 복잡해질수록, 해석하기가 점점 히들어질 것이다. 이를 위해서 `Recording Rule`을 사용하면, 다음과 같이 단순하게 표현 가능하다.

```
job:prometheus_nginxlog_exporter:avg_response_time_1m
```

다음과 같이 쿼리하게 되면 적어도, 패널이 쿼리하는 데이터가 1분간 집계된 평균 응답 시간이라는 것을 쉽게 알 수 있다. 사실은 이런 효과보다도, `InfluxDB`의 `Continuous Query` 같이, 복잡한 집계 쿼리를 배치를 돌듯 쿼리 결과를 하나의 시계열 데이터로 저장하여 쿼리 성능을 높이는 데 더 큰 의미가 있다고 생각한다.

## Recording Rule 사용법 (1) 작성 방법과 관례

## Recording Rule 사용법 (2) 권장되는 상황

## Recording Rule 사용법 (3) 권장되지 않는 상황