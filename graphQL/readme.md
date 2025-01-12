go bin을 등록하지 않아도 
바로 접근이 가느앟다
go bin 폴더면 


ex)
gqlgen init => go run github.com/99designs/gqlgen init

생기는 폴더 & 파일들
- gqlgen.yml: gqlgen 설정 파일입니다.
- graph/schema.graphqls: GraphQL 스키마 정의를 위한 파일입니다.
- graph/generated.go: gqlgen에 의해 자동 생성된 Go 코드입니다.
- graph/model/models_gen.go: 스키마에서 정의된 타입을 위한 Go 모델입니다.
- graph/resolver.go: 필드 리졸버를 구현하기 위한 기본 파일입니다.
- server.go: GraphQL 서버를 시작하는 기본 Go 파일입니다.

1. graph/schema.graphqls 파일을 열고 GraphQL 타입을 정의합니다. 예를 들어, 간단한 Todo 타입과 이를 조회할 수 있는 쿼리를 정의할 수 있습니다.

2. 스키마에 정의된 쿼리와 뮤테이션에 대한 실제 로직을 구현합니다. graph/resolver.go 파일에서 이를 수행할 수 있습니다. 예를 들어, Query 타입에 대한 리졸버를 구현할 수 있습니다.

추가 팁
- gqlgen generate 명령어를 사용하여 스키마가 변경될 때마다 Go 코드를 재생성할 수 있습니다.
- 스키마, 리졸버 로직, 또는 설정을 변경할 때마다 서버를 재시작해야 할 수도 있습니다.

실행법
