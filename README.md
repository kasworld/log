# 커스텀 로그 레벨을 사용하는 로그 생성기와 보조 패키지들 


## 특징 

로그 레벨의 종류, 이름을 커스터마이징 할 수 있다. (level.data를 입력으로 genlog)

사용 프로그램에서 각 로그 레벨을 독립적으로 켜고 끌수 있다. ( 로그레벨간 우선 순위같은 것이 없다.)

한개의 로그를 여러 곳(logdestination)으로 동시에 보낼 수 있다. (logdestinationgroup)

로그라인의 헤더를 커스터 마이징 할수있다. (logflags, logflagi)

## 사용법 

build.sh : install genlog

genlog :  loglevel.data 를 사용해서 log package를 생성합니다. 

loglevel data에는 사용할 level을 적으면 됩니다. (basicloglevel.data 참고)

다른 package들은 만들어진 log package가 사용할 라이브러리 입니다. 

## 패키지 설명 

genlog : 커스텀 로그 생성기 

    입력 받은 로그 레벨 데이터 를 사용해서 로그 생성기를 만든다. 
    예제에 basiclevel.data 를 사용하여 만들어진 basiclog 가 있다. 
    생성한 파일이름은 _gen.go로 끝난다. 

logdestination_file : 생성된 로그의 목적지가 file 인 경우 사용 

logdestination_stdio : 생성된 로그의 목적지가 표준 출력 , 표준 에러 인 경우 사용 

logdestinationi : 만들어진 로그 데이터가 출력/저장 될 목적지 interface 

logdestinationgroup : 로그 목적지의 묶음, 각 로그 레벨 에 대응 시켜 하나의 로그가 여러 곳으로 동시에 보내질수 있다. 

logflagi : 각 로그 라인의 header 를 위한 interface 

logflags : 기본으로 사용되는 로그 라인의 header 용 flag들 


### old document 
korean discription

http://kasw.blogspot.kr/2015/02/go-python-like-log.html
