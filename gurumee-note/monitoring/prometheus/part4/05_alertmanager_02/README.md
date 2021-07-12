# Alertmanager란 무엇인가 (2) 

![logo](../../logo.png)

## 개요

`Prometheus`의 알람은 크게 2가지 부분으로 나눌 수 있다.

* 알람 규칙을 정의하는 Alerting Rule
* 생성된 알람을 3자에 전달해주는 Alertmanager

이 문서에서는 `Prometheus`에서 전달된 알람을 제 3자, `Slack`, `Email` 등으로 전달하는 `Alertmanager`에 대해서 다룰 예정이다. 이번 장에서 다음 내용들을 살펴볼 것이다.

1. 특정 리시버로 알람 통보하기
2. 알람 통보 반복하기
3. 알람 통보 조절하기
4. 특정 알람 발생하면 알람 통보 억제하기
5. 특정 알람 통보 멈추기

또한 현재 문서에서 진행되는 실습들은 편의성을 위해 `Docker` 환경에서 진행하나, 실제 서버 환경에서도 크게 다르지 않으니 거의 동일하게 작업할 수 있다. 관련 코드는 다음 링크를 참고하길 바란다.

* 이번 장 코드 : [https://github.com/gurumee92/gurumee-prometheus-code/tree/master/part4/ch05](https://github.com/gurumee92/gurumee-prometheus-code/tree/master/part4/ch05)

## 알람 통보 더 자세히 알기 (1) 라우팅

## 알람 통보 더 자세히 알기 (2) 반복

## 알람 통보 더 자세히 알기 (3) 조절

## 알람 통보 더 자세히 알기 (4) 억제

## 알람 통보 더 자세히 알기 (5) 사일런싱


