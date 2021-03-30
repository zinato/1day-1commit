# 13장. Terraform으로 AWS VPC 관리하기

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

* 코드 링크 : [https://github.com/gurumee92/gurumee-terraform-code/tree/master/part3/ch11](https://github.com/gurumee92/gurumee-terraform-code/tree/master/part3/ch11)

## AWS VPC 리소스 생성

part3/ch11/provider.tf
```tf
provider "aws" {
  region = "us-east-1"
}
```

### VPC

part3/ch11/vpc.tf
```tf
# vpc
resource "aws_vpc" "vpc" {
    cidr_block = "10.10.0.0/16"

    tags = {
        Name = "vpc"
    }
}
```

### Subnet

part3/ch11/vpc.tf
```tf
# ...
# subnet (public)
resource "aws_subnet" "public_subnet_1a" {
    vpc_id     = aws_vpc.vpc.id
    cidr_block = "10.10.1.0/24"
    availability_zone = "us-east-1a"
    tags = {
        Name = "public_subnet_1a"
    }
}

resource "aws_subnet" "public_subnet_1b" {
    vpc_id     = aws_vpc.vpc.id
    cidr_block = "10.10.2.0/24"
    availability_zone = "us-east-1b"
    tags = {
        Name = "public_subnet_1b"
    }
}

# subnet (private)
resource "aws_subnet" "private_subnet_1a" {
    vpc_id     = aws_vpc.vpc.id
    cidr_block = "10.10.101.0/24"
    availability_zone = "us-east-1a"
    tags = {
        Name = "private_subnet_1a"
    }
}

resource "aws_subnet" "private_subnet_1b" {
    vpc_id     = aws_vpc.vpc.id
    cidr_block = "10.10.102.0/24"
    availability_zone = "us-east-1b"
    tags = {
        Name = "private_subnet_1b"
    }
}
```

### Internet Gateway

part3/ch11/vpc.tf
```tf
# ...
# igw
resource "aws_internet_gateway" "igw" {
    vpc_id = aws_vpc.vpc.id

    tags = {
        Name = "Internet Gateway"
    }
}
```

### NAT Gateway

part3/ch11/vpc.tf
```tf
# ...
# ngw
resource "aws_eip" "ngw_ip" {
    vpc   = true

    lifecycle {
        create_before_destroy = true
    }
}

resource "aws_nat_gateway" "ngw" {
    allocation_id = aws_eip.ngw_ip.id
    subnet_id     = aws_subnet.public_subnet_1a.id

    tags = {
        Name = "NAT Gateway"
    }
}
```

### Route Table

part3/ch11/vpc.tf
```tf
# ...
# route table (public)
resource "aws_default_route_table" "public_rt" {
    default_route_table_id = aws_vpc.vpc.default_route_table_id
    
    route {
        cidr_block = "0.0.0.0/0"
        gateway_id = aws_internet_gateway.igw.id
    }

    tags = {
        Name = "public route table"
    }
}

resource "aws_route_table_association" "public_rta_a" {
    subnet_id      = aws_subnet.public_subnet_1a.id
    route_table_id = aws_default_route_table.public_rt.id
}

resource "aws_route_table_association" "public_rta_b" {
    subnet_id      = aws_subnet.public_subnet_1b.id
    route_table_id = aws_default_route_table.public_rt.id
}

# route table (private)
resource "aws_route_table" "private_rt" {
    vpc_id = aws_vpc.vpc.id
    tags = {
        Name = "private route table"
    }
}

resource "aws_route_table_association" "private_rta_a" {
    subnet_id      = aws_subnet.private_subnet_1a.id
    route_table_id = aws_route_table.private_rt.id
}

resource "aws_route_table_association" "private_rta_b" {
    subnet_id      = aws_subnet.private_subnet_1b.id
    route_table_id = aws_route_table.private_rt.id
}

resource "aws_route" "private_rt_route" {
    route_table_id              = aws_route_table.private_rt.id
    destination_cidr_block      = "0.0.0.0/0"
    nat_gateway_id              = aws_nat_gateway.ngw.id
}
```

### Network ACL

part3/ch11/vpc.tf
```tf
# ...
# network acl
resource "aws_default_network_acl" "vpc_network_acl" {
    default_network_acl_id = aws_vpc.vpc.default_network_acl_id
    
    egress {
        protocol   = "tcp"
        rule_no    = 100
        action     = "allow"
        cidr_block = "0.0.0.0/0"
        from_port  = 0
        to_port    = 65535
    }

    ingress {
        protocol   = "-1"
        rule_no    = 100
        action     = "allow"
        cidr_block = "0.0.0.0/0"
        from_port  = 0
        to_port    = 0
    }

    tags = {
        Name = "network acl"
    }
}
```

### Securtity Group

part3/ch11/vpc.tf
```tf
# ...
# security group
resource "aws_default_security_group" "default_sg" {
    vpc_id = aws_vpc.vpc.id

    ingress {
        protocol    = "tcp"
        from_port = 0
        to_port   = 65535
        cidr_blocks = [aws_vpc.vpc.cidr_block]
    }

    egress {
        protocol    = "-1"
        from_port   = 0
        to_port     = 0
        cidr_blocks = ["0.0.0.0/0"]
    }

    tags = {
        Name = "default_sg"
        Description = "default security group"
    }
}

resource "aws_security_group" "inhouse_sg" {
    name        = "pinhouse_sg"
    description = "security group for inhouse"
    vpc_id      = aws_vpc.vpc.id

    ingress {
        description = "For Inhouse ingress"
        from_port   = 0
        to_port     = 65535
        protocol    = "tcp"
        cidr_blocks = [
            aws_vpc.vpc.cidr_block,
            "121.161.72.112/32",
        ]
    }

    egress {
        protocol    = "-1"
        from_port   = 0
        to_port     = 0
        cidr_blocks = ["0.0.0.0/0"]
    }

    tags = {
        Name = "inhouse_sg"
    }
}

resource "aws_security_group" "web_server_sg" {
    name        = "web_server_sg"
    description = "security group for web server"
    vpc_id      = aws_vpc.vpc.id

    ingress {
        description = "For http port"
        from_port   = 80
        to_port     = 80
        protocol    = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }

    ingress {
        description = "For https port"
        from_port   = 443
        to_port     = 443
        protocol    = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }

    egress {
        protocol    = "-1"
        from_port   = 0
        to_port     = 0
        cidr_blocks = ["0.0.0.0/0"]
    }

    tags = {
        Name = "web_server_sg"
    }
}
```