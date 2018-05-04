# Cron Monitor

## Setup

Golang Dependencies:
```
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
```
cockroach start --insecure --store=cronmon --host=localhost
```

Terminal B:
```
cockroach user set evilroach --insecure
cockroach sql --insecure -e 'CREATE DATABASE cronmon'
cockroach sql --insecure -e 'GRANT ALL ON DATABASE cronmon TO evilroach'
cronmon server
```
