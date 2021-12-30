## Simple REST-API Boilerplate

This is a simple implementation of REST-API using Golang and several packages (Echo and GORM). By default, I use PostgreSQL database from Docker Images. There are two ways to run this app, using Docker or run manually on your system. I suggest to use Docker if possible.

### 1. Docker
- Go to directory folder and type:
  <pre><code>docker-compose up -d</code></pre>
  >Note: The process will running into the background because we add option <code>-d (detach mode)</code>
- The server can be accessed through port <code>8080</code>

### 2. Local
> Prerequisites: Make sure you've already installed Golang and PostgreSQL on your system.
- Create a table on PostgreSQL by copying the SQL syntax inside file <code>docker_pg_init.sql</code>
- Go to directory folder and insert the environment variables below on terminal (bash/zsh):
  <pre>
  <code>export DB_HOST=localhost</code>
  <code>export DB_PORT=5432</code>
  <code>export DB_USER=postgres</code>
  <code>export DB_PASSWORD=postgres</code>
  <code>export DB_DATABASE=db_exercise</code>
  <code>export SERVER_PORT=8080</code>
  </pre>
  > Note: You can customize the value by desire.
- Install the dependencies:
  <pre><code>go mod tidy</code></pre>
- Type below to execute the app:
  <pre><code>go run main.go</code></pre>
- Server can be accessed through port <code>8080</code>

### Available Endpoints:
>Note: All endpoints start from prefix "/api"

| Endpoints | Method |    Request Body   |    Description   |
|-----------|:------:|:-----------------:|:----------------:|
| /student  | GET    |         -         | Get all students |
| /student  | POST   | {"name": "user1"} | Add a student    |
