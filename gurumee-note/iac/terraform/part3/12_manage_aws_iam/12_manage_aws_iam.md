# 12장. 테라폼으로 유저 관리하기 - IAM

## AWS IAM이란 무엇인가?

`AWS IAM`이란, AWS 리소스에 대한 액세스를 안전하게 제어할 수 있는 웹 서비스이다. IAM을 이용하여, 리소스를 사용하도록 인증 및 권한 부여된 대상을 제어한다.

그냥 쉽게 AWS 리소스를 제어할 수 있는 유저, 그룹, 정책 등을 관리하는 것이라고 생각하면 된다. 이에 대한 자세한 내용은 다음을 참고하라.

* [AWS 공식 문서 - IAM이란 무엇인가요](https://docs.aws.amazon.com/ko_kr/IAM/latest/UserGuide/introduction.html)

## 우리가 이번에 구성할 것은?

우리가 이번에 구성할 인프라는 바로, 유저와 그룹이다. `Cortex 클러스터`를 AWS에서 구성하기 위해서는, `IAM`, `S3`, `DynamoDB`, `AutoScaling` 등 권한이 필요하다. (물론 다른 방법도 있다!)

이들의 권한을 가진 정책들을 생성하고 유저와 그룹에 나누어서 할당할 것이다.

* 그룹
  * 이름 : cortex_cluster
  * 권한 : IAMFullAccess, AmazonS3FullAccess, AmazonDynamoDBFullAccess
* 유저 
  * 이름 : auto_scaler 
  * 속한 그룹 : cortex_cluster
  * 권한 : AutoScalingFullAccess

이번 장에서 쓰일 코드는 다음 URL에서 얻을 수 있다.

* 코드 링크 : [https://github.com/gurumee92/gurumee-terraform-code/tree/master/part3/ch12](https://github.com/gurumee92/gurumee-terraform-code/tree/master/part3/ch12)

## AWS Provider 구성

먼저 프로바이더를 구성한다.

part3/ch12/provider.tf
```tf
provider "aws" {
  region = "us-east-1"
}
```

끝이다. `region`의 경우는 적절하게 값을 주면 된다. 대한민국에서 코드를 작성하는 것이라면 `ap-northeast-2`를 추천한다.

나는 이미 이 지역에 리소스가 있어서 `us-east-1`로 지정했다. 별 차이는 없다. 다만 다른 지역에 AWS 리소스가 생성되어 속도 차이가 조금 있을 뿐..?

이제 이 `provider.tf`가 있는 위치에서 `terraform init` 명령어를 실행시킨다.

```bash
$ terraform init
```

## 유저와 그룹 생성

## 유저와 정책 연결

## 그룹과 정책 연결