## Running the application
We added basic project skeleton with docker-compose. (optional)
Feel free to refactor but provide us with good instructions to start the application
```
make docker-up
```

Update the `docker/mysql/dump.sql` to initialize the mysql database






# Introduction 

This repo is a very simple API for registering tables, guests and accompaning guests at a table for a fantastical celebration.

# How to run

## Dependencies

It depends on a few go libraries that are listed in ``` go.mod```. 

## Starting the app

To build the docker image:
``` make docker-up ```


## The API

This API consists of 6 endpoints:

### Add table

```
POST /tables
body: 
{
    "capacity": 10
}
response: 
{
    "id": 2,
    "capacity": 10
}
```

### Add a guest to the guestlist

```
POST /guest_list/name
body: 
{
    "table": int,
    "accompanying_guests": int
}
response: 
{
    "name": "string"
}
```

### Get the guest list

```
GET /guest_list
response: 
{
    "guests": [
        {
            "name": "string",
            "table": int,
            "accompanying_guests": int
        }, ...
    ]
}
```

### Guest Arrives

```
PUT /guests/name
body:
{
    "accompanying_guests": int
}
response:
{
    "name": "string"
}
```

### Guest Leaves

```
DELETE /guests/name
response code: 204
```

### Get arrived guests

```
GET /guests
response: 
{
    "guests": [
        {
            "name": "string",
            "accompanying_guests": int,
            "time_arrived": "string"
        }
    ]
}
```

### Count number of empty seats

```
GET /seats_empty
response:
{
    "seats_empty": int
}
```


## Further developments

1. Introduce unit testing for the web server and the application itself.
3. Add nginx.




