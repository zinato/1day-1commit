# 4xx Client Error

- 오류의 원인이 클라이언트에 있음 
- 똑같은 재시도시 실패가 항상 발생

## 400 Bad Request : 클라이언트가 잘못된 요청을 해서 서버가 요청을 처리할 수 없음

- 요청 파라미터가 잘못되거나, API 스펙이 맞지 않음(String 값이 와야하는데 숫자 값을 넣는 경우)

## 401 Unauthorized : 클라이언트가 해당 리소스에 대한 인증이 필요함 

- Authentication이 되지 않음
- 401 오류시 응답 WWW-Authenticate 헤더와 함께 인증 방법을 설명 

## 403 Forbidden : 서버가 요청을 이해했지만 승인을 거부함 

- Authentication OK, Authorization Fail
- 일반 사용자가 어드민 등급의 리소스에 접근하려고 하는 경우

## 404 Not Found : 요청 리소스를 찾을 수 없음

- 요청한 리소스가 서버에 없음 
- 클라이언트가 권한이 부족한 리소스에 접근할 때 해당 리소스를 숨기고 싶을 때도 사용

# 5xx Server Error

- 서버 오류
- 서버에 문제가 있기 때문에 재시도 하면 성공할 수도 있음 (ex: DB가 복구 되거나 하는 경우)

## 500 Internal Server Error : 서버 문제로 오류 발생

- 애매하면 500 오류 

## 503 Service Unavailable : 서비스 이용 불가 

- 서버 과부하 또는 예정된 작업으로 잠시 요청을 처리할 수 없음
- Retry-After 헤더 필드로 얼마뒤에 복구되는지 보낼 수도 있음 
- 현실적으로 사용하기 힘듦

