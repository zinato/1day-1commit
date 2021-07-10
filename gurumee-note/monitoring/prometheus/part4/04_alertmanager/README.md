# Alertmanager란 무엇인가

![logo](../../logo.png)

## 개요

`Prometheus`의 알람은 크게 2가지 부분으로 나눌 수 있다.

* 알람 규칙을 정의하는 Alerting Rule
* 생성된 알람을 3자에 전달해주는 Alertmanager

이 문서에서는 `Prometheus`에서 전달된 알람을 제 3자, `Slack`, `Email` 등으로 전달하는 `Alertmanager`에 대해서 다룰 예정이다. 또한 현재 문서에서 진행되는 실습들은 편의성을 위해 `Docker` 환경에서 진행하나, 실제 서버 환경에서도 크게 다르지 않으니 거의 동일하게 작업할 수 있다. 관련 코드는 다음 링크를 참고하길 바란다.

* 이번 장 코드 : [https://github.com/gurumee92/gurumee-prometheus-code/tree/master/part4/ch04](https://github.com/gurumee92/gurumee-prometheus-code/tree/master/part4/ch04)

## Alertmanager 설치

## 알람 라우팅

## 알람 조절과 반복

## 알람 억제와 사일런싱

## 알람 통보 (1) 슬랙에 전달하기

## 알람 통보 (2) API 서버에 전달하기