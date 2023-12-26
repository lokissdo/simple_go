# My Go App Name

## Prerequisites
Before running the app, ensure you have the following prerequisites installed:
- Go
- Bash
- MongoDB Compass

Make sure MongoDB is running locally, and you can configure the database connection in `config/database.go`.

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/lokissdo/simple_go.git
   cd my_app

2. Run the server:
    ```bash
    go run main.go
    Migrate data to your MongoDB:

3. Migrate data to your MongoDB:
    ```bash
    curl -X POST -H "Content-Type: application/json" -d @mock_data.json http://127.0.0.1:8080/products
4. Launch the website:
Open your browser and navigate to `http://127.0.0.1:8080`