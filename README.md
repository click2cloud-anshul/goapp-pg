# CRUD-GoApp:
Sample crud application in golang using Postgres database.


# API performs following tasks:
• REST API should take user information as an input and store in PostgreSQL/MySQL/any other database
• REST API should fetch all users information
• REST API should fetch single user information
• REST API should update single user information
• REST API should delete single user information

# Pre-requsites:
- API understanding
- Golang
- Postgres sql
# APP
1.API endpoints:

    - GET
    
    - POST
    
    - PUT
    
    - UPDATE
    
    - DELETE

2.app img name : app:1

3.pg img name : postgres:latest

4.Commands:

    -app container : docker run --net=my-net --name "appname" -it -p 8080:8080 -e POSTGRES_HOST=pgdb -e POSTGRES_PORT=5432 -e POSTGRES_USER=root -e POSTGRES_PASS=anshdb -e POSTGRES_NAME=users app:2     

    -pg container :  1. docker run --net=my-net --name pgserv -p 5444:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=anshdb -d postgres
5.Database:

    -docker exec -it pg psql -U root
	-db name : users
    -network bridge : docker network create my-net
    -create database : 
    CREATE TABLE users (
        userid SERIAL PRIMARY KEY,
        name TEXT,
        age INT,
        location TEXT
      );
   6.Environment variables:
   
    POSTGRES_HOST=localhost
    POSTGRES_PORT=5432
    POSTGRES_USER=root
    POSTGRES_PASSWORD=anshdb
    POSTGRES_NAME=users
    
#END