# openweathermap-cache-proxy
Simple golang http proxy service for openweathermap.org with caching. By default requests are
cached for 1 hour and only JSON is expected. Golang 1.9 is required.

This is a regular proxy so an api key is still needed. See more info at
http://openweathermap.org/appid.

## Build and run
```
go run main.go
```

or use [prebuilt](prebuilt) binaries
```
wget https://github.com/pkoretic/openweathermap-cache-proxy/blob/master/prebuilt/openweathermap-cache-proxy-x64-linux
chmod +x openweathermap-cache-proxy-x64-linux
./openweathermap-cache-proxy-x64-linux
```
## Use
```
# after starting this is a regular proxy for the 'http://api.openweathermap.org' listening at the
port 8000
http://localhost:8000/data/2.5/forecast?q=Zagreb,hr&cnt=5&appid=xxx
```
