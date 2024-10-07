FROM golang:1.22-alpine

# 필요한 패키지 설치
RUN apk add --no-cache git bash postgresql-client

# 작업 디렉토리 생성
WORKDIR /app

# 현재 디렉토리의 파일들을 컨테이너의 /app 디렉토리로 복사
COPY . .

# Go 모듈 다운로드
RUN go mod download

# Swagger CLI 설치
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Swagger 문서 생성
RUN swag init -g cmd/main.go

# 빌드
RUN go build -o main cmd/main.go

# 빌드 시 전달받을 ARG 정의
ARG DB_HOST
ARG DB_PORT
ARG DB_NAME
ARG DB_USER
ARG DB_SSLMODE
ARG DB_PASSWORD
ARG AWS_REGION

# 환경 변수 설정
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_NAME=${DB_NAME}
ENV DB_USER=${DB_USER}
ENV DB_SSLMODE=${DB_SSLMODE}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV AWS_REGION=${AWS_REGION}

# 컨테이너가 실행되면 실행할 명령어
CMD ["/app/main"]