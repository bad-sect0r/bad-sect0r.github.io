# bad-sect0r.github.io

This repository contains the source code for a simple web application that serves the `README.md` file from a GitHub repository and displays it in a styled HTML page. The application is built using Go and the Echo framework for the backend, and plain HTML, CSS, and JavaScript for the frontend.

## Note

This is just a mini app for testing and having fun. It serves as a playground for experimenting with different technologies and deployment strategies.

## Features

- Fetches and displays the `README.md` file from a specified GitHub repository.
- Supports multiple themes for displaying the markdown content.
- Provides an option to export the displayed content as a PDF.

## Deployment

This application is deployed using [Koyeb](https://www.koyeb.com/), a serverless platform that allows you to deploy applications effortlessly.

### Steps to Deploy on Koyeb

1. **Create a Koyeb Account**: If you don't have an account, sign up at [Koyeb](https://www.koyeb.com/).

2. **Create a New App**: In the Koyeb dashboard, create a new app and connect it to this GitHub repository.

3. **Configure Build and Run Commands**:

   - **Build Command**: `go build -o main .`
   - **Run Command**: `./main`

4. **Set Environment Variables**: If needed, set any environment variables required by your application.

5. **Deploy**: Click on the deploy button to start the deployment process. Koyeb will build and deploy your application.

6. **Access Your App**: Once deployed, you can access your application via the URL provided by Koyeb.

## Local Development

To run the application locally, follow these steps:

1. **Clone the Repository**:

   ```sh
   git clone https://github.com/bad-sect0r/bad-sect0r.github.io.git
   cd bad-sect0r.github.io
   ```

2. **Install Dependencies**:

   ```sh
   go mod tidy
   ```

3. **Run the Application**:

   ```sh
   go run main.go
   ```

4. **Open in Browser**: Open your browser and navigate to `http://localhost:8080` to see the application in action.
