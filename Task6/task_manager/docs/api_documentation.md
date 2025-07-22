# Task Management & User Authentication API Documentation

**Base URL:** `http://localhost:8080`

---

## **Overview**

This document describes the **Task Management API** and **User Authentication & Authorization API**, built with **Go, Gin, MongoDB, bcrypt, and JWT**.

---

# **Part 1: User Authentication API**

### **Features:**

* **User Registration**
* **User Login**
* **Promote user to admin**
* **JWT token generation and verification**

---

### **Models: User**

| Field    | Type     | Description              |
| -------- | -------- | ------------------------ |
| ID       | ObjectID | Unique user identifier   |
| Username | string   | Unique username          |
| Password | string   | Hashed password (bcrypt) |
| Isadmin  | bool     | Admin status             |

---

### **Endpoints**

#### **1. Register User**

* **URL:** `/register`
* **Method:** POST
* **Description:** Registers a new user. First user becomes admin.

**Request Body:**

```json
{
  "username": "string",
  "password": "string"
}
```

**Responses:**

| Status Code | Description                         |
| ----------- | ----------------------------------- |
| 201         | Successfully registered the user.   |
| 400         | Invalid input or username in use.   |
| 500         | Database or password hashing error. |

---

#### **2. Login User**

* **URL:** `/login`
* **Method:** POST
* **Description:** Authenticates user and returns JWT token.

**Request Body:**

```json
{
  "username": "string",
  "password": "string"
}
```

**Responses:**

| Status Code | Description                            |
| ----------- | -------------------------------------- |
| 200         | Successfully logged in, returns token. |
| 400         | Invalid input or wrong password.       |
| 404         | User not found.                        |
| 500         | Server error.                          |

---

#### **3. Promote User to Admin**

* **URL:** `/promote`
* **Method:** POST
* **Description:** Promotes a user to admin status.

**Request Body:**

```json
{
  "username": "string"
}
```

**Responses:**

| Status Code | Description                 |
| ----------- | --------------------------- |
| 200         | Successfully promoted user. |
| 400         | Invalid input.              |
| 404         | User not found.             |
| 500         | Database error.             |

---

### **JWT Token Handling**

* **Algorithm:** HS256
* **Claims:** `username`, `userid`, `isadmin`, `exp` (1 day)

| Function         | Description                          |
| ---------------- | ------------------------------------ |
| GetToken         | Generates a JWT token.               |
| VerifyToken      | Verifies validity of a JWT token.    |
| ExtractFromToken | Extracts username from token claims. |

---

### **Dependencies**

* github.com/gin-gonic/gin
* go.mongodb.org/mongo-driver
* golang.org/x/crypto/bcrypt
* github.com/golang-jwt/jwt/v5

---

# **Part 2: Task Management API**

### **Database**

* **Database name:** `taskdb`
* **Collection name:** `tasks`

**Task document structure:**

```json
{
  "_id": "ObjectId",
  "title": "string",
  "description": "string",
  "completed": true,
  "duedate": "ISODate"
}
```

---

### **Endpoints**

#### **1. Get All Tasks**

* **Endpoint:** `GET /tasks`
* **Description:** Retrieves a list of all tasks.

**Example cURL:**

```bash
curl --location 'http://localhost:8080/tasks'
```

---

#### **2. Get Task by ID**

* **Endpoint:** `GET /tasks/:id`
* **Description:** Fetches a specific task by ID.

**Example cURL:**

```bash
curl --location 'http://localhost:8080/tasks/2'
```

---

#### **3. Create Task**

* **Endpoint:** `POST /tasks`
* **Description:** Creates a new task.

**Request Body:**

```json
{
  "Title": "complete task 4",
  "Description": "on the way of completion",
  "Duedate": "2025-07-18T08:30:00Z",
  "Status": false
}
```

**Example cURL:**

```bash
curl --location --request POST 'http://localhost:8080/tasks' \
--header 'Content-Type: application/json' \
--data-raw '{
  "Title": "complete task 4",
  "Description": "on the way of completion",
  "Duedate": "2025-07-18T08:30:00Z",
  "Status": false
}'
```

---

#### **4. Update Task**

* **Endpoint:** `PUT /tasks/:id`
* **Description:** Updates an existing task by its ID.

**Request Body:**

```json
{
  "ID": "2",
  "Title": "complete task 4",
  "Description": "on the way of completion",
  "Duedate": "2025-07-18T08:30:00Z",
  "Status": false
}
```

**Example cURL:**

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

---

#### **5. Delete Task**

* **Endpoint:** `DELETE /tasks/:id`
* **Description:** Deletes a task by its ID.

**Example cURL:**

```bash
curl --location --request DELETE 'http://localhost:8080/tasks/4'
```

---

### **Common Status Codes**

| Code | Scenario                          |
| ---- | --------------------------------- |
| 200  | Success.                          |
| 201  | Resource created.                 |
| 400  | Bad request / validation error.   |
| 404  | Not found.                        |
| 500  | Internal server / database error. |

---

## **Author & Maintainer**

Abrham Yishak

---

