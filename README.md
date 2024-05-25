# Melody Chat Application

This project is a simple chat application demonstrating real-time communication using WebSockets in Go, with the `melody` library for managing WebSocket connections. The client side is built with HTML, CSS, and JavaScript.

## Features
- Real-time messaging between multiple clients.
- Simple and clean user interface.
- Each user is assigned a random guest name.

## Technologies Used
- Go: For the backend server.
- Melody: For WebSocket handling in Go.
- HTML/CSS: For the frontend structure and styling.
- JavaScript: For WebSocket communication on the client side.

## Getting Started
Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- Go (1.14 or later)
- `melody` library for Go

### Installation
1. Clone the repository:
   git clone https://github.com/yourusername/melody-chat.git
   cd melody-chat
   
2. Install the melody package:
 go get github.com/olahol/melody

3. Run the serber
  go run main.go
