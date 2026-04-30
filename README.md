# Library Management System

## Project Description

This Library Management System is a full-stack web application built with Angular, Go, and MySQL. It supports secure authentication, role-based access control, book management, borrowing and returning workflows, dashboard summaries, and book reviews with ratings.

## Tech Stack

- Front End: Angular
- Back End: Go
- Database: MySQL

## Main Features

- User registration and login with JWT authentication
- Role-based access control for admin and member flows
- Browse, search, and filter books
- Borrow and return books
- Admin book management
- Dashboard summary statistics
- Book reviews and ratings

## Project Structure

- `frontend/` - Angular client application
- `backend/` - Go REST API
- `schema.sql` - Database schema
- `Sprint1.md`, `Sprint2.md`, `Sprint3.md` - Sprint documentation

## Run Locally

### Backend

```powershell
cd backend
$env:GOSUMDB="off"
$env:GOMODCACHE="$PWD\.gomodcache"
$env:GOCACHE="$PWD\.gocache"
$env:GOTMPDIR="$PWD\.gotmp"
go mod tidy
go run .
```

### Frontend

```powershell
cd frontend
npm install
npm start
```

Frontend runs at `http://localhost:4200` and the backend runs at `http://localhost:3000`.
