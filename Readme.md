Valkey GLIDE Go Project

A simple Go application demonstrating basic operations with the Valkey database using the official GLIDE (General Lightweight Interface for Database Engines) driver.
🚀 Getting Started
Prerequisites

    Go: version 1.26.2 or higher.

    Valkey Server: Installed and running.

Installation

    Clone this repository.

    Install dependencies:
    Bash

    go mod download

Running the Server

Ensure your Valkey server is active before running the application:
Bash

valkey-server

Running the Application

You can configure the connection via environment variables (optional). By default, it connects to localhost:6379.
Bash

# Optional: Set custom host/port
export VALKEY_HOST=localhost
export VALKEY_PORT=6379

go run .

🛠 Features Demonstrated

The application performs the following operations:

    Connection: Establishes a client with a 5-second timeout and default database selection.

    Ping: Verifies server availability.

    Set / Get: Basic string operations for individual keys.

    MSet / MGet: Atomic multi-key insertion and retrieval.

    Del: Deletion of multiple keys.

📂 Project Structure

    main.go: Orchestrates the execution flow.

    connection.go: Handles client initialization and configuration using valkey-glide.

    operation.go: Contains the logic for CRUD operations.

    go.mod: Project dependencies, including valkey-io/valkey-glide/go/v2.

📝 Configuration Note

The client is configured by default to use non-TLS connections and Database ID 0. These settings can be modified in the connection() function within connection.go.
