# 01. InfluxDB 핵심 개념

![logo](../logo.png)

> Influx Data 공식 문서를 요약한 내용입니다. InfluxDB의 핵심 개념과 SQL과의 차이점을 비교합니다.


## InfluxDB란?

`InfluxDB`는 데이터를 저장하고 읽는데 높은 성능을 가진 `Time Series Database(시계열 데이터베이스)`이다. `TICK 스택`을 통합하는 컴포넌트로써, 센서 데이터, 애플리케이션 메트릭 등의 무수히 많은 양의 타임스탬프 데이터가 포함된 모든 데이터(시계열 데이터)에 대한 백업 저장소로 사용된다. `InfluxDB`의 주요 특징들은 다음과 같다.

- 시계열 데이터를 위한 데이터 저장소. `TSM`엔진을 사용하여, 높은 수집 속도와 데이터 압축을 할 수 있다.
- `Golang`으로 작성되었으며, 다른 외부 의존성 없이 단일 바이너리 파일로 컴파일되어 있다.
- HTTP API에서 제공하는 write/query는 간단하면서도 성능이 좋다.
- `Graphite`, `Collectd`, `OpenTSDB` 같은 다른 데이터 수집 저장소와 통합할 수 있는 플러그인을 제공한다.
- `SQL`과 문법이 비슷하여 쉽게 데이터를 쿼리 및 집계할 수 있다.
- `Tag`라는 것을 이용하여, 데이터를 더 빠르고 효율적으로 인덱싱할 수 있다. 
- `Retention Policy`라는 것을 이용하여, 오래된 데이터를 효율적으로 관리할 수 있다.
- `Continuous Query`라는 것을 이용하여, 자주 일어나는 집계 쿼리를 더 효율적으로 자동화할 수 있다.


## 핵심 개념

`InfluxDB`는 `Time Series Data(시계열 데이터)`를 전문적으로 다루는 데이터베이스이다. 주요 핵심 키워드는 다음과 같다.

* database
* retention policy
* measurement
* tag
* field
* timestamp
* series
* point

공식 문서에서 제공하는 예제와 함께 각각을 알아보자. 다음 데이터는 "특정 위치"의 "과학자별"로 "나비의 수"와 "꿀벌의 수"를 "시간"별로 센 것을 나타내고 있다. 이 데이터 셋들의 이름을 "census"라고 하자.

name = "census"
| time | butterflies | honeybees | location | scientist |
| :-- | :--: | :--: | :--: | :--: |
| 2015-08-18T00:00:00Z | 12 | 23 | 1 | langstroth |
| 2015-08-18T00:00:00Z | 1 | 30 | 1 | perpetua |
| 2015-08-18T00:06:00Z | 11 | 28 | 1 | langstroth |
| 2015-08-18T00:06:00Z | 3 | 28 | 1 | perpetua |
| 2015-08-18T05:54:00Z | 2 | 11 | 2 | langstroth |
| 2015-08-18T06:00:00Z | 1 | 10 | 2 | langstroth |
| 2015-08-18T06:06:00Z | 8 | 23 | 2 | perpetua |
| 2015-08-18T06:12:00Z | 7 | 22 | 2 | perpetua |

**measurement**

여기서 데이터 셋을 `measurement`라고 한다. `measurement`는 `measurement`의 이름과, `timestamp`, `fields`, `tags`로 이루어져 있다. "census"가 바로 `measurement`이다. 

**timestamp**

`InfluxDB`는 앞서 언급했듯이 시계열 데이터를 다루는데 주 목적이 있다. 따라서 시간은 매우 중요하다. `InfluxDB`에 저장되는 데이터는 모두 시간을 가지고 있어야 한다. 이 때 시간을 `timestamp`라고 한다.

**tag**

