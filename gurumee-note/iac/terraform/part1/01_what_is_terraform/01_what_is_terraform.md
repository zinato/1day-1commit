# 01장. Terraform과 설치

## Terraform이란 무엇인가

`Terraform`은 하시코프에서 만든 `IaC` 도구, 특히 인프라 선언 도구이다. `Terraform`은 안전하고 반복적으로 작업하더라도 인프라스트럭처를 구축, 변경할 수 있게 도와준다. 간혹 `Ansible`이나 `Puppet`과 같이 비교가 되곤 하는데, 엄밀히 말해서 `Ansible`과 `Puppet`은 설정 관리 도구로써 `Terraform`과는 다른 성격의 도구임을 분명히 알아두는 것이 좋다.  

### IaC와 Terraform

`IaC`란, `Infrastructure As Code`의 약자로써, 코드로 인프라를 관리하는 것을 말한다. 여기서 `IaC`가 관리하는 것은, 인프라를 이루는 서버, 미들웨어, 서비스 등 인프라를 구성하는 모든 요소들이 그 대상이다. 또한 `IaC` 도구는 크게 다음으로 분류된다.

* 서버 템플릿 ex) Packer, Docker
* 컨테이너 오케스트레이션 ex) Kubernetes
* 설정 관리 ex) Ansible, Puppet
* 인프라 선언 ex) Terraform

 `Terraform`은 0.x 버전임에도 불구하고 인프라 선언 도구 분야에서 표준으로 자리 잡았다. `HCL`이라는 도메인 특화 언어를 통해서 `AWS`, `GCP` 등의 퍼블릭 클라우드는 물론 여러 프로바이더의 원하는 인프라 구성을 자동화할 수 있다.

### Terraform의 장점

`Terraform`의 장점은 다음과 같다.

1. 코드로써 인프라를 관리하기 때문에, 생산성, 재사용성이 높아지며, 유지보수에 용이하다.
2. `Git`과 같은 `VCS`를 함께 쓰면 작업 이력이 남기 때문에, 문제 원인 파악이 보다 쉽다.
3. 업계 표준이기 때문에 예제가 풍부하여 배우기 쉽다.

### Terraform의 구성 요소

`Terraform`의 구성 요소는 크게 다음과 같다.

* provider : 테라폼으로 관리할 인프라의 종류를 의미 (ex) AWS, GCP, Azure ...)
* resource : 테라폼으로 관리할 인프라 자원을 의미 (ex) EC2, IAM, S3, RDS ...)
* state : 테라폼을 통해 생성된 리소스들의 상태를 의미. (= 테라폼 apply 명령어를 실행한 결과물)
* output : 테라폼으로 만든 리소스를 변수 형태로 state에 저장하는 것을 의미
* module : 공통적으로 활용할 수 있는 모듈을 정의하는 것을 의미
* remote : 다른 경로의 state를 참조하는 것을 의미하며, output 변수를 불러올 때 주로 사용

## Terraform 설치

이절에서는 `Mac OS` 환경과, `Linux` 환경 그 중에서도 `AWS Linux` 환경에서 `Terraform`을 어떻게 설치하는지를 다룬다.

### Mac OS 환경에서 Terraform 설치

`Mac OS`에서는 `HomeBrew` 패키지 매니저가 설치되었다고 가정하고 진행한다.

> 참고 HomeBrew 설치
> 
> Mac OS의 패키지 매니저 HomeBrew는 다음 [링크](https://brew.sh/index_ko)에서 다운로드 할 수 있습니다. 

`Terraform`을 메뉴얼하게 바로 설치해도 좋지만, 실제 개발 환경에 따라서 여러 버전을 사용할 가능성이 있다. 때문에 `Terraform` 버전 관리 도구인 `tfenv`를 먼저 설치하고, 이를 이용해서 `Terraform`을 설치한다. 터미널에 다음을 입력한다.

```bash
# tfenv 설치
$ brew install tfenv
```

설치가 완료되었으면, 다음을 입력하여 `tfenv`가 잘 설치되었는지 확인한다.

```bash
# tfenv 설치 확인
$ tfenv --version
tfenv 2.2.0
```

이제 `tfenv`를 이용해서 `Terraform`을 설치한다. 터미널에 다음 두 개의 명령어 중 하나를 선택해서 설치한다.

```bash
# terraform 최신 버전 설치
$ tfenv install

# terraform 특정 버전 설치를 원한다면 다음과 같은 형식으로 할 수 있다.
$ tfenv install 0.14.6
```

21년 2월 17일 기준, 0.14.6 버전이 설치된다. 이제 `Mac OS`에서 설치한 `Terraform` 버전을 사용하게끔 설정한다. 터미널에 다음을 입력한다.

```bash
# terraform 버전 사용
$ tfenv use 0.14.6
```

그 후 `Terraform` 버전을 확인해보자.

```bash 
# terraform 버전 확인
$ terraform version
Terraform v0.14.6
```

잘 설치되었다.

### AWS Linux 환경에서 Terraform 설치

이 절에서는 `AWS EC2` 인스턴스 1개를 이미 가동시킨 후라고 가정한다. `AWS Linux` 환경에서는 메뉴얼하게 `Terraform`을 설치할 것이다. 먼저 `EC2`에 접속한다. 그 후, 터미널에 다음 명령어를 차례대로 입력한다.

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

