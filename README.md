
# Golang In memory key-value store REST API

# Features
+ An endpoint to set the key
+ An endpoint to get the key
+ An endpoint to flush all data
+ An endpoint to fetch all the data
+ An endpoint to delete data by key
+ Save to file at a specified interval (every N minutes)
+ When the application stops and stands up again, it should load the existing data back into memory if there are any saved files ( /tmp/TIMESTAMP-data.json ?)

# Endpoints
+ set
+ get/{key}
+ flush
+ getAll
+ deleteSingle/{key}

# Heroku Link

https://go-case-deneme.herokuapp.com/

# Swagger Link
 
https://go-case-deneme.herokuapp.com/swagger/index.html

# Unit Test
+ (store_test.go, index_test.go)

+ All Run :
 
 ```
go test
 ```
 
+ Related Folder Run:
 ``` 
go test -v
 ```
 

+ Coverage Report:
 ```
go test -cover
 ```
 
+ Coverege Report on Browser:
```
go test --coverprofile=coverage.out && go tool cover -html=coverage.out  --> we can see coverage report on html
 ```

# Run

+ Standard
   ```
   go run main.go
  ```
+ Docker Container
    ``` 
    docker build -t go-in-memory
    ```
    
    ```
    docker run -p 4444:4444 go-in-memory
    ```
