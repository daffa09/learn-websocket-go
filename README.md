<!-- portfolio -->
<!-- slug: learn-websocket-go -->
<!-- title: WebSocket Chat Server -->
<!-- description: Real-time chat server built with Go and WebSocket for learning real-time communication -->
<!-- image: https://github.com/daffa09/learn-websocket-go/assets/68214221/af0a2c94-1882-4180-86c6-fc0e69c02926 -->
<!-- tags: golang, websocket, real-time, chat, server -->

# WebSocket Chat Server (Go)

A real-time chat server implementation built with Go and WebSocket protocol. This learning project demonstrates how to build WebSocket-based applications for real-time bidirectional communication between clients and server.

## 📋 Overview

This is a practice project focused on learning WebSocket implementation in Go. It creates a simple chat server where multiple clients can connect and send messages to each other in real-time. The project is inspired by [this excellent tutorial](https://youtu.be/BcuXtC4afzU?si=HPjwiARhpeWNKWOL).

**Note**: This repository contains only the **server-side** implementation. Client-side testing can be done using tools like Postman, wscat, or any WebSocket client.

## ✨ Features

- **Real-time Messaging**: Instant message delivery to all connected clients
- **WebSocket Protocol**: Efficient bidirectional communication
- **Multiple Clients**: Support for simultaneous connections
- **User Identification**: Name-based client identification
- **Broadcast System**: Messages sent to all connected users
- **Connection Management**: Handle client connect/disconnect events
- **Simple REST API**: WebSocket upgrade endpoint

## 🛠️ Technologies Used

- **Go (Golang)**: Server-side programming language
- **Gorilla WebSocket**: WebSocket library for Go
- **WebSocket Protocol**: Real-time communication protocol
- **HTTP**: For initial connection upgrade

## 📁 Project Structure

```
learn-websocket-go/
├── main.go              # Main server file with WebSocket handling
├── go.mod              # Go module dependencies
├── go.sum              # Dependency checksums
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- Go 1.18 or higher installed
- Basic understanding of WebSocket protocol
- WebSocket client for testing (Postman, wscat, or browser)

### Installation Steps

1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd learn-websocket-go
   ```

2. **Install Dependencies**
   ```bash
   go mod download
   ```
   
   Or let Go handle it automatically:
   ```bash
   go run main.go
   ```

3. **Run the Server**
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:9000`

## 💻 Usage Guide

### Testing with Postman

1. **Open Postman**
2. **Create New WebSocket Request**
3. **Connect to WebSocket**
   - URL: `ws://localhost:9000/ws/chat?name=yourname`
   - **Important**: `name` parameter must not be empty!
   
   Example URLs:
   ```
   ws://localhost:9000/ws/chat?name=John
   ws://localhost:9000/ws/chat?name=Alice
   ws://localhost:9000/ws/chat?name=Bob
   ```

4. **Click "Connect"**
   
   ![Postman Connection](https://github.com/daffa09/learn-websocket-go/assets/68214221/af0a2c94-1882-4180-86c6-fc0e69c02926)

5. **Send Messages**
   - Type your message in the message field
   - Click "Send"
   
   ![Type Message](https://github.com/daffa09/learn-websocket-go/assets/68214221/c7456f69-84f6-4476-b45d-39987170dbfc)

6. **Open Multiple Connections**
   - Repeat steps 2-4 with different names
   - Observe real-time message broadcasting

7. **See the Magic!**
   
   ![Message Broadcasting](https://github.com/daffa09/learn-websocket-go/assets/68214221/ae00dbbd-14cb-4c83-ab15-079ec4157683)
   
   Messages from User1 are broadcast to all connected users, including User2!

### Testing with wscat (Command Line)

Install wscat:
```bash
npm install -g wscat
```

Connect and chat:
```bash
# Terminal 1 - User 1
wscat -c "ws://localhost:9000/ws/chat?name=Alice"
> Hello everyone!

# Terminal 2 - User 2
wscat -c "ws://localhost:9000/ws/chat?name=Bob"
> Hi Alice!
```

### Testing with Browser JavaScript

```html
<!DOCTYPE html>
<html>
<body>
    <input type="text" id="message" placeholder="Type message...">
    <button onclick="sendMessage()">Send</button>
    <div id="messages"></div>

    <script>
        const ws = new WebSocket('ws://localhost:9000/ws/chat?name=WebUser');
        
        ws.onmessage = (event) => {
            const messagesDiv = document.getElementById('messages');
            messagesDiv.innerHTML += '<p>' + event.data + '</p>';
        };
        
        function sendMessage() {
            const input = document.getElementById('message');
            ws.send(input.value);
            input.value = '';
        }
    </script>
</body>
</html>
```

## 📡 WebSocket Endpoint

### Connection

```
WS /ws/chat
```

**Query Parameters:**
- `name` (required): Username for the chat participant

**Example:**
```
ws://localhost:9000/ws/chat?name=JohnDoe
```

### Message Format

Messages are sent as plain text strings. The server broadcasts each message to all connected clients.

**Send:**
```
Hello, everyone!
```

**Receive:**
```
JohnDoe: Hello, everyone!
```

## 🔧 Code Structure

### Main Components

1. **WebSocket Upgrade Handler**
   - Upgrades HTTP connection to WebSocket
   - Validates query parameters
   - Manages connections

2. **Message Broadcaster**
   - Distributes messages to all clients
   - Handles message formatting
   - Manages client list

3. **Connection Pool**
   - Tracks active connections
   - Handles user join/leave events
   - Cleanup on disconnect

## 🎓 Learning Concepts

This project demonstrates:

- **WebSocket Protocol**: Understanding upgrade mechanism
- **Concurrent Programming**: Using Go goroutines for multiple connections
- **Real-time Communication**: Bidirectional data flow
- **Connection Management**: Handling connect/disconnect events
- **Broadcasting**: Sending messages to multiple recipients
- **Query Parameters**: Extracting data from WebSocket URLs

## 💡 Examples

### Simple Chat Flow

1. Alice connects: `ws://localhost:9000/ws/chat?name=Alice`
2. Bob connects: `ws://localhost:9000/ws/chat?name=Bob`
3. Alice sends: "Hello!"
4. Bob receives: "Alice: Hello!"
5. Bob sends: "Hi Alice!"
6. Alice receives: "Bob: Hi Alice!"

## 🐛 Troubleshooting

**Connection Refused?**
- Ensure server is running
- Check port 9000 is not in use
- Verify correct URL format

**Name parameter missing error?**
- Always include `?name=yourname` in URL
- Name cannot be empty (username must not be empty)

**Messages not broadcasting?**
- Check multiple clients are connected
- Verify WebSocket connection is established
- Check server console for errors

## 🚀 Future Enhancements

Potential improvements:
- Private messaging between users
- Chat rooms/channels
- Message history/persistence
- User authentication
- Typing indicators
- Online users list
- File sharing
- React/Vue frontend client
- Message encryption

## 📚 Resources

- [Original Tutorial Video](https://youtu.be/BcuXtC4afzU?si=HPjwiARhpeWNKWOL)
- [Gorilla WebSocket Documentation](https://github.com/gorilla/websocket)
- [WebSocket Protocol RFC](https://tools.ietf.org/html/rfc6455)
- [Go Documentation](https://golang.org/doc/)

## 🙏 Acknowledgments

This project was inspired by and built following the tutorial at:
https://youtu.be/BcuXtC4afzU?si=HPjwiARhpeWNKWOL

Thanks to the creator for the excellent learning resource!

## 🤝 Contributing

This is a learning project. Feel free to fork, experiment, and enhance!

## 📄 License

Open source - available for educational purposes.

---

**Happy WebSocket Learning!** 🚀💬  
Real-time communication made simple with Go!
