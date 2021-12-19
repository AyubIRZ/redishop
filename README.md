# A Simple shop built with go and Redis search

### Prerequisites:
1. A working setup of Docker.
2. A working setup of Docker-compose.
3. Make!.

## Setup instructions

### Step 1:
To have the project working correctly, you should have an ".env" file, located in the root of the project.
For this you could copy the ".env.example" file to ".env":
```bash
$ cp .env.example .env
```

### Step 2:
Run the following command to setup the complete stack:
```bash
$ make build
```

### Step 3:
Use phpmyadmin(localhost:8080) with given username and password in the ".env" file to create product schema
(sorry for inconvenience, couldn't do better for now). Schema SQL code is located in 
"db/migrations/products_schema.up.sql" file.

### Step 4:
Import the given postman collection to be able to test the API endpoints. There are examples for each endpoint

##consideration
- There are some Test Doubles(fake implementations) to be used in tests. For example the "product_inmemory_repo".
- Unfortunately there aren't any test yet. There was a serious lack of time.

## TODO
- Tests