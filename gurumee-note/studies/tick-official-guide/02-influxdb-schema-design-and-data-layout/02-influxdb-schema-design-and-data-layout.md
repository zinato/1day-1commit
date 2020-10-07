# 02. InfluxDB 스키마 디자인과 데이터 레이아웃

![logo](../logo.png)

> Influx Data 공식 문서를 요약한 내용입니다. InfluxDB의 핵심 개념과  철학 그리고 기존 SQL과의 차이점을 비교합니다.


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



## Shard Group Duration 관리