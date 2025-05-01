run:
	@go test -bench=. -benchtime=1000000x


generate:
	@go generate ./structures