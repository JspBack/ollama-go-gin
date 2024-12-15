# Ollama-Go

Ollama-Go is an AI-powered application that provides chat and image processing functionalities using Golang, Docker, and Ollama models.

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install)

## Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/jspback/ollama-go-gin.git
   cd ollama-go-gin
   ```

2. **Build and run the services:**

   ```bash
   docker-compose up --build -d
   ```

   This command builds the Docker images and starts the services defined in `docker-compose.yml`.

## Services

- **Backend Service:** Runs on port `8000`.
- **Ollama Service:** Runs on port `11434`, hosting the AI models.

## API Endpoints

### Chat API

- **Endpoint:** `/api/chat`
- **Method:** `POST`
- **Description:** Send a message to the AI model and receive a response.
- **Parameters:**
  - `message` (string): The message to send.
  - `model` (string): The model to use (e.g., `llama3.2`).
  - `stream` (bool): Set to `true` to stream the response.

### Image API

- **Endpoint:** `/api/image`
- **Method:** `POST`
- **Description:** Process an image with the AI model.
- **Parameters:**
  - `message` (string): The message or prompt.
  - `image` (file): The image file to process.
  - `model` (string): The model to use (e.g., `llama3.2-vision`).
  - `stream` (bool): Set to `true` to stream the response.

### Health Check

- **Endpoint:** `/api/health`
- **Method:** `GET`
- **Description:** Check the health status of the application.

## Notes

- The Ollama service pulls the `llama3.2` and `llama3.2-vision` models on startup.
- Ensure your system has appropriate resources allocated for Docker, especially if using GPU acceleration.
- The `docker-compose.yml` is configured to use an NVIDIA GPU. Make sure your Docker setup supports GPU sharing.
- If you want to add otehr models configure it in `ollama/ollama.sh`.

## License

DO WHATEVER YOU LIKE.
