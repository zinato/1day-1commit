# 2장 Packer 용어

이번 장에서는 `Packer` 관련 용어들을 정리한다. `Packer` 관련 전반적으로 사용하는 용어들이니까 한 번쯤은 정리해볼 필요가 있다. 하지만 매우 간단하니까 쭉 훑는 느낌으로 공부해보자.

이 장을 진행하더라도 알쏭달쏭한 용어들이 있을 것이다. 하지만 추후 진행되는 장에서 `Packer`를 함께 사용하고 익혀나가다 보면 자연스럽게 용어들이 익혀질 것이니 걱정하지 말라.

## Artifact

`Artifact`는 단일 `Build`의 결과물이며, 보통 머신 이미지를 나타내는 ID 혹은 File 세트이다. 모든 `Builder`들은 한 개의 `Artifact`를 생성한다. 예를 들어 `Amazon EC2 Builder`의 경우, 결과물인 `Artifact`는 "AMI ID(리전 당 1개) 세트"이다. `VMware Builder`의 경우엔, 생성된 가상 머신을 구성하는 파일의 디렉토리가 `Artifact`가 된다.

## Build

`Build`는 단일 플랫폼에서 이미지를 생성하는 단일 작업이라고 볼 수 있다. `Packer`는 여러 개의 `Build`를 병렬적으로 실행할 수 있다. 즉, 단일 플랫폼에서 여러 개의 이미지를 동시에 생성할 수 있으며, 여러 플랫폼에서 각 이미지들을 동시에 생성할 수 있다.

## Builder

`Builder`는 단일 플랫폼에서 머신 이미지를 생성할 수 있는 `Packer`의 구성 요소이다. `Builder`는 몇몇 설정을 읽고, 그것을 머신 이미지를 생성하고 실행하는데 사용한다. `Builder`는 실제 결과 이미지를 생성하기 위해서, `Build` 작업의 일부로써 호출된다. `Builder`의 대표적인 예는 `VirtualBox`, `VMware`, `Amazon EC2`이다.

## Command(=명령어)

`Packer`는 `Command Line Interface Tool`에 속한다. 즉, 명령어를 통해서 여러 작업을 수행할 수 있다. `Packer`가 제공하는 명령어는 다음과 같다.

* init - Packer Plugin을 설치
* build - 머신 이미지를 생성
* console - Packer variable에 대해 테스트하는데 쓰임
* fix - 버전 호환성을 맞춰줌
* fmt - 표준 문법과 스타일로 맞춰줌.
* inspect - 템플릿에 정의된 다양한 구성요소를 출력함.
* validate - 설정/문법 확인
* hcl2_upgrade - .json 파일을 .hcl 파일로 변경

보통 많이 쓰는 명령어는 `Template`의 설정이나 문법에 이상이 없는지 확인하는 `validate` 그리고 실제 이미지를 만들어내는 `build`가 있다.

## Data Source

`Data Source`는 `Pakcer`의 구성요소로써, `Packer` 밖에서 정의된 데이터를 가져와서 `Template` 안에서 사용한다. 예를 들어 `amazon-ami`, `amazon-secretsmanager`가 있다.

## Post Processor

`Packer`의 구성 요소 중 하나로, 새로운 `Artifact`를 만들기 위해서 `Builder` 혹은 다른 `post-proceesor`의 결과를 가져오고 처리한다. 대표적인 `post-processor`의 예는 `Artifact`들을 압축하기 위한 압축, `Artifact`를 업로드 하는 것들이 있다.

## Provisioner

`Packer`의 구성 요소 중 하나로, 머신 이미지가 만들어지기 전에, 머신을 실행하여, 소프트웨어 설치 및 설정을 진행한다. 구성하는 머신 이미지에 대해서 유용한 소프트웨어를 설치하는데 필수적인 작업이다. 대표적인 `Provisoner`의 예는 `shell script`, `Chef`, `Ansible` 등이 있다.

## Template

`Template`은 `Packer`의 구성 요소로써, JSON(.json) 파일 혹은 HCL(.pkr.hcl) 파일로 만들어지며, 하나 이상의 `Build`가 설정된다. `Packer`는 `Template`을 읽을 수 있으며, 해당 정보를 이용하여 여러 머신 이미지를 병렬적으로 생성한다.

## 참고

* Packer 공식 문서 - Packer Terminology : [https://www.packer.io/docs/terminology](https://www.packer.io/docs/terminology)