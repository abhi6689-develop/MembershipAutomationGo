# Dynamic Cluster Membership Autoscaling Environments

This repository hosts a Go application designed to operate robustly within autoscaling environments, handling dynamic reassignment of worker nodes to new head nodes as scaling events occur. The system ensures continuous operation and state management through the use of Redis, even as head nodes are pruned during scale-down operations.

## Project Structure

- **Head/**
  - **main.go**: Initializes the server on the head node and handles sending of status messages.
  - **send_messages.go**: Manages periodic messaging to worker nodes to affirm active status.
  - **go.mod**: Specifies dependencies for the head node.

- **Workers/**
  - **main.go**: Entry point for worker nodes that manage incoming messages and head node tracking.
  - **communication.go**: Handles TCP communication logic.
  - **redis_actions.go**: Contains Redis utility functions for node state management.
  - **globals.go**: Manages global variables and synchronization mechanisms.
  - **reassignment.go**: Implements logic for reassigning workers to new heads when current heads are removed.
  - **watch_node.go**: Monitors the connectivity and status of the head node.
  - **go.mod**, **go.sum**: Manage dependencies for the worker nodes.

## Features

- **Dynamic Node Reassignment**: Automatically handles the reassignment of worker nodes to new head nodes during autoscaling events.
- **State Management via Redis**: Uses Redis to keep track of active head nodes and worker status, ensuring quick recovery and state synchronization.
- **Scalability and Fault Tolerance**: Designed to scale seamlessly with support for handling node failures and network partitions.

## Getting Started

These instructions will guide you through setting up the project on your local machine for development and testing purposes.

### Prerequisites

- Go 1.15 or higher
- Redis server

### Installation

Clone the repository and install dependencies for both head and worker components:

```bash
git clone https://github.com/yourusername/go-tcp-redis-server.git
cd go-tcp-redis-server/Head
go mod tidy
cd ../Workers
go mod tidy
```

### Running the Application

1. Start the Redis server:
   ```bash
   redis-server
   ```

2. Run the head node server:
   ```bash
   cd Head
   go run main.go
   ```

3. In a new terminal, start the worker node(s):
   ```bash
   cd Workers
   go run main.go
   ```

## Usage

This system is intended for use in environments where autoscaling of head and worker nodes is frequent, such as cloud deployments with elastic scaling capabilities. Customize the reassignment and failure handling logic in `reassignment.go` as per your specific operational requirements.

## Contributing

We welcome contributions! Please fork the repository, make changes, and submit a pull request with your improvements.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- Thanks to the Go community for continuous support and resources.
- Appreciation for the developers of the `github.com/go-redis/redis` package for providing a robust Redis client for Go.

