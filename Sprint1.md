# Sprint 1 — Library Management System

**Project:** Library Management System  
**Tech Stack:** Frontend (Angular), Backend (Golang), Database (MySQL)

---

## 1. User Stories for Sprint 1

### Frontend User Cases (5)

| ID | User Story | Demo Focus |
|----|------------|------------|
| **F-US-1** | **As a** guest, **I want to** see a login page with username and password fields **so that** I can sign in to access the library. | Login page UI, form validation |
| **F-US-2** | **As a** guest, **I want to** navigate to a registration page and create an account **so that** I can become a library member. | Register page, sign-up flow |
| **F-US-3** | **As a** logged-in user, **I want to** be redirected to a dashboard or home after login **so that** I can access the main app. | Routing, auth redirect, layout |
| **F-US-4** | **As a** user, **I want to** view a list of books (title, author, availability) **so that** I can browse the catalog. | Books list page, API integration (or mock) |
| **F-US-5** | **As a** user, **I want to** log out and be redirected to the login page **so that** my session ends securely. | Logout button, token clear, redirect |

### Backend User Cases (5)

| ID | User Story | Demo Focus |
|----|------------|------------|
| **B-US-1** | **As a** developer, **I want** the backend connected to MySQL with a schema (users, books, borrowings) **so that** data is persisted and structured. | DB connection, schema setup |
| **B-US-2** | **As a** guest, **I want to** register via `POST /api/register` with username and password **so that** I can create an account. | Register API, bcrypt, validation |
| **B-US-3** | **As a** user, **I want to** log in via `POST /api/login` and receive a JWT **so that** I can authenticate for protected endpoints. | Login API, JWT issuance |
| **B-US-4** | **As a** developer, **I want** protected routes to validate JWT and enforce role checks **so that** only authorized users access admin/borrow APIs. | Auth middleware, RBAC |
| **B-US-5** | **As a** user, **I want to** fetch the book catalog via `GET /api/books` **so that** the frontend can display available books. | List books API |

---

## 2. Issues Planned for Sprint 1

### Backend Issues

| Issue | Title | User Story |
|-------|--------|------------|
| B-1 | Set up Go project and MySQL connection | B-US-1 |
| B-2 | Create database schema (users, books, borrowingrecords) | B-US-1 |
| B-3 | Implement `POST /api/register` | B-US-2 |
| B-4 | Implement `POST /api/login` and JWT issuance | B-US-3 |
| B-5 | Implement auth middleware and RBAC | B-US-4 |
| B-6 | Implement `GET /api/books` | B-US-5 |

### Frontend Issues

| Issue | Title | User Story |
|-------|--------|------------|
| F-1 | Set up Angular project, routing (login, register, main) | F-US-1, F-US-2, F-US-3 |
| F-2 | Create login page (form, call API, store token) | F-US-1 |
| F-3 | Create register page and sign-up flow | F-US-2 |
| F-4 | Implement auth guard and redirect after login | F-US-3 |
| F-5 | Create books list page (display catalog) | F-US-4 |
| F-6 | Implement logout and token handling | F-US-5 |

---

## 3. Issues Successfully Completed

All Sprint 1 issues were implemented, tested, and integrated into the main branch via pull requests.

### Backend

| Issue | Title | Status | Notes |
|-------|--------|--------|-------|
| B-1 | Set up Go project and MySQL connection | ✅ Completed | |
| B-2 | Create database schema | ✅ Completed | |
| B-3 | Implement `POST /api/register` | ✅ Completed | |
| B-4 | Implement `POST /api/login` and JWT | ✅ Completed | |
| B-5 | Implement auth middleware and RBAC | ✅ Completed | |
| B-6 | Implement `GET /api/books` | ✅ Completed | |

### Frontend

| Issue | Title | Status | Notes |
|-------|--------|--------|-------|
| F-1 | Set up Angular project and routing | ✅ Completed | |
| F-2 | Create login page | ✅ Completed | |
| F-3 | Create register page | ✅ Completed | |
| F-4 | Auth guard and redirect | ✅ Completed | |
| F-5 | Books list page | ✅ Completed | |
| F-6 | Logout and token handling | ✅ Completed | |

---
