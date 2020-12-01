# InfluxDB 2.0 시작하기

![logo](../logo.png)

> InfluxDB 2.0 설치 및 시작하기 위한 설정에 대한 문서입니다.


## 요구 사항

요구사항은 다음과 같다.

- RedHat/CentOS 기반의 리눅스 운영체제가 설치된 머신 1대

나는 다음과 같은 환경에서 진행하였다.

- AWS EC2 Linux AMI(t2.micro)


## InfluxDB 2.0 설치하기

현재 안타깝지만, `InfluxDB 2.0`을 한 번에 설치할 수 있는 패키지가 준비되어 있지 않다.(내가 아는 한...) 그래서 바이너리 파일을 직접 설치할 것이다. 먼저 서버 내에서 적당한 디렉토리로 이동한 후, 다음을 순서대로 입력한다. 

```bash
# 현재 위치
$ pwd
/home/ec2-user

# 나의 경우 "apps"이란 디렉토리에 설치 파일들을 모아둔다. 그래서 apps 디렉토리로 이동한다.
$ cd apps

# 공식 문서에서 제공되는 influxdb 2.0 바이너리 파일이 압축된 tar 파일 다운로드 경로이다.
$ wget https://dl.influxdata.com/influxdb/releases/influxdb-2.0.0-rc.3_linux_amd64.tar.gz
--2020-11-03 01:59:49--  https://dl.influxdata.com/influxdb/releases/influxdb-2.0.0-rc.3_linux_amd64.tar.gz
Resolving dl.influxdata.com (dl.influxdata.com)... 13.249.43.98, 13.249.43.124, 13.249.43.59, ...
Connecting to dl.influxdata.com (dl.influxdata.com)|13.249.43.98|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 46135619 (44M) [application/x-gzip]
Saving to: ‘influxdb-2.0.0-rc.3_linux_amd64.tar.gz’

100%[===================================================================================================================>] 46,135,619  27.8MB/s   in 1.6s

2020-11-03 01:59:51 (27.8 MB/s) - ‘influxdb-2.0.0-rc.3_linux_amd64.tar.gz’ saved [46135619/46135619]

# 다운로드 확인
$ ll
합계 48660
-rw-rw-r-- 1 ec2-user ec2-user 49826885 10월 29 22:29 influxdb-2.0.0-rc.3_darwin_amd64.tar.gz

# 압축 파일 해제
$ tar zxvf influxdb-2.0.0-rc.3_linux_amd64.tar.gz
influxdb-2.0.0-rc.3_linux_amd64/LICENSE
influxdb-2.0.0-rc.3_linux_amd64/README.md
influxdb-2.0.0-rc.3_linux_amd64/influx
influxdb-2.0.0-rc.3_linux_amd64/influxd

# 실행 경로에 "influx", "influxd"를 옮긴다. 나의 실행 경로는 "/usr/local/bin"이다.
$ sudo cp influxdb-2.0.0-rc.3_linux_amd64/{influx,influxd} /usr/local/bin/

# influxdb 실행
$ influxd
2020-11-03T02:02:35.897826Z	info	Welcome to InfluxDB	{"log_id": "0QEyJTzG000", "version": "2.0.0-rc.3", "commit": "f46a3bd91e", "build_date": "2020-10-29T22:17:55Z"}
2020-11-03T02:02:35.902393Z	info	Resources opened	{"log_id": "0QEyJTzG000", "service": "bolt", "path": "/home/ec2-user/.influxdbv2/influxd.bolt"}
2020-11-03T02:02:35.903899Z	info	Migration "initial migration" started (up)	{"log_id": "0QEyJTzG000", "service": "migrations"}
2020-11-03T02:02:35.964158Z	info	Migration "initial migration" completed (up)	{"log_id": "0QEyJTzG000", "service": "migrations"}
...
```

그 후 브라우저에서, 현재 머신의 IP 주소 + 8086 포트에 접속하면, 다음 화면을 확인할 수 있다. 

![01](./01.png)


## InfluxDB 2.0 서비스로 만들기

## InfluxDB 2.0 설정하기 (UI)