FROM golang:1.19-alpine

# 필요한 패키지 설치
RUN apk add --no-cache git

# 작업 디렉토리 생성
RUN mkdir /app

# 현재 디렉토리의 파일들을 컨테이너의 /app 디렉토리로 복사
ADD . /app

# 작업 디렉토리 설정
WORKDIR /app

# Go 모듈 다운로드
RUN go mod download

# Swagger CLI 설치
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Swagger 문서 생성
RUN swag init -g cmd/main.go

# 빌드
RUN go build -o main cmd/main.go

# 컨테이너가 실행되면 실행할 명령어
CMD ["/app/main"]