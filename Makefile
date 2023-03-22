test:
	@echo "TEST shortener app"
	@go test ./...

run: 
	@echo "RUN shortener app"
	@go build -v -o url-shortener ./main.go
	@./url-shortener