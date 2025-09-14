# 🛒 Market Management System

A modular, microservices-based platform designed to streamline **supermarket and retail operations**. The system provides comprehensive management for products, inventory, debts, and analytics while leveraging modern technologies like gRPC, Kafka, PostgreSQL, MongoDB, MinIO, WebSocket and Swagger for documentation.

---

## ✨ Key Features

### 📦 Product Service
Handles all product-related operations including CRUD and stock management.

- **CreateProduct** – Add a new product to the inventory.  
- **GetProductById** – Retrieve product details by ID.  
- **GetProductByFilter** – Retrieve products using filters (category, price, color, etc.).  
- **UpdateStock** – Update the quantity of products in stock.  
- **UpdateProduct** – Modify product attributes such as name, price, or color.  
- **DeleteProduct** – Remove a product from the inventory.  

### 💰 Debt Service
Manages customer debts, payments, and financial tracking.

- **CreateDebt** – Record a new debt for a customer.  
- **UpdateDebt** – Modify existing debt records.  
- **DeleteDebt** – Remove a debt record.  
- **GetDebtById** – Retrieve details of a specific debt.  
- **GetDebtByFilter** – Retrieve debts using filters like customer, status, or date.  

### 📊 Dashboard Service
Provides real-time reporting and analytics for sales, stock, and debts.

- **UpsertProductSales** – Update product sales data.  
- **GetDashboardReport** – Retrieve an aggregated dashboard report including key metrics.  

---

## 🛠️ Technology Stack

- **Backend:** Go (Golang)  
- **Communication:** gRPC, HTTPS  
- **Databases:** PostgreSQL, MongoDB  
- **Messaging:** Apache Kafka, WebSocket
- **File Storage:** MinIO (for product images)  
- **Logging:** Logrus  
- **API Documentation:** Swagger  

---