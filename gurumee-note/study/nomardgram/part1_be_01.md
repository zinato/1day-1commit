BE #1 프로젝트 세팅
================

> "노마드 코더"에서 제공하는 인스타그램 클론 코딩 강의를 따라하며, 배운 것들을 정리한 문서입니다.

Contents
------------

1. 시작하며...
2. 백엔드란 무엇인가?
3. 장고란 무엇인가?
4. 가상환경이란 무엇인가?
5. 쿠키 커터로 프로젝트 빠르게 설치하기
6. 데이터 베이스 설치하기
7. 유저 모델 작업하기
8. 이미지 앱 및 모델들 작업하기
9. 어드민 패널 작업하기
10. 마치며...

## 시작하며...

자 지금부터 백엔드를 만들어볼 것입니다. 시작하죠!


## 백엔드란 무엇인가?

먼저 백엔드란 무엇인가요? 백엔드는 우리가 만들 애플리케이션의 두뇌입니다. 사용자에게 보여줄 데이터를 읽거나 저장하는 역할을 하죠. 또한 어떻게 보여줄지 어떤 식으로 저장할지를 알고 있습니다. 백엔드는 크게 다음의 3가지로 구분할 수 있습니다.

* 서버
* 애플리케이션
* 데이터 베이스

서버는 엄청나게 좋은 컴퓨터를 뜻합니다. 인터넷과 연결되어 있으며, 애플리케이션 데이터베이스 등이 적재되어 있는 컴퓨터라고 생각하면 편합니다. 주로 리눅스 OS를 사용하지만 게임 업계에서는 윈도우즈 OS도 사용하곤 합니다. 항상 켜져 있어 클라이언트의 요청을 기다리고 요청이 오면 애플리케이션을 통해서 적절한 응답을 보내줍니다.

애플리케이션은 WAS(Web Application Server)라고도 말하며, 주로 파이썬 같은 프로그래밍 언어로 만들어진 프로그램을 뜻합니다. 실제적으로 서버가 요청을 받으면 응답을 처리하기 위한 행동들이 정의됩니다. 많은 애플리케이션들이 사용자 요청에 따라 데이터베이스의 데이터를 가공하여 보여주는 역할을 합니다.

데이터베이스는 애플리케이션에서 사용자의 데이터를 저장하는 곳입니다. 일반적으로 RDBMS 라고 해서 테이블 형식으로 데이터를 저장하는 SQL 데이터베이스가 많았지만 요즘에는 No-SQL 데이터베이스, GraphQL 데이터 베이스 등 더 좋은 데이터베이스들이 나오고 있습니다. 하지만 이 프로젝트에서는 RDBMS 중 하나인 PostgreSQL을 사용할 것입니다.


## 장고란 무엇인가?

우리는 이 프로젝트의 백엔드를 `장고`란 파이썬 웹 프레임워크를 이용하여 만들것입니다. 장고 프레임워크는 처음 "로렌스 저널 월드" 신문사에서 자신들의 웹 서비스를 만들다가 만든 프레임워크입니다. 요즘 범용 프로그래밍 언어로써 급 부상한 파이썬으로 만들어졌으며, 구글이 쓰게 되면서 전 세계적으로 인기를 끌게 되었습니다. 간략히 말하면, 손쉽게 백엔드 서버를 만드는 최고의 췝 프레임워크 중 하나입니다. 장고 프레임워크는 크게 다음의 3가지로 이루어져 있습니다.

* Settings
* Urls
* Apps

먼저 프로젝트의 시간대 설정, 언어 설정, 데이터 베이스 설정 등 모든 설정들은 `settigns`에 의해 정의됩니다. 또한 우리가 만든 앱, 설치한 모듈들을 연결하는 지점이기도 합니다. 애플리케이션 구동시 `settings.py` 혹은 나뉘어진 설정 파일을 읽어서 프로젝트에 필요한 설정들을 설정합니다.

`urls`는 웹 상에서 클라이언트가 요청할 URL과 연결될 `뷰`를 매칭시켜 정의합니다. 간단하게 "/"에 해당하는 URL에 대해서 "IndexView"를 연결할 수 있습니다. 여기서 뷰는 해당 URL에 접속하면 애플리케이션에서 클라이언트에게 보여지는 데이터들을 어떻게 보여줄지(비지니스 로직을) 정의하는 녀석입니다. 일반적으로 MVC 패턴에서 **컨트롤러**라고 생각하시면 됩니다. 

