# Gook-oauth-API
This is one of the [Gook](https://github.com/Armingodiz/Gook) services for authorizing users and login and creating access token which follow domain driven deployment design pattern .

## Features 

* Get created access Token
* Create access token 

## Dependencies

name     | repo
------------- | -------------
  gin-gonic   | https://github.com/gin-gonic/gin
  redis       | https://github.com/go-redis/redis
  

## Installation 

First make sure you have installed all dependencies ,
make sure you have docker then pull redis image and connect your 8282 port to redis with `sudo docker run --name redis-usdb -p 8282:6379 -d redis`.
Then just simply clone this repository and start service with `go run main.go` (your service will be running on `localhost:2222`)


## EndPoints 

	GET ==> /oauth/access_token/:access_token_id (Get Access token by id)
	POST ==> /oauth/access_token (Create Access token)





