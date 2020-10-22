# 야! 너도 TICK 할 수 있어 (1) - TICKscript 소개

![logo](../logo.png)

> Kapacitor 공식 가이드 "TICKscript language reference > introduction"를 읽고 정리한 문서입니다. 자세한 내용은 [여기](https://docs.influxdata.com/kapacitor/v1.5/tick/introduction/)를 참고하세요.

## 목차

  - [개요](#개요)
  - [노드](#노드)
  - [파이프라인](#파이프라인)
  - [기본 예제](#기본-예제)


## 개요

`Kapacitor`는 `InfluxData Platform`의 `TICK 스택`의 구성 중 하나로 데이터 프로세싱 및 알림을 처리하는 컴포넌트이다. `Kapacitor`는 `TICKscript`라는 도메인 특화 언어(Domain Specific Language 이하 "DSL")를 사용하여, `InfluxDB`에 저장된 데이터를 추출하여 원하는 형태로 변환할 수 있다. 심지어 변환된 데이터에 대해서 특정 임계값 혹은 조건을 만족시키지 않을 때, `Slack` 등의 서드 파티로 알림을 보낼 수 있다. 

`Kapacitor`의 DSL인 `TICKscript`는 다음과 같은 특징이 있다.

* `.tick` 형식으로 파일을 만들어야 한다.
* "노드" 단위로 데이터를 처리할 수 있다.
* 노드와 노드를 잇는 "파이프라인"을 만들 수 있다.


## 노드

노드는 `TICKscript`를 구성하는 기본 요소이다. 노드는 "Property Method"와 "Chaining Method"를 가지고 있다. 프로퍼티 메서드(Property Method)는 "."으로 호출할 수 있으며, 노드 안의 데이터를 조작하는데 쓰인다. 반환 타입은 같은 노드이다. 반면에, 체이닝 메서드(Chaining Method)는 "|"로 호출할 수 있으며 노드에서 다른 노드로 변경할 때 쓰인다.

즉, 노드는 부모 혹은 형제 노드의 체이닝 메소드로부터 만들어진다. 자주 쓰이는 노드 타입은 다음과 같다.

* stream
* batch
* query
* from
* eval
* alert

`TICKscript`의 노드들 중 최상위 노드는 `batch`와 `stream`이다. 특별한 아규먼트 없이 만들 수 있으며, 프로퍼티 메서드들을 통해서, 노드 내부의 데이터를 처리할 수 있다. 

최상위 노드라는 것은 `Kapacitor`에서 어떤 데이터를 처리하기 위해서 태스크를 정의해야 한다고 할 때, 데이터를 가져올 때 `batch` 혹은 `stream` 노드를 사용해야 한다는 뜻이다. 모든 `TICKscript`는 데이터베이스, 리텐션 폴리시 선언, 변수 선언을 제외하고, `batch` 노드 혹은 `stream` 노드가 최상단에 선언되게 된다. 추후, [기본 예제](#기본-예제) 절에서 이를 살펴보자.

최상위 노드를 제외한 노드들은, 모두 상위 노드의 체이닝 메서드 호출로부터 생성된다. 즉, 각 노드는 `stream` 혹은 `batch` 데이터 스트림이 필요("wants")하다. 또한, `stream` 혹은 `batch` 형태의 데이터 스트림을 만들 수 있다("provides"). 공식 가이드에서는 일반적인 "wants-provides" 패턴  4가지를 다음과 같이 정의하고 있다.

* wants batch, provide a stream - 예: 평균, 최대, 최솟값을 계산할 때
* wants batch, provide a batch - 예: 배치 데이터에서 지나치게 높거나 낮은 값(outlier)을 식별할 때
* wants stream, provide a batch - 예: 유사 데이터 포인트들을 그루핑할 때
* wants stream, provide a stream - 예: 개별 포인트마다 값에 logarithm같은 수학적 함수를 적용할 때


## 파이프라인

모든 `TICKscript`는, 하나 혹은 그 이사의 파이프라인으로 구성되어 있다. 노드와 노드가 논리적으로 연결되어 있는 것을 `edge`라고 표현하고, 이 `edge`들의 집합을 `파이프라인`이라고 볼 수 있겠다. 파이프라인이 실행될 때 연결된 이전 노드로 돌아가지는 못한다.

또한, 별 개의 파이프라인은 `join` 혹은 `union` 노드로 결합하여 사용할 수 있다. 이 것이 `TICKscript`의 장점인데, 기본적으로 `InfluxQL`에서 메저먼트끼리, "JOIN"이 지원되지 않는다. 메저먼트는 서로 독립적인데, 실제 시스템을 운영하다 보면, 여러 지표들의 조합이 한 에러를 나타내는 경우들이 있다. 이럴 때, `TICKscript`의 `join` 혹은 `union` 노드를 사용하여, 독립된 메저먼트로부터 생성된 데이터 스트림을 결합시킬 수 있다.

> 참고!
> 
> InfluxDB 1.8에서부터 지원되는 "Flux"로 인하여 장점이 많이 사라진 느낌입니다. 실제로 InfluxDB 2.0에서는 데이터 처리 및 변환 같은 경우 "Flux"로만 할 수 있습니다.("InfluxDB-Kapacitor-Chronograf"가 통합됨) 따라서 "계속 Kapacitor를 써야 하는 것인가"에 대한 의문이 남습니다.

**stream인가 batch인가**

`stream`은 포인트 별로 `InfluxDB`에 도착할 때마다, 해당 스트림을 읽어온다. `Kapacitor` 설정에서 `InfluxDB`를 연결했을 때, subscription이 생성되는 것이 바로 이 때문인 것 같다. `batch`는 `InfluxDB` 내에서 처리되어 저장된 데이터 프레임을 읽어온다. 즉, `stream`은 원본 데이터 스트림을, `batch`는 이미 `InfluxDB` 내에 저장된 데이터 포인트를 가져온다.

이 때문인지, 둘의 쓰임새가 다르다. 공식 문서에 따르면, 긴 기간 동안의 거대한 데이터 포인트들은 `batch`로, 짧은 기간 동안의 데이터 포인트를 가져오는 것은 `stream`이 선호된다고 적혀 있다. 

`stream`으로 긴 기간의 데이터 포인트를 가져오게 되면, 잠재적으로 필요 없는 거대한 데이터가 메모리에 잡히기 때문에, 메모리에 과부하가 일어난다. 반면에 `batch`는 조회할 때, disk에 저장된 데이터를 가져오기 때문에, 데이터베이스에 과부하가 일어난다. 따라서 짧은 기간의 데이터를 `batch`로 여러 번 조회하는 것을 권장하지 않고 있다.

**파이프라인을 조금만 더 알아보자**

`Kapacitor`의 파이프라인들은 모두 방향성, 비순환 그래프와 같다. 각 `edge`는 데이터 플로우를 각 방향으로 흐르게 하며, 파이프라인에 순환될 수 없게 한다. (코드로 강제적으로 순환하게끔은 만들 수 있지만, 공식 문서에서는 절대 그렇게 사용해선 안된다고 경고하고 있다.) 보통은 다음의 `edge`들로 시작하게 된다.

* `stream -> from()` - 한 번에 단일 데이터 포인트만 전송한다.
* `batch -> query()` - 한 번에 데이터 뭉치(쿼리된 데이터 포인트 집합)를 전송한다.

이 때 주의할 점이 있다. `Kapacitor`는 `TICKscript` 형식이 맞는지 체크하고, `edge`들을 실행한다. 그러나 간혹 형식엔 맞지만, 파이프라인이 유효하지 않을 때가 있다. 이 때 다음과 같은 Runtime 에러가 발생한다.

```
...
[cpu_alert:alert4] 2017/10/24 14:42:59 E! error evaluating expression for level CRITICAL: left reference value "usage_idle" is missing value
[cpu_alert:alert4] 2017/10/24 14:42:59 E! error evaluating expression for level CRITICAL: left reference value "usage_idle" is missing value
...
```

위 에러는 파이프라인에서, 데이터 포인트의 `field` 값(usage_idle)이 소실된 경우를 보여준다. 이는 `eval` 노드를 사용할 때, 다음 노드로 데이터 스트림을 전달하면서 필드 "usage_idle"이 유지되지 못해서 필드 자체가 소실 된 경우이다. 이 경우, 프로퍼티 메서드인 `keep`을 사용하는 것이 좋다.

> 참고!
> 태그는 keep을 쓰지 않아도 무조건 파이프라인에서 유지된다고 합니다. 다만, 필드들은, 꼭 유지해야할 목록을 추려서 keep을 호출하는 것이 좋을 것 같습니다.


## 기본 예제

이번 절에서는 다음의 기본적인 2가지 스크립트를 확인해 볼 것이다. 

1. `stream -> from` pipeline
2. `batch -> query` pipeline

바로 시작하자.

**stream -> from pipeline**

ex01.tick
```
// database retension policy 설정
dbrp "telegraf"."autogen"

// 스트림 노드 선언
stream
    // from 노드로 스트림 전송
    |from()
        // 해당 포인트에서 "cpu"라는 메저먼트 추출
        .measurement('cpu')
    // httpOut 노드로 스트림 전송
    |httpOut('dump')
```

기본적으로 `TICKscript`는 데이터베이스와, 리텐션 폴리시의 정보를 최상단에 적어주어야 한다.

```
dbrp "telegraf"."autogen"

...
```

이 문구는 연결된 `InfluxDB`에서 "telegraf"란 데이터베이스에서 "autogen" 리텐션 폴리시를 갖는 데이터 스트림을 가져온다. 

```
...

stream
    |...
```

이 문구는 최상위 노드인 `Stream` 노드로 선언한 것이다. 그러면, "telegraf"."autogen"에 데이터 포인트가 쌓일 때마다 한 포인트씩, 해당 `TICKscript`를 사용하는 태스크는 이를 가져오게 된다


```
...

stream
    |from()
        .measurement('cpu')
    |...
```

`from()` 체이닝 메서드는 `Stream` 노드에서 `From` 노드로 변경한다. `From` 노드는 데이터 스트림에서 원하는 조건을 선택해서 추출할 수 있다. 즉 데이터 프로세싱이 발생한다. 참고로 데이터베이스, 리텐션 폴리시 설정도 가능한데 그러면 위의 "dbrp" 구문은 생략해도 좋다. 여기선, measurement "cpu"인 데이터 포인트만 추출한다.

```
...
stream
    |from()...
    |httpOut('dump')
```

이제 `httpOut()` 체이닝 메서드는 `From` 노드에서 `HttpOut`노드로 변경한다. `HttpOut` 노드는 데이터를 `kapacitor/v1/task/:task_id/:field`로 확인할 수 있다. 만약 태스크 이름 "example"이라고 할 때, 위 추출한 값은 다음과 같이 API로 접근할 수 있다.

```
$ curl -XGET http://<kapacitor address>/kapacitor/v1/task/example/dump
```

이 스크립트는 2개의 `edge`가 존재한다.

* `Strean` -> `From`
* `From` -> `HttpOut`


**batch -> query pipeline**

ex02.tick
```
// 배치 노드 생성
batch
    // 쿼리 노드 생성
    |query('SELECT * FROM "telegraf"."autogen".cpu WHERE time > now() - 10s')
        // 10 초 간격의
        .period(10s)
        // 매 10초마다.
        .every(10s)
    // HttpOut 노드 생성
    |httpOut('dump')
```

위의 예를 살펴봤기 때문에 읽는데 무리가 줄을 것이다. 먼저 `Batch` 노드를 생성한다.

```
batch
    |...
```

그 후 체이닝 메서드인 `query()`를 사용해서 `Query` 노드로 데이터 스트림을 변환한다.

```
batch
    |query('SELECT * FROM "telegraf"."autogen".cpu WHERE time > now() - 10s')
        .period(10s)
        .every(10s)
```

이 때, `Query`노드는 연결된 `InfluxDB`에 해당 쿼리를 발생시킨다.

```
'SELECT * FROM "telegraf"."autogen".cpu WHERE time > now() - 10s
```

`period`는 가져오는 데이터 간격을, `every` 노드를 실행하는 시간을 나타낸다.