`apps`는 하나의 모듈입니다. 프로젝트의 작은 기능 단위라고 생각하면 됩니다. 우리가 `쇼핑몰`이란 장고 프로젝트를 만든다고 가정해봅시다. 우리는 다음의 기능들이 있을 것입니다.

* 상품 관리
* 유저 관리
* 카트 관리
* 결재 관리

이러한 기능들을 다음의 앱으로 만드는 것입니다. 

* Products
* Users
* Carts
* Payments

여기서 해당하는 기능에 CRUD 그러니까 상품을 예를 들면 상품 정보를 만들고, 읽고 수정하고 삭제하는 기능들을 앱에 정의하는 것이지요. 장고는 프로젝트를 진행하면서 더 알아보도록 하겠습니다.


## 가상환경이란 무엇인가?

파이썬으로 개발할 때, (최근 많은 언어들이 이런 추세긴 합니다만..) 가상 환경이란 것을 이용합니다. 여기서 가상 환경이란 일종의 **비눗 방울**이라고 생각하시면 됩니다. 우리는 하나의 노트북으로 하나의 프로젝트만 만들지 않습니다. 여러 프로젝트를 만들죠. 

근데 꼭 프로젝트마다 같은 파이썬/장고 버전을 이용하는 것은 아닙니다. 오히려, 프로젝트마다 의존성 모듈들의 버전이 다를 때가 많습니다. 이 때, 글로벌하게 의존성을 관리하게 되면 여러 프로젝트를 실행하기가 번거로워집니다. 이 때를 위해서 가상 환경을 만듭니다. 가상 환경을 구성하면 프로젝트마다 파이썬 및 모듈들의 버전을 각각 정의할 수 있습니다. 우리는 **pipenv**라는 것을 통해 가상 환경을 만들 것입니다.


## 쿠키 커터로 프로젝트 빠르게 설치하기

쿠키 커터는 어떤 좋은 개발자가 만든 장고 프로젝트를 빠르게 설정할 수 있는 모듈입니다. 먼저 쿠키커터를 글로벌하게 설치합니다.

```bash
$ pip3 install cookiecutter
```

그 후 다음을 입력합니다.

```bash
$ cookiecutter https://github.com/pydanny/cookiecutter-django
```

터미널에서 어떻게 설정할지 계속 묻는데 잘 읽고 설정해주면 됩니다. 그러면 프로젝트가 설치됩니다. 이제 가상환경을 키고 의존성을 설치하도록 하겠습니다.

```bash
# python3 가상 환경을 설치합니다.
$ pipenv --three

# 가상환경을 활성화합니다.
$ pipenv shell

# requirements/local.txt에 정의된 의존성을 모두 설치합니다. 이 때, base.txt에 있는 모듈들도 설치됩니다.
$ pipenv install -r requirements/local.txt

# 가상 환경 비활성화
$ exit

# 가상환경을 활성화합니다.
$ pipenv shell
```

이제 프로젝트 설치가 끝났습니다. 만약 컴퓨터를 끄거나 가상 환경에서 나갔다면 장고 프로젝트를 구동시키거나 의존성을 설치할 때 반드시 켜야 하는 것을 잊지 마세요. 앞으로 가정할 때 가상 환경이 켜져있다고 가정하겠습니다.


## 데이터 베이스 설치하기

자 이제 데이터 베이스 문제를 고쳐보도록 하겠습니다. 먼저, 우리의 문제는 다음과 같습니다.

1. 연결할 PostgreSQL이 없다.
2. 접속 정보가 없다.

우리는 `PostgreSQL`을 설치해야 하지만 저는 이전에 도커로 한다고 말씀드렸죠? 여기서는 도커, 도커-컴포즈가 설치되어 있다고 가정하겠습니다. 먼저 프로젝트 디렉토리 하단에 `docker-compose.yaml`을 작성하겠습니다.

nomadgram/docker-compose.yaml
```yaml
version: '3.6'

services:
  postgresql:
    hostname: postgresqldb
    image: postgres
    restart: always
    ports: 
      - 5432:5432
    volumes:
      - postgresql-data:/var/lib/postgresql/data
      - ./postgresql:/docker-entrypoint-initdb.d
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: nomadgram
      POSTGRES_INITDB_ARGS: --encoding=UTF-8
    
volumes:
  postgresql-data:
    name: test-postgresql-data
```

