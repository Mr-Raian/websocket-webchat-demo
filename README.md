# Simple Chat Website

This is a simple chat website built using the [Echo Go HTTP framework](https://echo.labstack.com/), the [ent library](https://github.com/ent/ent), and the [Gorilla WebSocket library](https://github.com/gorilla/websocket). The app allows users to send and receive messages in real-time through a WebSocket connection in JSON format. All active clients are able to receive the messages and the messages are also stored in a Postgres DB.

The project is not intended for production use and is simply a demonstration of the functionality of WebSockets. The license for this project is MIT.

## Features

- Real-time messaging using WebSockets
- JSON message format
- Message broadcasting to all active clients
- Saving of messages to a Postgres DB

## Dependencies

- go 1.19
- entgo.io/ent v0.11.6
- github.com/gorilla/websocket v1.5.0
- github.com/labstack/echo/v4 v4.10.0
- github.com/labstack/gommon v0.4.0

Note: The Setup instructions and requirements are not mentioned in the README.md as the project is not production ready, and it's for learning purpose only.
