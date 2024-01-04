# FloorPlanXpert

This is my submission for the Backend Developer role at MoveInSync.

## Problem Statement

The problem statement is to develop an Intelligent Floor Plan Management System for a Seamless Workspace Experience, provding features for users to upload room status, addressing potential conflicts during simultaneous updates.

## Solution

I have implemented a CLI tool that can be used to provide the required functionalities from a given set of inputs. The tool is written in Python and uses the requests library to make the requests to the Go server.

## Pre-requisites

Make sure Go, python and postgresql are installed on your system.

- [Go](https://golang.org/doc/install)
- [Python](https://www.python.org/downloads/)
- [PostgreSQL](https://www.postgresql.org/download/)

## Getting Started

1. Clone the repository

```bash
git clone https://github.com/0xt3j4s/FloorPlanXpert.git
```

2. Navigate to the repository

```bash
cd FloorPlanXpert
```

3. Initialize Go Modules

```bash
go mod init
```

4. Download the dependencies

```bash
go get .
go mod tidy
```

5. Create a database in PostgreSQL

```bash
sudo -u postgres psql
```

```sql
CREATE DATABASE floorplan;
```

6. Create a username and password for the database

```sql
CREATE USER <username> WITH PASSWORD '<password>';
```
7. Grant Permissions to the user

```sql
GRANT ALL PRIVILEGES ON DATABASE floorplan TO <username>;
```

<!-- write this file location as a link to the local file in the project directory -->
8. Update the credentials of the database in the [/internal/db/db.go](/internal/db/db.go) file.

9. Run the server

```bash
go run cmd/main.go
```
10. Open another terminal and run the CLI Tool

```bash
python app.py
```

11. Follow the instructions on the CLI Tool to create a new room or book a room.

## Outputs
1. Running the Go server
   ![image](https://github.com/0xt3j4s/FloorPlanXpert/assets/75741089/cb354f5c-3d05-4b8f-afdb-22dd618498c2)

2. CLI Tool
   ![image](https://github.com/0xt3j4s/FloorPlanXpert/assets/75741089/d6b268c8-4cab-47bf-8ce2-a4c9792ea0d1)
