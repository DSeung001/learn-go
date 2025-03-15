```
cd test
go test -bench=BenchmarkGrpcRegisterParallel -benchmem -benchtime=30s ./
go test -bench=BenchmarkRestRegisterParallel -benchmem -benchtime=30s ./
```