태그는 키와 값으로 이루어져 있다. `measurement` 내 공통된 속성들이라고 말할 수 있다. 위에 "census"라는 `measurment`에서는 `location`과 `scientist`가 태그 키, 그리고 1, 2가 `location` 태그의 값들, "langstroth", "perpetua"가 `scientist` 태그의 값들이다.

태그는 인덱싱이 된다. 쉽게 생각하면, `InfluxDB` 데이터를 쿼리할 때, `SQL`처럼 `WHERE` 절로 데이터를 필터링할 수 있는데 이 때, `tag`로 그 값을 필터링하면 인덱싱되어 더 빠른 검색 결과를 얻을 수 있다.

**field**

필드는 키와 값으로 이루어져 있다. `measurement`에서 측정되는 실제 값이라고 말할 수 있다. 위에 "census"라는 `measurement`에서는 `butterflies`와 `honeybees`가 필드의 키이다. 그리고 해당 값들이 필드의 값들이다. 

"2015-08-18T00:00:00Z 시간에, 위치 1에 있는 langstroth 과학자가 센 나비의 수는 12이다." 이 때 나비의 수가 실제 유의미한 값을 표현한다. 그것이 바로 필드이다. 

**series**

`series`는 데이터 뭉치이다. `series`는 `measurement`, `tag set`, `field key`로 구성되어 있다. 위의 "census"라는 `measurement`는 총 8개의 `seiries`를 가지고 있다. 

| Series number | Measurement | Tag set | Field key |
| :--: | :--: | :-- | :--: |
| series 1 | census | location = 1,scientist = langstroth | butterflies |
| series 2 | census | location = 2,scientist = langstroth | butterflies |
| series 3 | census | location = 1,scientist = perpetua | butterflies |
| series 4 | census | location = 2,scientist = perpetua | butterflies |
| series 5 | census | location = 1,scientist = langstroth | honeybees |
| series 6 | census | location = 2,scientist = langstroth | honeybees |
| series 7 | census | location = 1,scientist = perpetua | honeybees |
| series 8 | census | location = 2,scientist = perpetua | honeybees |

**point**

`point`는 한 `timestamp`에 찍힌 데이터를 표현한다. 이를 테면 다음 처럼 말이다.

```
name: census
-----------------
time                    butterflies honeybees   location    scientist
2015-08-18T00:00:00Z    1           30          1           perpetua
```

마지막으로, `measurement`가 모이면 `database`가 된다. 또한 각 `database`는 데이터를 저장하는 기간을 설정해야 하는데, 그것이 바로 `retention policy`이다. 아무것도 설정하지 않으면 "autogen"으로 설정된다.


## SQL과의 차이점

`InfluxDB`는 시계열 데이터베이스임에도 `SQL`을 무척 닮아있다. 개인적인 생각으로는 그래서 쿼리 이름도 `InfluxQL`이다. 뭐 1.8부터는 `Flux`라는 새로운 쿼리 언어를 사용할 수 있으며 2.0 에서는 `InfluxQL`은 deprecated되고 `Flux`로 데이터를 쿼리한다. 근데 `Flux`는 `SQL`과는 사뭇 달라서, 일단 여기서는 배제한다. 일단 `InfluxQL`을 잠깐 할펴보자.

```
SELECT * FROM "my_db"."autogen"."census" WHERE "location"=1 AND time > 2015-08-18T00:06:00Z
```

이런 식으로 쿼리할 수 있다. `SQL`과 정말 많이 닮았다. 뭐 다양한 집계 함수도 제공한다. 무엇보다 시간에 대해서 쿼리하는 것이 좋다.

```
SELECT * FROM "my_db"."autogen"."census" WHERE "location"=1 AND time > now() - 3h
```

이런 식으로 현재 시간 기준으로 3시간 이전 데이터를 쿼리할 수도 있다. 그러나 `InfluxQL`는 단점도 존재한다. `InfluxQL`은 서브쿼리, JOIN 기능이 없다. 그래서 여러 메트릭에 대해 알림을 처리하고 싶을 때가 있다. 이 때, `InfluxDB`와 `InfluxQL`만으로는 해결이 되지 않아서 `Kapacitor`의 `TICKscript`로 데이터를 JOIN을 흉내내곤 한다. 

