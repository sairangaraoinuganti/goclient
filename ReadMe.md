# GoClient

## About 

  go Client sends the request to the server. Captures the response and validate it. It implements the following features 

 - valitdates path before sending the request
 - vaildates Response to see if matches our expected schema (struct)
 - validates Response UID(32 digit alphanumeric characters) 
 - have unit test 


## Prerequisites
  - Go
  - Docker 

## Build  
 ``` docker build  -t  goclient:v1 . ```

## RUN
- Default Run

    ```docker run -it goclient:v1 ```

  2021/12/01 08:55:20 Response Body:      {"uid":"187ef4436122d1cc2f40dc2b92f0eba0"} 

- Passing valid arguments with short endpoint 

    ```docker run -it goclient:v1 /short/a3 ```

    2021/12/01 08:15:18 Response Body: {"uid":"9d607a663f3e9b0a90c3c8d4426640dc"}

 - Passing valid arguments with long endpoint 

    ```docker run -it goclient:v1 /long/ac1 ```

     2021/12/01 03:16:37 Response Body: {"uid":"182a2e5271287357ea1966e47f0b427d"} 

  - Passing Invalid arguments with Long endpoint

    ```docker run -it goclient:v1 /long/ac ```

    2021/12/01 03:18:18 invalid path: /long/ac



## Dev Setup
- Download go modules
    
      go mod download

- Running tests 

      go test -v 

- Building Binary

      go build -o client

- Running  Binary
   
      ./client /short/ac
