ZPE Systems technical test
---                            

This project is made using zord-microframework, my project to create golang projects faster. It is based on hexagonal architecture and using casbin RBAC from roles control, the idea is to manage and enforce permissions for another projects.

---

To find the http layer implementation go to /cmd/http
To find all business rules go to /internal/application/services
To find casbin calls and other external libs and reqs go to /pkg

### to start the project is needed to create a .env from .env.example and after start all containers run the migrate command inside the container
### after run the migration script run the seed.sql script on your database to set up the admin user and the 3 predefined roles
user:
```
email: admin@admin.com
password: 123456
role: admin
```

predefined roles:
```
admin: can do everything
modifier: can list and edit anything
watcher: can list and detail
```

the swagger documentation is in /swagger/index.html route, if you`re using the default port and docker will be http://localhost:9000/swagger/index.html


# Development
> Remember to create your .env file based on .env.example

### 1. Using Docker Compose
Up mysql and zord project:

``` SHELL
docker compose up
```

<br />

#### 2. Using raw go build

You will need to build the http/main.go file:

``` SHELL
go build -o server cmd/http/main.go
```

Then run the server

``` SHELL
./server
```

<br />

#### 3. Running from go file

``` SHELL
go run cmd/http/main.go
```

<br />

**Note:** To run the local build as described in the second or third option, a MySQL server must be running. This is necessary for the application to interact with its database. The easiest way to set up a MySQL server locally is by using Docker. Below is a command to start a MySQL server container using Docker:

``` SHELL
docker compose up mysql -d
```
This command will ensure that a MySQL server is running in the background, allowing you to execute the local build successfully.

---

### Cli

#### build cli

to build cli into binary file run
``` SHELL
go build -o cli cmd/cli/main.go
```

then you can run all cli commands with the binary file
``` SHELL
./cli -h
```

if you`re developing something in the cli the best way is run it directly to all changes 
``` SHELL
go run cmd/cli/main.go
```

---

#### Cli Commands

create new domain (crud):
``` SHELL
./cli create-domain {{domain}}
```

destroy domain:
``` SHELL
./cli destroy-domain {{domain}}
```

migrate all domains:
``` SHELL
./cli migrate
```

**Obs:** If you`re generating code inside docker container you need to change generated folder and file permissions to code out of docker container.

run the follow command to edit generated files:
``` SHELL
sudo chown $USER:$USER -R .
```

if you have a group name different from username change the command accordingly

---

#### Run tests
Run all tests:
``` SHELL
go test ./...
```

Verify code coverage:
``` SHELL
// Generate coverage output
go test ./... -coverprofile=coverage.out

// Generate HTML file
go tool cover -html=coverage.out
```

### Docs (WIP):
https://github.com/not-empty/zord-microframework-golang/wiki

### Development

Want to contribute? Great!

The project using a simple code.
Make a change in your file and be careful with your updates!
**Any new code will only be accepted with all validations.**


**Not Empty Foundation - Free codes, full minds**