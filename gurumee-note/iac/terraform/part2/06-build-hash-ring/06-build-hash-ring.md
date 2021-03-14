# 6장. 테라폼으로 Hash Ring 구성하기

이번 장은 사실 지난 장에 복습이다. 먼저 문서에 도움 없이 진행해보라. 그리고 막히는 부분이 있으면, 그 때 문서를 참고해보라. `Terraform`이 익숙해지는데 도움이 될 것이다.

## 무엇을 구성할 것인가

우리는 이번 장에서 무엇을 구성할 것인가. 다음 그림을 살펴보자.

![01](./01.png)

그림에서 살펴보면 `Consul`이란 것이 맨 중앙에서 `Distributor`, `Ingester`, `Querier`, `Ruler`를 연결하고 있다. 그 외에도 `AlertManager`, `QueryFrontend` 가 있는데 이들이 바로 `Cortex`이다.

즉 저 컴포넌트들이 모여서 우리가 만들 `Cortex 클러스터`를 구성하는 것이다. 클러스터 내에서 구성된 `Cortex`들은 서로 돌아가며 데이터를 처리하기 위해서 같은 클러스터에 묶인 `Cortex`들을 알아야 한다. 그 기능을 해주는 것이 바로 "Hash Ring"이다. 이 "Hash Ring"은 대표적으로 `etcd`나 `Consul`로 구성할 수 있다. 

이번 장에서 우리는 `Terraform`을 이용해서 `Consul`을 "Hash Ring"으로 구성한다. 역시 이번 장에서도 이미지와 컨테이너를 `Terraform`으로 구성해보자.

이번 장에서 쓰일 코드는 다음 링크에서 확인할 수 있다.

