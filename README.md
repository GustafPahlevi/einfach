# Einfach-Message
_Einfach_ meaning __simple__ in English. This is a simple API service application.

## Prerequisites
- [Golang 1.12.7][go] as a Programming Language
- [MongoDB Go Driver][mongodb] as persistence storage

## Usage

### 1. Clone this repository to any desired directory.

```sh
git clone https://github.com/GustafPahlevi/einfach-msg.git
```

### 2. Install all required dependencies.

```sh
go mod tidy && go mod vendor
```

### 4. Setup database
1. Install mongodb in your local machine
2. setup the database configuration with details below:
```sh
database name: "message"
collection: "message_collection"
```

### 5. Run the application

```sh
go run main.go serve
```

## References
- Repository structure style is adopted from [Ardan Labs Article](ardan)

[go]: https://golang.org/dl/ 
[mongodb]: https://github.com/mongodb/mongo-go-driver