이것은 단순히 PostgreSQL만을 키는 도커 컴포즈 파일입니다. 원래는 여러 개를 한꺼번에 띄울 때 자주 쓰이지만 저는 docker run 도 귀찮아서 보통 도커 컴포즈로 작업합니다. 굉장히 쉽거든요. 이제 PostgreSQL을 키겠습니다. 터미널에 다음을 입력하세요.

```bash
$ docker-compose up --build -d
```

이렇게 하면, 이미지 설치 및 컨테이너를 데몬으로 띄웁니다, 여기서 데몬으로 띄운다는 것은 백그라운드에서 PostgreSQL을 실행한다고 생각하시면 됩니다. 이제 PostgreSQL도 켰으니 우리 앱에 접속 정보를 알려주어야 합니다. `config/settings/base.py`에 DATABASE 설정 값을 다음처럼 수정해주세요.

nomadgram/config/settings/base.py
```python

# 이전 코드와 동일

DATABASES = {
    # "default": env.db("DATABASE_URL", default="postgres://localhost:5432/nomadgram")
    "default": {
        "ENGINE": 'django.db.backends.postgresql',
        "NAME": "nomadgram",
        "USER": "user",
        "PASSWORD": "password",
        "HOST": "localhost",
        "PORT": "5432"
    }
}
DATABASES["default"]["ATOMIC_REQUESTS"] = True

# 이전 코드와 동일
```

쿠키 커터로 설정하는 방법이 따로 있는 것 같은데, 저는 잘 모르니 일단 기본적인 장고 설정으로 하겠습니다. 잘 보시면, docker-compose.yaml에 설정한 DB값들을 넣은 것을 확인할 수 있습니다. 이제 디비가 잘 돌아가는지 확인해볼까요? 터미널에 다음을 입력하세요.

```bash
$ python manage.py migrate
```

위 명령어가 잘 수행되면 설정이 잘 된 것입니다. 그리고 이번 절의 내용과는 상관 없지만 관리자 계정을 생성합시다.

```bash
$ python manage.py createsuperuser
```

그 후 서버를 킨 후, /admin 으로 들어가 우리가 만든 계정으로 접속해보세요.


## 유저 모델 작업하기

자 이번에는 유저 모델을 정의합시다. 쿠키 커터로 만들어진 장고 앱은 기본적으로 프로젝트와 동일한 이름의 앱을 하나 가지고 있습니다. 그리고 여기에서 users 앱이 정의되어 있죠. `nomadgram/users/models.py`에 우리가 원하느 User 클래스가 존재합니다.

nomadgram/nomadgram/users/models.py
```python
from django.contrib.auth.models import AbstractUser
from django.db.models import CharField
from django.urls import reverse
from django.utils.translation import ugettext_lazy as _


class User(AbstractUser):

    # First Name and Last Name do not cover name patterns
    # around the globe.
    name = CharField(_("Name of User"), blank=True, max_length=255)

    def get_absolute_url(self):
        return reverse("users:detail", kwargs={"username": self.username})
```

여기 `AbstractUser`는 장고에서 기본적으로 제공하는 계정을 관리할 수 있는 클래스입니다. 우리는 이것을 확장해서, 계정 정보를 추가하면 됩니다. 우리가 추가할 것은 website, bio, phone, gender 입니다. username, email 등 기타 정보는 AbstractUser가 갖고 있기 때문에 따로 적으시지 않아도 됩니다. 다음 처럼 수정해주세요.

nomadgram/nomadgram/users/models.py
```python
from django.contrib.auth.models import AbstractUser
from django.db.models import CharField, URLField, TextField
from django.urls import reverse
from django.utils.translation import ugettext_lazy as _


class User(AbstractUser):

    GENDER_CHOICES = (
        ('male', 'Male'),
        ('female', 'Female'),
        ('not-specified', 'Not specified')
    )

    # First Name and Last Name do not cover name patterns
    # around the globe.
    name = CharField(_("Name of User"), blank=True, max_length=255)
    # 아래가 추가 코드입니다.
    website = URLField(null=True)
    bio = TextField(null=True)
    phone = CharField(max_length=140, null=True)
    gender = CharField(max_length=80, choices=GENDER_CHOICES, null=True)
    followers = ManyToManyField("self")
    following = ManyToManyField("self")

    def get_absolute_url(self):
        return reverse("users:detail", kwargs={"username": self.username})
```

followers, following 부분을 유의해서 보세요. 이것은 다대다 관계를 표현한 것입니다. 한 유저는 자신이 팔로우하는 친구들이 있겠죠? 그리고 자신을 팔로우하는 친구들이 있을 겁니다. 이 때 각 유저는 유저들의 목록들을 가질 수 있습니다. 이것이 다대다 관계 `ManyToMany`입니다.

