# Respire Backend API

This is the backend API for the Respire learning platform, built with Go.

## Features

- User authentication and management
- Post management (create, view, delete)
- Friend requests and management
- Chat functionality
- Course management system
- Category management
- Section and content management
- Notice board functionality

## API Endpoints

### Authentication
- `POST /login` - User login
- `POST /register` - User registration
- `POST /start-verification` - Start user verification
- `POST /complete-registration` - Complete user registration
- `GET /me` - Get current user
- `POST /me` - Update current user

### Posts
- `GET /posts` - Get all posts
- `POST /posts` - Create a post
- `DELETE /posts/{id}` - Delete a post
- `POST /posts/{id}/like` - Like a post
- `POST /posts/{id}/unlike` - Unlike a post
- `GET /posts/{id}/comment` - Get comments on a post
- `POST /posts/{id}/comment` - Add a comment to a post

### Friends
- `GET /friends` - Get friends
- `GET /notfriends` - Get non-friends
- `GET /friend-requests` - Get friend requests
- `POST /friend-requests` - Send a friend request
- `POST /friend-requests/{id}/accept` - Accept a friend request
- `POST /friend-requests/{id}/reject` - Reject a friend request

### Courses
- `GET /courses` - Get all courses
- `POST /courses` - Create a course (admin)
- `GET /courses/{id}` - Get course details
- `PUT /courses/{id}` - Update a course (admin)
- `DELETE /courses/{id}` - Delete a course (admin)
- `POST /courses/{id}/subscribe` - Subscribe to a course
- `POST /courses/{id}/unsubscribe` - Unsubscribe from a course
- `GET /subscriptions` - Get user's subscriptions

### Categories
- `GET /categories` - Get all categories
- `POST /categories` - Create a category (admin)
- `GET /categories/{id}` - Get category details
- `PUT /categories/{id}` - Update a category (admin)
- `DELETE /categories/{id}` - Delete a category (admin)

### Sections
- `GET /sections` - Get all sections
- `POST /sections` - Create a section (admin)
- `GET /sections/{id}` - Get section details
- `POST /content` - Create content for a section

### Notices
- `GET /notices` - Get all notices
- `POST /notices` - Create a notice (admin)
- `DELETE /notices/{id}` - Delete a notice (admin)

### Other
- `GET /articles` - Get all articles
- `POST /articles` - Create an article
- `GET /articles/{id}` - Get article details
- `POST /articles/{id}` - Update an article
- `GET /chats` - Get chats
- `POST /add-chat` - Add a chat
- `POST /upload` - Upload a file
- `POST /uploadlink` - Get upload link
- `GET /assets/{name}` - Get an asset

## Technologies Used
- Go
- MongoDB
- WebSockets for real-time chat
- JWT for authentication

## Database
MongoDB connection string: mongodb+srv://softsolspak:SFWZ9evKS69CdQSx@respire.9xsja.mongodb.net/