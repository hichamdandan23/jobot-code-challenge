<details> <summary><strong>Click to reveal the README.md content</strong></summary>
# GraphQL Ping Project

A simple demonstration of a **GraphQL** application using **Go (gqlgen)** for the backend and **Next.js** with **Apollo Client** for the frontend. This project showcases queries, mutations, and subscriptions in real-time.

## Table of Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Tech Stack](#tech-stack)
4. [Project Structure](#project-structure)
5. [Setup & Installation](#setup--installation)
    - [Backend](#backend)
    - [Frontend](#frontend)
6. [Running the Project](#running-the-project)
7. [GraphQL Endpoints](#graphql-endpoints)
8. [How It Works](#how-it-works)
9. [Troubleshooting](#troubleshooting)
10. [License](#license)

---

## Overview

This project maintains a **last ping timestamp** on a Go GraphQL server and allows multiple clients to:

- **Query** the latest ping timestamp on page load.
- **Send** a new ping mutation every 1–3 seconds.
- **Subscribe** to timestamp updates so all clients receive real-time updates when any client pings.

---

## Features

- **Go (gqlgen)** for GraphQL schema and resolvers
- **Next.js** for the React-based frontend
- **Apollo Client** to handle queries, mutations, and subscriptions over WebSockets
- **Real-time updates**: Each time a client pings, all connected clients get the new timestamp

---

## Tech Stack

- **Backend**: Go, gqlgen, Gorilla WebSocket
- **Frontend**: Next.js (React), Apollo Client
- **Transport**: GraphQL over HTTP & WebSockets (subscriptions)

---

. ├── README.md ├── backend/ │ ├── go.mod │ ├── go.sum │ ├── server.go │ └── graph/ │ ├── schema.graphql │ ├── generated/ // gqlgen auto-generated code │ ├── resolver.go // main resolver setup │ └── schema.resolvers.go └── frontend/ ├── package.json ├── pages/ │ └── index.js ├── lib/ │ └── apolloClient.js └── ...

## Project Structure


---

## Setup & Installation

### Backend

1. **Install Go** (1.18+ recommended).
2. **Navigate** to the `backend` folder:
   ```bash
   cd backend
3. Initialize the Go module (if not done already):
go mod tidy

4. Run the server:
   ```bash
   go run server.go

### Frontend
1. Install Node.js (v16+ recommended).
2. Navigate to the frontend folder:
    ```bash
    cd frontend
3. Install dependencies:
    ```bash
    npm install
    or
    ```bash
    yarn
4. Start the Next.js dev server:
    ```bash
    npm run dev
This starts the frontend on http://localhost:3000 by default.


## Running the Project
### Start the Go Server:

bash
Copy
cd backend
go run server.go
The GraphQL Playground is available at http://localhost:8081/.
The GraphQL endpoint (queries/mutations/subscriptions) is at /query.
Start the Next.js App:

    ```bash
    cd frontend
    npm run dev
Open http://localhost:3000 in your browser to view the application.
GraphQL Endpoints
Query: lastPing

    ```bash
    query {
      lastPing
    }
Returns the most recent ping timestamp (string).

Mutation: ping

    ```bash
    mutation {
      ping
    }
Sets the new timestamp and broadcasts to all subscribed clients.

Subscription: pingReceived

    ```bash
    subscription {
      pingReceived
    }
Listens for new pings from any client.

### How It Works
Query (lastPing):
  When the frontend loads, it requests the most recent ping from the Go server.

Mutation (ping):
  Each tab (or client) sends a ping every 1–3 seconds. The server updates the timestamp and notifies all subscribers.

Subscription (pingReceived):
  All clients subscribe to pingReceived. Whenever a ping occurs, they get a real-time update with the new timestamp.

Troubleshooting
  CORS Errors:

    Ensure your Go server allows requests from http://localhost:3000 if you’re testing locally.
    If you see “blocked by CORS policy,” update your server’s CORS settings.
    WebSocket Connection Fails:
    
    Make sure your server includes the WebSocket transport and CheckOrigin is set to true (or properly configured).
    Verify your frontend WebSocket URL is correct (e.g., ws://localhost:8081/query).
    Port Conflicts:
    
    If something else is running on 8081, change the port in server.go or use an environment variable.
    Subscriptions Not Receiving Updates:
    
    Check that your ping mutation sends updates to all subscribers.
    Verify your useSubscription or GraphQL Playground subscription is pointed at the correct endpoint.



</details>
