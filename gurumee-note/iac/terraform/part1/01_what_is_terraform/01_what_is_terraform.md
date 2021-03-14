# 1장 테라폼이란 무엇인가

## 테라폼이란 무엇인가

테라폼은 하시코프에서 만든 IaaC 도구이다. 테라폼은 안전하고 반복적으로 작업하더라도 인프라스트럭처를 구축, 변경할 수 있게 도와준다. 프로비저닝이라는 큰 틀에서는 Ansible과 같이 비교가 되곤 하는데, Ansible은 설정 관리 도구로써 테라폼과는 약간 성격이 다르다는 것을 알아두는 것이 좋다.  

### IaaC란 무엇인가

IaaC란, "Infrastructure As A Code"의 약자로써, 코드로 인프라를 관리하는 것을 말한다. 인프라를 이루는 서버, 미들웨어, 서비스 등 인프라를 구성하는 모든 요소들이 이 대상에 속한다. 테라폼은 0.x 버전임에도 불구하고 이 분야에서 표준으로 자리 잡았다. HCL이라는 도메인 특화 언어를 통해서 AWS, GCP 등의 퍼블릭 클라우드는 물론 여러 프로바이더의 원하는 리소스 설정을 자동화할 수 있다.

### 테라폼의 장점

1. 코드로써 인프라를 관리하기 때문에, 생산성, 재사용성이 높아지며, 유지보수도 쉬워진다.
2. 업계 표준이기 때문에 예제가 풍부하여 배우기 쉽다.

### 테라폼의 구성 요소

* provider : 테라폼으로 관리할 인프라의 종류를 의미 (ex) AWS, GCP, Azure ...)
* resource : 테라폼으로 관리할 인프라 자원을 의미 (ex) EC2, IAM, S3, RDS ...)
* state : 테라폼을 통해 생성된 리소스들의 상태를 의미. (= 테라폼 apply 명령어를 실행한 결과물)
* output : 테라폼으로 만든 리소스를 변수 형태로 state에 저장하는 것을 의미
* module : 공통적으로 활용할 수 있는 모듈을 정의하는 것을 의미
* remote : 다른 경로의 state를 참조하는 것을 의미하며, output 변수를 불러올 때 주로 사용

## 테라폼 설치

### 맥에서 테라폼 설치

맥에서는 "HomeBrew" 패키지 매니저가 설치되었다고 가정하고 진행한다.

> HomeBrew는 다음 [링크](https://brew.sh/index_ko)에서 다운로드 할 수 있습니다. 

테라폼을 바로 설치해도 좋지만, 여러 테라폼 버전을 사용할 수 있으니 버전 관리할 수 있는 도구인 "tfenv"를 먼저 설치하겠다.

```bash
# tfenv 설치
$ brew install tfenv
```

설치가 완료되었으면, 다음을 입력하여 tfenv가 잘 설치되었는지 확인한다.

```bash
# tfenv 설치 확인
$ tfenv --version
tfenv 2.2.0
```

이제 이를 이용해서 테라폼을 설치한다.

```bash
# terraform 최신 버전 설치
$ tfenv install

# terraform 특정 버전 설치를 원한다면 다음과 같은 형식으로 할 수 있다.
$ tfenv install 0.13.2
```

21년 2월 17일 기준, 0.14.6 버전이 설치된다. 이제 "tfenv"가 위의 명령어로 설치한 테라폼을 사용하게끔 설정하자.

```bash
# terraform 버전 사용
$ tfenv use 0.14.6
```

그 후 테라폼 버전을 확인해보자.

```bash 
# terraform 버전 확인
$ terraform version
Terraform v0.14.6
```

### 리눅스에서 테라폼 설치

AWS Linux는 다음과 같이 설치하면 된다.

```bash
# 필요 모듈 설치
$ sudo yum install -y yum-utils

# yum 패키지 업데이트
$ sudo yum-config-manager --add-repo https://rpm.releases.hashicorp.com/AmazonLinux/hashicorp.repo

# 테라폼 설치
$ sudo yum -y install terraform
```

역시 설치가 완료되면, 다음을 입력하여, 제대로 설치가 되어 있는지 확인한다.

```bash
# terraform 버전 확인
$ terraform version
Terraform v0.14.2
```

> 만약 위 명령어가 잘 먹히지 않는다면 EC2 접속을 해제한 후 다시 접속을 해보세요. 그러면 정상적으로 실행이 될 것입니다.

