# Websocket Practice

## Description

This is a simple websocket practice project. It is a simple chat application that allows users to send messages to each other in real time. and it build just the Server, the client side is not included.

## Installation

make sure you already have go installed on your machine.

1. Clone the repository
2. run this command

```php
go run main.go
```

3. go to postman or any other tool that can send websocket request and send a request to, and make sure name is not empty

```php
ws://localhost:9000/ws?name=yourname
```

4. after conected then you can send message in field message

5. repeat step 3 and 4 to add more user and then see the magic

message form user 1 is send to all user that connected to the server
