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



