# **📊 Dashboard App**  

Welcome to the **Dashboard App**! This application tracks API metrics and stores them in **InfluxDB**, providing real-time insights into API usage.  

---

## **🚀 Getting Started**  

### **1️⃣ Prerequisites**  

Ensure you have the following installed on your system:  

- **Docker** – To build and run the application using `docker-compose`  
- **Git** – To clone the repository  

---

### **2️⃣ Clone the Repository**  

Clone this repository to your local machine and navigate to the project directory:  

```bash
git clone https://github.com/drive-deep/waf.git
cd waf/dashboard_app
```

---

### **3️⃣ Configure Environment Variables**  
update your `/etc/hosts` file with the following entries:  

```text
127.0.0.1 osto-assignment.com
127.0.0.1 www.osto-assignment.com
```

This ensures local domain resolution for API requests.

---

### **4️⃣ Start the Application**  

Run the following command to build and start the app:  

```bash
docker-compose up --build
```

This command will:  

✔ Build the application and required containers  
✔ Start **InfluxDB**, **Nginx**, and the **Go Web App**  

Once the setup is complete, the app will be accessible at:  

👉 **`http://localhost:8090`**  

---

## **🛠 Testing the API**  

Once the app is running, you can test its functionality and verify the metrics being recorded.  

### **1️⃣ Make API Requests**  

You can test the API by making requests to different endpoints:  

```bash
curl -i http://osto-assignment.com/users
```
✔ Returns `application/json` in response headers  

```bash
curl -i http://osto-assignment.com/orders
```
✔ Returns `application/json` in response headers  

```bash
curl -i http://osto-assignment.com/hello
```
✔ Returns `application/txt` in response headers  

---

### **2️⃣ Fetch Metrics from the API**  

Retrieve API usage metrics with:  duration is in second

```bash
curl -i "http://localhost:8090/?interval=5000"
```

📌 **Example Response:**  

```json
{
  "data": [
    {"api_endpoint": "/orders", "hits": 1},
    {"api_endpoint": "/users", "hits": 1}
  ],
  "timeframe": "5000"
}
```

To get daily and weekly aggregated metrics:  

```bash
curl -i "http://localhost:8090/"
```

📌 **Example Response:**  

```json
{
  "day": [
    {"api_endpoint": "/orders", "hits": 3},
    {"api_endpoint": "/users", "hits": 4}
  ],
  "week": [
    {"api_endpoint": "/orders", "hits": 3},
    {"api_endpoint": "/users", "hits": 4}
  ]
}
```

---

### **3️⃣ Verify Metrics in InfluxDB**  

You can check the recorded API metrics in **InfluxDB** using either:  

- The **InfluxDB UI**  
- The **InfluxDB CLI**  

Query the `api_metrics` bucket to see real-time data.

---

## **📜 License**  

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.  

---

💡 **Need Help?** If you run into any issues, feel free to reach out! 🚀  