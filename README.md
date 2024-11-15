# Go API Backend

This is a simple API backend built with Go that provides user authentication and manages users, events and registrations. The application uses SQLite for database and [Gin Gonic](https://gin-gonic.com/)) as the web framework.

## Features

- User Authentication: Secure login and registration for users.
- Event Management: Create, view, update, and delete events.
- Event Registration: Users can register for events.
- Built with Go: Lightweight and efficient backend design.
- SQLite Database: Easy-to-use, file-based relational database.

## Technologies Used

- Go: A statically typed, compiled programming language known for its simplicity and performance.
- SQLite: A lightweight relational database for storing data.
- Gin Gonic: A high-performance HTTP web framework for Go.

## Getting Started

### Prerequisites

Make sure you have the following installed on your machine:

- Go (install from [gin-gonic](https://gin-gonic.com/))
- SQLite (Pre-installed on most operating systems)


### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/your-repo-name.git
   cd your-repo-name

2. Install dependencies
   - go mod tidy
    
3. Run the application:
  - go run main.go

4. The API will be available at http://localhost:8080

   

# API Endpoints

## Authentication
- POST /signup: Register a new user
- POST /login: Log in an existing user

## Events
- GET /events: Retrieve a list of all events
- POST /events: Add a new event (needs authentication)
- GET /events/{id}: Retrieve a specific event by ID
- PUT /events/{id}: Update a specific event by ID (needs authentication)
- DELETE /events/{id}: Delete a specific event by ID (needs authentication)
- POST /events/{id}/register: Registers a specific event for a user (needs authentication)
- DELETE /events/{id}/register: Cancels registration of a sepcific event for a user (needs authentication)
