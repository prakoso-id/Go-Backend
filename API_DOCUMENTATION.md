# Go Backend API Documentation

This document provides comprehensive documentation for the REST API endpoints available in the Go Backend project.

## Base URL

All API endpoints are prefixed with `/api`.

## Authentication

The API uses JWT (JSON Web Token) for authentication. Protected routes require a valid access token.

### Authentication Endpoints

#### Login

- **URL**: `/api/auth/login`
- **Method**: `POST`
- **Auth Required**: No
- **Description**: Authenticates a user and returns access and refresh tokens.
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "user": {
        "id": "uuid",
        "email": "user@example.com",
        "name": "User Name"
      }
    }
    ```

#### Logout

- **URL**: `/api/auth/logout`
- **Method**: `POST`
- **Auth Required**: Yes (Access Token)
- **Description**: Invalidates the current user's tokens.
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "message": "Successfully logged out"
    }
    ```

#### Reset Password

- **URL**: `/api/auth/reset-password`
- **Method**: `POST`
- **Auth Required**: No
- **Description**: Initiates password reset process.
- **Request Body**:
  ```json
  {
    "email": "user@example.com"
  }
  ```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "message": "Password reset instructions sent to email"
    }
    ```

## User Endpoints

### List Users

- **URL**: `/api/users`
- **Method**: `GET`
- **Auth Required**: No
- **Description**: Returns a list of all users.
- **Query Parameters**:
  - `page` (optional): Page number for pagination
  - `limit` (optional): Number of items per page
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "data": [
        {
          "id": "uuid1",
          "name": "User 1",
          "email": "user1@example.com"
        },
        {
          "id": "uuid2",
          "name": "User 2",
          "email": "user2@example.com"
        }
      ],
      "pagination": {
        "total": 100,
        "page": 1,
        "limit": 10
      }
    }
    ```

### Create User

- **URL**: `/api/users`
- **Method**: `POST`
- **Auth Required**: Yes (Access Token)
- **Description**: Creates a new user.
- **Request Body**:
  ```json
  {
    "name": "New User",
    "email": "newuser@example.com",
    "password": "password123"
  }
  ```
- **Success Response**:
  - **Code**: 201 Created
  - **Content**:
    ```json
    {
      "id": "uuid",
      "name": "New User",
      "email": "newuser@example.com",
      "created_at": "2023-01-01T00:00:00Z"
    }
    ```

### Get User by ID

- **URL**: `/api/users/:id`
- **Method**: `GET`
- **Auth Required**: Yes (Access Token)
- **Description**: Returns details of a specific user.
- **URL Parameters**:
  - `id`: User ID
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "id": "uuid",
      "name": "User Name",
      "email": "user@example.com",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
    ```

### Update User

- **URL**: `/api/users/:id`
- **Method**: `PUT`
- **Auth Required**: Yes (Access Token)
- **Description**: Updates an existing user.
- **URL Parameters**:
  - `id`: User ID
- **Request Body**:
  ```json
  {
    "name": "Updated Name",
    "email": "updated@example.com"
  }
  ```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "id": "uuid",
      "name": "Updated Name",
      "email": "updated@example.com",
      "updated_at": "2023-01-02T00:00:00Z"
    }
    ```

### Delete User

- **URL**: `/api/users/:id`
- **Method**: `DELETE`
- **Auth Required**: Yes (Access Token)
- **Description**: Deletes a user.
- **URL Parameters**:
  - `id`: User ID
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "message": "User successfully deleted"
    }
    ```

## Post Endpoints

### List Posts

- **URL**: `/api/posts`
- **Method**: `GET`
- **Auth Required**: No
- **Description**: Returns a list of all posts.
- **Query Parameters**:
  - `page` (optional): Page number for pagination
  - `limit` (optional): Number of items per page
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "data": [
        {
          "id": "uuid1",
          "title": "Post Title 1",
          "content": "Post content...",
          "user_id": "user_uuid1",
          "created_at": "2023-01-01T00:00:00Z"
        },
        {
          "id": "uuid2",
          "title": "Post Title 2",
          "content": "Post content...",
          "user_id": "user_uuid2",
          "created_at": "2023-01-02T00:00:00Z"
        }
      ],
      "pagination": {
        "total": 50,
        "page": 1,
        "limit": 10
      }
    }
    ```

### Get Post by ID

