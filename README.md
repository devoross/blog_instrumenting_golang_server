# Simple Go HTTP Web Server

This is a simple HTTP Web Server written in Go for the purposes of a blog I am writing, to give a simple introduction to recording RED style metrics from a simple web application written in Go

## What is this web server?

This server will retrieve a random piece of advice from an in memory store, as well as provides Create and Delete capabilities. 

In `server/server.go` we run a method on the store which will populate it. This takes one single `int` parameter which represents the amount of quotes you wish to populate the store with when the application starts. Below are the endpoints and parameters for those if you wish to add your own

## Endpoints

### /api/v1/advice

The `/api/v1/advice` endpoint is used to retrieving random pieces of advice, as well as the capability of creating your own or removing any

#### GET

We use the `GET` endpoint to retrieve a random piece of advice

##### Status Codes:

* `200` - Success - The resource was retrieved successfully
* `404` - Not Found - There were no pieces of advice to provide

```bash
curl -XGET http://localhost:8080/api/v1/advice
```

Response

```json
{
    "advice": "If you've nothing nice to say, say nothing."
}
```

#### POST

We use the `POST` method to create a resource, in this instance a piece of advice

##### Status Codes:

* `200` - Success - The resource was created successfully
* `400` - Bad Request - The payload you have provided was not correct
* `409` - Conflict - The resource already exists

```bash
curl -XPOST -H "Content-Type: application/json" -d '{"advice": "Make your bed in the morning."}' http://localhost:8080/api/v1/advice -v
```

Response

```json
{
    "success": true
}
```

#### DELETE

We use the `DELETE` method to delete a resource that was previously added

##### Status Codes:

* `200` - Success - The resource was deleted successfully
* `400` - Bad Request - The payload you have provided was not correct
* `404` - Not Found - The resource does not exist

```bash
curl -XDELETE -H "Content-Type: application/json" -d '{"advice": "Make your bed in the morning."}' http://localhost:8080/api/v1/advice -v
```

Response

```json
{
    "success": true
}
```