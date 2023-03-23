# Go URL Shortener

Go URL Shortener is a very simple web application built with Go that allows users to shorten long URLs.

## Features
- Shorten long URLs to more manageable lengths
- Redirect shortened URLs to their original destination

## Tech
- Go
- Redis

## Package/Frameworks
- Gin

## Getting Started
To get started with Go URL Shortener, you will need to have Go installed on your computer. You can download Go from the official website: https://golang.org/dl/

Once you have Go installed, follow these steps:
- Clone the repository to your local machine using
```bash
git clone https://github.com/Alitindrawan24/go-url-shortener.git.
```
- Navigate to the project directory using the command line.
- Create .env file from .env.example and setup .env with your environment configuration
```bash
cp .env.example .env
```
- Create .env.testing file from .env.example  and setup .env.testing with your testing environment configuration (change APP_ENV into "testing")
```bash
cp .env.example .env.testing
```
- Run the app using
```bash
make run
```
Or using
```bash
go run main.go
```
- Open your web browser and navigate to http://localhost:8080 to access the application.

## Endpoints
- ```GET/```  get welcome message
- ```POST/``` create short url using parameter ```initial_url``` and ```user_id``` in body
- ```GET/:short_url``` redirect from short url into original url

## Test
You can test the app using command :
```bash
make test
```
Or using
```bash
go test ./...
```

## What's Next ?
- Allow spesific custom short url without generate using base58 encode
- Add SQL Database as backup to the Redis main store, that should be used for store of cold values/least frequently fetched data.
- Build a minimal front-end that would be consuming this service.
- Record any user click the short url (statistic)

## Reference Tutorial
- https://www.eddywm.com/lets-build-a-url-shortener-in-go/