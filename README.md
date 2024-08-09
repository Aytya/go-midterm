# Go - Service for compiling task lists

### RESTful API for the Todo List microservice.

## Getting Started

### Prerequisites:
    - Go 1.22 or later
    - React 18.2.0 or later

### Installation:
1. Clone the repository:
   ```bash
   https://github.com/Aytya/todoApp-wails
   ```
2. Navigate into the project directory:
   ```bash
    cd todoApp-wails
   ```
3. Install dependencies:
   ```bash
    go get -u "github.com/go-chi/chi"
    go get -u "github.com/jmoiron/sqlx"
    go get -u "github.com/lib/pq"
    go get -u "github.com/rs/cors"
    go get -u "github.com/spf13/viper"
    go get -u "github.com/go-chi/chi/v5/middleware"
    go get -u "github.com/joho/godotenv"
   ```

##  Build and Run Locally:
### Build the application:
   ```bash
   wails dev
   ```
- Wails port:
   ```bash
      wails://wails.localhost:34115
  ```
## API Endpoints:
### Create a New Task:
- URL: http://localhost:8080
- Method: POST
- Request Body:
 ```bash
    {
       "title": "Получить работу",
       "date":"12/12/12",
       "time":"12:30",
       "priority":"high"
    }
 ```
- Response Body:
 ```bash
    {
       "id":"cbb255c3-5f79-4da6-8669-70ebde53e1f6",
       "title":"Получить работу",
       "date":"12/12/12",
       "time":"12:30",
       "active_at":"2024-08-08T17:05:09.228715+05:00",
       "status":false,
       "priority":"high"
    }
 ```
### Update an Existing Task:
- URL: http://localhost:8080/{id}
- Method: PUT
- Request Body:
 ```bash
    {
        "title": "Получить работу",
        "date":"12/12/12",
        "time":"12:30",
        "priority":"hign"
    }
 ```
### Delete an Existing Task:
- URL: http://localhost:8080/{id}
- Method: DELETE

### Change Task Status From Active To Done:
- URL: http://localhost:8080/check/{id}
- Method: PUT

### Get By Id:
- URL: http://localhost:8080/{id}
- Method: GET

### Get By All:
- URL: http://localhost:8080/
- Method: GET