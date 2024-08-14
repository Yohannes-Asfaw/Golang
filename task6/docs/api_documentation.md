# Task Management API Documentation

## MongoDB Integration
- MongoDB is used as the persistent data storage for tasks.
- The application connects to MongoDB using the MongoDB Go Driver.

### Configuration
- Connection URI: `mongodb://localhost:27017`
- Database: `task_manager`
- Collection: `tasks`

## Endpoints

### GET /tasks
- Retrieves a list of all tasks.

### GET /tasks/:id
- Retrieves details of a specific task by its ID.

### POST /tasks
- Creates a new task.
- **Request Body**:
  ```json
  {
    "title": "Task title",
    "description": "Task description",
    "due_date": "YYYY-MM-DD",
    "status": "pending"
  }