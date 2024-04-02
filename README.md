# Websocket Practice

## Description

This is a simple websocket practice project. It is a simple chat application that allows users to send messages to each other in real time. and it build just the Server, the client side is not included. And also this project is inspired by this video (https://youtu.be/BcuXtC4afzU?si=HPjwiARhpeWNKWOL)

## Installation

make sure you already have go installed on your machine.

1. Clone the repository
2. run this command

```php
go run main.go
```

3. go to postman or any other tool that can send websocket request and send a request to, and make sure name is not empty

```php
ws://localhost:9000/ws/chat?name=yourname
```

4. then click connect, after conected then you can send message in field message
  <div style="text-align:center;">
  <img src="https://github.com/daffa09/learn-websocket-go/assets/68214221/af0a2c94-1882-4180-86c6-fc0e69c02926" alt="First Image" width="600" height="auto">
</div>


   type the message in here
  <div style="text-align:center;">
  <img src="https://github.com/daffa09/learn-websocket-go/assets/68214221/c7456f69-84f6-4476-b45d-39987170dbfc" alt="Second Image" width="600" height="auto">
</div>




6. repeat step 3 and 4 to add more user and then see the magic
7. <div style="text-align:center;">
  <img src="https://github.com/daffa09/learn-websocket-go/assets/68214221/ae00dbbd-14cb-4c83-ab15-079ec4157683" alt="Third Image" width="600" height="auto">
</div>

message form user 1 is send to all user that connected to the server, as you can see i was in user2 request screen and then i get message from user1

