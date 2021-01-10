
# Github Action을 통한 CI/CD 전사 도입

- 기간 : 2019.09 ~ 10 : Github Survey 및 CI/CD 구축
- 사용 기술 : Github action. Git
- 각자 개별화 테스트, 배포 하는 것을 서버부터 먼저 도입 후 클라이언트(Flutter -> Web) 등으로 점차적으로 도입
- 메뉴얼 문서 작성
- 각 브랜치 별 액션 정의
- 슬랙, 이메일 연결 기능 공유

# 토큰 스마트 컨트랙트 작성 및 블록체인 서비스 설계

- 기간 : 2020.03 ~ 05
- 사용 기술 : Solidity, Ethereum, Javascript 
- 메인 토큰과 팬덤 토큰의 스마트 컨트랙트 작성
- OpenZepplin을 통한 ERC20 표준 컨트랙트 구현
- 테스트 코드 구현 (로컬, 테스트 서버)
- 실제 Xpeare 서비스를 위한 토큰 및 루니버스 API를 통한 자체 Asset API 설계 / DB 설계
- 보안 가이드 라인 설계

# Xp-soultion-Asset API - Proxy Server 구축, 통합 테스트 코드 작성 

- 기간 : 2020.05 ~ 06
- 사용 기술 : Java8, GCP AppEngine, GCP Compute Engine, JUnit, Luniverse API, nodejs proxy server opensource, javascript, mocha, chai
- 루니버스 API 를 사용한 블록체인 토큰 지급을 위한 API 개발
- GCP AppEngine은 고정 IP를 제공하지 않아 Compute Engine을 활용하여 Proxy Server를 구축하여 루니버스의
WhiteIP 등록하여 연결 (with Chakun, Now)
- 개발 기간이 부족하여 통합 테스트 코드가 없어 mocha, chai 를 통해 javascript 테스트 코드를 2일에 걸쳐 작성 및 테스트 정상적으로 완료. 

# 2020 데이터 바우처 사업

- 2020.05 ~ 12월 (아직 진행 중 마무리 단계)
- 실무 책임자로 참여 
- 수호와 데이터 바우처 가공 업체로 지원 '적격'
- 코로나로 인해 발표질의응답을 비디오 녹화로 대체
- 2020.06.08 발표자료 제출 및 선정 완료  
- 수호와 5회에 걸쳐 사업 진행 (09.21~11.16일)
- 2020.11.19 이행점검 진행 및 비대면 미팅 진행 


# 쿠버네티스 학습 및 테스트 진행 

- 기간 : 2020.07~ 08월
- 사용기술 : GCP, Kubernetes, Docker, Locust
- 쿠버네티스 학습 및 사내 서버 구축을 위한 테스트 진행
- 로드 밸런싱, 롤링 업데이트등 테스트, 동접 성능 테스트

# 개톡 문화 도입 (개발자 모임 문화)

- 기간 : 2020.08 ~
- 업무의 효율을 높이고 개발자들끼리 문화를 공유하고 자신의 감정을 공유할 수 있는 개발자 문화 도입, 컨퍼런스보 약간 가벼운 개념으로 시작
- 블로그 글을 작성하고 회사 블로그로 선정되면 소정의 혜택 추가
- FreeTalk 시간에 공유하고 싶은 기술 내용 공유
- 코드 리뷰로 발전해 나갈 예정

# 회고 문화 도입

- 기간 : 2020.08 ~ 
- 하나의 스크럼 또는 한달 기준으로 정기적으로 회고를 하는 것을 갖기 위해 진행자 및 회고에 대한 설명 공유
- 팀 별로 회고를 진행
- action plan 도출 및 대표님의 문제점을 논의할 수 있는 시간을 정기적으로 가짐.
- 자신들의 스크럼을 다시 한번 돌아볼 수 있는 계기를 통해 장점과 잘 못했던 점과 앞으로 개선할 점을 파악할 수 있음
- 하나의 프로젝트를 완료 했다는 끝맺음의 마음을 가질 수 있음
 
 # Hazelcast 성능 테스트 진행 =>  Ignite 로 변경 
 - 기간 : 2020.09.25 ~ (ProjectD로 인해 홀딩)
 - 사용기술 : Java, Hazelcast => Ignite, GCP
 - 추천 엔진에서 Serving을 위해 Ignite 사용 가능한지에 대한 테스트
 
 # Project D 단비 프로젝트 진행
 - 기간 : 2020.10.21 ~ 2020.11.30 
 - 사용기술 : GCP ComputeEngine, CloudSQL, Java, Javascript, FireStore 
 - 내용 :
 - Cron 프로그램 개발 
    - Java를 사용하여 Myfan 관련 데이터의 ETL 작업을 통해 일별, 주별, 월별, 데이터를 합산하어 매월 1일, 매주 월요일, 매일 특정시간에
    동작하는 프로그램 개발
 - API 개발 
    - 기존 Javascript로 작성된 API에 Report 관련 API 2개 추가 (book 기준 , author 기준)
 - 특이사항 
    - 외할머니 장례식 5일 참여. 
    - 기존 DB 설계가 잘못되어 있어 프로그램 수정과 DB 타입 변경등이 많아 시간이 예상보다 오래 걸림
    
 # KTR 인증 테스트 진행 
 
 - 기간 : 2020. 11.12 ~ 2020.11.13
 - 사용기술 : 쿠버네티스 , Locust, GCP ComputeEngine, LoadBalancer 
 - 내용 : 
  - KTR 인증 테스트를 만족하는 서버 구축 및 테스트 결과 도출 
      - 트렌드 서버 1000 TPS 
      - 트렌드 서버 800ms 이내
  - GCP ComputeEngine 과 LoadBalancer를 튜닝하여 Locust를 통해 테스트를 진행했을 때 1000 TPS를 만족하여 테스트 통과 및 결과 만족 
 
 # 오벤터스 관련 Google RecommendationAI 결과 프로그램 개발
 
 - 기간 : 2020.12.01 ~ 2020.12.03
 - 사용 기술 : Google RecommendationAI, Java 11
 - 내용 : 
    - 조아라 UserId 약 19만개를 가지고 Google RecommendationAI를 활용하여 트레이닝 된 모델 (joara-s-read30-7d-user-to-item-model)에서 추천결과를 받아
    ETL 처리를 하여 {"UserID ": "추천결과"} JSON 파일을 생성 

 # Myfan ReportMaker Production 배포 작업 및 Push Cron 추가 개발 
 
 - 기간 : 2020.12.07 ~ 2020.12.30(배포 예정)
 - 사용 기술 : GCP ComputeEngine, CloudSQL, Java, Javascript, FireStore 
 - 내용 : 
    - 단비프로젝트에 이어 개발이 미구현된 ReportMaker 완성과 API 추가 테스트, Push 크론 프로그램 추가 개발 
    - Prd 환경 구축  
    - Dev 테스트, Stg 테스트 추가 진행
    - 배포 모니터링 
 
