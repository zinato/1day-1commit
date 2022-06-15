# 3xx 번대 에러 코드

## 3xx (Redirection) : 요청을 완료하기 위해 유저 에이전트의 추가적인 조치 필요

    - 300 : Multiple Choices 301 Moved Permanently 302 Found
    - 303 : See Other
    - 304 : Not Modified
    - 307 : Temporary Redirect
    - 308 : Permanent Redirect

### 리다이렉션이란 ?

- Redirection : 웹 브라우저는 3xx의 응답결과에 Location 헤더가 있으면 Location 위치로 자동으로 이동
    - 영구 리다이렉션 : 특정 리소스의 URI가 영구적으로 이동
        - 예시) 기존 이벤트에서 새로운 이벤트 페이지로 보내고 싶을 때
        - /event -> /new-event
    - **일시 리다이렉션** : 일시적 이동, **실무에서 가장 많이 사용**
        - 예시) 주문 완료후 주문 내용 화면
        - PRG : Post, Redirection, Get
    - 특수 리다이렉션
        - 결과 대신 캐시를 사용

#### 영구 리다이렉션 : 301, 308

- 기존 URL을 사용하지 않고 검색엔진에서 새로 지정된 URL로 이동
- 301 : Moved Permanently
    - 리다이렉트 시 요청 메서드가 GET으로 변함
    - 본문이 제거될 수도 있음 : **MAY**
    - <그림 추가 해야 함 >
- 308 : Permanent Redirect
    - 301과 기능은 유사하지만 GET으로 변환되지 않고 요청 메서드와 본문을 유지 (POST로 보내면 리다이렉트도 POST 유지)
    - <그림 추가 해야 함 >

#### 일시적 리다이렉션 : 302, 307, 303

- 리소스의 URI가 일시적으로 변경이 됨 
- 위 status 코드를 사용시 검색 엔진등에서 URL을 변경하면 안됨

- 302 : Found
  - 리다이렉트시 GET으로 변경되고 값이 본문이 제거될 수 있음 : **MAY**
  - 현재 실무에서 302로 개발된 것들이 많아 많이 사용됨 
  
- 307 : Temporary Redirect
  - 302와 기능은 동일 
  - 리다이렉트시 요청 메서드와 본문 유지 : **MUST NOT**
  - 302를 보완하기위해 나온 스펙 

- 303 : See Other
  - 302와 기능은 동일 
  - 리다이렉트시 요청메서드가 GET으로 변경 

#### PRG : Post, Redirect, Get

  - 현재 결제 시스템등에서 많이 사용 
  - PRG를 사용하지 않을 경우 서버나 클라이언트에서 미리 처리를 해야 하지만 status code 만 가지고 봤을 때 
  주문 API라고 가정하고 사용자가 재요청시 주문이 중복될 수 있는 문제등이 있음 
  - < 그림 추가 P.214 >
  - PRG 사용시  
    - POST 새로 고침을 통한 중복 방지
    - POST로 주문 후 GET 메서드로 주문 결과 화면으로 Redirect
    - 새로고침을 하더라도 GET 화면을 새로고침을 하게 됨 -> 중복 방지!
    - < 그림 추가>
    
#### 그럼 어떤 것을 사용 하는 것이 좋을까? 302? 303? 307?

  - 302 : Found -> GET으로 변할 수 있음 (거의 그렇게 됨, 브라우저에 따라)
    - 원래 302의 의도는 메서드를 **유지**하는 것이었으나 브라우저들이 개발하면서 원래 의도와는 다르게 개발되어(G**ET으로 변경되도록**)스펙이 변경된 경우
    - 그래서 302를 보완하기 위해 303, 307 스펙을 정의 
  - 307 : 메서드가 변하면 안됨
  - 303 : 메서드가 GET으로 변경 
  - **결론** : 
    - 303, 307을 권장하지만 실제로는 302를 사용하는 애플리케이션 라이브러리들이 많이 있음 
    - 자동 리다이렉션시 GET으로 변경해도 이슈가 없으면 302를 사용해도 무방 


#### 기타 다이렉션 

- 300 : Multiple Choice -> 거의 사용 안함
- 304 : Modified
  - 캐시를 목적으로 사용 
  - 클라이언트에게 서버에서 변경된 것이 없으니 로컬에 있는 캐시를 사용하라고 알려주는 용도 
  - 304 응답은 메시지 **바디 사용 X** (로컬 캐시를 사용해야 함으로)
  - GET, HEAD 요청시 사용 (조건부)