- **URL**: `/api/posts/:id`
- **Method**: `GET`
- **Auth Required**: No
- **Description**: Returns details of a specific post.
- **URL Parameters**:
  - `id`: Post ID
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "id": "uuid",
      "title": "Post Title",
      "content": "Post content...",
      "user_id": "user_uuid",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
    ```

### List Posts by User ID

- **URL**: `/api/posts/user/:user_id`
- **Method**: `GET`
- **Auth Required**: No
- **Description**: Returns all posts by a specific user.
- **URL Parameters**:
  - `user_id`: User ID
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "data": [
        {
          "id": "uuid1",
          "title": "Post Title 1",
          "content": "Post content...",
          "user_id": "user_uuid",
          "created_at": "2023-01-01T00:00:00Z"
        },
        {
          "id": "uuid2",
          "title": "Post Title 2",
          "content": "Post content...",
          "user_id": "user_uuid",
          "created_at": "2023-01-02T00:00:00Z"
        }
      ],
      "pagination": {
        "total": 20,
        "page": 1,
        "limit": 10
      }
    }
    ```

### Create Post

- **URL**: `/api/posts`
- **Method**: `POST`
- **Auth Required**: Yes (Access Token)
- **Description**: Creates a new post.
- **Request Body**:
  ```json
  {
    "title": "New Post Title",
    "content": "Post content..."
  }
  ```
- **Success Response**:
  - **Code**: 201 Created
  - **Content**:
    ```json
    {
      "id": "uuid",
      "title": "New Post Title",
      "content": "Post content...",
      "user_id": "user_uuid",
      "created_at": "2023-01-01T00:00:00Z"
    }
    ```

### Update Post

- **URL**: `/api/posts/:id`
- **Method**: `PUT`
- **Auth Required**: Yes (Access Token)
- **Description**: Updates an existing post. User can only update their own posts.
- **URL Parameters**:
  - `id`: Post ID
- **Request Body**:
  ```json
  {
    "title": "Updated Post Title",
    "content": "Updated post content..."
  }
  ```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "id": "uuid",
      "title": "Updated Post Title",
      "content": "Updated post content...",
      "user_id": "user_uuid",
      "updated_at": "2023-01-02T00:00:00Z"
    }
    ```

### Delete Post

- **URL**: `/api/posts/:id`
- **Method**: `DELETE`
- **Auth Required**: Yes (Access Token)
- **Description**: Deletes a post. User can only delete their own posts.
- **URL Parameters**:
  - `id`: Post ID
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "message": "Post successfully deleted"
    }
    ```

## Project Endpoints

### List Projects

- **URL**: `/api/projects`
- **Method**: `GET`
- **Auth Required**: No
- **Description**: Returns a list of all projects.
- **Query Parameters**:
  - `page` (optional): Page number for pagination
  - `limit` (optional): Number of items per page
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "data": [
        {
          "id": "uuid1",
          "name": "Project 1",
          "description": "Project description...",
          "user_id": "user_uuid1",
          "created_at": "2023-01-01T00:00:00Z"
        },
        {
          "id": "uuid2",
          "name": "Project 2",
          "description": "Project description...",
          "user_id": "user_uuid2",
          "created_at": "2023-01-02T00:00:00Z"
        }
      ],
      "pagination": {
        "total": 30,
        "page": 1,
        "limit": 10
      }
    }
    ```

### Get Project by ID

- **URL**: `/api/projects/:id`
- **Method**: `GET`
- **Auth Required**: No
- **Description**: Returns details of a specific project.
- **URL Parameters**:
  - `id`: Project ID
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "id": "uuid",
      "name": "Project Name",
      "description": "Project description...",
      "user_id": "user_uuid",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
    ```

### Create Project

- **URL**: `/api/projects`
- **Method**: `POST`
- **Auth Required**: Yes (Access Token)
- **Description**: Creates a new project.
- **Request Body**:
  ```json
  {
    "name": "New Project",
    "description": "Project description..."
  }
  ```
- **Success Response**:
  - **Code**: 201 Created
  - **Content**:
    ```json
    {
      "id": "uuid",
      "name": "New Project",
      "description": "Project description...",
      "user_id": "user_uuid",
      "created_at": "2023-01-01T00:00:00Z"
    }
    ```

### Update Project

- **URL**: `/api/projects/:id`
- **Method**: `PUT`
- **Auth Required**: Yes (Access Token)
- **Description**: Updates an existing project. User can only update their own projects.
- **URL Parameters**:
  - `id`: Project ID
