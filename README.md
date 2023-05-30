![Go Reference](https://pkg.go.dev/badge/github.com/go-telegram-bot-api/telegram-bot-api/v5.svg)
![Test](https://github.com/go-telegram-bot-api/telegram-bot-api/actions/workflows/test.yml/badge.svg)
<div style="width:100%; display: flex; align-items: center;">
  <h1>Api Go REST - GIN Framework
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" height="60" width="70" style="margin-bottom: -15px; z-index: -10; margin-left: 1.25rem"/>
  </h1> 
</div>


### 🛠  Description   

</br>

This project aims to create a robust REST API using the Go programming language (Golang) and the gin framework.. It takes advantage of several Go libraries to handle different tasks:

Gin: This HTTP web framework is used for API routing. Gin, optimized for speed, offers an elegant and simple interface for creating routes to the REST API and includes useful built-in features such as JSON validation and error handling.

Tests: Go's native testing library, in conjunction with the github.com/stretchr/testify/assert library, is used to write tests for the API. The assert library provides a set of useful assertion functions that can help make tests more readable and easier to write.

Database: Gorm, an Object-Relational Mapper (ORM) in Go, is used for database interactions. Gorm provides a simple interface for creating, retrieving, updating, and deleting records from a database. It also offers advanced features like automatic transactions and migrations.

Validation: The gopkg.in/validator.v2 library is used to validate the data before sending it to the database. This helps ensure that the data is in the correct format and meets any necessary requirements before it is stored.

With the combination of these libraries, this project provides a solid foundation for creating Go REST APIs that are easy to develop and test. It demonstrates best practices for structuring a Go application, including separation of responsibilities between routing, request handling, database interactions, and data validation.


## Preview Tests

</br>

<p align="center">
  <kbd>
 <img width="800" style="border-radius: 10px" height="400" src="https://github.com/JuanCampbsi/Preview_README/blob/0be0214da84b4c28d01b289b29b13a021a767e62/assets/api_gin.gif" alt="Intro"> 
  </kbd>
  </br>
</p>

</br>

### ⌨ Database configuration
Create an .env file in the root of the project and define your database user and password in it and place the file in db.go and docker-compose.yml
For example:

```bash
# .env
  DATABASE_USER=use
  DATABASE_PASSWORD=password
  POSTGRES_USER=user
  POSTGRES_PASSWORD=password
  POSTGRES_DB=db
  STRING_CONNECTDB="host=localhost..."
 

# use in docker-compose.yml 
  POSTGRES_USER=${POSTGRES_USER}
  POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
  POSTGRES_DB=${POSTGRES_DB} 

  PGADMIN_DEFAULT_EMAIL: ${DATABASE_USER}
  PGADMIN_DEFAULT_PASSWORD: ${DATABASE_PASSWORD}   

# install lib godotenv and use in db.go
  stringConnect := os.Getenv("STRING_CONNECTDB")
```

### ⌨ Installation
To use it, you need to clone the repository, install the dependencies and run the project.

```bash
# Open terminal/cmd and then Clone this repository
$ git clone https://github.com/JuanCampbsi/api_go_rest.git

# Access project folder in terminal/cmd
$ cd api_go_rest

# Install the dependencies
$ go mod tidy

# Run the application in development mode
$ go run main.go

# Run test all 
$ go test    

# Run test one
$ go test -run TestVerifyStatusCodeSearchStudentsWithParams   

# In order to create a container for this application, you'll need to run Docker command. 
#Please ensure that Docker is already installed on your machine before proceeding. 
#If you don't have Docker installed, you can download it from the 
#[official website](https://www.docker.com/products/docker-desktop).

$ docker compose up                                 

```

</br>	

### ⌨ Stack of technologies and libraries

-   [Golang](https://go.dev/doc/) - version 1.20
-   [GIN](https://github.com/gin-gonic/gin) - version 1.9.0
-   [Validator v2](https://gopkg.in/validator.v2) - version 2.0.1
-   [Testify](https://github.com/stretchr/testify) - version 1.8.3
-   [GORM](https://gorm.io/gorm ) - version 1.25.1
-   [Driver PostgreSQL](https://gorm.io/driver/postgres) - version 1.5.2 
-   [Godotenv](https://github.com/joho/godotenv) - version 1.5.1
 
</br>

👨‍💻 **Author** 💻

Developed by [_Juan Campos_](https://www.linkedin.com/in/juancampos-ferreira/)