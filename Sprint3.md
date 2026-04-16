# Sprint 3 - Library Management System

---
## Video Links: 
Drushtant Patil : https://drive.google.com/file/d/1DYEIkFmqN8m7fEhBs9iDZKMnctOIJXlE/view?usp=drive_link
Kanishkka Dhaundiyal : https://drive.google.com/file/d/1DzxRs1N2ogZ4Be_PL4Dtezc_k1x82X3C/view?usp=drive_link

## Sprint 3 Goals

- Continue Sprint 2 work by extending the integrated application with new user-facing functionality.
- Add unit tests for the new functionality while retaining the Sprint 2 test inventory.
- Update the backend API documentation to reflect the latest endpoints and behaviors.

## Sprint 3 User Stories

### Frontend User Stories

- As a reader, I want to see reviews on the book detail page so that I can decide whether a book is worth borrowing.
- As a logged-in member, I want to rate a book from 1 to 5 so that I can share my opinion with other users.
- As a logged-in member, I want to write a review comment for a book so that I can describe what I liked or disliked.
- As a reviewer, I want my previous review to appear in the form so that I can edit it without starting over.
- As a user, I want to see the average rating and total review count so that I can quickly understand overall feedback.
- As a member, I want borrowing and reviewing to work together on the same book page so that I can use both features in one place.

### Backend User Stories

- As a developer, I want a `Reviews` table linked to users and books so that ratings and comments are stored persistently.
- As a user, I want an endpoint to fetch all reviews for a book so that the frontend can display them.
- As an authenticated member, I want an endpoint to submit a review for a book so that my rating and comment are saved.
- As a reviewer, I want the backend to update my existing review instead of creating duplicates so that each user has one review per book.
- As a developer, I want review input validation so that invalid ratings or empty comments are rejected cleanly.
- As a team member, I want updated API documentation and unit tests for the review feature so that Sprint 3 submission is complete.

## Sprint 3 User Issues

### Frontend Issues

| ID | User Issue | Description |
|----|------------|-------------|
| F3-1 | Add book review section to the book detail page | As a reader, I want to view reviews for a book so I can judge whether it is worth borrowing. |
| F3-2 | Allow members to submit ratings and comments | As a logged-in member, I want to rate and review a book so I can share feedback with other readers. |
| F3-3 | Show average rating and review count on the detail page | As a user, I want to see review summaries so I can quickly understand community sentiment. |
| F3-4 | Prefill my existing review when editing | As a reviewer, I want my existing review to load into the form so I can update it without rewriting everything. |
| F3-5 | Add Angular unit tests for the review API integration | As a developer, I want tests for the review service so API regressions are easier to detect. |
| F3-6 | Keep the borrowing flow working beside reviews | As a member, I want book borrowing and reviewing to coexist on the detail page without breaking the existing flow. |

### Backend Issues

| ID | User Issue | Description |
|----|------------|-------------|
| B3-1 | Create review schema for books and users | As a developer, I want a review table so ratings and comments are stored persistently. |
| B3-2 | Implement `GET /api/books/{id}/reviews` | As a user, I want to fetch reviews for a book so the frontend can display them. |
| B3-3 | Implement `POST /api/books/{id}/reviews` | As an authenticated member, I want to submit a review for a book so my feedback is saved. |
| B3-4 | Upsert a review when the same user reviews again | As a reviewer, I want to update my earlier review instead of creating duplicates. |
| B3-5 | Validate review rating and comment input | As a developer, I want review requests validated so only clean review data is stored. |
| B3-6 | Add backend unit tests and updated API docs for reviews | As a team member, I want tests and documentation for the review feature so Sprint 3 submission is complete. |

## New Functionality Implemented

### Book Reviews and Ratings

Sprint 3 introduces a full reviews feature across the stack:

- Users can open a book detail page and see existing reader reviews.
- Logged-in non-admin members can submit a rating from 1 to 5 and a written comment.
- If the same user reviews the same book again, the previous review is updated instead of duplicated.
- The book detail page shows the average rating and the total review count.

## Completed Sprint 3 Work

### Frontend

