# 7월 29일 작업 : 소셜 로그인 기능

## 개요

현재, 블로그 애플리케이션 개발 중이다. 구현 스펙은 다음과 같다.

* 언어 : Golang
* 라이브러리, 프레임워크 : Gorm(Persistance), Echo(http Framework)
* 데이터베이스 : PostgreSQL
* 컨테이너 : 도커
* CI/CD : Github Action
* 클라우드 프로바이더 : Heroku

이전까지는 Post 등록/수정/삭제/조회 기능을 완료하였다. 현재는 소셜 로그인 기능을 구현하고 있다. 그런데 막혀 있다..


## 어제 작업 요약

어제 소셜 로그인 기능을 위해서 Configuration 설정 + 구글/네이버 로그인 버튼 구현 및 Callback 엔드 포인트를 작성했다. 그런데 왠걸... 로그인이 정상적으로 작동하지 않았다.


## 오늘 작업

어제 로그인이 왜 안됐을까? 두 가지 이유가 있다.

**첫 번째, 구글/네이버 애플리케이션 등록 시, 한 앱만 접근이 가능하다?**

이것은 더 확인이 필요하다. 이 프로젝트와 쌍둥이 프로젝트로 스프링 애플리케이션을 같이 만들고 있다. 그래서 같은 구글/네이버 애플리케이션을 이용한다. 아무래도 같은 ClientID와 ClientSecret을 공유해서 생기는 문제인 듯 싶다. 

왜 이렇게 결론을 냈냐면, Go 애플리케이션 소셜 로그인 기능을 만들기 이전에 잘 동작했던 스프링 애플리케이션의 소셜 로그인 기능이 망가졌기 때문이다. Go 애플리케이션에서 한 번 구글/네이버 로그인 기능을 이용하면, 스프링 애플리케이션에서 더 이상 사용할 수 없게 되었다. 그래서 각 애플리케이션 별 구글/네이버 애플리케이션을 등록하였다.

**두 번째, 소셜 로그인 프로바이더에 의해 리다이렉트되면, 쿠키를 잃어버린다?**

현재는 쿠키를 이용하여, state를 저장하고 이를 비교한다. 다음은 구글 로그인 버튼을 눌렀을 때, "/oauth2/authorization/google" 이동하게 되는데, 이 엔드포인트에서 소셜 로그인 전처리를 담당하는 핸들러의 코드이다.

```go
// GoogleLogin is
func (h *Handler) GoogleLogin(c echo.Context) error {
    // state 생성
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
    url := h.config.NaverOAuth.AuthCodeURL(state)
    // 쿠키 저장.
    cookie := new(http.Cookie)
	cookie.Name = "oauthstate"
	cookie.Value = state
	cookie.Expires = time.Now().Add(1 * 24 * time.Hour)
    c.SetCookie(cookie)
    // 구글 로그인 엔드 포인트로 리다이렉트
	return c.Redirect(http.StatusPermanentRedirect, url)
}
```

현재는 공식 문서를 토대로, http.Cookie를 생성하여, echo 컨텍스트에서 쿠키의 값을 저장하게 하고 있다. 그 후, 구글 애플리케이션에서 등록한 리다이렉트 URL "/login/oauth2/code/google" 을 담당하는 핸들러, 즉 소셜 로그인 후 처리를 담당하는 핸들러 코드를 다음과 같이 작성하였다.

```go
// GoogleCallback is
func (h *Handler) GoogleCallback(c echo.Context) error {
    // 저장된 쿠키 불러옴
	cookie, err := c.Cookie("oauthstate")

    // 이 단계에서 에러가 남.
	if err != nil {
		log.Println("error: ", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	if c.FormValue("state") != cookie.Value {
		log.Printf("invalid google oauth state cookie:%s state:%s\n", cookie.Value, c.FormValue("state"))
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
```

이제 리다이렉트된 URL에서 쿠키에서 저장한 "oauthstate"를 불러오는데, 이렇게 하면, 해당 쿠키가 없다면서 에러가 난다. 따라서 추론할 수 있는 것은, 다음과 같다.

1) 소셜 로그인 버튼을 누른다.
2) "/oauth2/authorization/:social_provider"로 이동한다.
3) SocialLogin 핸들러에서 "state"를 만들고 그 것을 쿠키에 저장한다. 그 후 state와 함께 소셜 로그인 엔드포인트로 리다이렉트한다.
4) 유저 정보 입력 후, 소셜 로그인 서비스에 등록한 리다이렉트 URL "/login/oauth2/code/:social_provider"로 이동한다.
5) 해당 URL를 담당하는 핸들러가 호출된다. 이 때, 이미 쿠키가 소실된다.

실제 GoogleLogin 핸들러 가장 끝에서 쿠키를 가져와봤는데, 잘 가져온 것을 확인하였다. 이를 어떻게 해결할 수 있을까.s


## 내일 작업

내일은 이어서 "쿠키"가 아닌 "세션"을 이용해서, state를 해볼까 한다. 이것도 안되면 어떡하지...