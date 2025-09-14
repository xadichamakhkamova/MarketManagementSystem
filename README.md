# 🛒 Market Management System

A modular, **microservices-based platform** designed to streamline **supermarket and retail operations**.  
The system provides management for **products**, **inventory**, **debts**, and **analytics**, leveraging modern technologies like **gRPC**, **Kafka**, **PostgreSQL**, **MongoDB**, **MinIO**, **WebSocket**, and **Swagger**.

---

## ✨ Key Features

### 📦 Product Service
Handles all product-related operations including CRUD and stock management.

- ➕ **CreateProduct** – Add a new product to the inventory  
- 🔍 **GetProductById** – Retrieve product details by ID  
- 🎯 **GetProductByFilter** – Filter products (category, price, color, etc.)  
- 📉 **UpdateStock** – Update product quantity in stock  
- ✏️ **UpdateProduct** – Modify product attributes (name, price, color)  
- 🗑️ **DeleteProduct** – Remove a product from inventory  

---

### 💰 Debt Service
Manages customer debts, payments, and financial tracking.

- ➕ **CreateDebt** – Record a new customer debt  
- ✏️ **UpdateDebt** – Modify existing debt records  
- 🗑️ **DeleteDebt** – Remove a debt record  
- 🔍 **GetDebtById** – Retrieve details of a specific debt  
- 🎯 **GetDebtByFilter** – Filter debts (customer, status, date)  

---

### 📊 Dashboard Service
Provides **real-time reporting and analytics** for sales, stock, and debts.

- 📈 **UpsertProductSales** – Update product sales data  
- 📊 **GetDashboardReport** – Retrieve an aggregated dashboard report with key metrics  

---

## 🛠️ Technology Stack

- 🐹 **Backend:** Go (Golang)  
- 🔗 **Communication:** gRPC, HTTPS  
- 🗄️ **Databases:** PostgreSQL, MongoDB  
- 📩 **Messaging:** Apache Kafka, WebSocket  
- 🖼️ **File Storage:** MinIO (for product images)  
- 📝 **Logging:** Logrus  
- 📑 **API Documentation:** Swagger  

---
