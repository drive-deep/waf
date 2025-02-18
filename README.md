

# **Dashboard App**

Welcome to the **Dashboard App**! This app tracks API metrics and stores them in InfluxDB.

---

## **📦 Getting Started**

### **1️⃣ Prerequisites**

Ensure you have the following installed on your machine:

- **Docker**: To build and run the application with `docker-compose`
- **Git**: To clone the repository

### **2️⃣ Clone the Repository**

Clone this repository to your local machine:

```bash
git clone https://github.com/drive-deep/waf.git
cd waf/dashboard_app
```

### **3️⃣ Set Up Environment Variables**

Create a `.env` file in the `dashboard_app` directory with the following content:

```text
INFLUXDB_URL=http://localhost:8086
INFLUXDB_ORG=my-org
INFLUXDB_BUCKET=api_metrics
```

### **4️⃣ Start the App Using Docker Compose**

Build and start the app with the following command:

```bash
docker-compose up --build
```

This will:

- Build the app and required containers
- Start **InfluxDB**, **Nginx**, and the **Go Web App**

The app will be available at `http://localhost:8090`.

---

## **🧪 Test the API**

Once the app is running, you can test the API and check the metrics.

### **1️⃣ Make an API Request**

Use `curl` to make an API request to the app:

```bash
curl http://localhost:8090/
```

You should receive a JSON response like:

```json
[
  {
    "APIEndpoint": "/orders",
    "HitsToday": 120,
    "HitsWeek": 800
  },
  {
    "APIEndpoint": "/users",
    "HitsToday": 90,
    "HitsWeek": 600
  }
]
```

### **2️⃣ Verify Metrics in InfluxDB**

To verify the metrics, you can query the InfluxDB instance. You can use the InfluxDB UI or CLI to check the data in the `api_metrics` bucket.

---

## **📄 License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Let me know if you need more details! 🚀