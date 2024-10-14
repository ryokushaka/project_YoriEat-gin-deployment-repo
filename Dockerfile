FROM golang:1.22-alpine AS builder

# 필요한 패키지 설치
RUN apk add --no-cache git bash

# 작업 디렉토리 생성
WORKDIR /app

# go.mod와 go.sum 파일을 먼저 복사
COPY go.mod go.sum ./

# Go 모듈 다운로드 및 정리
RUN go mod download
RUN go mod tidy

# 나머지 파일들을 복사
COPY . .

# Swagger CLI 설치 및 문서 생성
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/main.go

# 빌드
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# 실행 스테이지
FROM alpine:latest

RUN apk --no-cache add ca-certificates postgresql-client

WORKDIR /root/

# 빌드 스테이지에서 빌드한 파일과 필요한 파일들만 복사
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# 컨테이너가 실행되면 실행할 명령어
CMD ["./main"]