* 코드 링크 : [https://github.com/gurumee92/gurumee-terraform-code/tree/master/part2/ch06](https://github.com/gurumee92/gurumee-terraform-code/tree/master/part2/ch06)

> 참고! 이 장을 진행하기 전에
> 
> 이 장은 5장을 먼저 진행했음을 가정하고 합니다. 즉 이전 장에서 구성했던 tf 파일들(provider, network)이 이미 존재한다고 가정합니다. 없다면 [https://github.com/gurumee92/gurumee-terraform-code/tree/master/part2/ch05](https://github.com/gurumee92/gurumee-terraform-code/tree/master/part2/ch05)를 참고하세요.


## Consul 이미지 구성하기

이제 `Consul` 이미지를 구성해보자. 이전 장과 동일하다. `ring.tf`를 다음과 같이 입력한다.

part2/ch06/ring.tf
```tf
resource "docker_image" "consul" {
  name         = "consul:latest"
  keep_locally = false
}
```

이제 인프라 리소스를 추가했으니 `terraform plan` 명령어를 입력하자.

> 참고! 만약 5장을 건너 뛰었다면?
> 
> 5장을 건너뛰었다면 ring.tf가 없는 상태에서 terraform init 명령어와 terraform apply 명령어를 먼저 실행해주세요.

```bash
$ terraform plan
docker_image.cassandra: Refreshing state... [id=sha256:9ea2636247a5c934c61c4332f5fdf7c2dde7feb43b316dc3a7a5de4d656aef6ccassandra:latest]
docker_network.cortex-cluster: Refreshing state... [id=690b961cfa4ae2429dc0544fe21d006aabd2dafd338c92b63b5201b907231c36]
docker_container.cortex-cluster-store: Refreshing state... [id=49d99af2f8f0ca0221ac6e2dcf30c44a6f85d126cf4d8c17865179d77ba1096c]

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # docker_image.consul will be created
  + resource "docker_image" "consul" {
      + id           = (known after apply)
      + keep_locally = false
      + latest       = (known after apply)
      + name         = "consul:latest"
      + output       = (known after apply)
    }

Plan: 1 to add, 0 to change, 0 to destroy.

------------------------------------------------------------------------

Note: You didn't specify an "-out" parameter to save this plan, so Terraform
can't guarantee that exactly these actions will be performed if
"terraform apply" is subsequently run.
```

아직 `terraform plan`만 했기 때문에 바뀌는 것은 없다. 역시 로그를 확인하면, `Consul` 이미지가 생성될 거라는 것을 확인할 수 있다. 이제 인프라에 적용해보자. `terraform apply` 명령어를 입력한다.

```bash
$ terraform apply
docker_image.cassandra: Refreshing state... [id=sha256:9ea2636247a5c934c61c4332f5fdf7c2dde7feb43b316dc3a7a5de4d656aef6ccassandra:latest]
docker_network.cortex-cluster: Refreshing state... [id=690b961cfa4ae2429dc0544fe21d006aabd2dafd338c92b63b5201b907231c36]
docker_container.cortex-cluster-store: Refreshing state... [id=49d99af2f8f0ca0221ac6e2dcf30c44a6f85d126cf4d8c17865179d77ba1096c]

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # docker_image.consul will be created
  + resource "docker_image" "consul" {
      + id           = (known after apply)
      + keep_locally = false
      + latest       = (known after apply)
      + name         = "consul:latest"
      + output       = (known after apply)
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: 
```

여기서 "yes"를 입력하고 엔터를 치면 이제 도커 이미지가 설치된다. 터미널에 다음처럼 로그가 보일 것이다.

```bash
docker_image.consul: Creating...
docker_image.consul: Still creating... [10s elapsed]
docker_image.consul: Creation complete after 13s [id=sha256:d544f4c4e87c388d3535d758860bbc15cc6369ed977d6d8d361936e79e913576consul:latest]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
```

`Consul` 이미지가 생성되었음을 알 수 있다. 이제 도커 이미지 목록에 `consul:latest`가 있는지 확인해보자. 터미널에 `docker images`를 입력한다.

```bash
$ docker images
REPOSITORY           TAG                 IMAGE ID       CREATED         SIZE
consul               latest              d544f4c4e87c   2 days ago      120MB
...
```

## Consul 컨테이너 구성하기

이제 컨테이너를 구성한다. 구성하기 전에, `Cortex`에서 `Consul`을 사용할 때 주는 옵션이 있는데 다음과 같다.

```bash
$ docker run -d --name=consul --network=cortex -e CONSUL_BIND_INTERFACE=eth0 consul
```

바로 환경 변수에 "CONSUL_BIND_INTERFACE=eth0" 값을 주어야 한다. 어떻게 해야 할까? 답은 테라폼 레지스트리에서 확인할 수 있다.

* 테라폼 레지스트리 env 문서 : [https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/resources/container#env](https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/resources/container#env)

결국 `env` 값에 `문자열 Set`의 값을 주면 된다. `ring.tf`에 다음을 추가한다.

part2/ch06/ring.tf
```tf
# ...

resource "docker_container" "cortex-cluster-hash-ring" {
  image = docker_image.consul.latest
  name  = "cortex-cluster-hash-ring"
  network_mode = "bridge"
  networks_advanced {
      name = docker_network.cortex-cluster.name
  }
  env = ["CONSUL_BIND_INTERFACE=eth0"]
}
```

`Terraform`에서 `List/Set` 타입은 대괄호`[]` 를 사용해서 값을 지정해주면 된다. 이 외에도 `networks_advanced`처럼 "json" 형식으로 표현할 수 있는 `Map/Object/Block` 타입도 존재한다. 그 외 문자열인 `String`, 숫자 `Number`, 부울 값인 `Bool` 그리고 Null을 표현하는 `Null` 타입이 있다.

이제 리소스를 추가했으니 `terraform plan` 명령어를 입력한다.

```bash
$ terraform plan
docker_image.cassandra: Refreshing state... [id=sha256:9ea2636247a5c934c61c4332f5fdf7c2dde7feb43b316dc3a7a5de4d656aef6ccassandra:latest]
docker_image.consul: Refreshing state... [id=sha256:d544f4c4e87c388d3535d758860bbc15cc6369ed977d6d8d361936e79e913576consul:latest]
docker_network.cortex-cluster: Refreshing state... [id=690b961cfa4ae2429dc0544fe21d006aabd2dafd338c92b63b5201b907231c36]
docker_container.cortex-cluster-store: Refreshing state... [id=49d99af2f8f0ca0221ac6e2dcf30c44a6f85d126cf4d8c17865179d77ba1096c]

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # docker_container.cortex-cluster-hash-ring will be created
  + resource "docker_container" "cortex-cluster-hash-ring" {
      + attach           = false
      + bridge           = (known after apply)
      + command          = (known after apply)
      + container_logs   = (known after apply)
      + entrypoint       = (known after apply)
      + env              = [
          + "CONSUL_BIND_INTERFACE=eth0",
        ]
      + exit_code        = (known after apply)
      + gateway          = (known after apply)
      + hostname         = (known after apply)
      + id               = (known after apply)
      + image            = "sha256:d544f4c4e87c388d3535d758860bbc15cc6369ed977d6d8d361936e79e913576"
      + init             = (known after apply)
      + ip_address       = (known after apply)
      + ip_prefix_length = (known after apply)
      + ipc_mode         = (known after apply)
      + log_driver       = "json-file"
      + logs             = false
      + must_run         = true
      + name             = "cortex-cluster-hash-ring"
      + network_data     = (known after apply)
      + network_mode     = "bridge"
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

      + networks_advanced {
          + aliases = []
          + name    = "cortex-cluster"
        }
    }

Plan: 1 to add, 0 to change, 0 to destroy.

------------------------------------------------------------------------

Note: You didn't specify an "-out" parameter to save this plan, so Terraform
can't guarantee that exactly these actions will be performed if
"terraform apply" is subsequently run.
```

`Consul` 도커 컨테이너가 구성되는 것을 볼 수 있다. 이제 도커 컨테이너를 띄워보자. `terraform apply` 명령어를 입력한다.

```bash
$ terraform apply
docker_image.consul: Refreshing state... [id=sha256:d544f4c4e87c388d3535d758860bbc15cc6369ed977d6d8d361936e79e913576consul:latest]
docker_image.cassandra: Refreshing state... [id=sha256:9ea2636247a5c934c61c4332f5fdf7c2dde7feb43b316dc3a7a5de4d656aef6ccassandra:latest]
docker_network.cortex-cluster: Refreshing state... [id=690b961cfa4ae2429dc0544fe21d006aabd2dafd338c92b63b5201b907231c36]
docker_container.cortex-cluster-store: Refreshing state... [id=49d99af2f8f0ca0221ac6e2dcf30c44a6f85d126cf4d8c17865179d77ba1096c]

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # docker_container.cortex-cluster-hash-ring will be created
  + resource "docker_container" "cortex-cluster-hash-ring" {
      + attach           = false
      + bridge           = (known after apply)
      + command          = (known after apply)
      + container_logs   = (known after apply)
      + entrypoint       = (known after apply)
      + env              = [
          + "CONSUL_BIND_INTERFACE=eth0",
        ]
      + exit_code        = (known after apply)
      + gateway          = (known after apply)
      + hostname         = (known after apply)
      + id               = (known after apply)
      + image            = "sha256:d544f4c4e87c388d3535d758860bbc15cc6369ed977d6d8d361936e79e913576"
      + init             = (known after apply)
      + ip_address       = (known after apply)
      + ip_prefix_length = (known after apply)
      + ipc_mode         = (known after apply)
      + log_driver       = "json-file"
      + logs             = false
      + must_run         = true
      + name             = "cortex-cluster-hash-ring"
      + network_data     = (known after apply)
      + network_mode     = "bridge"
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

      + networks_advanced {
          + aliases = []
          + name    = "cortex-cluster"
        }
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: 
```

역시 `terraform plan` 명령어에서 확인했던 추가되는 리소스를 확인할 수 있다. "yes"를 입력한다. 

```bash
docker_container.cortex-cluster-hash-ring: Creating...
docker_container.cortex-cluster-hash-ring: Creation complete after 0s [id=2fcae265254ba7dcb1a7e45d360b8e6e0bb2145ebbebae1eeffbbea6d524cf2b]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
```

`Consul` 도커 컨테이너가 생성되었다. 이제 `docker ps` 명령어를 통해서, 실행이 되고 있는지 확인해보자.

```bash
$ docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS          PORTS                                                        NAMES
2fcae265254b   d544f4c4e87c   "docker-entrypoint.s…"   49 seconds ago   Up 48 seconds   8300-8302/tcp, 8500/tcp, 8301-8302/udp, 8600/tcp, 8600/udp   cortex-cluster-hash-ring
49d99af2f8f0   9ea2636247a5   "docker-entrypoint.s…"   19 minutes ago   Up 19 minutes   7000-7001/tcp, 7199/tcp, 9042/tcp, 9160/tcp                  cortex-cluster-store
```

`Consul` 도커 컨테이너가 실행되고 있음을 확인할 수 있다. 지난 장 스토리지에 이어서 "Hash Ring"까지 구성을 마쳤다. 