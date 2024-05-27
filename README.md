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

### Library and tools used
I have used belows pretty cool library and tools, follow installation steps on their official docs.

**Golang-migrator**

**Go-Swagger**

**SQLC**

**Docker**

**Postgresql**

**Goland**


## Documentation
 
### Checkout the [swagger docs](http://localhost:8082/api/v1/swagger/index.html) please!)

### Available endpoints
#### User 
 * Create user
 * Get
 * Delete

#### Location
* Create location
* Get
* Delete

#### Price Estimation

* create
* List
* Get
* Update
* Delete

## Reminder
I had create a [hexagonal-architecture](https://github.com/yinebebt/hexagonal-architecture) project boilerplate, 
please check out and let me know.
