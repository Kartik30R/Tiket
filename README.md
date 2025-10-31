# 🎟️ Ticket Booking API — QR-Based Validation System

A modular Go backend for managing **events**, **tickets**, and **QR-based entry validation**.

---

## 🚀 Endpoints

### **POST /api/auth/register**

**Description:** Register a new user

**Request Body:**

```json
{
  "email": "john@example.com",
  "password": "strongpassword123"
}
```

**Response:**

```json
{
  "status": "success",
  "message": "Successfully registered",
  "data": {
    "token": "<JWT_TOKEN>",
    "user": {
      "id": 1,
      "email": "john@example.com",
      "role": "manager"
    }
  }
}
```

---

### **POST /api/auth/login**

**Description:** Log in an existing user

**Request Body:**

```json
{
  "email": "john@example.com",
  "password": "strongpassword123"
}
```

**Response:**

```json
{
  "status": "success",
  "message": "Successfully logged in",
  "data": {
    "token": "<JWT_TOKEN>",
    "user": {
      "id": 1,
      "email": "john@example.com",
      "role": "manager"
    }
  }
}
```

---

### **GET /api/event/**

**Description:** Fetch all events

**Headers:**

```
Authorization: Bearer <JWT_TOKEN>
```

**Response:**

```json
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "name": "TechFest 2025",
      "location": "Delhi",
      "date": "2025-12-25T19:00:00Z"
    }
  ]
}
```

---

### **POST /api/event/**

**Description:** Create a new event

**Headers:**

```
Authorization: Bearer <JWT_TOKEN>
```

**Request Body:**

```json
{
  "name": "TechFest 2025",
  "location": "Delhi",
  "date": "2025-12-25T19:00:00Z"
}
```

**Response:**

```json
{
  "status": "success",
  "message": "Event created successfully",
  "data": {
    "id": 1,
    "name": "TechFest 2025",
    "location": "Delhi",
    "date": "2025-12-25T19:00:00Z"
  }
}
```

---

### **GET /api/event/:eventId**

**Description:** Fetch event by ID

**Example:** `/api/event/1`

**Response:**

```json
{
  "status": "success",
  "data": {
    "id": 1,
    "name": "TechFest 2025",
    "location": "Delhi",
    "date": "2025-12-25T19:00:00Z"
  }
}
```

---

### **PUT /api/event/:eventId**

**Description:** Update event details

**Example:** `/api/event/1`

**Request Body:**

```json
{
  "location": "Mumbai"
}
```

**Response:**

```json
{
  "status": "success",
  "message": "Event updated successfully",
  "data": {
    "id": 1,
    "name": "TechFest 2025",
    "location": "Mumbai",
    "date": "2025-12-25T19:00:00Z"
  }
}
```

---

### **DELETE /api/event/:eventId**

**Description:** Delete event by ID

**Example:** `/api/event/1`

**Response:**

```json
{
  "status": "success",
  "message": "Event deleted successfully"
}
```

---

### **GET /api/ticket/**

**Description:** Fetch all tickets for logged-in user

**Headers:**

```
Authorization: Bearer <JWT_TOKEN>
```

**Response:**

```json
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "eventId": 1,
      "userId": 1,
      "qrCode": "<QR_BASE64>"
    }
  ]
}
```

---

### **POST /api/ticket/**

**Description:** Create a new ticket for a specific event

**Request Body:**

```json
{
  "eventId": 1
}
```

**Response:**

```json
{
  "status": "success",
  "message": "Ticket created successfully",
  "data": {
    "id": 1,
    "eventId": 1,
    "userId": 1,
    "qrCode": "<QR_BASE64>"
  }
}
```

---

### **GET /api/ticket/:ticketId**

**Description:** Get ticket by ID with QR code

**Example:** `/api/ticket/1`

**Response:**

```json
{
  "status": "success",
  "data": {
    "id": 1,
    "eventId": 1,
    "userId": 1,
    "qrCode": "<QR_BASE64>"
  }
}
```

---

### **POST /api/ticket/validate**

**Description:** Validate ticket entry (QR-based)

**Request Body:**

```json
{
  "ticketId": 1
}
```

**Response:**

```json
{
  "status": "success",
  "message": "Welcome to the show!",
  "data": {
    "ticketId": 1,
    "validatedAt": "2025-10-31T19:00:00Z"
  }
}
```

---

### ✅ **Common Header for Protected Routes**

```
Authorization: Bearer <JWT_TOKEN>
```

---

## ⚙️ Features

* 🔐 **JWT Authentication** — Secure login & registration
* 🎟️ **Event & Ticket Management** — Full CRUD for events and tickets
* 📦 **Modular Go Architecture** — Clean handlers and services separation
* 🧾 **QR-Based Validation** — Scan & verify tickets instantly
 * 📊 **RESTful API Design** — Structured request/response patterns

---

## 💻 Tech Stack

* **Go (Gin Framework)** — High-performance web framework
* **GORM** — ORM for PostgreSQL
* **JWT** — Authentication and authorization
*  * **PostgreSQL** — Primary database
