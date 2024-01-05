# udacity-go-crm

Code for Udacity's Go course final project.

## Dependencies

Run `go get .`

Dependencies are listed in `go.mod`:

```
	github.com/google/uuid v1.5.0
	github.com/gorilla/mux v1.8.1
```

## Starting the server

Run `go run .`

## Testing

Run `go test`

## API documentation

For the HTML version, visit `http://localhost:8000` while running the server.

## `GET /customers`

Returns a list of all customers

### Response

Returns a `200 OK` response, and a JSON array of objects with the following properties:

- `id`: String. The unique identifier of the customer.
- `name`: String. The name of the customer.
- `role`: String. The role of the customer.
- `email`: String. The email of the customer
- `phone`: String. The phone of the customer
- `contacted`: Boolean. If the customer has been contacted or not

### Example

Request:

```
GET /customers
```

Response:

```json
{
  "17b0f3c0-7148-4e7a-8b91-71c22ca1105c": {
    "name": "Nicholas Runolfsdottir V",
    "role": "Security Analyst",
    "email": "Sherwood@rosamond.me",
    "phone": "(586)493-6943",
    "contacted": true
  },
  "e1827a7d-1acd-46ef-9d92-2a4d78bd7669": {
    "id": "e1827a7d-1acd-46ef-9d92-2a4d78bd7669",
    "name": "Clementina DuBuque",
    "role": "CEO",
    "email": "Rey.Padberg@karina.biz",
    "phone": "(024)648-3804",
    "contacted": true
  },
  "fb871ddf-ad69-40b9-966d-ab8e29504438": {
    "id": "fb871ddf-ad69-40b9-966d-ab8e29504438",
    "name": "Glenna Reichert",
    "role": "Software Engineer",
    "email": "Chaim_McDermott@dana.io",
    "phone": "(775)976-6794"
  }
}
```

### Errors

This endpoint doesn’t return any error codes

## `GET /customers/{customer_id}`

Returns a specific customer

### Response

Returns a `200 OK` response, and a JSON object with the following properties:

- `id`: String. The unique identifier of the customer.
- `name`: String. The name of the customer.
- `role`: String. The role of the customer.
- `email`: String. The email of the customer
- `phone`: String. The phone of the customer
- `contacted`: Boolean. If the customer has been contacted or not

### Example

Request:

```
GET /customers/e1827a7d-1acd-46ef-9d92-2a4d78bd7669
```

Response:

```json
{
  "id": "e1827a7d-1acd-46ef-9d92-2a4d78bd7669",
  "name": "Clementina DuBuque",
  "role": "CEO",
  "email": "Rey.Padberg@karina.biz",
  "phone": "(024)648-3804",
  "contacted": true
}
```

## Errors

This endpoint uses the following error codes:

- `404 Not Found`: The requested customer was not found.

## `POST /customers`

Creates a customer and adds it to the customer list

### Response

Returns a `201 Created` response, and a JSON object with the following properties:

- `id`: String. The unique identifier of the customer.
- `name`: String. The name of the customer.
- `role`: String. The role of the customer.
- `email`: String. The email of the customer
- `phone`: String. The phone of the customer
- `contacted`: Boolean. If the customer has been contacted or not

### Example

Request:

```json
POST /customers

Body:
{
  "name": "John Dorian, MD", // Required
  "role": "Resident",
  "email": "john@dorian.md",
  "phone": "(532)746-3955",
  "contacted": true
}
```

Response:

```json
{
  "id": "e1827a7d-1acd-46ef-9d92-2a4d78bd7669",
  "name": "John Dorian, MD", // Required
  "role": "Resident",
  "email": "john@dorian.md",
  "phone": "(532)746-3955",
  "contacted": true
}
```

## Errors

This endpoint uses the following error codes:

- `400 Bad Request`: When name is empty

## `PUT /customers/{customer_id}`

Returns a specific customer

### Response

Returns a `200 OK` response, and a JSON object with the following properties:

- `id`: String. The unique identifier of the customer.
- `name`: String. The name of the customer.
- `role`: String. The role of the customer.
- `email`: String. The email of the customer
- `phone`: String. The phone of the customer
- `contacted`: Boolean. If the customer has been contacted or not

### Example

Request:

```json
POST /customers/e1827a7d-1acd-46ef-9d92-2a4d78bd7669

Body:
{
  "id": "e1827a7d-1acd-46ef-9d92-2a4d78bd7669",
  "name": "Clementina DuBuque",
  "role": "CEO",
  "email": "Rey.Padberg@karina.biz",
  "phone": "(024)648-3804",
  "contacted": true
}
```

Response:

```json
{
  "id": "e1827a7d-1acd-46ef-9d92-2a4d78bd7669",
  "name": "Clementina DuBuque",
  "role": "CEO",
  "email": "Rey.Padberg@karina.biz",
  "phone": "(024)648-3804",
  "contacted": true
}
```

## Errors

This endpoint uses the following error codes:

- `404 Not Found`: The requested customer to be updated was not found.
- `400 Bad Request`: The request body didn’t include name, or name is empty.

## `DELETE /customers/{customer_id}`

Returns a specific customer

### Response

After a successful deletion, and empty response with `204 No Content` status code is returned

### Example

Request:

```
DELETE /customers/e1827a7d-1acd-46ef-9d92-2a4d78bd7669
```

Response:

```json
""
```

## Errors

This endpoint uses the following error codes:

- `404 Not Found`: The requested customer to be deleted was not found.
