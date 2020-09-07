# Hugo + Github으로 개인 블로그 만들기

- [Hugo + Github으로 개인 블로그 만들기](#hugo--github으로-개인-블로그-만들기)
  - [개요](#개요)
  - [Hugo 설치](#hugo-설치)
  - [Github 레포지토리 2개 파기](#github-레포지토리-2개-파기)
  - [Hugo로 프로젝트 만들기](#hugo로-프로젝트-만들기)
  - [Git 레포지토리 연결 및 쉘 스크립트 작성](#git-레포지토리-연결-및-쉘-스크립트-작성)
  - [포스트 작성하기](#포스트-작성하기)
  - [Utterences (Github 댓글 위젯) 추가하기](#utterences-github-댓글-위젯-추가하기)
  - [참고](#참고)

## 개요 

개인적으로, 원래 블로그 플랫폼으로 `Tistory`를 이용했었다. 언젠가 사람들이 `Github` 페이지로 개발 블로그를 만드는 것을 보고 멋있다고 생각한 것 같다. 올해 어떤 계기로 티스토리 블로그를 닫았는데, 새로운 마음으로 `Github` 페이지룰 이용하여, 개발 블로그를 만들 생각이다.

보통 `Github` 페이지로 개발 블로그를 만들 때 사용하는 것은 `Jekyll`, `Hexo`, `Hugo` 이 3가지 중 하나를 사용한다. 이들은 `Static Site Generator`로써 직역하면 정적 페이지 생성기이다. 웹 페이지는 크게 `정적 페이지`와 `동적 페이지`로 나눌 수 있다. `동적 페이지`란, 클라이언트 액션에 반응해서 웹 서버, DB 등 여러 컴포넌트들이 상호 작용하여 HTML 페이지를 동적으로 만들어주는 페이지를 말한다. 뭐 개발자라면, `Java` 진영에서는 `Spring MVC`로 만든 웹 서비스라고 생각하면 편하다. 반면, `정적 페이지`는 HTML/CSS/JS를 미리 올려서 서버가 바뀌지 않는 HTML 페이지를 보여주는 것을 말한다. 

블로그를 방문하는 유저 입장에서는 모든 컨텐츠들이 언제나 `Read Only`이다. 즉 상호 작용하는 건이 별로 없다는 뜻이다. 그래서 정적 페이지 생성기를 이용해, 개발 블로그를 만들 것이다. 그 중 `Hugo`를 사용한다. `Jekyll`이 참고할 수 있는 문서량이 제일 많지만, 페이지를 많이 올릴수록 빌드 속도가 매우 저하된다는 소문이 있다. 반면 `Hugo`는 굉장히 빠르다고 한다. 그리고 `Golang`을 이용해서 만들었다는게 참 친근하게 다가왔다. 이제 잡소리는 그만하고 바로 시작하자.


## Hugo 설치

준비물은 다음과 같다.

* Mac - Windows 사용자들은 Hugo 설치 방법을 따로 찾길 바란다. 그리고 `HomeBrew`가 설치 되어 있어야 한다. 설치가 되지 않았다면, [이 곳](https://brew.sh/index_ko)을 참조하자.
* Github ID - 회원가입이 안되어 있다면, [이 곳](https://github.com/)으로 가서 회원가입하자.
* Golang - Hugo는 Golang으로 만들어져 있다. Golang을 설치해야 한다. [이 곳](https://golang.org/)을 참고하자.

터미널에 다음을 입력한다.

```bash
# hugo 설치
$ brew install hugo
```

설치가 완료되면, 다음을 터미널에 입력해보자. 

```bash
# hugo 버전 확인
$ hugo version
# 다음이 출력되면 정상적으로 설치된 것이다.
Hugo Static Site Generator v0.74.3/extended darwin/amd64 BuildDate: unknown
```


## Github 레포지토리 2개 파기

깃헙 레포지토리를 2개 파야 한다. 둘 다 아무렇게 지어도 상관은 없지만, 잘 모르겠다면, 다음처럼 만들어두자.

* `blog`
* `<USERNAME>.github.io`

다음과 같이 말이다.

![깃 헙 레포지토리](./01.png)


## Hugo로 프로젝트 만들기

먼저, 프로젝트를 만들어야 한다. 적당한 위치에서 다음을 입력한다.

```bash
# 현재 위치 확인
$ pwd
# 현재 경로
/Users/gurumee/Workspaces

# hugo 프로젝트 생성
$ hugo new site blog
# 생성 완료되면 출력 문구
Congratulations! Your new Hugo site is created in /Users/gurumee/Workspaces/blog.

Just a few more steps and you're ready to go:

1. Download a theme into the same-named folder.
   Choose a theme from https://themes.gohugo.io/ or
   create your own with the "hugo new theme <THEMENAME>" command.
2. Perhaps you want to add some content. You can add single files
   with "hugo new <SECTIONNAME>/<FILENAME>.<FORMAT>".
3. Start the built-in live server via "hugo server".

Visit https://gohugo.io/ for quickstart guide and full documentation.
```

만들어진 프로젝트 구조는 다음과 같다.

```
├── archetypes
│   └── default.md
├── config.toml
├── content
├── data
├── layouts
├── static
└── themes
```

이제, 테마를 설정해야 한다. 테마는 자신이 만들 수도 있고, 남이 만든 것을 이용할 수 있다. 나는 남이 만든 것을 이용한다. [이 곳](https://themes.gohugo.io/)에서 공개된 테마를 설치할 수 있다. 내가 설치한 테마는 [이 곳](https://themes.gohugo.io/hugo-tranquilpeak-theme/)에서 확인할 수 있다. 보통 `Downloads`를 눌러보면 해당 `Github` 페이지로 이동되는데, 설치할 수 있는 방법이 적혀져 있다.

나의 경우는 `Hugo` 공식 문서를 참고해서, 설치하였다. 자 따라해보자. 프로젝트 루트에서 터미널을 열고 다음을 입력한다. 

```bash
$ pwd
# 프로젝트 루트 경로
/Users/gurumee/Workspaces/blog

# git submodule add https://github.com/<theme 경로>.git themes/<theme 이름>
$ git submodule add https://github.com/kakawait/hugo-tranquilpeak-theme.git themes/hugo-tranquilpeak-theme
```

그럼 `blog/themes` 디렉토리의 구조는 다음과 같다.

```
└── hugo-tranquilpeak-theme
    ├── archetypes
    ├── docs
    │   └── img
    ├── exampleSite
    │   ├── content
    │   │   └── post
    │   └── static
    │       └── img
    ├── i18n
    ├── images
    ├── layouts
    │   ├── _default
    │   ├── partials
    │   │   ├── internal
    │   │   └── post
    │   ├── shortcodes
    │   └── taxonomy
    ├── src
    │   ├── images
    │   ├── js
    │   └── scss
    │       ├── base
    │       ├── components
    │       ├── layouts
    │       ├── pages
    │       ├── themes
    │       └── utils
    │           └── mixins
    ├── static
    │   ├── css
    │   ├── images
    │   └── js
    ├── tasks
    │   ├── config
    │   └── register
    ├─ theme.toml
    ├─ ...
```

여기서 `blog/themes/hugo-tranquilpeak-theme/exampleSite/config.toml`을 복사해서 `blog/config.toml`에 옮겨 놓자. 그 후 자신에게 맞게 커스텀하면 된다. 제일 중요한 것은 다음을 바꿔줘야 한다는 것이다. 

blog/config.toml
```toml
# ...
# 여기가 바뀌어야 합니다.
# https://<username>.github.io/
baseURL = "https://gurumee92.github.io/"

# ...
languageCode = "en-us"
defaultContentLanguage = "en-us"
title = "Gurumee92's Devlog"
theme = "hugo-tranquilpeak-theme"
disqusShortname = "hugo-tranquilpeak-theme"
# googleAnalytics = "UA-123-45"
paginate = 7
canonifyurls = true

[permalinks]
  post = "/:year/:month/:slug/"
```

이제 터미널에 다음을 입력해보자.

```bash
# 로컬 서버 실행
$ hugo server
                   | EN-US  
-------------------+--------
  Pages            |     9  
  Paginator pages  |     0  
  Non-page files   |     0  
  Static files     |     4  
  Processed images |     0  
  Aliases          |     1  
  Sitemaps         |     1  
  Cleaned          |     0  

Built in 7 ms
Watching for changes in /Users/gurumee/Workspaces/blog/{archetypes,content,data,layouts,static,themes}
Watching for config changes in /Users/gurumee/Workspaces/blog/config.toml
Environment: "development"
Serving pages from memory
Running in Fast Render Mode. For full rebuilds on change: hugo server --disableFastRender
Web Server is available at http://localhost:1313/ (bind address 127.0.0.1)
Press Ctrl+C to stop
```

그 후, "localhost:1313"에 들어가면 다음과 같은 화면이 뜰 것이다.

![로컬 띄어놓은 화면](./02.png)


## Git 레포지토리 연결 및 쉘 스크립트 작성

터미널에 다음을 입력한다. `blog`에 `Github Repository blog`, `blog/public`에 `Github Repositroy <username>.github.io`를 연결한다.

```bash
# 현재 위치 확인
$ pwd
/Users/gurumee/Workspaces/blog

# blog -> blog 레포지토리 연결
# git remote add origin http://github.com/<username>/blog.git
$ git remote add origin http://github.com/gurumee92/blog.git

# blog/public -> <username>.github.io 연결
# git submodule add -b master http://github.com/<username>/<username>.github.io.git public
$ git submodule add -b master http://github.com/gurumee92/gurumee92.github.io.git public
```

이제 배포를 보다 쉽게 하기 위해서 쉘 스크립트를 작성한다. `blog` 프로젝트 루트 디렉토리에, `deploy.sh`를 만들고 다음을 입력한다.

blog/deploy.sh
```bash
#!/bin/bash

echo -e "\033[0;32mDeploying updates to GitHub...\033[0m"

# Build the project.
# hugo -t <여러분의 테마>
hugo -t hugo-tranquilpeak-theme

# Go To Public folder, sub module commit
cd public
# Add changes to git.
git add .

# Commit changes.
msg="rebuilding site `date`"
if [ $# -eq 1 ]
  then msg="$1"
fi
git commit -m "$msg"

# Push source and build repos.
git push origin master

# Come Back up to the Project Root
cd ..


# blog 저장소 Commit & Push
git add .

msg="rebuilding site `date`"
if [ $# -eq 1 ]
  then msg="$1"
fi
git commit -m "$msg"

git push origin master
```

그 후, 터미널에 다음을 입력한다.

```bash
# deploy.sh 실행 파일 권한 부여
$ chmod 777 deploy.sh

# 배포 실행
$ ./deploy.sh
```

조금 이따가 `https://<username>.github.io`를 접속하면, 블로그가 뜬 것을 확인할 수 있다.


## 포스트 작성하기

포스트는 다음과 같이 만들면 된다. 터미널에 다음을 입력한다.

```bash
# hugo new post/<원하는 패쓰>/파일 이름.md
$ hugo new post/test.md                 
/Users/gurumee/Workspaces/blog/content/post/test.md created
```

그러면, `blog/content/post/` 경로에 `test.md`가 생성된다. 그리고 간단하게 다음을 적어본다.

blog/content/post/test.md
```md
---
title: "Test"
date: 2020-08-28T21:48:13+09:00
categories:
- category
- subcategory
tags:
- tag1
- tag2
keywords:
- tech
#thumbnailImage: //example.com/image.jpg
---

<!--more-->
테스트입니다.
```

이제 터미널에 다음을 입력하여 로컬 서버를 킨다.

```bash
# 로컬 서버 실행
$ hugo server -D
```

그럼 `http://localhost:1313` 에 접속하면 다음 화면이 출력된다.

![포스트 화면](./03.png)

컨텐츠 작성이 완료되면, 다시 터미널에 다음을 입력해서 배포하면 된다.

```bash
$ ./deploy.sh
```


## Utterences (Github 댓글 위젯) 추가하기

이제 블로그에 댓글을 달아보자. 이 때, 레포지토리가 하나 더 필요하다. 나는 `blog-comments`로 지었다. 그 후 [이 곳](https://utteranc.es/)으로 이동하자. 그리고 `configuration` 절로 이동해서 "repo" 입력 창에, 자신에 레포지토리를 입력하자. 나의 경우는 `gurumee92/blog-comments`이다.

![Utterences 1](./04.png)

그리고 `Blog Post ↔️ Issue Mapping` 절로 이동해서 다음 6개 중 하나를 고른다. 나는 3번째 것을 선택했다.

![Utterences 2](./05.png)

그러면 밑에 다음 스크립트가 생성된다.

```html
<script src="https://utteranc.es/client.js"
        repo="gurumee92/blog-comments"
        issue-term="title"
        theme="github-light"
        crossorigin="anonymous"
        async>
</script>
```

이 스크립트를 복사하여, `blog/themes/<자신의 테마>/layouts/partials/post/disqus.html` 에다 다음 코드를 수정하면 된다.

> 참고!
> 이건 내 테마에만 그럴 수 있다. 하지만 보통 테마 밑에 "layouts" 디렉토리는 있으니 차분히, 디스커스 혹은 footer 쪽을 찾다 보면, 적당한 위치를 찾을 수 있을 것이다.

blog/themes/hugo-tranquilpeak-theme/layouts/partials/post/disqus.html
```html
<div id="disqus_thread">
  <noscript>Please enable JavaScript to view the <a href="//disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>
</div>
<!-- 추가 코드 -->
<script src="https://utteranc.es/client.js"
        repo="gurumee92/blog-comments"
        issue-term="title"
        theme="github-light"
        crossorigin="anonymous"
        async>
</script>
```

실제 내 블로그의 포스트는 다음과 같이 댓글이 연결되어 있다.

![Utterences 3](./06.png)

> 참고!
> 댓글 달 때, `Utterences`앱을 설치하지 않으면, 설치하라는 문구가 뜬다 `blog-comments`에만 설치하도록 하자. 자신이 블로그를 개설할 때, 댓글을 먼저 달아보도록 하자.


## 참고

* [Hugo 공식 홈페이지 Quick Start](https://gohugo.io/getting-started/quick-start/)
* [Hugo로 Github.io 블로그 만들기](https://github.com/Integerous/Integerous.github.io)



