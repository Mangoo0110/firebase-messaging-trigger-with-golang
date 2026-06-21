# Push Notification Demo (Go + Firebase Cloud Messaging)

## Project Overview

This project demonstrates how to trigger push notifications for business events such as order placement, order delivery, and other order lifecycle updates.

The backend is implemented in Go and uses Firebase Cloud Messaging (FCM) to deliver notifications to the target mobile device.

For simplicity, device tokens are stored in an in-memory cache. No database has been used in this demo.

---

# Purpose

The goal of this demo is to:

* Register mobile devices with the backend
* Store device FCM tokens
* Trigger push notifications when an order-related event occurs
* Deliver notifications to the correct user device

---

# Tech Stack

| Component          | Technology                     |
| ------------------ | ------------------------------ |
| Language           | Go                             |
| HTTP Server        | Gin                            |
| Push Notifications | Firebase Cloud Messaging (FCM) |
| Storage            | In-Memory Cache (Go Map)       |
| Mobile Client      | Flutter (FCM Token Provider)   |

---

# Architecture

```text
Flutter Mobile App
        │
        │ Register Device
        ▼
POST /devices/register
        │
        ▼
Go Backend
        │
        │ Store FCM Token
        ▼
In-Memory Token Store

------------------------------------------------

Flutter Mobile App
        │
        │ Place Order
        ▼
POST /order/place
        │
        ▼
Go Backend
        │
        │ Find Merchant Device Token
        ▼
Firebase Admin SDK
        │
        ▼
Firebase Cloud Messaging (FCM)
        │
        ▼
Target Device
        │
        ▼
Notification Tray
```

---

# Setup

## 1. Install Go

Install Go on your machine and verify the installation:

```bash
go version
```

---

## 2. Create Firebase Project

Create a Firebase project from:

https://console.firebase.google.com

---

## 3. Generate Service Account Key

Navigate to:

```text
Firebase Console
→ Project Settings
→ Service Accounts
→ Generate New Private Key
```

Download the generated JSON file and place it in the project root:

```text
project-root/
│
├── serviceAccountKey.json
├── main.go
├── go.mod
└── ...
```

> Do not commit this file to version control.

---

## 4. Install Dependencies

```bash
go get firebase.google.com/go/v4
go get github.com/gin-gonic/gin
go get google.golang.org/api/option

go mod tidy
```

---

# Available APIs

## Register Device

Registers a mobile device and stores its Firebase Cloud Messaging token.

### Endpoint

```http
POST /devices/register
```

### Request

```json
{
  "user_id": "merchant_1",
  "token": "FCM_DEVICE_TOKEN"
}
```

### Response

```json
{
  "message": "Device registered!"
}
```

---

## Place Order

Simulates an order placement event and triggers a push notification to the merchant.

### Endpoint

```http
POST /order/place
```

### Request

```json
{
  "merchant_id": "merchant_1",
  "customer_id": "customer_1",
  "order_id": "order_101"
}
```

### Response

```json
{
  "message": "Order placed and notification sent"
}
```

---

# How It Works

### Device Registration

The Flutter mobile application retrieves an FCM token from Firebase Messaging and sends it to the backend using the `devices/register` API.

Example:

```json
{
  "user_id": "merchant_1",
  "token": "FCM_DEVICE_TOKEN"
}
```

The backend stores the token using the user identifier:

```go
deviceTokens["merchant_1"] = "FCM_DEVICE_TOKEN"
```

---

### Order Placement

When the backend receives an order placement request:

```json
{
  "merchant_id": "merchant_1",
  "customer_id": "customer_1",
  "order_id": "order_101"
}
```

it:

1. Finds the merchant's registered FCM token
2. Creates a notification payload
3. Sends the notification using Firebase Admin SDK
4. Firebase Cloud Messaging delivers the notification to the target device

---

# Notification Delivery

A successful response from:

```go
fcmClient.Send(...)
```

returns a Firebase Message ID.

This indicates that Firebase Cloud Messaging has accepted the message for delivery.

Example:

```text
projects/project-id/messages/0:1750434956439502
```

Note that acceptance by Firebase does not guarantee that the user has already seen the notification. Actual delivery depends on:

* Device connectivity
* Operating system behavior
* Notification permissions
* Device state (foreground/background/terminated)

---

# Limitations

This demo intentionally uses an in-memory cache:

```go
var deviceTokens = map[string]string{}
```

As a result:

* Tokens are lost when the server restarts
* Tokens are not persisted
* Multiple backend instances cannot share token data

In a production environment, device tokens should be stored in a database such as:

* PostgreSQL
* MySQL
* MongoDB
* Redis

---

# Future Improvements

* Persistent database storage
* User authentication
* Notification history
* Retry mechanism for failed notifications
* Topic-based notifications
* Batch notification support
* Notification analytics and delivery tracking

---

# Demo Flow

1. Flutter app obtains FCM token.
2. Flutter app calls `POST /devices/register`.
3. Backend stores the token.
4. Order is placed via `POST /order/place`.
5. Backend retrieves merchant token.
6. Backend sends notification through Firebase Cloud Messaging.
7. Notification appears on the merchant's device.
