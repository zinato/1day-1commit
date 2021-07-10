# Alertmanager란 무엇인가

![logo](../../logo.png)

## 개요

`Prometheus`의 알람은 크게 2가지 부분으로 나눌 수 있다.

* 알람 규칙을 정의하는 Alerting Rule
* 생성된 알람을 3자에 전달해주는 Alertmanager

이 문서에서는 `Prometheus`에서 전달된 알람을 제 3자, `Slack`, `Email` 등으로 전달하는 `Alertmanager`에 대해서 다룰 예정이다. 또한 현재 문서에서 진행되는 실습들은 편의성을 위해 `Docker` 환경에서 진행하나, 실제 서버 환경에서도 크게 다르지 않으니 거의 동일하게 작업할 수 있다. 관련 코드는 다음 링크를 참고하길 바란다.

* 이번 장 코드 : [https://github.com/gurumee92/gurumee-prometheus-code/tree/master/part4/ch04](https://github.com/gurumee92/gurumee-prometheus-code/tree/master/part4/ch04)

## Alertmanager 설치

`AWS EC2 Amazon Linux2` 서버 환경에서 `Alertmanager`는 다음 명령어로 설치할 수 있다.

```bash
# 압축 파일 설치
$ wget https://github.com/prometheus/alertmanager/releases/download/v0.22.2/alertmanager-0.22.2.linux-amd64.tar.gz

# 압축 파일 해제
$ tar -xf alertmanager-0.22.2.linux-amd64.tar.gz

# 해제된 파일 /home/ec2-user/apps/alertmanager 경로로 변경
$ mv alertmanager-0.22.2.linux-amd64.tar.gz ~/apps/alertmanager
```

실행을 하고 싶다면 터미널에 다음을 입력하면 된다.

```bash
# 설치 경로 이동
$ cd ~/apps/alertmanager

# alertmanager 실행
$ ./alertmanager
```

`Prometheus`에서 설정을 하고 싶다면, `Prometheus`가 실행되는 서버에서 다음과 같이 설정하면 된다.

```yml
# my global config
global:
  scrape_interval:     15s # By default, scrape targets every 15 seconds.
  evaluation_interval: 15s # By default, scrape targets every 15 seconds.

# ...

# alert
alerting:
  alertmanagers:
  - scheme: http
    static_configs:
    - targets:
      # alertmanager ip:port
      - "alertmanager:9093"
```

## 알람 라우팅

## 알람 조절과 반복

## 알람 억제와 사일런싱

## 알람 통보 (1) 슬랙에 전달하기

## 알람 통보 (2) API 서버에 전달하기