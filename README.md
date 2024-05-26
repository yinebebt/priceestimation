# priceestimation

price estimation REST API

## Installation
 ### Postgresql

 Suppose a postgresql installation is found on you machine,
 Open postgresql shell via `sudo -u postgres psql` and run below queries

```sql
create database price_est;

create user <username> with encrypted password '<password>';

grant all privileges on database price_est to <username>; 
```

## Available endpoints

### User 
 * Create user
```curl
curl --location 'http://localhost:8082/api/v1/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name":"Yinebeb",
    "last_name":"Tariku",
    "email":"yintar5@gmail.com",
    "password":"Abcd1234@"
}'
```

* Login
```curl
curl --location 'http://localhost:8082/api/v1/users/auth' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"yintar5@gmail.com",
    "password":"Abcd1234@"
}'
```

* Get
```curl
curl --location 'http://localhost:8082/api/v1/users/c35963eb-7aa4-4e37-80a0-edb0fec93ff5'
```

* Delete
```curl
curl --location --request DELETE 'http://localhost:8082/api/v1/users/c35963eb-7aa4-4e37-80a0-edb0fec93ff5'
```

## Reminder
I had create a [hexagonal-architecture](https://github.com/yinebebt/hexagonal-architecture) project boilerplate, 
please check out.
