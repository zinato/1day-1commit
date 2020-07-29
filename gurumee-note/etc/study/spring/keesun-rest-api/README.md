백기선의 스프링 기반 REST API 개발
===================

> 백기선님의 인프런 강좌 "스프링 기반 REST API 개발"을 보고 정리한 문서들입니다. 강의는 개발자의 피와 땀의 결실이라고 생각합니다. 꼭 인프런 강의를 듣고, 참고 사항으로 보셨으면 좋겠습니다.

Contents
-----------------
1. 학습 목표
2. 요구 사항
3. 목차


## 학습 목표

백기선님의 이번 강좌는 TDD 기반으로 Spring HATEOAS, Spring REST Docs 를 활용하여 Self-Description Message + HATEOAS 한 API 서버를 만드는 것입니다. 


## 요구 사항

1. Java 8
2. Maven 3.6
3. Spring Boot 2.x
    - Spring Web
    - Spring HATEOAS
    - Sprint REST Docs
    - Spring Data JPA
    - Spring Security OAuth


## 목차

1. [REST API 및 프로젝트 소개](./ch01.md)
2. [이벤트 생성 API 개발 #1](./ch02_01.md)
3. [이벤트 생성 API 개발 #2](./ch02_02.md)
4. [HATEOAS와 Self-Descriptive Message 적용 #1](./ch03_01.md)
5. [HATEOAS와 Self-Descriptive Message 적용 #2](./ch03_02.md)
6. [이벤트 조회 및 수정 REST API 개발](./ch04.md)
7. [REST API 보안 적용](./ch05.md)

프로젝트 소스 코드는 다음 URL을 참고하세요. 각 챕터마다 브랜치를 남겨놓았으니 비교해보면서 보시면 훨씬 편할 거에요!

* [소스 코드](https://github.com/gurumee92/keesun-rest-api)