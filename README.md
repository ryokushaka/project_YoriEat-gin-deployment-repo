# YoriEat 백엔드 서버 리포지토리

![YoriEat Architecture](https://outcomeschool.com/_next/image?url=%2Fstatic%2Fimages%2Fblog%2Fgo-backend-arch-diagram.png&w=1920&q=75)

[아키텍처 참조 링크](https://outcomeschool.com/blog/go-backend-clean-architecture)

## 프로젝트 개요
YoriEat 서비스의 백엔드 서버 리포지토리입니다. 이 프로젝트는 `gin gonic`, `MongoDB`, `Docker`, `Docker Compose`를 사용하여 구성되었습니다.

## 프로젝트 구조

```markdown
.
├── Dockerfile
├── README.md
├── api
│   ├── controller
│   │   ├── login_controller.go
│   │   ├── profile_controller.go
│   │   ├── profile_controller_test.go
│   │   ├── refresh_token_controller.go
│   │   ├── signup_controller.go
│   │   └── task_controller.go
│   ├── middleware
│   │   └── jwt_auth_middleware.go
│   └── route
│       ├── login_route.go
│       ├── profile_route.go
│       ├── refresh_token_route.go
│       ├── route.go
│       ├── signup_route.go
│       └── task_route.go
├── bootstrap
│   ├── app.go
│   ├── database.go
│   └── env.go
├── cmd
│   └── main.go
├── docker-compose.yaml
├── domain
│   ├── error_response.go
│   ├── jwt_custom.go
│   ├── login.go
│   ├── profile.go
│   ├── refresh_token.go
│   ├── signup.go
│   ├── success_response.go
│   ├── task.go
│   └── user.go
├── go.mod
├── go.sum
├── internal
│   └── tokenutil
│       └── tokenutil.go
├── mongo
│   └── mongo.go
├── repository
│   ├── task_repository.go
│   ├── user_repository.go
│   └── user_repository_test.go
└── usecase
    ├── login_usecase.go
    ├── profile_usecase.go
    ├── refresh_token_usecase.go
    ├── signup_usecase.go
    ├── task_usecase.go
    └── task_usecase_test.go
```

## 테스트 실행 방법

### 모든 테스트 실행
```sh
go test ./...
```

## 모의(Mock) 코드 생성 방법

이 프로젝트에서는 use-case, repository, database의 테스트를 위해 모의 코드를 생성해야 합니다.

### use-case 및 repository 모의 코드 생성
```sh
mockery --dir=domain --output=domain/mocks --outpkg=mocks --all
```

### database 모의 코드 생성
```sh
mockery --dir=mongo --output=mongo/mocks --outpkg=mocks --all
```

해당 use-case, repository, 또는 database의 인터페이스에 변경 사항이 있을 때마다 테스트를 위해 모의 코드를 재생성해야 합니다.

## 추가 정보

### 환경 변수 설정
프로젝트의 환경 변수를 설정하려면 `bootstrap/env.go` 파일을 참고하세요.

### 데이터베이스 초기화
MongoDB 데이터베이스 초기화는 `bootstrap/database.go` 파일을 통해 이루어집니다.

### 애플리케이션 시작
애플리케이션을 시작하려면 `cmd/main.go` 파일을 사용하세요.

### Docker 및 Docker Compose 사용
이 프로젝트는 Docker 및 Docker Compose를 사용하여 쉽게 배포할 수 있습니다. `Dockerfile`과 `docker-compose.yaml` 파일을 참고하여 컨테이너를 빌드하고 실행하세요.