# 프로메테우스 시작하기

![logo](../logo.png)

> 책 "프로메테우스 오픈소스 모니터링 시스템"을 읽고 정리한 문서입니다.

## 요구 사항

이번 문서를 진행하기 위한 나의 환경은 다음과 같다.

- RedHat/CentOS 기반 운영체제
- 포트 9090에, http 접근 허용

레드햇 계열 리눅스와 `CentOS`는 이 문서가 `Prometheus`를 포함한 모니터링 시스템을 구축하는데 도움이 될 것이다. (아마 다른 운영체제도 비슷할 것이다.) 자 시작해보자.


## 프로메테우스 설치

`Promethues`는 [공식 홈페이지](https://prometheus.io/download/)에서 단일 파일로 설치할 수 있다. 터미널에 다음을 입력한다.

```bash
$ pwd
/home/ec2-user

# 나는 설치한 앱들을 이 곳에서 관리한다.
$ cd apps

# 프로메테우스 바이너리 파일이 들어 있는 압축 파일 설치
$ wget https://github.com/prometheus/prometheus/releases/download/v2.22.0/prometheus-2.22.0.linux-amd64.tar.gz

# 압축 파일 해제
$ tar zxvf prometheus-2.22.0.linux-amd64.tar.gz

# 프로메테우스 디렉토리 간소화
$ mv prometheus-2.22.0.linux-amd64 prometheus

# 프로젝트 디렉토리 이동
$ cd prometheus

# 프로메테우스 실행
$ ./prometheus
```

이제 `서버 IP 주소:포트`로 이동하면, 다음 화면을 확인할 수 있다.

![01](./01.png)

이제 프로메테우스를 서비스로 등록해보자. 먼저 `vim` 에디터로 다음 파일을 만든다.

```bash
$ sudo vim /etc/systemd/system/prometheus.service
```

이제 "i"를 누르고 다음을 입력한다.

```service
[Unit]
Description=Prometheus Server

[Service]
Restart=on-failure

#Change this line if you download the
#Prometheus on different path user
ExecStart=/home/ec2-user/apps/prometheus/prometheus \
  --config.file=/home/ec2-user/apps/prometheus/prometheus.yml \
  --storage.tsdb.path=/home/ec2-user/apps/prometheus/data

[Install]
WantedBy=multi-user.target
```

이제 "esc"를 누르고 ":wq!"를 입력한다. 그 후 터미널에 다음을 입력한다.

```bash
# 서비스 데몬 리로딩
$ sudo systemctl daemon-reload

# prometheus 서비스 시작
$ sudo systemctl start prometheus

# prometheus 서비스 상태 확인
$ sudo systemctl status prometheus
● prometheus.service - Prometheus Server
   Loaded: loaded (/etc/systemd/system/prometheus.service; disabled; vendor preset: disabled)
   Active: active (running) since 수 2020-11-04 09:04:14 UTC; 4s ago
...
```

다시 `서버 IP 주소:포트`로 이동하면, 다음 화면을 확인할 수 있다.

![01](./01.png)

`Prometheus` 서버는 지정한 `prometheus.yml`을 참조하여, 기본 값들을 설정하고 서버를 실행시킨다. 기본 설정값은 다음과 같다.

```yml
# my global config
global:
  scrape_interval:     15s 
  evaluation_interval: 15s 

# Alertmanager configuration
alerting:
  alertmanagers:
  - static_configs:
    - targets:
      # - alertmanager:9093

rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"


scrape_configs:
  - job_name: 'prometheus'
    static_configs:
    - targets: ['localhost:9090']
```

이는 15초에 1번씩 프로메테우스의 상태 데이터를 수집한다는 뜻과 같다.

![02](./02.png)

실제로 이를 확인하기 위헤서 위 그림과 같이 웹 브라우저 상에서 Status > Targets를 클릭해보자.

![03](./03.png)

그럼 위의 그림과 같은 화면을 확인할 수 있다. 이번에는 "Endpoint"를 클릭해보자. (안된다면, "IP 주소:port/metrics"로 이동하면 된다.)

![04](./04.png)

그럼 위의 화면을 확인할 수 있다. 이제 다시 메인페이지로 돌아와서 입력 창에 "up"을 쳐보자.

![05](./05.png)

이런 식으로 수식 브라우저를 통해서 몇 개의 메트릭들의 값을 확인할 수 있다. 또한, "Table" 오른쪽에 "Graph"를 눌러보면, 해당 메트릭에 대한 그래프를 확인할 수 있다.

`Prometheus`는 자체적으로 UI를 제공하긴 하지만, `Grafana`
를 연동하면, 더 멋진 대시보드를 꾸밀 수 있다. 그래서 애는 이 정도로만 알고 가자.


## 노드 익스포터 설치 및 프로메테우스 연동

노드 익스포터를 통해서 리눅스, 유닉스 계열의 시스템에서 커널이나 머신 레벨의 메트릭을 수집할 수 있다. 이를 설치하고 프로메테우스와 연동하여, 메트릭들을 확인해보자.

역시, [공식 홈페이지](https://prometheus.io/download/)에서 다운로드할 수 있다. 터미널에 다음을 입력한다.

```bash
# 다시 apps로 이동
$ cd ~/apps

# 노드 익스포터 바이너리 파일이 있는 압축 파일 설치
$ wget https://github.com/prometheus/node_exporter/releases/download/v1.0.1/node_exporter-1.0.1.linux-amd64.tar.gz

# 압축 파일 해제
$ tar zxvf node_exporter-1.0.1.linux-amd64.tar.gz

# 좀 더 쉽게 접근을 위해 디렉토리 이름 변경
$ mv node_exporter-1.0.1.linux-amd64 node_exporter

# 노드 익스포터 실행
$ ./node_exporter/node_exporter
...
```

실행되는 것을 확인하고 다시 종료하자. 서비스로 등록하자. `vim` 에디터로 파일을 생성하자.

```bash
$ sudo vim /etc/systemd/system/node-exporter.service
```

이제 아까 전과 같은 요령으로 다음을 입력하자.

```
Description=Node Exporter

[Service]
Restart=on-failure

#Change this line if you download the
#Prometheus on different path user
ExecStart=/home/ec2-user/apps/node_exporter/node_exporter

[Install]
WantedBy=multi-user.target
```

그 후 터미널에 다음을 입력한다.

```bash
# 서비스 데몬 리로드
$ sudo systemctl daemon-reload

# 노드 익스포터 서비스 시작
$ sudo systemctl start node-exporter

# 노드 익스포터 서비스 상태 확인
$ sudo systemctl status node-exporter
● node-exporter.service
   Loaded: loaded (/etc/systemd/system/node-exporter.service; disabled; vendor preset: disabled)
   Active: active (running) since 수 2020-11-04 09:38:13 UTC; 9s ago
...
```

이제 프로메테우스와 노드 익스포터를 연동한다. `apps/prometheus/promethues.yml`을 `vim` 에디터로 연다.

```bash
$ cd ~/apps

$ vim prometheus/prometheus.yml
```

그 후 맨 밑에 다음을 추가한다.

```yml
# 이전과 동일
  - job_name: 'node_exporter'
    static_configs:
    - targets:
        - localhost:9100
```

그 다음 프로메테우스 서비스를 다시 시작한다.

```bash
$ sudo systemctl restart prometheus
```

이제 다시 프로메테우스 UI에 접속한 후, "UP"을 입력해보자.

![06](./06.png)

그럼 위 화면처럼, 인스턴스가 하나 추가된 것을 확인할 수 있다. 여러 개의 잡 중 하나를 선택하려면 이런 식으로 칠 수 있다.

```
up{job="node_exporter"}
```


## 알림 만들어서 내보내기

이제 알림을 만들어보자. 알림을 만들기 위해서는 크게 2가지가 필요하다.

1. prometheus.yml에 알림 파일 설정하기
2. 알림 규칙을 설정한느 rules.yml 설정하기

먼저 프로메테우스가 설치된 경로로 이동하여 `prometheus.yml`을 다음과 같이 수정한다.

```yml
global:
  scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.

alerting:
  alertmanagers:
  - static_configs:
    - targets:
      # - alertmanager:9093

rule_files:
  - "rules.yml"

scrape_configs:
  - job_name: 'prometheus'

    static_configs:
    - targets: ['localhost:9090']

  - job_name: 'node_exporter'
    static_configs:
    - targets:
        - localhost:9100
```

`rule_files`에 알림을 설정하는 `rules.yml` 경로가 들어있다. 이는 현재 경로에, `rules.yml`을 참조한다는 뜻이다. 그 후 같은 경로에, `rules.yml`을 다음과 같이 생성한다.

```yml
groups:
    - name: example
      rules:
          - alert: InstanceDown
            expr: up == 0
            for: 1m
```

이 알림은 프로메테우스가 알고 있는 시스템 중 1분 동안 꺼져 있을 때, 알림을 보낸다. 프로메테우스를 다시 시작하고, 알림을 보내기 위해서 노드 익스포터를 중지한다.

```bash
# 프로메테우스 재시작
$ sudo systemctl restart prometheus

# 노드 익스포터 중지
$ sudo systemctl stop node-exporter
```

이제 프로메테우스 UI를 접속한다.

![07](./07.png)

그 후 "Alerts"를 클릭한다.

![08](./08.png)

알림 탭에 들어가면, 설정된 알림을 확인할 수 있다. 현재는 알림이 Pending된 상태이다. 조금만 기다려보자.

![09](./09.png)

알림이 발생하였다. 이제 이 알림을 "알림 매니저"를 통해서 슬랙으로 알림을 보내보자. 먼저 알림 매니저를 설치한다. 터미널에 다음을 입력한다.

```bash
# 앱 설치 경로로 이동
$ cd ~/apps

# 알림 매니저 설치
$ wget https://github.com/prometheus/alertmanager/releases/download/v0.21.0/alertmanager-0.21.0.linux-amd64.tar.gz

# 압축 파일 해제
$ tar zxvf alertmanager-0.21.0.linux-amd64.tar.gz

# 압축 파일 삭제
$ rm alertmanager-0.21.0.linux-amd64.tar.gz

# 이름 변경
$ mv alertmanager-0.21.0.linux-amd64 alertmanager

# 알림 매니저 실행
$ ./alertmanager/alertmanager
...
```

실행되는 것을 확인했으면, 알림 매니저를 서비스로 만든다. 터미널에 다음을 입력하여, vim 에디터를 실행시킨다.

```bash
$ sudo vi /etc/systemd/system/alertmanager.service
```

그 후 다음을 입력한다.

```
[Unit]
Description=Alert Manager
Wants=network-online.target
After=network-online.target

[Service]
ExecStart=/home/ec2-user/apps/alertmanager/alertmanager --config.file=/home/ec2-user/apps/alertmanager/alertmanager.yml --storage.path=/var/lib/alertmanager
Restart=on-failure
StartLimitBurst=2
StartLimitInterval=30

[Install]
WantedBy=multi-user.target
```

`:wq!`를 눌러 저장하고, 터미널에 다음을 입력한다.

```bash
# 데몬 리로드
$ sudo systemctl daemon-reload

# 알림 매니저 실행
$ sudo systemctl start alertmanager

# 알림 매니저 상태 확인
$ sudo systemctl status alertmanager
● alertmanager.service - Alert Manager
   Loaded: loaded (/etc/systemd/system/alertmanager.service; enabled; vendor preset: disabled)
   Active: active (running) since 수 2020-11-18 23:40:42 UTC; 9s ago
 Main PID: 6631 (alertmanager)
   CGroup: /system.slice/alertmanager.service
           └─6631 /home/ec2-user/apps/alertmanager/alertmanager --config.file=/home/ec2-user/apps/alertmanager/alertmanager.yml --storage.path=/var/lib/alertmanager
...
```