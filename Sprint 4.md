# Sprint 4 - Library Management System

**Project:** Library Management System  
**Tech Stack:** Frontend (Angular), Backend (Golang), Database (MySQL)

---

## 1. User Stories for Sprint 4

### Frontend User Cases (6)

| ID | User Story | Demo Focus |
|----|------------|------------|
| **F-US-10** | **As a** user, **I want to** see live dashboard summary cards **so that** I can quickly understand catalog health, loan activity, and reviews. | Dashboard statistics, live summary cards |
| **F-US-11** | **As a** user, **I want to** see ratings and review counts while browsing books **so that** I can choose books with more confidence. | Catalog rating badges, review counts |
| **F-US-12** | **As a** borrower, **I want to** submit or update a review from the book detail page **so that** I can share my reading experience. | Review form, validation, success/error feedback |
| **F-US-13** | **As a** user, **I want to** read reviews on each book detail page **so that** I can learn from other readers before borrowing. | Review list, review dates, reviewer names |
| **F-US-14** | **As a** user, **I want to** see overdue and due-soon indicators in my borrowing history **so that** I can return books on time. | Borrowing status badges, due-date messaging |
| **F-US-15** | **As a** user, **I want to** see a richer book detail page with ratings, notes, availability, and actions **so that** the website feels complete and polished. | Enhanced detail page UI |

### Backend User Cases (6)

| ID | User Story | Demo Focus |
|----|------------|------------|
| **B-US-10** | **As a** frontend developer, **I want** a dashboard summary API **so that** the dashboard can display live statistics from the database. | `/api/dashboard/summary` |
| **B-US-11** | **As a** borrower, **I want to** submit a review through an API **so that** my feedback is stored permanently. | `POST /api/books/{id}/reviews` |
| **B-US-12** | **As a** user, **I want to** fetch reviews for a specific book **so that** the frontend can display community feedback. | `GET /api/books/{id}/reviews` |
| **B-US-13** | **As a** frontend developer, **I want** book responses to include rating aggregates **so that** catalog and detail views can show review data. | Review count, average rating |
| **B-US-14** | **As a** user, **I want** borrowing history responses to include overdue status and due-date insights **so that** I can manage active loans better. | Overdue flag, days until due |
| **B-US-15** | **As a** system, **I want** review data stored in a dedicated table with validation rules **so that** ratings and comments remain structured and reliable. | Review schema, validation |

---

## 2. Issues Planned for Sprint 4

### Backend Issues

| Issue | Title | User Story |
|-------|--------|------------|
| B-11 | Add dashboard summary endpoint | B-US-10 |
| B-12 | Add submit review endpoint | B-US-11 |
| B-13 | Add get reviews by book endpoint | B-US-12 |
| B-14 | Extend book responses with review aggregates | B-US-13 |
| B-15 | Add overdue metadata to borrowing history | B-US-14 |
| B-16 | Add review schema and validation rules | B-US-15 |

### Frontend Issues

| Issue | Title | User Story |
|-------|--------|------------|
| F-11 | Add live dashboard summary cards | F-US-10 |
| F-12 | Show ratings in catalog cards | F-US-11 |
| F-13 | Add review form on book detail page | F-US-12 |
| F-14 | Show reviews on book detail page | F-US-13 |
| F-15 | Add overdue and due-soon borrowing indicators | F-US-14 |
| F-16 | Upgrade book detail page experience | F-US-15 |

---

## 3. Issues Successfully Completed

Sprint 4 focused on turning the project into a more complete final submission with reviews, live dashboard data, and richer borrowing/book experiences.

### Backend

| Issue | Title | Status | Notes |
|-------|--------|--------|-------|
| B-11 | Add dashboard summary endpoint | ✅ Completed | Added `/api/dashboard/summary` with live counts for books, loans, reviews, and user-specific stats. |
| B-12 | Add submit review endpoint | ✅ Completed | Borrowers can create or update one review per book through the API. |
| B-13 | Add get reviews by book endpoint | ✅ Completed | Book detail pages can fetch reader reviews directly from the backend. |
| B-14 | Extend book responses with review aggregates | ✅ Completed | Book list/detail responses now include `review_count` and `average_rating`. |
| B-15 | Add overdue metadata to borrowing history | ✅ Completed | Borrowing history now includes overdue state and due-date insight fields. |
| B-16 | Add review schema and validation rules | ✅ Completed | Review table added with validation for rating range, comment length, and borrower-only review permissions. |

### Frontend

| Issue | Title | Status | Notes |
|-------|--------|--------|-------|
| F-11 | Add live dashboard summary cards | ✅ Completed | Dashboard now shows live stats for catalog, loans, reviews, and user activity. |
| F-12 | Show ratings in catalog cards | ✅ Completed | Catalog cards display average rating and review count. |
| F-13 | Add review form on book detail page | ✅ Completed | Logged-in readers can submit or update a review with rating and comment. |
| F-14 | Show reviews on book detail page | ✅ Completed | Book detail page lists reader reviews with names, stars, and dates. |
| F-15 | Add overdue and due-soon borrowing indicators | ✅ Completed | Borrowing history highlights overdue and due-soon items clearly. |
| F-16 | Upgrade book detail page experience | ✅ Completed | Detail page now combines notes, ratings, reviews, availability, and actions in one polished layout. |

---

## 4. Implementation Notes

### Backend

- Added review-related models, controllers, and routes for creating and reading reviews.
- Added a dashboard summary controller that aggregates book, borrowing, and review statistics.
- Extended existing book and borrowing APIs to support richer frontend displays.
- Updated seeding so final demos include sample readers, borrowings, and reviews.

### Frontend

- Added dashboard and review services for the new Sprint 4 APIs.
- Enhanced catalog cards with rating data and book detail pages with review workflows.
- Improved borrowing history with user-friendly due-date messaging.
- Strengthened the overall demo quality by making more screens feel complete and data-driven.

---

## 5. Sprint 4 Summary

Sprint 4 completed the project by adding live dashboard metrics, persistent book reviews, richer catalog metadata, and clearer borrowing history states.  
These features make the application feel like a full final submission rather than a partial prototype.

---
