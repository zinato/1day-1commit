# 야! 너두 TICK 할 수 있어 (1) - TICKscript 문법

![logo](../logo.png)

> Kapacitor 공식 가이드 "TICKscript language reference > syntax"를 읽고 정리한 문서입니다. 자세한 내용은 [여기](https://docs.influxdata.com/kapacitor/v1.5/tick/syntax/)를 참고하세요.

## 개요

## 문법

키워드는 다음과 같다.

* TURE - "true"와 같다
* FALSE - "false"와 같다
* AND - 논리 연산자 AND
* OR - 논리 연산자 OR
* lambda: - 람다 식 선언
* var - 변수 선언
* dbrp - 데이터베이스 선언

연산자는 다음과 같다.

* 산술 연산자 : +, -, *, /
* 비교 연산자 : ==, !=, <, <=, >, >=, =~(정규 표현식 일치하는가), !~(정규 표현식 불일치하는가)
* 논리 연산자 : !, AND, OR
* 체이닝 연산자 : |(체이닝 메서드), .(프로퍼티 메서드), @(UDF 함수)

타입은 다음과 같다.

* string
* duration
* int
* float
* lambda
  
변수 선언은 다음과 같다.

```
// 문자열
var region1 = 'EMEA'
var old_standby = 'foo' + 'bar'
var query1 = 'SELECT 100 - mean(usage_idle) AS stat FROM "telegraf"."autogen"."cpu" WHERE cpu = \'cpu-total\' '
var query2 = '''SELECT 100 - mean(usage_idle) AS stat FROM "telegraf"."autogen"."cpu" WHERE cpu = 'cpu-total' '''
var query3 = '
    SELECT 100 - mean(usage_idle)
    AS stat
    FROM "telegraf"."autogen"."cpu"
    WHERE cpu = \'cpu-total\'
    '

// 숫자 
var my_int = 6
var my_float = 2.71828
var my_octal = 0400

// BOOL 
var true_bool = TRUE

// 기간
var span = 10s
var frequency = 10s
```

대 소문자 식별, 변수에 변수 대입, 무엇보다 타입 선언 및, 변수 할당까지 가능한 것이 눈에 띈다. 또한, 기간 역시 설정할 수 있다. 기간 리터럴은 다음과 같다.

* u - microsecond
* ms - millisecond
* s - second
* m - minute
* h - hour
* d - day
* w - week


## 노드 타입의 분류

노드 타입 종류는 다음과 같다.

* 데이터 소스 정의 노드 - batch, stream
* 데이터 정의 노드 - from ,query
* 데이터 조작 노드 - default, sample, shift, where, window
* 데이터 처리 노드 
  * 데이터 구조 변경 혹은, 결합 처리 노드: combine, eval, groupBy, join, union
  * 데이터 변환 노드 : delete, derivative, flatten, influxQL, stateCount, stateDuration, stats
* 알림 노드 - alert, deadman, httpOut, httpPost, influxDBOut, k8sAutoscale, kapacitorLoopback, log
* UDF 노드 - UDF
* 사용해선 안되는 노드 - noOp

## TICKscript 안에서의 InfluxQL

## Lambda 식