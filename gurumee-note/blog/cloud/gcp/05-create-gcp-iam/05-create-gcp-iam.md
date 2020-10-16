# GCP 프로젝트에 IAM 키 생성하기

![로고](../logo.png)

> 실제 GCP를 이용하면서 정리한 내용입니다. 이 문서는 GCP에서 만든 프로젝트에 서비스 계정 외, 다른 사용자가 접속할 수 있도록 IAM 키를 생성하는 방법에 대하여 다루고 있습니다.


## 요구 사항

이 문서를 진행하기 위해서는 적어도 `GCP` 내에서 하나 이상의 프로젝트가 만들어져야 한다. 아직 안 만들었다면, 다음을 참고하라.

* [GCP 가입하기](https://gurumee92.github.io/2020/09/gcp-%EA%B0%80%EC%9E%85%ED%95%98%EA%B8%B0/)
* [GCE 인스턴스 생성 및 설정하기](https://gurumee92.github.io/2020/09/gce-%EC%9D%B8%EC%8A%A4%ED%84%B4%EC%8A%A4-%EC%83%9D%EC%84%B1-%EB%B0%8F-%EC%84%A4%EC%A0%95%ED%95%98%EA%B8%B0/)

자 `GCP` 가입과, 프로젝트 생성을 완료했다면, 시작해보자.


## 서비스 계정과 IAM 키

먼저 `IAM 키`를 알기 전에 `서비스 계정`이 무엇인지 알아야 한다. 공식 문서에 따르면, 정의는 다음과 같다.

    서비스 계정은 사용자가 아닌 애플리케이션 또는 가상 머신(VM) 인스턴스에서 사용하는 특별한 유형의 계정입니다. 
    애플리케이션은 서비스 계정 자체 또는 G Suite로 승인되거나, 도메인 전체 위임을 통해 Cloud ID 사용자로 승인된 
    API 호출을 수행하기 위해 서비스 계정을 사용합니다.

쉽게 말해서, 프로젝트를 만든 "사용자"가 아니라, 사용자가 특별히 사용자의 클라우드 리소스를 접근을 허용한 "계정"이라고 보면 된다. 그리고 이 계정들은 비밀번호로 접근하는 것이 아니라, 일반적으로 `IAM 키`를 가지고 접근하게 된다.

즉, `IAM 키`란 `Identity and Access 키`의 약자로, 서버 접근 권한을 가진 키로 이해하면 된다.


## GCP 웹 콘솔로 IAM 키 관리하기

`IAM 키` 생성은 여러가지 방법이 있다. 이 문서에서는 먼저 웹 브라우저를 통해서 UI로 `IAM 키`를 관리하는 것을 다룬다. 먼저 [이 곳](https://console.cloud.google.com/projectselector2/iam-admin/serviceaccounts)을 클릭하여 이동한다.

![01](./01.png)

그럼 위 화면처럼 프로젝트 목록을 확인할 수 있는 화면이 나온다. 원하는 프로젝트를 클릭한다. 

![02](./02.png)

나는 설명을 위해 가계정을 만들어놨는데, 이미 만들어진 서비스 계정이 있을 것이다. 이를 클릭한다. 

> 참고!
> 
> 저는 문서 작성 후 계정을 삭제해서 해당 정보를 제거하였습니다. 이런 키들은 절대 노출이 되면 안됩니다!! 만약 계정을 만들고 싶다면, "+ 서비스 계정 만들기" 버튼을 클릭하세요. 계정 권한 설정 등은 공식 문서를 확인하고 작업해주어야 합니다.

![03](./03.png)

그리고 "키 추가" 버튼을 누른다. 

![04](./04.png)

그러면 위 화면처럼 드롭다운 UI로 "새 키 만들기", "기존 키 업로드" 가 나오는데 "새 키 만들기"를 클릭한다. 그럼 다음과 같은 팝업이 뜬다.

![05](./05.png)

이제 "만들기" 버튼을 클릭한다. 그럼 다음 화면처럼 키가 생성된 것을 확인할 수 있다.

![06](./06.png)

또한, 다음 삭제 버튼을 누르면, 만든 키를 삭제할 수 있다.

![07](./07.png)


## gcloud 도구로 IAM 키 관리하기

> 참고!
> 
> 이 절은 로컬 머신에 gcloud 도구가 설치되어 있어야 진행이 가능합니다.

이번에는 `gcloud`도구로 더 쉽게 `IAM 키`를 생성하고 삭제해보자. 먼저 키를 생성해본다. 터미널을 열고 다음을 입력한다.

```bash
# gcloud iam service-accounts keys create ~/key.json --iam-account <서비스 계정 이름>@<프로젝트 id>.iam.gserviceaccount.com
$ gcloud iam service-accounts keys create ~/key.json --iam-account  test-983@geerio.iam.gserviceaccount.com
```

그럼 local 머신에 `key.json`이 설치된 것을 확인할 수 있다. 또한, UI에서도 확인 가능하다.

![08](./08.png)

아까와 다른 키가 생성된 것을 확인할 수 있다. 이제 키를 삭제해보자.

```bash
# gcloud iam service-accounts keys delete <IAM 키-id> --iam-account <서비스 계정 이름>@<프로젝트 id>.iam.gserviceaccount.com
$ gcloud iam service-accounts keys delete ee928fa0c8957000117235236fa2e9d1e3f157d0 --iam-account test-983@geerio.iam.gserviceaccount.com
ou are about to delete key [ee928fa0c8957000117235236fa2e9d1e3f157d0]
 for service account [test-983@geerio.iam.gserviceaccount.com].

Do you want to continue (Y/n)?  y # y를 입력하셔야 합니다.

deleted key [ee928fa0c8957000117235236fa2e9d1e3f157d0] for service account [test-983@geerio.iam.gserviceaccount.com]
```

끝이다.


> 참고!
> 공개 키 업로드도 가능합니다만, 일단 저는 필요 없기 때문에 이는 따로 설명하지 않겠습니다. 원하신다면, 아래 참고 문서 절의 "서비스 계정 키 생성 및 관리" 링크를 참고하세요.


## 참고 문서

* [GCP 공식 문서 - 서비스 계정](https://cloud.google.com/iam/docs/service-accounts?hl=ko)
* [GCP 공식 문서 - 서비스 계정 키 생성 및 관리](https://cloud.google.com/iam/docs/creating-managing-service-account-keys?hl=ko#iam-service-account-keys-create-console)