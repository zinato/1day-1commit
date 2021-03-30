# 13장. 테라폼으로 AWS VPC 관리하기

## AWS VPC이란 무엇인가?

`Amazon Virtual Private Cloud`(이하 Amazon VPC 혹은 VPC)는 AWS 사용자가 정의한 "가상 네트워크"로써 cidr 블록 방식으로 IP 대역을 설정한다.

`VPC`의 대표적인 구성 요소는 다음과 같다.

### Subnet

`Subnet`은 `VPC` 대역 안의 IP 주소 범위를 cidr 블록 방식으로 지정한다. 보통은 하나의 `Availability Zone`을 외부에서 통신할 수 있는 public subnet과 외부에서 통신이 불가능한 private subnet으로 나눈다.

또한 HA 구성을 위해 `VPC` 내부에 public-pirvate subnet 구조가 두 쌍이 되도록 만든다. 

### Internet Gateway(IGW)

`VPC` 내부의 `public subnet` 상의 인스턴스들과 외부 인터넷 간의 통신할 수 있도록 `VPC`에 연결하는 게이트웨이를 뜻한다. `route table`을 통해서, 연결된다.

### NAT Gateway(NGW)

`NAT`는 네트워크 주소를 변환하는 장치를 뜻한다. `NGW`는 `VPC` 내부의 `prviate subnet`의 인스턴스들과 인터넷/AWS 서비스에 연결하는 게이트웨이이다. 

내부적으로 `public subnet`에 위치하고 있으며, `elastic IP`를 할당한 상태로 구성된다. 역시, `route table`을 통해서 연결된다.

참고적으로 IPv6는 지원하지 않으므로 아웃바운드 전용 인터넷 게이트웨이를 사용해야 한다.

### Routing Table

`route table`은 네트워크 트래픽을 전달할 위치를 결정하는데 사용되는 라우팅 규칙들의 집합이다. `subnet`과 `subnet` 간 통신, `subnet`과 `gateway`들의 통신을 결정한다.

### Network ACL

`Security Group`과 함께 `AWS VPC`의 보안 장치 중 하나이다. 1개 이상의 `subnet`과 외부의 트래픽을 제어할 수 있다. 쉽게 말해서, `subnet` 단위의 보안 계층이라고 보면 된다. "stateless"하다라는 특징을 가지고 있다.

`Network ACL`은 여러 `subnet`과 연결이 가능하지만 반대로 `subnet`은 딱 1개의 `Network ACL`을 가져야만 한다.

### Security Group 

`Network ACL`과 함께 `AWS VPC`의 보안 장치 중 하나이다. `Security Group`은 인스턴스 단위의 보안 계층이며, 인스턴스에 대한 인바운드 및 아웃바운드 트래픽을 제어하는 가상 방화벽 역할을 한다. `Network ACL`과는 달리 "stateful"하다.

`Security Group`은 각 인스턴스 당 최대 5개까지 적용이 가능하다. `Network ACL`과 `Security Group`을 혼용해서 쓰게 되면 보안 이슈 해결 시 복잡도가 매우 높아지기 떄문에 보통은 둘 중 한 가지를 선택해서 `VPC`에 보안을 적용한다.

우리는 `Security Group`을 기반으로 `VPC` 보안을 설정할 것이다.

## 우리가 이번에 구성할 것은?

우리가 구성할 인프라스트럭처는 다음과 같다.

![01](./01.png)

각 구성 요소를 나열하면 다음과 같다.

* VPC 1개
* subnet 4개 (public 2개, private 2개)
* internet gateway(igw) 1개
* nat gateway(ngw) 1개
    * eip 1개
* route table 2개 (public - igw, private - ngw)
    * route table assosiation 5개 (public 2개, private 3개)
* network acl 1개
* security group 3개
    * default (private subnet 인스턴스만 적용)
    * VPN (여기서는 자신의 IP를 VPN 취급)
    * 웹 서버용 (public subnet 인스턴스만 적용 80, 443번 port만 허용)

이번 장에서 만들 코드는 다음에서 확인할 수 있다.

* 코드 링크 : [https://github.com/gurumee92/gurumee-terraform-code/tree/master/part3/ch13](https://github.com/gurumee92/gurumee-terraform-code/tree/master/part3/ch13)

## AWS VPC 리소스 생성

### VPC

### Subnet

### Internet Gateway

### NAT Gateway

### Route Table

### Network ACL

### Securtity Group

## AWS VPC 구성 테스트