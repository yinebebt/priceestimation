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

## Reminder
I had create a [hexagonal-architecture](https://github.com/yinebebt/hexagonal-architecture) project boilerplate, 
please check out.
