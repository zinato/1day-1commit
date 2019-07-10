Vue Router 기초
=================

Contents
---------
1. Install vue-router 
2. How to use vue-router
3. Make Header Navigation
4. Practice #1

Install vue-router
---------

`vue-router`는 해당 URL 뿌려질 컴포넌트들 그러니까 그 페이지를 설정할 수 있는 라이브러리이다. 설치 방법은 프로젝트 루트 디렉토리에서 다음 명령어를 터미널에 입력한다.

```bash
# install vue-router
$ npm i vue-router --save
```

그 후, `package.json` 에서 "dependencies" 프로퍼티 부분을 보면 의존성이 추가 된 것을 확인할 수 있다.

```json5
{
  // ...
  "dependencies": {
    "core-js": "^2.6.5",
    "vue": "^2.6.10",
    "vue-router": "^3.0.7"
  },
  //...
}
```

How to use vue-router
---------

먼저 `src` 디렉토리에서 `routes` 라는 디렉토리를 만들고 `index.js`를 만든 후 다음을 입력한다.

vue-news/src/routes/index.js
```js
import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

export const router = new VueRouter({
  routes: [
    {
        path: '',
        component: ''
    },   
  ]
});
```

여기서 `router` 변수는 "export" 구문에 의해서, 다른 외부 파일에 쓸 수 있도록 만들어준다. json 형식으로 만든 `routes` 필드는 프로젝트에 사용될 라우트의 리스트를 넣어주면 된다. 형식은 위와 같 path, component 프로퍼티를 넣어준다.

* path - url
* component - url 에 뿌려질 컴포넌트

우리가 사용할 URL 은 다음과 같다.

* /news ( "/" 도 이 URL 과 연결된다.)
* /jobs
* /ask

`path`와 `component`를 연결하기 전에 해당 컴포넌트를 만들어보자. `src`밑에 `views` 디렉토리를 만들고, 다음 파일들을 만들어보자.

* NewsView.vue
* JobsView.vue
* AskView.vue

그리고 `vscode` 설치 시에 "Vetur" 플러그인을 설치한다면, `sfc`라는 것을 치기만 해도 자동으로 템플릿이 만들어진다. 이제 각 파일 마다 다음을 입력하자.

vue-news/src/views/NewsView.vue
```vue
<template>
    <div>
        news
    </div>
</template>

<script>
export default {

}
</script>

<style>

</style>
```

vue-news/src/views/JobsView.vue
```vue
<template>
    <div>
        jobs
    </div>
</template>

<script>
export default {

}
</script>

<style>

</style>
```

vue-news/src/views/AskView.vue
```vue
<template>
    <div>
        ask
    </div>
</template>

<script>
export default {

}
</script>

<style>

</style>
```

이제 만든 view 들을 router 에 연결하자. `routes/index.js` 파일에서 다음을 입력한다.

vue-news/src/routes/index.js
```js
import Vue from 'vue';
import VueRouter from 'vue-router';
//추가된 부분 해당 뷰들을 불러온다.
import NewsView from '../views/NewsView.vue';
import JobsView from '../views/JobsView.vue';
import AskView from '../views/AskView.vue';

Vue.use(VueRouter);

export const router = new VueRouter({
  // 원래 뷰 라우터에 연결되면 localhost:8080/#/news 이런 식으로 연결되는데, localhost:8080/news 으로 바꿔준다. 즉 # 을 없애준다.
  mode: 'history',   
  routes: [
    // "/" 에 연결되면, "/news"로 리다이렉트
    {
      path: '/',
      redirect: '/news',
    },
    // "/news" 로 접속하면 NewsView 를 보여준다.
    {
      // url
      path: '/news',
      // rendering component
      component: NewsView,
    },
    // "/ask" 로 접속하면 AskView 를 보여준다.
    {
      path: '/ask',
      component: AskView,
    },
    // "/jobs" 로 접속하면 JobsView 를 보여준다.
    {
      path: '/jobs',
      component: JobsView,
    },
  ]
});
```
  
이제 `main.js` 에서 우리가 만든 `router`를 연결해준다.

vue-news/src/main.js
```js
import Vue from 'vue'
import App from './App.vue'
// routes/index.js 에 만든 router 를 불러옴
import { router } from './routes/index.js'

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  //router: router 축약
  router,
}).$mount('#app')
```

