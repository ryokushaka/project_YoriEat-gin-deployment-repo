FROM golang:1.22-alpine AS builder

# 필요한 패키지 설치
RUN apk add --no-cache git bash

# 작업 디렉토리 생성
WORKDIR /app

# 현재 디렉토리의 파일들을 컨테이너의 /app 디렉토리로 복사
COPY . .

# Go 모듈 다운로드
RUN go mod download

# Swagger CLI 설치 및 문서 생성
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/main.go

# 빌드
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# 실행 스테이지
FROM alpine:latest

RUN apk --no-cache add ca-certificates postgresql-client

WORKDIR /root/

# 빌드 스테이지에서 빌드한 파일과 필요한 파일들만 복사
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# 환경 변수 설정
ARG DB_HOST
ARG DB_PORT
ARG DB_NAME
ARG DB_USER
ARG DB_SSLMODE
ARG DB_PASSWORD
ARG AWS_REGION
ARG ACCESS_TOKEN_EXPIRY_HOUR
ARG REFRESH_TOKEN_EXPIRY_HOUR
ARG ACCESS_TOKEN_SECRET
ARG REFRESH_TOKEN_SECRET

ENV DB_HOST=$DB_HOST \
    DB_PORT=$DB_PORT \
    DB_NAME=$DB_NAME \
    DB_USER=$DB_USER \
    DB_SSLMODE=$DB_SSLMODE \
    DB_PASSWORD=$DB_PASSWORD \
    AWS_REGION=$AWS_REGION \
    ACCESS_TOKEN_EXPIRY_HOUR=$ACCESS_TOKEN_EXPIRY_HOUR \
    REFRESH_TOKEN_EXPIRY_HOUR=$REFRESH_TOKEN_EXPIRY_HOUR \
    ACCESS_TOKEN_SECRET=$ACCESS_TOKEN_SECRET \
    REFRESH_TOKEN_SECRET=$REFRESH_TOKEN_SECRET

# 컨테이너가 실행되면 실행할 명령어
CMD ["/app/main"]