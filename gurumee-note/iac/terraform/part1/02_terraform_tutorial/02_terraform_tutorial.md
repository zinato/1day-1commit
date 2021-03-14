# 2장 테라폼 맛보기

## 도커 설치하기

이 장의 실습을 진행하기 위해서는 `Docker`가 설치되어 있어야 한다. 여기서는 MacOS와 AWS Linux 상에서 설치하는 방법을 알아본다.

### MacOS에서 도커 설치

먼저 MacOS의 경우에는 다음 문서에서 다운로드할 수 있다.

* [도커 공식 문서 - 다운로드](https://docs.docker.com/docker-for-mac/install/)

만일 `Home Brew` 패키지 매니저가 있다면, 다음 명령어로 간단하게 설치할 수 있다.

```bash
# 도커 설치
$ brew cask install docker

# 도커 설치 확인
$ docker --version
Docker version 20.10.2, build 2291f61
```

### AWS Linux에서 도커 설치

터미널에 다음을 입력한다.

```bash
# yum 패키지 업데이트
$ sudo yum -y upgrade

# 도커 설치
$ sudo yum -y install docker
```

이제 잘 설치되었는지 확인한다.

```bash
# 도커 버전 확인
$ docker --version
Docker version 19.03.13-ce, build 4484c46
```

그 후, 도커를 서비스 형태로 실행시켜야 한다.

```bash
# 도커 서비스 실행
$ sudo systemctl start docker

# 도커 서비스 실행 상태 확인
$ sudo systemctl status docker
● docker.service - Docker Application Container Engine
   Loaded: loaded (/usr/lib/systemd/system/docker.service; disabled; vendor preset: disabled)
   Active: active (running) since 일 2021-02-21 13:17:26 UTC; 2min 48s ago
     Docs: https://docs.docker.com
     ...
```

근데 리눅스에서는 바로 쓸 수가 없다. 다음 명령어를 쳐보자.

```bash
$ docker ps
docker: Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Post http://%2Fvar%2Frun%2Fdocker.sock/v1.38/containers/create: dial unix /var/run/docker.sock: connect:
```

그럼 위와 같은 에러 문구가 뜨는데 `ec2-user`가 도커에 접근할 권한이 없기 때문이다. 권한을 주면 손쉽게 해결된다. 터미널에 다음을 입력한다.

```bash
# ec2-user에 도커 접근 권한 부여
$ sudo usermod -aG docker ec2-user
```

그 후 터미널을 껐다가 다시 접속한다. 그 후 `docker ps` 명령어를 한 번 더 친다.

```bash
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
```

위와 같이 뜨면, 무사히 설치가 완료된 것이다.
 
## 테라폼으로 도커 컨테이너 구성하기

이 절의 전체 코드는 다음에 존재한다.

* [https://github.com/gurumee92/gurumee-terraform-code/tree/master/part1/ch02](https://github.com/gurumee92/gurumee-terraform-code/tree/master/part1/ch02)

자 이제 적당한 위치에 다음 파일을 만든다. 

code/part1/ch02/main.tf
```terraform
terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
    }
  }
}

provider "docker" {}

resource "docker_image" "nginx" {
  name         = "nginx:latest"
  keep_locally = false
}

resource "docker_container" "nginx" {
  image = docker_image.nginx.latest
  name  = "tutorial"
  ports {
    internal = 80
    external = 8000
  }
}
```

일단 "terraform"은 추후에 다룰 것이니 일단 건너 뛴다. 우선 살펴볼 것은 다음 코드이다.

```terraform
provider "docker" {}
```

위 코드는 내가 관리할 인프라의 대상이 "Docker"라는 것을 명시한다. 이 프로바이더 위치에 "aws", "gcp" 등 다양한 퍼블릭 프로바이더를 쓸 수 있다. 이는 2부, 3부에서 더 자세하게 다루도록 하겠다. 그 다음으로 볼 코드는 다음 코드이다.

```terraform
resource "docker_image" "nginx" {
  name         = "nginx:latest"
  keep_locally = false
}
```

먼저 위의 resource는 "nginx" 도커 이미지를 사용할 것을 명시한다. 그리고 테라폼은 관리할 리소스에 대해서 상세 정보를 적을 수 있다. 해당 도커 이미지의 `name`은 "nginx:latest"이며, `keep_locally`는 추후 테라폼으로 리소스들을 삭제할 때 이미지까지 삭제 여부이다. true면 삭제 명령에도 도커 이미지는 삭제되지 않는다.

상세 정보는 필수적일수도 있고, 아닐 수도 있다. 잘 모를 때는 [테라폼 레지스트리](https://registry.terraform.io/)라는 곳에서 프로바이더 및 리소스에 대한 상세 정보들을 확인할 수 있다. 예를 들어 `name`은 반드시 적어 주어야 하는 필수 정보이다.

```terraform
resource "docker_container" "nginx" {
  image = docker_image.nginx.latest
  name  = "tutorial"
  ports {
    internal = 80
    external = 8000
  }
}
```

두 번째 resource는 `nginx`를 도커 컨테이너로써 실행할 것을 명시한다. 역시 필수 정보들이 있다. 여러 정보들이 있는데, `image`를 유심히 살펴보자.

```
image = docker_image.nginx.latest
```

이 `docker_image.nginx`는 우리가 정의한 첫 번째 리소스를 가리킨다. 이런 식으로 변수로써 사용할 수 있다. 이제 정말로 테라폼으로 도커 컨테이너를 띄어보자. 먼저 `main.tf`가 있는 위치에서 다음 명령어를 입력한다.

```bash
$ terraform init

Initializing the backend...

Initializing provider plugins...
- Finding latest version of kreuzwerker/docker...
- Installing kreuzwerker/docker v2.11.0...
- Installed kreuzwerker/docker v2.11.0 (self-signed, key ID 24E54F214569A8A5)

Partner and community providers are signed by their developers.
If you'd like to know more about provider signing, you can read about it here:
...
```

`terraform init` 명령어는 해당 디렉토리에서, 테라폼으로 인프라를 관리할 수 있는 상태로 만든다. 명령어를 실행하고 나면 `.terraform` 폴더와, `.terraform.lock.hcl`이 생성되는 것을 볼 수 있다. 이렇게 초기화를 진행했으면 `terraform plan`이라는 명령어로 현재 코드로 구축된 인프라가 잘 설정된 것인지 확인할 수 있다.

```bash
$ terraform plan

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # docker_container.nginx will be created
  + resource "docker_container" "nginx" {
      + attach           = false
      + bridge           = (known after apply)
      + command          = (known after apply)
      + container_logs   = (known after apply)
      + entrypoint       = (known after apply)
      + env              = (known after apply)
      + exit_code        = (known after apply)
      + gateway          = (known after apply)
      + hostname         = (known after apply)
      + id               = (known after apply)
      + image            = (known after apply)
      + init             = (known after apply)
      + ip_address       = (known after apply)
      + ip_prefix_length = (known after apply)
      + ipc_mode         = (known after apply)
      + log_driver       = "json-file"
      + logs             = false
      + must_run         = true
      + name             = "tutorial"
      + network_data     = (known after apply)
      + read_only        = false
      + remove_volumes   = true
      + restart          = "no"
      + rm               = false
      + security_opts    = (known after apply)
      + shm_size         = (known after apply)
      + start            = true
      + stdin_open       = false
      + tty              = false

      + healthcheck {
          + interval     = (known after apply)
          + retries      = (known after apply)
          + start_period = (known after apply)
          + test         = (known after apply)
          + timeout      = (known after apply)
        }

      + labels {
          + label = (known after apply)
          + value = (known after apply)
        }

      + ports {
          + external = 8000
          + internal = 80
          + ip       = "0.0.0.0"
          + protocol = "tcp"
        }
    }

  # docker_image.nginx will be created
  + resource "docker_image" "nginx" {
      + id           = (known after apply)
      + keep_locally = false
      + latest       = (known after apply)
      + name         = "nginx:latest"
      + output       = (known after apply)
    }

Plan: 2 to add, 0 to change, 0 to destroy.

------------------------------------------------------------------------

Note: You didn't specify an "-out" parameter to save this plan, so Terraform
can't guarantee that exactly these actions will be performed if
"terraform apply" is subsequently run.
```

이번에는 명령어에서 출력 문구를 생략하지 않았다. 자세히 보면 어떤 이미지가 생성되었는지 어떤 컨테이너가 포트매핑이 어떻게 되어있는가까지 확인할 수 있다. 그러나 아직까지 인프라는 구성되지 않은 상태이다. `docker ps` 명령어를 입력해보자.

```bash
$ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```

`terraform plan` 명령어는 말 그대로 코드로써 계획된 인프라가 문제가 없는지 확인하는 작업이다. 실제로 인프라를 구축하려면 `terraform apply` 명령어를 입력한다.

```bash
$ terraform apply

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

....

# 여기서는 yes라고 입력해야 한다.
Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

# 실제로 도커 이미지 생성 및 컨테이너 실행이 되는 것을 확인할 수 있다.
docker_image.nginx: Creating...
docker_image.nginx: Creation complete after 10s [id=sha256:35c43ace9216212c0f0e546a65eec93fa9fc8e96b25880ee222b7ed2ca1d2151nginx:latest]
docker_container.nginx: Creating...
docker_container.nginx: Creation complete after 0s [id=27ef04292942440ba1e9dbdb9d48d729e1fd3251808ed66e49f6d53e73c74025]

Apply complete! Resources: 2 added, 0 changed, 0 destroyed.
```

이제 터미널에 `docker ps` 입력해보자.

```bash
$ docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED              STATUS              PORTS                  NAMES
27ef04292942   35c43ace9216   "/docker-entrypoint.…"   About a minute ago   Up About a minute   0.0.0.0:8000->80/tcp   tutorial
```

`nginx`가 계획대로, 80번 포트가 머신의 8080 포트에 매핑되었고 컨테이너 이름 역시 계획한대로 "tutorials"가 되었다. 이제부터 테라폼 파일들이 추가/수정/삭제에 따라 `terraform plan`, `terraform apply` 명령어를 입력해서 수정된 인프라를 테스트 및 구축할 수 있다. 마지막으로 `terraform destroy` 명령어를 이용하면, 구축된 인프라를 모두 삭제할 수 있다.

```bash
$ terraform destroy

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  - destroy

Terraform will perform the following actions:

 ...

Plan: 0 to add, 0 to change, 2 to destroy.

# 역시 yes를 입력해야 한다.
Do you really want to destroy all resources?
  Terraform will destroy all your managed infrastructure, as shown above.
  There is no undo. Only 'yes' will be accepted to confirm.

  Enter a value: yes

# 도커 컨테이너 및 이미지가 삭제되는 것을 확인할 수 있다.
docker_container.nginx: Destroying... [id=27ef04292942440ba1e9dbdb9d48d729e1fd3251808ed66e49f6d53e73c74025]
docker_container.nginx: Destruction complete after 0s
docker_image.nginx: Destroying... [id=sha256:35c43ace9216212c0f0e546a65eec93fa9fc8e96b25880ee222b7ed2ca1d2151nginx:latest]
docker_image.nginx: Destruction complete after 0s

Destroy complete! Resources: 2 destroyed.
```

이제 터미널에 `docker ps` 입력해보자.

```bash
$ docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED              STATUS              PORTS                  NAMES
```

컨테이너가 제거된 것을 확인할 수 있다. 이렇게 도커를 이용해서 테라폼으로 인프라 초기화/테스트/구축/제거 하는 방법을 살짝 맛보았다. 이어지는 다음 장에서는 AWS 혹은 GCP 퍼블릭 프로바이더들의 리소스들을 관리하는 방법에 대해서 배워보도록 하자.