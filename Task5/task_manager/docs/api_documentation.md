# Task Management API Documentation

**Base URL:** `http://localhost:8080`

---

## Database

The application uses **MongoDB** as the database to store task data.

- **Database name:** `taskdb`  
- **Collection name:** `tasks`

Each task document in MongoDB has the following structure:

```json
{
  "_id": "ObjectId",           /* Automatically generated unique identifier by MongoDB */
  "title": "string",           /* Title of the task */
  "description": "string",     /* Description/details about the task */
  "completed": true,           /* Status indicating if the task is completed */
  "duedate": "ISODate"         /* Due date/time of the task */
}

```
## endpoints

---

###  get all tasks

**endpoint**:  
`get /tasks`

**description**:  
retrieves a list of all tasks from the database.

**request headers**:
- `content-type: application/json`

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
Fetches a specific task by its ID from databse.

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

---
### Create Task


**Endpoint**:  
`POST /tasks/`

**Description**:  
create and inset to the database based.

**Request Headers**:
- `Content-Type: application/json`

**Request Body** (JSON):
```json
{
  "Title": "complete task 4",
  "Description": "on the way of completion",
  "Duedate": "2025-07-18T08:30:00Z",
  "Status": false
}
```bash
curl --location --request PUT 'http://localhost:8080/tasks/2' \
--header 'Content-Type: application/json' \
--data-raw '{  
  "title": "complete task 4",
  "Description": "on the way of completion",
  "Duedate": "2025-07-18T08:30:00Z",
  "Status": false
}'
```
**Success Responses**:
- `201 Created`: Created a new resource (less common)
### Update Task

**Endpoint**:  
`PUT /tasks/:id`

**Description**:  
Updates an existing task by its ID. This request **replaces** the entire task object with the new task based on the given id.

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

**Success Responses**:
- `200 OK`: Successfully updated
- `201 Created`: Created a new resource (less common)
- `204 No Content`: No content returned but updated

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
Deletes a task by its ID from the database.

**Request Headers**:
- `Content-Type: application/json`

**Request Body**:  
None

**Success Responses**:
- `200 OK`: Task deleted
- `202 Accepted`: Request accepted
- `204 No Content`: Deleted, no content returned

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