또한 각 필드마다 null=True 구문을 쓴 이유는 쿠키커터가 이미 사용자 앱에 대한 정보를 데이터베이스에 기록했기 때문입니다. 만약 우리가 계정에 대한 정보를 저장한 후, 모델을 건드렸을 상황에 대해서 예방차원인 것이지요. 이렇게 하면, 이전에 데이터들까지 마이그레이션할 수 있습니다. 이들 필드가 null 이 될 수 있기 때문이죠. 이제 터미널에 다음을 입력하세요.

```bash
# DB에 필드 정보를 알려줍니다.
$ python manage.py makemigrations

# 실제 수정된 정보를 데이터베이스에 적용합니다. 이 때 테이블의 구조가 바뀝니다!
$ python magage.py migrate
```


## 이미지 앱 및 모델들 작업하기

이번에는 이미지 앱을 만들도록 하겠습니다. 터미널에 다음을 입력하세요.

```bash
$ django-admin startapp images
```

이제 images 앱을 만들었습니다. 그 후 앱을 프로젝트와 연결하겠습니다. `config/settings/base.py`에서 LOCAL_APPS 를 다음처럼 수정해주세요.

nomadgram/config/settings/base.py
```python
# 이전 코드와 동일

LOCAL_APPS = [
    "nomadgram.users.apps.UsersConfig",
    # Your stuff: custom apps go here
    "images",
]
# https://docs.djangoproject.com/en/dev/ref/settings/#installed-apps
INSTALLED_APPS = DJANGO_APPS + THIRD_PARTY_APPS + LOCAL_APPS

# 이전 코드와 동일
```

원래 기본적인 장고는 INSTALLED_APPS에 모두 정의하고 우리가 만든 앱도 여기에 둡니다. 그러나 쿠키커터로 만든 장고 앱은 장고 모듈, 서드 파티 모듈, 개발자가 만든 로컬 모듈로 나누어서 모듈들을 연결합니다. 추후에 더해지지만요. 

이제 모델 작업을 해봅시다. 인스타그램에서, 사진을 올리며, 그 사진에 대한 설명을 씁니다. 그리고 우리 친구들은 댓글을 달고 좋아요를 누르지요. 이 때 사진을 업로드하거나 댓글을 달거나 좋아요를 누르거나 모두 시간 정보가 들어가 있습니다. 이들을 차례대로 구현해보겠습니다.
`images/models.py`에 다음 코드를 넣어주세요.

nomadgram/images/models.py
```python
from django.db import models

# Create your models here.

class TimeStampedModel(models.Model):
    """
    시간 정보에 대한 Abstract Class입니다. 처음 만들어진 날짜, 수정된 날짜들에 대한 정보를 가집니다.
    Image, Comment, Like의 기본 클래스가 됩니다.
    """
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    class Meta:
        abstract = True


class Image(TimeStampedModel):
    """
    인스타그램에서 사진 정보를 저장하는 모델입니다. 사진 경로, 위치, 설명에 대한 정보가 들어갑니다.
    """

    file = models.ImageField()
    location = models.CharField(max_length=140)
    caption = models.TextField()
    creator = models.ForeignKey(user_models.User, on_delete=models.CASCADE, null=True)


class Comment(TimeStampedModel):
    """
    댓글 정보를 저장하는 모델입니다.
    """
    message = models.TextField()
    creator = models.ForeignKey(user_models.User, on_delete=models.CASCADE, null=True)
    image = models.ForeignKey(Image, on_delete=models.CASCADE, null=True)


class Like(TimeStampedModel):
    """
    사진에 좋아요 정보를 저장하는 모델입니다.
    """
    creator = models.ForeignKey(user_models.User, on_delete=models.CASCADE, null=True)
    image = models.ForeignKey(Image, on_delete=models.CASCADE, null=True)
```

`TimeStampedModel`은 생성된 날짜, 수정된 날짜를 자동으로 기록하는 모델입니다. 사진, 좋아요, 댓글들은 모두 생성된 날짜, 수정된 날짜 정보가 들어 있지요. 중복을 제거하기 위해 추상 모델을 정의한 것입니다. 그것을 `Meta` 클래스로 정의한 후 `abstract=True` 값을 주어 선언하였습니다.

