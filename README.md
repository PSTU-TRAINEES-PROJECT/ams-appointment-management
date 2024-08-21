# AMS-schedule-management

AMS Schedule Management is a tool built using Go to help manage and automate scheduling tasks within the AMS (Appointment Management System). The project aims to simplify the management of appointments, events, and other time-based activities.


# Initial Folder Structure

 ```bash
ams-schedule-management/
│   docker-compose.yml
│   Dockerfile
│   main.go
│   Makefile
│   README.md
│
└───app
    ├───cmd
    │       root.go
    │       serve.go
    │
    ├───common
    │   ├───consts
    │   │       consts.go
    │   │
    │   ├───logger
    │   │       logger.go
    │   │
    │   ├───pagination
    │   │       pagination.go
    │   │
    │   └───utils
    │           utils.go
    │
    ├───config
    │       config.go
    │
    ├───conn
    │   └───db
    │           db.go
    │
    ├───controller
    │   │   scheduler.go
    │   │
    │   └───response
    │           response.go
    │
    ├───domain
    │   ├───models
    │   │       scheduler.go
    │   │
    │   └───repository
    │           scheduler.go
    │
    ├───middlewares
    │       middlewares.go
    │
    ├───routes
    │       scheduler.go
    │
    ├───serializer
    │       scheduler.go
    │
    ├───server
    │       server.go
    │
    └───service
            scheduler.go

```

# Setup and Installation
## Prerequisites
Before you begin, ensure you have the following installed:

- **Go** (version 1.19 or higher): [Install Go](https://golang.org/dl/)
- **Git**: [Install Git](https://git-scm.com/)

## Installation

Follow these steps to set up and install the "ams-schedule-management" project on your local machine:

### 1. Clone the Repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/ams-schedule-management.git
cd ams-schedule-management
```

### 2. Install Dependencies

If your project requires any Go dependencies, use 'go mod' to download them. Run the following commands in the project directory:

```bash
go mod tidy
```

### 3. Build the Project

Compile the application by running:

```bash
go build -o ams-schedule-management ./cmd/ams-schedule-management
```