> Flux에서는 JOIN을 지원한다. 그러나 형태가 SQL보다는 함수형 프로그래밍에 가깝다.

그리고 용어나 쓰는 개념이 살짝 살짝 틀리다. `SQL`에서는 보통 테이블, 열, 행으로 데이터를 나눈다. 다음처럼 말이다.

```
+---------+---------+---------------------+--------------+
| park_id | planet  | time                | #_foodships  |
+---------+---------+---------------------+--------------+
|       1 | Earth   | 1429185600000000000 |            0 |
|       1 | Earth   | 1429185601000000000 |            3 |
|       1 | Earth   | 1429185602000000000 |           15 |
|       1 | Earth   | 1429185603000000000 |           15 |
|       2 | Saturn  | 1429185600000000000 |            5 |
|       2 | Saturn  | 1429185601000000000 |            9 |
|       2 | Saturn  | 1429185602000000000 |           10 |
|       2 | Saturn  | 1429185603000000000 |           14 |
|       3 | Jupiter | 1429185600000000000 |           20 |
|       3 | Jupiter | 1429185601000000000 |           21 |
|       3 | Jupiter | 1429185602000000000 |           21 |
|       3 | Jupiter | 1429185603000000000 |           20 |
|       4 | Saturn  | 1429185600000000000 |            5 |
|       4 | Saturn  | 1429185601000000000 |            5 |
|       4 | Saturn  | 1429185602000000000 |            6 |
|       4 | Saturn  | 1429185603000000000 |            5 |
+---------+---------+---------------------+--------------+
```

`InfluxQL`은 시리즈 별로 나뉜다. 다음처럼 말이다.

```
name: foodships
tags: park_id=1, planet=Earth
time			               #_foodships
----			               ------------
2015-04-16T12:00:00Z	 0
2015-04-16T12:00:01Z	 3
2015-04-16T12:00:02Z	 15
2015-04-16T12:00:03Z	 15

name: foodships
tags: park_id=2, planet=Saturn
time			               #_foodships
----			               ------------
2015-04-16T12:00:00Z	 5
2015-04-16T12:00:01Z	 9
2015-04-16T12:00:02Z	 10
2015-04-16T12:00:03Z	 14

name: foodships
tags: park_id=3, planet=Jupiter
time			               #_foodships
----			               ------------
2015-04-16T12:00:00Z	 20
2015-04-16T12:00:01Z	 21
2015-04-16T12:00:02Z	 21
2015-04-16T12:00:03Z	 20

name: foodships
tags: park_id=4, planet=Saturn
time			               #_foodships
----			               ------------
2015-04-16T12:00:00Z	 5
2015-04-16T12:00:01Z	 5
2015-04-16T12:00:02Z	 6
2015-04-16T12:00:03Z	 5
```

차이를 알겠는가. 제일 중요한 차이는 `InfluxQL`에서는 "시간"이란 것이 반드시 데이터에 포함되어 있어야 한다는 것이다. 

그리고 제일 중요한 다른 한 가지는 철학에 있다. `SQL`은 보통 데이터 CRUD 즉 만들고, 읽고, 수정하고, 삭제하는 기능에 초점을 맞춘다. 그러나 **`InfluxDB`는 읽고 쓰는데 초점을 맞춘다.** 업데이트, 삭제도 할 수야 있겠지만 그런 용도로 쓰면 성능 혹은 효율이 좋지 않다. 심지어 삭제의 경우는 데이터 단일 삭제가 아닌 `measurment`, `database` 단위로 `drop` 시키는 것이 오히려 권장 사항이다. 

대용량 시계열 데이터를 읽고 쓰는 것. 그것이 `InfluxDB`의 중요 포인트라고 할 수 있다.