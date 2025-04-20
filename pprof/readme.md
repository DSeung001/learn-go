# Go Profiling Guide

## 1. 실행 방법
프로파일링을 실행하려면 다음 명령어를 사용하세요.

### 1.1 기본 실행 (파일 디스크립터 프로파일링)
```sh
go run main.go
```
이 명령을 실행하면 `fd.pprof` 파일이 생성됩니다.

또는 빌드 후 실행하려면:
```sh
go build -o myapp main.go
./myapp
```

## 2. 프로파일링 데이터 분석
생성된 `fd.pprof` 파일을 분석하려면 `go tool pprof` 명령어를 사용합니다.

### 2.1 프로파일링 데이터 요약 확인
```sh
go tool pprof fd.pprof
```
이후 pprof 인터페이스에서 다음 명령어를 사용할 수 있습니다:
- `top` : 가장 많은 연산을 차지한 함수 확인
- `list main.main` : 특정 함수별 실행 정보 확인
- `web` : 호출 그래프 시각화 (Graphviz 필요)

### 2.2 호출 그래프 보기
```sh
go tool pprof -http=:8080 fd.pprof
```
브라우저에서 `http://localhost:8080`에 접속하면 시각적으로 분석할 수 있습니다.

## 3. CPU 프로파일링 추가 실행 (선택 사항)
CPU 프로파일링을 활성화하려면 `main.go`를 수정하여 `runtime/pprof`를 추가해야 합니다.

### 3.1 CPU 프로파일링 실행
```sh
go run main.go
```
이제 `cpu.pprof` 파일이 생성됩니다.

### 3.2 CPU 프로파일링 데이터 분석
```sh
go tool pprof cpu.pprof
```
마찬가지로 `top`, `list`, `web` 등의 명령어를 활용해 분석할 수 있습니다.

## 4. 필요 패키지 설치
`Graphviz`를 설치하면 `web` 명령어를 사용하여 호출 그래프를 볼 수 있습니다.

### 4.1 Graphviz 설치 (선택 사항)
#### MacOS
```sh
brew install graphviz
```
#### Ubuntu/Linux
```sh
sudo apt install graphviz
```
#### Windows (scoop 사용)
```sh
scoop install graphviz
```

## 5. 참고 사항
- `fd.pprof`, `cpu.pprof` 같은 프로파일링 파일은 실행할 때마다 갱신됩니다.
- `pprof` 데이터를 분석하여 성능 최적화에 활용하세요.

이제 Go 프로파일링을 실행하고 분석하는 방법을 알게 되었습니다! 🚀

