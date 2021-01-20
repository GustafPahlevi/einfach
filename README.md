# Einfach-Message
_Einfach_ meaning __simple__ in English. Einfach-Message is a simple API service application ‍💻

__Cases :__  
1. User could post message as sending the message. 
2. User could get all delivered message.

## Prerequisites
- [Golang 1.12.7][go] as a Programming Language
- [MongoDB Go Driver][mongodb] as persistence storage

## Usage

### 1. Clone this repository to any desired directory.
```sh
git clone https://github.com/GustafPahlevi/go-simple-svc.git
```

### 2. Install all required dependencies.
```sh
go mod tidy && go mod vendor
```

### 3. Setup database
1. Install mongodb on your local machine
2. setup the database configurations with details below:
```sh
create database name = message
create database collection = message_collection
```

### 4. Run the application
```sh
go run main.go serve
```

### 5. Use the application
Endpoint is available:
```sh
http://localhost:8080/v1/message
```

Method are available:
- GET
- POST

Or, you could import these cURL:
- __Get__ message action: 
```sh
curl --location --request GET 'http://localhost:8080/v1/message' \
--header 'Content-Type: application/json'
```
- __Create__ message action:
```sh
curl --location --request POST 'http://localhost:8080/v1/message' \
--header 'Content-Type: application/json' \
--data-raw '{
	"sender_id":"12345",
	"receiver_id":"67890",
	"subject":"Hello",
	"message":"Hello World!"
}'
```

## References
- Repository structure style is adopted from [Ardan Labs Article](ardan)

[go]: https://golang.org/dl/ 
[mongodb]: https://github.com/mongodb/mongo-go-driver