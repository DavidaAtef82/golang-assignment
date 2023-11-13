#Running Steps

#Run the following commands
 ```
go mod tidy
go mod download
```

#Run the app
```
change database configuration in .env file
go run main.go
```

#Docker
```
docker build -t my-golang-app .  
docker run -p 9090:9090 my-golang-app
```