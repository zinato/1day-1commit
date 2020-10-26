# 배쉬 쉘 스크립트에서 어떻게 하면 이전 명령어를 실행할 수 있을까요?

답은 다음과 같다.

```bash
$ !!
```

검증을 해보자. 먼저 "ls" 명령어를 입력한다. 내 맥묵에서 실행하면 다음과 같다.

```bash
$ ls
CODE_OF_CONDUCT.md      LICENSE                 _config.yml             index.js                zinato-note
COMMIT_MESSAGE.md       README.md               gurumee-note            package.json
```

이제 터미널에 다시 다음을 입력해보자.

```bash
$ !!
ls
CODE_OF_CONDUCT.md      LICENSE                 _config.yml             index.js                zinato-note
COMMIT_MESSAGE.md       README.md               gurumee-note            package.json
```

아예 똑같이는 안 나와도, 현재 디렉토리에 존재하는 모든 디렉토리를 출력하는 것을 확인할 수 있다.