`Image`는 사진 정보를 의미합니다. 이미지, 장소, 설명 정보가 있으며, 생성한 유저를 담고 있습니다. 이 때, 한 유저는 여러 개의 사진을 가질 수 있지만, 사진의 소유자는 딱 한 유저뿐입니다. 이를 OneToMany 라고 표현하며, 장고에서는 ForeignKey 로 표현할 수 있습니다.

`Comment`, `Like`는 각각 댓글, 좋아요 정보를 표시한 것입니다. 마찬가지로 한 유저는 여러 개의 댓글과 좋아요를 할 수 있습니다. 또한 사진에는 여러 개의 댓글과 좋아요가 남겨질 수 있죠. 하지만 댓글, 좋아요 입장에서 봤을 때, 하나의 유저, 하나의 그림 속에서만 존재합니다. 이들 역시 다대다 관계라는 것이지요. 이들을 표현해 두었습니다.


## 어드민 패널 작업하기

장고의 최대 장점 중 하나는 손쉽게 어드민 패널, 관리자 페이지를 꾸밀 수 있다는 것입니다. 먼저 이해를 위하여 쿠키커터가 만든 유저의 admin은 한 번 살펴볼까요? 

nomadgram/nomdgram/users/admin.py
```python
from django.contrib import admin
from django.contrib.auth import admin as auth_admin
from django.contrib.auth import get_user_model

from nomadgram.users.forms import UserChangeForm, UserCreationForm

User = get_user_model()


@admin.register(User)
class UserAdmin(auth_admin.UserAdmin):

    form = UserChangeForm
    add_form = UserCreationForm
    fieldsets = (("User", {"fields": ("name", "followers", 'followings')}),) + auth_admin.UserAdmin.fieldsets
    list_display = ["username", "name", "is_superuser"]
    search_fields = ["name"]
```

장고 프레임워크를 모르시면, 무슨 소린지 잘 모르실 수 있습니다. 차근 찬근 살펴보죠.

`@admin.register(User)`

이 부분은 `데코레이션`이라고 자바의 `애노테이션`과 유사한 역할을 합니다. 쉽게 생각하면, 관리자 페이지에 우리 유저를 등록하는 일을 합니다.

`form = UserChangeForm`
`add_form = UserCreationForm`

이 부분은 쿠키 커터가 만들어둔 장고 폼 클래스를 가져다 쓰는 것입니다. 각각 유저 정보를 수정, 유저를 추가할 때 쓰이는 폼 클래스들입니다.

`fieldsets = (("User", {"fields": ("name", "followers", 'followings')}),) + auth_admin.UserAdmin.fieldsets`

이 부분은 쉽게 생각해서 수정할 수 있는 필드를 표시해 둔 곳이라고 생각하면 됩니다. 이름, followers, followings를 수정할 수 있는 것이죠.

`list_display = ["username", "name", "is_superuser"]`

관리자 페이지에서 모델을 표시할 때 보이는 필드 목록입니다.

`search_fields = ["name"]`

관리자 페이지에서 모델을 검색할 때 적용할 필드 이름입니다. 이번에는 우리 모델들의 `ModelAdmin`을 적용해 보겠습니다. images의 admin.py에 다음 코드를 입력하세요

nomadgram/images/admin.py
```python
from django.contrib import admin
from .models import Image, Comment, Like
# Register your models here.

@admin.register(Image)
class ImageAdmin(admin.ModelAdmin):

    list_display_links = (
        'location',
    )

    search_fields = (
        'location',
        'caption',
    )

    list_filter = (
        'location',
        'creator',
    )

    list_display = (
        'file',
        'location',
        'caption',
        'creator',
        'created_at',
        'updated_at'
    )


@admin.register(Comment)
class CommentAdmin(admin.ModelAdmin):
    list_display = (
        'message',
        'creator',
        'image',
        'created_at',
        'updated_at'
    )



@admin.register(Like)
class LikeAdmin(admin.ModelAdmin):
    list_display = (
        'creator',
        'image',
        'created_at',
        'updated_at'
    )
```            

우리의 모델들도 이렇게 관리자 페이지에 등록을 해 두었습니다. 각 필드에 대해서는 장고 문서를 보고 필드를 적용하고 관리자 페이지에 적용되는 것을 보면 금방 익힐 수 있을 것입니다.


## 마치며..

이렇게 해서 백엔드 파트1이 끝났습니다. 파트2에서는 API 작업을 도와주는 DRF 적용이 주를 이룰 것으로 보입니다. 이상 감사합니다.