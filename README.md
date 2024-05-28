
# go-openai: Chat with an AI Assistant using Golang and OpenAI

## Description

This Go project empowers you to interact with an AI assistant in a chat thread format. It seamlessly integrates the OpenAI API to provide chat completion, fostering engaging and informative conversations.

##  Installation

**Prerequisites:**

- Ensure you have Git and Golang installed on your system.
- Replace `OPENAI_API_KEY` with your actual OpenAI API key in the project's configuration file (refer to OpenAI documentation for obtaining an API key).
- Consider using a `.env` file to securely store sensitive information like your API key, preventing accidental exposure in version control.
- You can customize the chat experience by configuring the OpenAI model used for chat completion and other parameters available in the OpenAI API.

**Steps:**

1. **Clone the Repository:**
   ```bash
   git clone git@github.com:pi-prakhar/go-openai.git
   ```

2. **Navigate to the Project Directory:**
   ```bash
   cd go-openai
   ```

**Running the Project**

**A. Using Docker**

1. Build and start the application in detached mode:
   ```bash
   docker-compose up --build -d
   ```

**B. Running Locally (Windows)**

1. Build the executable:
   ```bash
   go build -o main.exe ./cmd/go-openai
   ```

2. Run the application:
   ```bash
   ./main.exe
   ```

## Accessing the Application

Once the application is running, launch it in your web browser using the following URL:

```
http://localhost:8000/
``` 

## Getting Started

For detailed instructions on configuration and further customization, consult the project's source code and any additional documentation provided within the repository.


## API Endpoints

* **`/api/chat/send` (POST):** Send messages to the AI assistant. The request body must be in JSON format with a `message` field.

* **`/api/chat/messages` (GET):** Retrieve all conversation messages exchanged between you and the AI assistant during the current session.

The application provides an API endpoint for sending messages to the AI assistant:

**Endpoint:** `/api/chat/send`
**Method:** POST
**Body:** JSON (containing a `message` field)

**JSON Example:**

```json
{
  "message": "Hi! How can I help you today?"
}
```

The response from the AI assistant will be displayed in the chat thread format using:

**Endpoint:** `/api/chat/messages`
**Method:** POST
**Body:** JSON (containing a `message` field)

**API Response Formats**

* **Success Response (All Messages):**

  ```json
  {
      "code": 200,  // HTTP status code for success
      "message": "Successfully fetched all messages",
      "data": [     // Required data in the response
          {
              "role": "user",
              "content": "What new model of chat gpt is coming?"
          },
          {
              "role": "assistant",
              "content": "OpenAI recently announced the release of GPT-4, the latest iteration of their chatbot model. GPT-4 is said to have even more advanced capabilities, including improved natural language processing and the ability to generate more coherent and contextually relevant responses. This new model is expected to further push the boundaries of AI-powered conversational agents."
          }
      ]
  }
  ```

* **Error Response Schema:**

  ```json
  {
      "code": 400,  // Example error code (can vary)
      "message": "Bad request"  // Human-readable error description
  }
  ```


**Please refer to the project's source code for the exact error response format and details on the specific error codes used.**
