# Cron Monitor

## Setup

Golang Dependencies:
```shell
go get -u -v github.com/gorhill/cronexpr
go get -u -v github.com/prometheus/client_golang/prometheus
go get -u -v github.com/jmoiron/sqlx
go get -u -v github.com/lib/pq
go get -u -v github.com/sirupsen/logrus
go get -u -v github.com/labstack/echo
go get -u -v github.com/labstack/echo/middleware
go get -u -v github.com/prometheus/client_golang/prometheus/promhttp
go get -u -v github.com/spf13/cobra
```

Terminal A:
```shell
cockroach start --insecure --store=cronmon --host=localhost
```

Terminal B:
```shell
cockroach user set evilroach --insecure
cockroach sql --insecure -e 'CREATE DATABASE cronmon'
cockroach sql --insecure -e 'GRANT ALL ON DATABASE cronmon TO evilroach'
cronmon server
```

## Usage

### Create Task to Monitor

#### cURL

Request:
```shell
curl -XPOST -H 'Content-Type: application/json' -d '{ "name": "Example Task", "description": "This is an example task for the cronmon README file", "cron_expression": "* * * * *"}' http://localhost:8080/api/v1/tasks
```

Response (formatted for readability via jq):
```json
{
  "id": 345302135986061313,
  "name": "Example Task",
  "description": "This is an example task for the cronmon README file",
  "cron_expression": "* * * * *",
  "created_at": {
    "Time": "2018-05-04T15:35:36.949293868Z",
    "Valid": false
  },
  "created_by": "system",
  "updated_at": {
    "Time": "2018-05-04T15:35:36.949309482Z",
    "Valid": false
  },
  "updated_by": "system",
  "last_started": {
    "Time": "0001-01-01T00:00:00Z",
    "Valid": false
  },
  "last_completed": {
    "Time": "0001-01-01T00:00:00Z",
    "Valid": false
  }
}
```

### Get a task

#### cURL

Request:
```shell
curl -XGET -H 'Content-Type: application/json' http://localhost:8080/api/v1/tasks/345302135986061313
```

Response (formatted for readability via jq):
```json
{
  "id": 345302135986061300,
  "name": "Example Task",
  "description": "This is an example task for the cronmon README file",
  "cron_expression": "* * * * *",
  "created_at": {
    "Time": "2018-05-04T15:37:29.109985Z",
    "Valid": true
  },
  "created_by": "system",
  "updated_at": {
    "Time": "2018-05-04T15:37:29.109985Z",
    "Valid": true
  },
  "updated_by": "system",
  "last_started": {
    "Time": "0001-01-01T00:00:00Z",
    "Valid": false
  },
  "last_completed": {
    "Time": "0001-01-01T00:00:00Z",
    "Valid": false
  }
}
```