- **Request Body**:
  ```json
  {
    "name": "Updated Project Name",
    "description": "Updated project description..."
  }
  ```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "id": "uuid",
      "name": "Updated Project Name",
      "description": "Updated project description...",
      "user_id": "user_uuid",
      "updated_at": "2023-01-02T00:00:00Z"
    }
    ```

### Delete Project

- **URL**: `/api/projects/:id`
- **Method**: `DELETE`
- **Auth Required**: Yes (Access Token)
- **Description**: Deletes a project. User can only delete their own projects.
- **URL Parameters**:
  - `id`: Project ID
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "message": "Project successfully deleted"
    }
    ```

## Profile Endpoints

### List Profiles

- **URL**: `/api/profiles`
- **Method**: `GET`
- **Auth Required**: No
- **Description**: Returns a list of all profiles.
- **Query Parameters**:
  - `page` (optional): Page number for pagination
  - `limit` (optional): Number of items per page
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "data": [
        {
          "id": "uuid1",
          "user_id": "user_uuid1",
          "bio": "User bio...",
          "avatar_url": "https://example.com/avatar1.jpg",
          "created_at": "2023-01-01T00:00:00Z"
        },
        {
          "id": "uuid2",
          "user_id": "user_uuid2",
          "bio": "User bio...",
          "avatar_url": "https://example.com/avatar2.jpg",
          "created_at": "2023-01-02T00:00:00Z"
        }
      ],
      "pagination": {
        "total": 100,
        "page": 1,
        "limit": 10
      }
    }
    ```

### Get Profile by ID

- **URL**: `/api/profiles/:id`
- **Method**: `GET`
- **Auth Required**: No
- **Description**: Returns details of a specific profile.
- **URL Parameters**:
  - `id`: Profile ID
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "id": "uuid",
      "user_id": "user_uuid",
      "bio": "User bio...",
      "avatar_url": "https://example.com/avatar.jpg",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
    ```

### Create Profile

- **URL**: `/api/profiles`
- **Method**: `POST`
- **Auth Required**: Yes (Access Token)
- **Description**: Creates a new profile.
- **Request Body**:
  ```json
  {
    "bio": "User bio...",
    "avatar_url": "https://example.com/avatar.jpg"
  }
  ```
- **Success Response**:
  - **Code**: 201 Created
  - **Content**:
    ```json
    {
      "id": "uuid",
      "user_id": "user_uuid",
      "bio": "User bio...",
      "avatar_url": "https://example.com/avatar.jpg",
      "created_at": "2023-01-01T00:00:00Z"
    }
    ```

### Update Profile

- **URL**: `/api/profiles/:id`
- **Method**: `PUT`
- **Auth Required**: Yes (Access Token)
- **Description**: Updates an existing profile. User can only update their own profile.
- **URL Parameters**:
  - `id`: Profile ID
- **Request Body**:
  ```json
  {
    "bio": "Updated user bio...",
    "avatar_url": "https://example.com/new-avatar.jpg"
  }
  ```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "id": "uuid",
      "user_id": "user_uuid",
      "bio": "Updated user bio...",
      "avatar_url": "https://example.com/new-avatar.jpg",
      "updated_at": "2023-01-02T00:00:00Z"
    }
    ```

### Delete Profile

- **URL**: `/api/profiles/:id`
- **Method**: `DELETE`
- **Auth Required**: Yes (Access Token)
- **Description**: Deletes a profile. User can only delete their own profile.
- **URL Parameters**:
  - `id`: Profile ID
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "message": "Profile successfully deleted"
    }
    ```

## Health Check Endpoint

### Health Check

- **URL**: `/api/health`
- **Method**: `GET`
- **Auth Required**: No
- **Description**: Returns the health status of the API.
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "status": "ok",
      "version": "1.0.0",
      "timestamp": "2023-01-01T00:00:00Z"
    }
    ```

## Error Responses

All endpoints may return the following error responses:

### Bad Request

- **Code**: 400 Bad Request
- **Content**:
  ```json
  {
    "error": "Bad Request",
    "message": "Invalid request parameters",
    "details": ["Field 'email' is required"]
  }
  ```

### Unauthorized

- **Code**: 401 Unauthorized
- **Content**:
  ```json
  {
    "error": "Unauthorized",
    "message": "Authentication required"
  }
  ```

### Forbidden

- **Code**: 403 Forbidden
- **Content**:
  ```json
  {
    "error": "Forbidden",
    "message": "You don't have permission to access this resource"
  }
  ```

### Not Found

- **Code**: 404 Not Found
- **Content**:
  ```json
  {
    "error": "Not Found",
    "message": "Resource not found"
  }
  ```

### Internal Server Error

- **Code**: 500 Internal Server Error
- **Content**:
  ```json
  {
    "error": "Internal Server Error",
    "message": "Something went wrong"
  }
  ```
