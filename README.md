# Cron Monitor

## Setup


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
