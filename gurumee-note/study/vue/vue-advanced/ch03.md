Project Setup
================

Contents
-------------

1. Create Project By Vue Cli3 
2. ES Lint turn off

Create Project By Vue Cli3
-------------

먼저 nodejs, npm, vue-cli 가 설치되어 있어야 한다. 다음은 리눅스(ubuntu-19.04) 기준 nodejs, npm vue-cli 설치 방법이다.

```bash
# nodejs, npm install
$ curl -sL https://deb.nodesource.com/setup_10.x | sudo bash -
$ sudo apt install nodejs

# vue-cli install 
$ sudo npm install -g @vue/cli
```

다음은 위 명령어로 설치된 nodejs, npm, vue-cli 버전들이다.

```bash
# nodejs version
$ node -v
v16.10.0

# npm version
$ npm -v
6.9.0

# vue-cli version
$ vue --version
3.9.2
``` 

이제 원하는 디렉토리에다 다음 명령어를 치면 된다.

```bash
# create project
$ vue create vue-news
```

나머지는 그냥 디폴트 설정을 따른다. 그럼 프로젝트 설정은 다음과 같이 된다.

- node_modules(D)
- public(D)
- src(D)
    - assets(D)
        - logo.png
    - components(D)
        - HelloWorld.vue
    - App.vue
    - main.js
- .gitignore
- babel.config.js
- package-lock.json
- package.json
- README.md

방금 만든 프로젝트를 실행하려면 다음 명령어를 입력한다. 다른 명령어들은 `pacakge.json` 에서 "scripts" 부분을 확인한다.

```bash
# run
$ npm run serve
```

ES Lint turn off
-------------

default 설정을 해두면 `ES Lint`가 기본적으로 설정되어 있다. `ES Lint`는 개발자가 작성한 자바스크립트 코드 문법을 체크하는 용도라고 보면 된다. 이것들을 지키면서 개발하는 것이 베스트겠지만, 안될 수도 있다. 하지만 `Vue`는 기본적으로 `ES Lint`를 지키지 않으면 Error 를 띄운다. 
이것을 프로젝트마다 지키기 어려운 사람들을 위해서, `ES Lint` 끄는 법을 작성한다.

먼저 `Vue`에서는 3가지 방식이 존재한다.

1. Vue 파일에서 script 태그 위에 /* eslint-disable */ 을 명시한다.
   ```vue
   <template>
    ...
   </template>
   
   /* eslint-disable */
   <script>
   ...
   </script>
   ```
2. Vue 파일에서 ES Lint 에 걸린 명령어 바로 위에 // eslint-disable-next-line 을 명시한다.
    ```vue
   <template>
    ...
   </template>
   
   /* eslint-disable */
   <script>
   // eslint-disable-next-line
   // 파괴 js 문법
   </script>
   ```  
   
3. 프로젝트 루트 디렉토리에 `vue.config.js` 를 다음과 같이 작성한다.
    ```js
    module.exports = {
       lintOnSave: false
    }
    ```
   
이렇게 하고 프로그램을 실행하면 `ES Lint`를 무시하는 문법이 나와도 정상적으로 컴파일되는 것을 볼 수 있다. 