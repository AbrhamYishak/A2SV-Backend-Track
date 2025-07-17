# Task Management API Documentation

**Base URL:** `http://localhost:8080`

---

## Endpoints

---

###  Get All Tasks

**Endpoint**:  
`GET /tasks`

**Description**:  
Retrieves a list of all tasks.

**Request Headers**:
- `Content-Type: application/json`

**Query Parameters**: *(optional)*  
- `id`: Use `id=1` to get a specific task (alternative to `/tasks/:id`)

**Request Body**:  
None

**Success Response**:
- `200 OK`: Returns an array of task objects

**Example cURL**:
```bash
curl --location 'http://localhost:8080/tasks'
```

**Example Response**:
```json
[
  {
    "ID": "2",
    "Title": "complete task 4",
    "Description": "on the way of completion",
    "Status": true,
    "Duedate": "2025-07-18T08:30:00Z"
  },
  {
    "ID": "1",
    "Title": "complete task 4",
    "Description": "on the way of completion",
    "Status": true,
    "Duedate": "2025-07-18T08:30:00Z"
  },
  {
    "ID": "3",
    "Title": "complete task 4",
    "Description": "on the way of completion",
    "Status": true,
    "Duedate": "2025-07-18T08:30:00Z"
  }
]
```

---

###  Get Task by ID

**Endpoint**:  
`GET /tasks/:id`

**Description**:  
Fetches a specific task by its ID.

**Request Headers**:
- `Content-Type: application/json`

**Request Body**:  
None

**Success Response**:
- `200 OK`: Task data returned

**Example cURL**:
```bash
curl --location 'http://localhost:8080/tasks/2'
```

**Example Response**:
```json
{
  "ID": "2",
  "Title": "complete task 4",
  "Description": "on the way of completion",
  "Status": false,
  "Duedate": "2025-07-18T08:30:00Z"
}
```
### Create Task

**Endpoint**:  
`POST /tasks/`

**Description**:  
Creates a new task. This request submits data to the server via the request body in JSON format.

**Request Headers**:
- `Content-Type: application/json`

**Request Body** (JSON):
```json
{
  "ID": "4",
  "Title": "complete task 4",
  "Description": "on the way of completion",
  "Duedate": "2025-07-18T08:30:00Z",
  "Status": true
}
```

**Success Responses**:
- `201 Created`: Successfully created the task
- `200 OK`: Request processed successfully (less common)

**Example cURL**:
```bash
curl --location 'http://localhost:8080/tasks/' \
--header 'Content-Type: application/json' \
--data-raw '{
  "ID":"4",
  "Title":"complete task 4",
  "Description": "on the way of completion",
  "Duedate": "2025-07-18T08:30:00Z",
  "Status": true
}'
```

**Example Response**:
```json
{
  "message": "Successfully created the task"
}
---

### Update Task

**Endpoint**:  
`PUT /tasks/:id`

**Description**:  
Updates an existing task by its ID. This request **replaces** the entire task object.

**Request Headers**:
- `Content-Type: application/json`

**Request Body** (JSON):
```json
{
  "ID": "2",
  "Title": "complete task 4",
  "Description": "on the way of completion",
  "Duedate": "2025-07-18T08:30:00Z",
  "Status": false
}
```

**Success Responses**:
- `200 OK`: Successfully updated

**Example cURL**:
```bash
curl --location --request PUT 'http://localhost:8080/tasks/2' \
--header 'Content-Type: application/json' \
--data-raw '{
  "ID":"2",
  "Title":"complete task 4",
  "Description":"on the way of completion",
  "Duedate":"2025-07-18T08:30:00Z",
  "Status": false
}'
```

**Example Response**:
```json
{
  "message": "Successfully updated the task"
}
```

---

###  Delete Task

**Endpoint**:  
`DELETE /tasks/:id`

**Description**:  
Deletes a task by its ID.

**Request Headers**:
- `Content-Type: application/json`

**Request Body**:  
None

**Success Responses**:
- `200 OK`: Task deleted

**Example cURL**:
```bash
curl --location --request DELETE 'http://localhost:8080/tasks/4'
```

**Example Response**:
```json
{
  "message": "Successfully deleted the task"
}
```

---
