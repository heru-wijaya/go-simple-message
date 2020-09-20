# go-simple-message


this project will have 2 part
1. REST API with post and get message
2. long live connection with websocket

## REST API
if want to test API, you can do

`go run main.go` 

then import warpin.postman_collection.json


## Websocket
if want to test websocket, you can do

`go run server/main.go`

open 2 or more terminal and do

`go run client/main.go`

try typing on terminal client, and you can see the messages on other client
every post message from client will be save to database. you can see on API GET all message to check