그리고 이제 `App.vue` 에서 등록한 `router` 를 쓰면 된다.

vue-news/src/App.vue
```vue
<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
}
</script>

<style>
body {
  padding: 0;
  margin: 0;
}
</style>
```

`router-view` 태그를 이용하면 우리가 등록한 router 에 따라서 해당 url 에 대해서 component 를 뿌려준다. 뷰를 실행시킨 후 다음 URL 에 접속해보자.

```bash
# run vue app
$ npm run serve
```

* "/", "/news" - NewsView 가 뿌려짐
* "/ask" - AskView 가 뿌려짐
* "/jobs" - JobsView 가 뿌려짐

Make Header Navigation
---------

자, 이제 `vue-router`를 이용해서 "Header Navigation"을 만들어보자. 보통의 사이트에서 헤더 영역에 보통 네비게이션을 뿌려주지 않던가? 그 작업을 하는 것이다. 먼저 `components` 디렉토리에 `ToolBar.vue`를 만들고 다음을 입력하자.

vue-news/src/components/ToolBar.vue
```vue
<template>
    <div class="header">
        <router-link to="/news">News</router-link> |
        <router-link to="/ask">Ask</router-link> |
        <router-link to="/jobs">Jobs</router-link>
    </div>
</template>

<script>
export default {

}
</script>

<style scoped>
.header{
    color: white;
    background: #42b883;
    display: flex;
    padding: 8px 8px 8px 8px;
}

.header a:active {
    color: #35495e;
}

.header a {
    color: white
}
</style>
```

`router-link` 태그를 이영하면 해당하는 URL 로 이동하는 하이퍼링크 태그를 손쉽게 만들 수 있다. 이제 이것을 `App.vue`에 연결하기만 하면 된다.

vue-news/src/App.vue
```vue
<template>
  <div id="app">
    <tool-bar></tool-bar>
    <router-view></router-view>
  </div>
</template>

<script>
//ToolBar 불러옴
import ToolBar from './components/ToolBar.vue';

export default {
  // 현재 앱에 ToolBar 컴포넌트드를 등록
  components: {
    ToolBar,
  }
}
</script>

<style>
body {
  padding: 0;
  margin: 0;
}
</style>
```

이제 앱을 켜보면 우리가 등록한 네비게이션이 맨 위에 헤더에 뜨는 것을 확인할 수 있다. 

> CSS 작업은 웹 콘솔에서 띄운 후, 해당 컴포넌트에 스타일을 이것 저것 넣어보고 뜨는 것을 본 후, 코드로 옮겨오는 것이 훨~~씬 편하다! 개꿀팁

Practice #1
---------

> 지금 배운 내용들을 이용하여 해당 URL에 뿌려질 컴포넌트들을 만들고 등록시켜보자. 
> * "/user" -> UserView 
> * "/item" -> ItemView 

해당 실습에서는 `views` 디렉토리에 `UserView.vue, ItemView.vue`를 만들고 `routes/index.js`에 만든 뷰들을 등록시키기만 하면 된다. 해당 파일들의 수정된 코드는 다음과 같다.

vue-news/src/views/UserView.vue
```vue
<template>
    <div>
        User View
    </div>
</template>

<script>
export default {

}
</script>

<style>

</style>
```

vue-news/src/views/ItemView.vue
```vue
<template>
    <div>
        Item View
    </div>
</template>

<script>
export default {

}
</script>

<style>

</style>
```

vue-news/src/routes/index.js
```js
import Vue from 'vue';
import VueRouter from 'vue-router';
import NewsView from '../views/NewsView.vue';
import JobsView from '../views/JobsView.vue';
import AskView from '../views/AskView.vue';
// UserView, ItemView 불러옴
import UserView from '../views/UserView.vue';
import ItemView from '../views/ItemView.vue';

Vue.use(VueRouter);

export const router = new VueRouter({
  mode: 'history', //delete #
  routes: [
    {
      path: '/',
      redirect: '/news',
    },
    {
      // url
      path: '/news',
      // rendering component
      component: NewsView,
    },
    {
      path: '/ask',
      component: AskView,
    },
    {
      path: '/jobs',
      component: JobsView,
    },
    // /user -> UserView 연결
    {
      path: '/user',
      component: UserView,
    },
    // /item -> ItemView 연결
    {
      path: '/item',
      component: ItemView,
    },
  ]
});
```
