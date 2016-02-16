# Go Web

**Go-web** is a tiny HTTP web server library which leverages [go-micro](https://github.com/micro/go-micro) to create 
micro web services as first class citizens in a microservice ecosystem. It's merely a wrapper around registration, 
heartbeating and initialistion of the go-micro client. In the future go-platform features may be included.

## Getting Started

## Prerequisites

Go-web uses a similar pattern to go-micro. Look at the go-micro [readme](https://github.com/micro/go-micro) for 
starting up the registry.

### Usage


```golang
service := web.NewService(
	web.Name("go.micro.web.example"),
	web.Version("latest"),
)

service.HandleFunc("/foo", fooHandler)

if err := service.Init(); err != nil {
	log.Fatal(err)
}

if err := service.Run(); err != nil {
	log.Fatal(err)
}
```
