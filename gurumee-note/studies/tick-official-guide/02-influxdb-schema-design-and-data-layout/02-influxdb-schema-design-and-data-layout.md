# 02. InfluxDB 스키마 디자인과 데이터 레이아웃

![logo](../logo.png)

> Influx Data 공식 문서를 요약한 내용입니다. InfluxDB에서 추천하는 스키마 디자인과 추천하지 않는 스키마 디자인에 대해서 알아봅니다.
>
> 참고 : InfluxDB 공식 문서 [스키마 디자인과 데이터 레이아웃](https://docs.influxdata.com/influxdb/v1.8/concepts/schema_and_data_layout/)


## 추천하는 스키마 디자인

공식 문서에 따르면, 두 가지 스키마 디자인을 추천하고 있다.

1. `tag`에 메타 데이터를 저장하는 스키마 디자인
2. `tag`, `field` 이름으로 키워드 사용하지 않는 스키마 디자인

먼저, `tag`에 메타 데이터를 넣어야 하는 이유는 **tag는 인덱싱이 가능하기 때문**이다. 인덱싱되지 않는 `field`보다 훨씬 빠르게 쿼리가 가능하다. 가이드에서는 다음의 사항들을 일반적인 상황에서 사용할 것을 권고하고 있다.

* 공통적으로 쿼리하는 메타 데이터는 `tag`에 저장한다.
* `GROUP BY` 절에 쓰일 데이터라면 `tag`에 저장한다.
* `function`에 쓰일 데이터라면 `fields`에 저장한다.
* 숫자 타입의 값은 `fields`에 저장한다. `tag`는 오로지 문자열 타입만 저장할 수 있다.

두 번째 추천 사항은 필수적이진 않다. 다만, `tag` 혹은 `field`에 키워드를 이름으로 저장하면 쿼리할 때 반드시 `""`로 감싸주어야 한다. 키워드가 아니라면 감쌀 필요는 없다. `InfluxQL`의 키워드는 다음과 같다.

```
ALL           ALTER         ANALYZE       ANY           AS            ASC
BEGIN         BY            CREATE        CONTINUOUS    DATABASE      DATABASES
DEFAULT       DELETE        DESC          DESTINATIONS  DIAGNOSTICS   DISTINCT
DROP          DURATION      END           EVERY         EXPLAIN       FIELD
FOR           FROM          GRANT         GRANTS        GROUP         GROUPS
IN            INF           INSERT        INTO          KEY           KEYS
KILL          LIMIT         SHOW          MEASUREMENT   MEASUREMENTS  NAME
OFFSET        ON            ORDER         PASSWORD      POLICY        POLICIES
PRIVILEGES    QUERIES       QUERY         READ          REPLICATION   RESAMPLE
RETENTION     REVOKE        SELECT        SERIES        SET           SHARD
SHARDS        SLIMIT        SOFFSET       STATS         SUBSCRIPTION  SUBSCRIPTIONS
TAG           TO            USER          USERS         VALUES        WHERE
WITH          WRITE
```

또한 `Flux`의 키워드는 다음과 같다.

```
and    import  not  return   option   test
empty  in      or   package  builtin
```

참고적으로, `tag` 혹은 `field` 이름에 영대소문자, _를 제외한 문자가 들어가는 녀석들을 쿼리해야 할 경우에, `InfluxQL`은 `""`로 묶어주어야 하고 `Flux`는 "bracket"이라고 불리우는,  `[""]` 이 녀석으로 묶어주어야 한다.


## 추천하지 않는 스키마 디자인

공식 문서에 따르면 다음의 스키마 디자인은 추천하지 않는다.

1. 너무 많은 시리즈를 사용하는 스키마 디자인
2. 같은 이름의 tag와 field를 사용하는 스키마 디자인
3. measurement의 이름으로 인코딩된 데이터를 저장하는 스키마 디자인
4. 하나의 태그에 하나 이상의 데이터를 저장하는 스키마 디자인

`tag`에 UUID, 해시, 등 많은 정보를 담는 것은 데이터베이스가 감당할 수 없는 시리즈 카디널리티를 야기할 수 있다.

Series Cardinality란 database, measurement, tag set 그리고 field key 조합의 수를 나타낸다. 예를 들어서 다음의 표를 보자.

user.login
| email | status |
| :-- | :-- |
| lorr@influxdata.com | start |
| lorr@influxdata.com | finish |
| marv@influxdata.com | start |
| marv@influxdata.com | finish |
| cliff@influxdata.com | start |
| cliff@influxdata.com | finish |

database=user, measurement=login, tag=[email, status], 여기서, 필드의 값은 이메일이 3개(lorr, marv, cliff), 상태가 2개(start, finish)이다. 이 login measurement의 시리즈 카니ㅓㄹ리티는 3 * 2 = 6 개가 된다.

근데, 태그에 많은 정보가 포함될 수록, 이 카디널리티 수는 증가된다. `InfluxDB` 공식 가이드에 따르면, 10,000,000 개 이상의 시리즈 카디널리티는 OOM을 발생시킬 수 있다고 한다. 

태그와 필드가 같은 스키마 디자인은 기대하지 않은 동작이 발생할 수 있다. 실제로, 필자가 그런 경우에 속했는데, 알림을 위해서 `Kapacitor`를 사용했는데, `StreamNode`를 이용해서 데이터를 조회할 때 태그와 필드가 같을 때 에러가 발생하였다. 그래서 결국 `BatchNode`로 바꿔서 조회하는 것으로 해결을 보았는데 이 때 절실하게 이러한 스키마 디자인이 왜 추천하지 않는지 공감하게 되었다.

세 번째, 네 번째 얘기는 비슷한 맥락이다. 하나의 태그에 인코딩된 정보나, 여러 개의 데이터 정보를 엮어서 저장하는 경우, 쿼리할 때 매우 어렵게 만든다. 예를 살펴보자.

```
# Schema 1 - Query for multiple data encoded in a single tag
> SELECT mean("temp") FROM "weather_sensor" WHERE location =~ /\.north$/

# Schema 2 - Query for data encoded in multiple tags
> SELECT mean("temp") FROM "weather_sensor" WHERE region = 'north'
```

첫 번째 쿼리가 태그에 인코딩 된 정보를 가지고 있는 스키마를 쿼리한 것이다. `WHERE location =~ /\.north$/` 이런 식으로 어렵게 조회를 해야 한다. 여기서는 정규 표현식을 같이 사용해서 쿼리를 한다. 

반면 태그에 하나의 정보만 있는 두 번째 쿼리는 매우 명확하게 쿼리하는 것을 볼 수 있다.


## Shard Group Duration 관리

샤드 그룹이란, `Retention Policy(RP)`에 따라 구성되며, 샤드 기간이라고 하는 특정 기간 간격 내에 있는 타임스탬프와 함께 저장된다. 기본적으로 `RP`에 따라 다음과 같이 잡힌다.

| RP Duration |	Shard Group Duration |
| :-- | :-- |
| < 2 days	| 1 hour |
| >= 2 days and <= 6 months	| 1 day |
| > 6 months | 7 days |

샤드 그룹은 아직까지 깊게 안 봐서 이 정도만 정리한다. 추천 사항만 정리한다. RP 기간에 따라 샤드 그룹 기간을 다음과 같이 잡는 것을 추천하고 있다.

| RP Duration | Shard Group Duration |
| :-- | :-- |
| <= 1 day | 6 hours |
| > 1 day and <= 7 days | 1 day |
| > 7 days and <= 3 months | 7 days |         
| > 3 months | 30 days |
| infinite | 52 weeks or longer |
