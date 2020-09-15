# TICK 스택 소개

![logo](../logo.png)

> Influx Data 공식 문서를 요약한 내용입니다.

## InfluxData Platform

`InfluxData Platform`은 현재 메트릭과 이벤트를 위한 타임 시리즈 플랫폼 중 하나이다. 크게 1.x와 2.x로 나눌 수 있다.

**InfluxData 1.x**

이른 바, `TICK 스택`이라 불리우고 있다. 구성 요소는 다음과 같다.

![TICK 스택](./01.png)

* Telegraf : 데이터 수집기이다.
* InfluxDB : 데이터 저장소이다.
* Chronograf : 데이터를 시각화한다. 유저 인터페이스로 보면 된다.
* Kapacitor : 데이터를 처리하고 알림을 만든다. 위의 그림에서는 안보이지만 InfluxDB와 통신하면서, 데이터를 처리한다.

시스템 모니터링 플랫폼에서 큰 사랑을 받았다. 주관적인 생각이지만 현재는 시스템 메트릭 모니터링은 `Prometheus`, 로그 시스템 분석 등은 `Elastic 스택`에게 빼앗긴 듯 하다. 그래도 상용 서비스에는 많이 쓰이는 것 같다. 나도 회사 업무에 이 스택을 안썼다면, 이렇게 공부하고 있지 않을 것이다.

**InfluxData 2.x**

2.x로 들어서부터 `InfluxDB`, `Chronograf`, `Kapacitor`가 합쳐졌다. 알림부터 데이터 프로세싱, UI까지 하나의 컴포넌트로 다 할 수 있다. (다만, 메트릭 수집기인 `Telegraf`는 여전히 존재한다.) 

스크립트 언어도 `InfluxQL`에서 더 강력한 `Flux`로 바뀌었다. 

개인적으로, 1.x는 `Elastic 스택`과, 2.x는 `Prometheus`와 구성 요소가 비슷하다.


## TICK 스택 설치

`TICK 스택` 설치 방법에는 여러가지가 있지만 이 문서에서는, `InfluxData`가 공식적으로 제공하는 샌드박스를 이용해서 설치할 것이다. 이 설치 절차를 진행하려면 다음이 먼저 설치되어 있어야 한다.

* Git
* Docker
* docker-compose

터미널에 다음을 입력한다.

```bash
$ git clone https://github.com/influxdata/sandbox.git
```

프로젝트를 복제 했으면 다음 명령어를 실행하여, `TICK 스택`의 구성 요소들의 도커 컨테이너 설치 및 실행한다.

```bash
# cd into the sandbox directory
$ cd sandbox

# Start the sandbox
$ ./sandbox up
```

명령어를 실행하면 잠시 후에 브라우저에서, "localhost:8888", "localhost:3010"이 띄어질 것이다. 이 둘은 크로노그래프와 튜토리얼이다. 터미널에서 정상적으로 동작하는지 확인하고 싶다면 다음 명령어를 치면 된다.

```bash
$ docker ps
CONTAINER ID        IMAGE                   COMMAND                  CREATED             STATUS              PORTS                                                                    NAMES
81a4760a458b        chrono_config           "/entrypoint.sh chro…"   3 minutes ago       Up 2 minutes        0.0.0.0:8888->8888/tcp                                                   sandbox_chronograf_1
28216e9a1396        kapacitor               "/entrypoint.sh kapa…"   3 minutes ago       Up 3 minutes        0.0.0.0:9092->9092/tcp                                                   sandbox_kapacitor_1
4d36ff50a0c9        telegraf                "/entrypoint.sh tele…"   3 minutes ago       Up 3 minutes        8092/udp, 8125/udp, 8094/tcp                                             sandbox_telegraf_1
2eb6f24a3575        influxdb                "/entrypoint.sh infl…"   3 minutes ago       Up 3 minutes        0.0.0.0:8082->8082/tcp, 0.0.0.0:8086->8086/tcp, 0.0.0.0:8089->8089/udp   sandbox_influxdb_1
a9c37eb82b38        sandbox_documentation   "/documentation/docu…"   3 minutes ago       Up 3 minutes        0.0.0.0:3010->3000/tcp                                                   sandbox_documentation_1
```

이제 `TICK 스택`을 종료해보자.

```bash
$ ./sandbox down
```

[이 곳](https://github.com/influxdata/sandbox)으로 가면 더 많은 샌드 박스 명령어에 대한 정보를 얻을 수 있다. 


## InfluxDB로 Telegraf 수집 데이터 쿼리해보기

**이 절은 샌드 박스를 실행 시켜두어야 한다.**

## Kapacitor + Chronograf 알림 만들기

**이 절은 샌드 박스를 실행 시켜두어야 한다.**