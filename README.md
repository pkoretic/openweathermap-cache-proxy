# openweathermap-cache-proxy
Simple golang http proxy service for openweathermap.org with caching. By default requests are
cached for 1 hour and only JSON is expected. Golang 1.9 is required.

## Build and run
```
go run main.go
```

or use [prebuilt](prebuilt) binaries.

## Use
```
# after starting this is a regular proxy for the 'http://api.openweathermap.org'
http://localhost:8000/data/2.5/forecast?q=Zagreb,hr&cnt=5&appid=xxx
```
