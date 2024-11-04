# GoGrub

GoGrub is a web application where users can create profiles, save favorite recipes, and browse through a growing collection of user-generated content. Built with Go, SQLite, and standard web technologies, this app serves as a personal recipe keeper and social food-sharing platform.

## Getting Started

Follow the instructions below to get GoGrub running locally on your machine.

### Prerequisites

- **Go**: Make sure you have Go installed (version 1.20 or later recommended).
- **SQLite**: The app uses SQLite for the database, and dependencies are managed using Go modules.

### Installation

1. **Clone the Repository**

   Start by cloning the repository and navigating to the project folder:

   ```bash
   git clone https://github.com/alliecostner27/API-Project-PGL.git
   cd API-Project-PGL
   ```

2. **Pull the Latest Changes**

   Ensure you are on the `main` branch and have the latest updates:

   ```bash
   git fetch origin
   git pull origin main
   ```

3. **Install Dependencies**

   Run the following command to install any missing dependencies:

   ```bash
   go mod tidy
   ```

### Running the Application

After setting up, you can run the app with:

```bash
go run .
```

This will start the server on `http://localhost:8000`.

### Usage

1. Open a web browser and navigate to [http://localhost:8000](http://localhost:8000).
2. You should see the GoGrub website, where you can sign up, log in, and start using the app.

### Project Structure

- **`main.go`**: Entry point of the application.
- **`db/`**: Contains database initialization and query logic.
- **`public/`**: Stores static assets like images and stylesheets.
- **`views/`**: Contains HTML templates for the website.

### Troubleshooting

If you encounter issues with dependencies or database connections, ensure:
- Your Go version is up to date.
- You have followed the steps above without any missing configuration.

### Contributing

Contributions are welcome! Please fork the repository and make a pull request with your changes.

---

Enjoy cooking and sharing with GoGrub! 🍲
