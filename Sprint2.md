# Sprint 2 - Library Management System

## Sprint 2 Goals

- Continue improving the integrated library workflow delivered in Sprint 1.
- Strengthen frontend-backend integration with better API contracts and live dashboard data.
- Add framework-specific unit tests for frontend and backend.
- Add one simple Cypress end-to-end test.
- Document the backend API in detail for submission.

## Sprint 2 User Issues

### Frontend Issues

| ID | User Issue | Description |
|----|------------|-------------|
| F2-1 | Integrate live dashboard summary cards | As a logged-in user, I want the dashboard to show live catalog and borrowing statistics so I can understand library activity at a glance. |
| F2-2 | Improve catalog search and availability filtering | As a user, I want to search books and optionally show only available titles so I can find borrowable books faster. |
| F2-3 | Harden authentication UX for integrated APIs | As a user, I want login flows and API errors to be handled consistently so I can understand failures and continue safely. |
| F2-4 | Prepare the UI for automated browser testing | As a developer, I want stable page selectors and a simple browser flow so I can demonstrate the integrated app with Cypress. |
| F2-5 | Add Angular unit tests for frontend services | As a developer, I want unit tests around API-facing Angular services so frontend integration regressions are caught early. |
| F2-6 | Validate member browsing workflow | As a user, I want the browse flow to remain smooth after integration so I can log in, search, and open the catalog successfully. |

### Backend Issues

| ID | User Issue | Description |
|----|------------|-------------|
| B2-1 | Fix JWT verification fallback mismatch | As a user, I want a token issued during login to also work on protected routes even when default environment values are used. |
| B2-2 | Standardize JSON API responses for integration | As a frontend developer, I want API errors in JSON format so the UI can display backend messages consistently. |
| B2-3 | Add dashboard summary endpoint | As a logged-in user, I want library summary metrics from the backend so the dashboard can show live information. |
| B2-4 | Strengthen request validation for auth, books, and borrowing | As a developer, I want invalid requests rejected clearly so the backend behaves predictably. |
| B2-5 | Prevent duplicate active borrowings for the same user and book | As a library user, I want borrowing rules enforced so I cannot accidentally borrow the same book twice at once. |
| B2-6 | Add Go unit tests and backend API documentation | As a team member, I want tests and endpoint documentation so Sprint 2 submission is complete and maintainable. |

## Completed Sprint 2 Work

### Frontend

- Added a live dashboard summary section backed by `GET /api/dashboard/summary`.
- Added an availability-only filter and result summary to the catalog page.
- Added stable `data-testid` hooks for login and book search flows.
- Added Angular unit tests for `AuthService`, `BookService`, and `DashboardService`.
- Added Cypress configuration and a simple login-and-browse end-to-end test.

### Backend

- Fixed JWT secret fallback so login tokens and middleware validation use the same secret behavior.
- Added reusable JSON response helpers for consistent API responses.
- Added request validation helpers for credentials, books, borrowing, and pagination.
- Added a dashboard summary API endpoint.
- Added duplicate active borrowing protection.
- Added Go unit tests for validation helpers and auth middleware.

## Frontend Tests

### Angular Unit Tests

- `src/app/core/services/auth.service.spec.ts`
  - Verifies login stores token and user state.
  - Verifies logout clears auth state and redirects to `/login`.
- `src/app/core/services/book.service.spec.ts`
  - Verifies search uses the `q` query parameter.
  - Verifies book detail requests target the correct endpoint.
- `src/app/core/services/dashboard.service.spec.ts`
  - Verifies dashboard summary requests target the correct backend endpoint.

### Cypress Test

- `cypress/e2e/login-and-browse.cy.ts`
  - Loads demo data.
  - Logs in as `admin`.
  - Navigates to the catalog.
  - Searches for a book and verifies catalog results render.

## Backend Tests

- `backend/controllers/helpers_test.go`
  - Tests pagination normalization.
  - Tests credential validation.
  - Tests book validation.
  - Tests borrowing validation.
- `backend/middleware/auth_test.go`
  - Tests JWT verification with the default secret fallback.
  - Tests that admin-only middleware blocks non-admin users.

## Backend API Documentation

### Base URL

- Development frontend proxy: `/api`
- Direct backend access: `http://localhost:3000/api`

### Authentication

#### `POST /register`

- Purpose: Create a new user account.
- Auth required: No

Request body:

```json
{
  "username": "reader1",
  "password": "secret123",
  "role": "user"
}
```

Success response:

```json
{
  "message": "User registered successfully"
}
```

Validation notes:

- `username` is required.
- `password` is required and must be at least 6 characters.
- `role` defaults to `user`.

#### `POST /login`

- Purpose: Authenticate a user and issue a JWT.
- Auth required: No

Request body:

```json
{
  "username": "reader1",
  "password": "secret123"
}
```

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

- Purpose: Return summary metrics for the dashboard.
- Auth required: No in current implementation

Success response:

```json
{
  "totalBooks": 12,
  "availableBooks": 24,
  "activeBorrowings": 4
}
```

### Books

#### `GET /books`

- Purpose: List catalog books.
- Auth required: No
- Query parameters:
  - `q`: optional search term matched against title, author, and genre
  - `page`: optional page number, minimum `1`
  - `limit`: optional page size, default `50`, max `100`

Success response:

```json
[
  {
    "id": 1,
    "title": "Clean Code",
    "author": "Robert C. Martin",
    "isbn": "9780132350884",
    "genre": "Programming",
    "language": "English",
    "shelf_number": "A-12",
    "available_copies": 3
  }
]
```

#### `GET /books/{id}`

- Purpose: Return one book by ID.
- Auth required: No

#### `POST /books`

- Purpose: Add a new book.
- Auth required: Yes
- Role required: `admin`

Request body:

```json
{
  "title": "The Pragmatic Programmer",
  "author": "Andrew Hunt",
  "isbn": "9780201616224",
  "genre": "Programming",
  "language": "English",
  "shelf_number": "B-02",
  "available_copies": 4
}
```

#### `PUT /books/{id}`

- Purpose: Update an existing book.
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

Request body:

```json
{
  "bookId": 3
}
```

Success response:

```json
{
  "message": "Book borrowed successfully"
}
```

Business rules:

- `bookId` must be a positive integer.
- The book must exist.
- At least one copy must be available.
- A user cannot keep two active borrowings for the same book.
- Due date is set to 14 days from borrowing time.

#### `POST /borrow/return`

- Purpose: Return a borrowed book for the authenticated user.
- Auth required: Yes

Request body:

```json
{
  "bookId": 3
}
```

### Borrowing History

#### `GET /users/{id}/borrowings`

- Purpose: Fetch borrowing history for one user.
- Auth required: Yes
- Authorization:
  - Users can fetch their own borrowing history.
  - Admins can fetch any user's borrowing history.

Success response:

```json
[
  {
    "book_id": 3,
    "title": "Clean Code",
    "author": "Robert C. Martin",
    "isbn": "9780132350884",
    "borrowed_at": "2026-03-26T10:00:00Z",
    "due_date": "2026-04-09T10:00:00Z",
    "returned_at": null,
    "status": "Borrowing"
  }
]
```

### Demo Data

#### `GET /seed`

- Purpose: Seed demo books and an admin account.
- Auth required: No
- Notes:
  - Creates `admin / admin123` if missing.
  - Uses Open Library when available and falls back to hardcoded books.
  - `?force=1` clears existing books and reseeds them.

## Submission Notes

- `Sprint2.md` now includes completed work, frontend tests, backend tests, and backend API documentation as required.
- The Cypress test expects the Angular frontend on port `4200` and the Go backend on port `3000`.
