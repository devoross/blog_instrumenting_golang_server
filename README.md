# Simple Go HTTP Web Server

This is a simple HTTP Web Server written in Go for the purposes of a blog I am writing, to give a simple introduction to recording RED style metrics from a simple web application written in Go

## What is this web server?

This server will retrieve a random piece of advice from an in memory store, as well as provides CRUD capabilities. 

In `server/server.go` we run a method on the store which will populate it. This takes one single `int` parameter which represents the amount of quotes you wish to populate the store with when the application starts. Below are the endpoints and parameters for those if you wish to add your own

The server has a built in store, which is responsible for storing a slice of advices of which we can pick at random as and when that is requested

We are using the [Prometheus Go Client Library](https://github.com/prometheus/client_golang) in this project to record our metrics

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

#### PUT

We use the `PUT` method to update a resource, in this instance a piece of advice

##### Status Codes:

* `200` - Success - The resource was updated successfully
* `400` - Bad Request - The payload you have provided was not correct
* `404` - Not Found - The resource you are trying to update could not be found
* `409` - Conflict - The resource already exists

```bash
curl -XPUT -H "Content-Type: application/json" -d '{"advice": "Make your bed in the morning.", "updated_advice": "Make your bed in the morning"}' http://localhost:8080/api/v1/advice -v
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

### Metrics

RED metrics are what we are collecting as part of this simple Go HTTP Web Server

* Request - Total amount of requests (We should be able to use this with the errors metric to easily determine success/fail ratio)
* Errors - Errors that we experience within our application
* Duration - The latency of our requests, how long are they taking?

#### Metrics we are collecting

* Total Requests - `iwebserver_requests_total` - This will be a counter to count the total amount of requests - Labels: `code, route`
* Total Errors - `iwebserver_requests_errors_total` - This will be a counter to count the total amount of failed requests - Labels: `code, route`
* Duration - `iwebserver_requests_duration` - This will be a histogram to record the total duration of a request - Labels: `code, route`
* Items action total - `iwebserver_store_writes_total` - This will be a counter to count the amount of times that a piece of advice has been removed from the store - Labels: `action (deleted/created)`
* Items action duration - `iwebserver_store_writes_duration` - This will be a histogram that records the total time taken to remove/create a piece of advice - Labels: `action (deleted/created)`
* Total items in the store - `iwebserver_store_items_total` - This will be a histogram which states the current total amount of items in the store
* Total errors to the store - `iwebserver_store_errors_total` - This will be a counter which records the total amount of errors when hitting the store

Prometheus metrics found here : http://localhost:8080/metrics