- Added a review model and review service for the new backend endpoints.
- Extended the book detail page to display review summaries, existing reviews, and a review form.
- Added form prefilling when the signed-in user already reviewed the selected book.
- Added Angular unit tests for the review service.

### Backend

- Added the `Reviews` table to automatic schema setup and `schema.sql`.
- Added `GET /api/books/{id}/reviews`.
- Added authenticated `POST /api/books/{id}/reviews`.
- Added review validation helpers and review sanitization.
- Added update-in-place review behavior using MySQL upsert logic.
- Added backend tests for review payload validation.

## Frontend Unit Tests

### Sprint 2 Tests

- `src/app/core/services/auth.service.spec.ts`
- `src/app/core/services/book.service.spec.ts`
- `src/app/core/services/dashboard.service.spec.ts`

### Sprint 3 Tests

- `src/app/core/services/review.service.spec.ts`
  - Verifies review loading requests.
  - Verifies review submission requests.

## Backend Unit Tests

### Sprint 2 Tests

- `backend/controllers/helpers_test.go`
  - Pagination validation
  - Credential validation
  - Book validation
  - Borrow request validation
- `backend/middleware/auth_test.go`
  - JWT validation with fallback secret
  - Admin middleware protection

### Sprint 3 Tests

- `backend/controllers/helpers_test.go`
  - Review rating validation
  - Review comment validation

## Updated Backend API Documentation

### Base URL

- Frontend proxy: `/api`
- Direct backend access: `http://localhost:3000/api`

### Authentication

#### `POST /register`

- Purpose: Register a new user.
- Auth required: No

#### `POST /login`

- Purpose: Authenticate a user and return a JWT.
- Auth required: No

Success response:

```json
{
  "message": "Login successful",
  "token": "jwt-token",
  "username": "reader1",
  "role": "user",
  "userId": 1
}
```

### Dashboard

#### `GET /dashboard/summary`

- Purpose: Return dashboard summary statistics.
- Auth required: No

### Books

#### `GET /books`

- Purpose: Return the catalog with optional search and pagination.
- Query parameters:
  - `q`: search term for title, author, or genre
  - `page`: optional page number
  - `limit`: optional page size, max `100`

#### `GET /books/{id}`

- Purpose: Return one book by ID.

#### `POST /books`

- Purpose: Create a book.
- Auth required: Yes
- Role required: `admin`

#### `PUT /books/{id}`

- Purpose: Update a book.
- Auth required: Yes
- Role required: `admin`

#### `DELETE /books/{id}`

- Purpose: Delete a book.
- Auth required: Yes
- Role required: `admin`

### Borrowing

#### `POST /borrow`

- Purpose: Borrow a book for the authenticated user.
- Auth required: Yes

#### `POST /borrow/return`

- Purpose: Return a borrowed book for the authenticated user.
- Auth required: Yes

#### `GET /users/{id}/borrowings`

- Purpose: Return borrowing history for a user.
- Auth required: Yes

### Reviews

#### `GET /books/{id}/reviews`

- Purpose: Fetch all reviews for a book.
- Auth required: No

Success response:

```json
[
  {
    "id": 1,
    "bookId": 5,
    "userId": 2,
    "username": "reader1",
    "rating": 5,
    "comment": "Excellent practical reference.",
    "createdAt": "2026-04-13T09:00:00Z",
    "updatedAt": "2026-04-13T09:00:00Z"
  }
]
```

#### `POST /books/{id}/reviews`

- Purpose: Create or update the authenticated user's review for a book.
- Auth required: Yes

Request body:

```json
{
  "rating": 4,
  "comment": "Helpful examples and clear explanations."
}
```

Validation rules:

- `rating` must be between `1` and `5`
- `comment` is required
- `comment` must be at most `500` characters

Behavior:

- If the user has not reviewed the book before, the backend creates a new review.
- If the user already reviewed the same book, the backend updates the existing review.

## Submission Notes

- `Sprint3.md` includes Sprint 3 work completed, frontend unit tests, backend unit tests, and updated backend API documentation.
- The application now supports both borrowing and reviewing books from the detail page.
- Previous Sprint 2 tests remain part of the overall test inventory for the project.
