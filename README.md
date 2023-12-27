# Simple Go

## Prerequisites
Before running the app, ensure you have the following prerequisites installed:
- Go
- Bash (required for local MongoDB)
- MongoDB Compass (required for local MongoDB)

Make sure MongoDB is running locally, and you can configure the database connection in `config/database.go`.

## Getting Started


1. Clone the repository:
   ```bash
   git clone https://github.com/lokissdo/simple_go.git
   cd simple_go
   cd my_app

2. Run the server:
    ```bash
    go run main.go

3. Migrate data to your MongoDB:
    ```bash
    curl -X POST -H "Content-Type: application/json" -d @mock_data.json http://127.0.0.1:8080/products
    ```
   Here, since my MongoDB is hosted on a cloud cluster, I migrated the data earlier. If you want to run it locally, just uncomment the `clientOptions` for the local configuration and comment out my `clientOptions` in  `config/database.go`

4. Launch the website:
Open your browser and navigate to `http://127.0.0.1:8080`