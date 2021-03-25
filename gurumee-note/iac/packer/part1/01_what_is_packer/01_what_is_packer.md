# 1장. Packer란 무엇인가?

## Packer란 무엇인가

`Hashicorp`가 제공하는 `Packer` 튜토리얼 문서에 따르면, "`Packer`는 사용하기 쉽고 모든 유형의 머신 이미지 생성을 자동화하는 도구"라고 말하고 있다. 자동화된 스크립트를 활용하여, `Packer`에서 만든 이미지 내에서 필요한 소프트웨어들을 설치하고 구성할 수 있다. 크게 말해서 `IaC(Infrastructure As a Code)` 그 중 `Provisioning` 도구이고, 세분화해서 보면 이미지 빌드 도구라고 보면 된다.

`Packer`는 기본적으로 프로비저닝 시, 쉘 스크립트를 사용한다. 하지만 쉘 스크립트 뿐 아니라 `Ansible`과 같은 "설정 관리 도구"를 활용하여 보다 강력하고 쉬운 프로비저닝도 가능하다. 또한 `Vagrant`를 이용해서 이미지 후처리도 가능하다. 그리고 여러 이미지를 병렬적으로 생성하는 것도 가능하다. 

## Packer 설치 (로컬 Mac OS)

이번 절에서는 Mac OS 로컬 환경에서 `Packer`를 설치한다. 이 때 조금 더 쉬운 설치와 관리를 위해서 `pkenv`라는 것을 설치한다. 이를 위해서는 Mac OS에 `Home Brew` 패키지 매니저가 설치되어 있어야 한다.

터미널에 다음을 입력한다.

```bash
$ brew tap kwilczynski/homebrew-pkenv
$ brew install pkenv
```

그럼 설치가 된다. `pkenv` 설치가 제대로 되었는지 확인하기 위해서 터미널에 다음을 입력한다.

```bash
$ pkenv
pkenv 0.5.2
Usage: pkenv <command> [<options>]

Commands:
   install       Install a specific version of Packer
   use           Switch a version to use
   uninstall     Uninstall a specific version of Packer
   list          List all installed versions
   list-remote   List all installable versions
```

자 이제 `Packer`를 설치해보자. 현재 최신 버전인 `1.7.0`을 설치한다. 터미널에 다음을 입력한다.

```bash
$ pkenv install 1.7.0
```

이제 `Packer`가 잘 설치 되어 있는지 확인해보자. 터미널에 다음을 입력한다.

```bash
$ packer
Usage: packer [--version] [--help] <command> [<args>]

Available commands are:
    build           build image(s) from template
    console         creates a console for testing variable interpolation
    fix             fixes templates from old versions of packer
    fmt             Rewrites HCL2 config files to canonical format
    hcl2_upgrade    transform a JSON template into an HCL2 configuration
    init            Install missing plugins or upgrade plugins
    inspect         see components of a template
    validate        check that a template is valid
    version         Prints the Packer version
```

만약 안되어 있다면 다음을 확인해보자.

```bash
$ pkenv list
# 이렇게 안나온다면 다시 "pkenv install 1.7.0"을 입력한다.
* 1.7.0 (set by /usr/local/Cellar/pkenv/0.5.0/version)
```

## Packer 설치 (서버 AWS Linux)

이번에는 서버 환경에서 설치를 진행한다. 서버 환경은 `AWS Linux`에서 진행되었다. 터미널에 다음을 순서대로 입력한다.

```bash
$ sudo yum install -y yum-utils
$ sudo yum-config-manager --add-repo https://rpm.releases.hashicorp.com/AmazonLinux/hashicorp.repo
$ sudo yum -y install packer
```

그럼 설치가 된다. 터미널에 다음을 입력해서 제대로 설치가 되었는지 확인해보자.

```bash
$ packer
Usage: packer [--version] [--help] <command> [<args>]

Available commands are:
    build           build image(s) from template
    console         creates a console for testing variable interpolation
    fix             fixes templates from old versions of packer
    fmt             Rewrites HCL2 config files to canonical format
    hcl2_upgrade    transform a JSON template into an HCL2 configuration
    init            Install missing plugins or upgrade plugins
    inspect         see components of a template
    validate        check that a template is valid
    version         Prints the Packer version
```

## 참고

* Packer 튜토리얼, Getting Started : [https://learn.hashicorp.com/collections/packer/getting-started](https://learn.hashicorp.com/collections/packer/getting-started)
* Packer 튜토리얼, Install Packer : [https://learn.hashicorp.com/tutorials/packer/getting-started-install?in=packer/getting-started](https://learn.hashicorp.com/tutorials/packer/getting-started-install?in=packer/getting-started)
* pkenv 설치: [https://github.com/iamhsa/pkenv](https://github.com/iamhsa/pkenv)
