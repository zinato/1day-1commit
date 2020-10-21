# 야! 너도 TICK 할 수 있어 (1) - TICKscript 소개

![logo](../logo.png)

> Kapacitor 공식 가이드 "TICKscript language reference > introduction"를 읽고 정리한 문서입니다. 자세한 내용은 [여기](https://docs.influxdata.com/kapacitor/v1.5/tick/introduction/)를 참고하세요.

## 개요

`Kapacitor`는 `InfluxData Platform TICK 스택`의 구성 중 하나로 데이터 프로세싱 및 알림을 처리하는 컴포넌트이다. `Kapacitor`는 `TICKscript`라는 도메인 특화 언어(Domain Specific Language 이하 "DSL")를 사용하여, `InfluxDB`에 저장된 데이터를 추출하여 원하는 형태로 변환할 수 있다. 심지어 변환된 데이터에 대해서 특정 임계값 혹은 조건을 만족시키지 않을 때, `Slack` 등의 서드 파티로 알림을 보낼 수 있다. 

`Kapacitor`의 DSL인 `TICKscript`는 다음과 같은 특징이 있다.

* `.tick` 형식으로 파일을 만들어야 한다.
* Node 단위로 데이터를 처리할 수 있다.
* Node와 Node 간 Pipeline을 연결할 수 있다.


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

최상위 노드를 제외한 노드들은, 모두 상위 노드의 체이닝 메서드 호출로부터 생성된다. 즉, 각 노드는 `stream` 혹은 `batch` 데이터 스트림으로부터 생성("wants")되며 또한, `stream` 혹은 `batch` 형태의 데이터 스트림을 제공("provides")할 수 있다. 공식 가이드에서는 일반적인 "wants-provides" 패턴  4가지를 다음과 같이 정의하고 있다.



## 파이프라인

## 기본 예제