# Sprint 2 - Library Management System

**Project:** Library Management System  
**Tech Stack:** Frontend (Angular), Backend (Golang), Database (MySQL)

---
## Video Links: 
Drushtant Patil : https://drive.google.com/file/d/1CcVYY4KhjImTgqOSdisLwnCFMS-HBD67/view?usp=sharing

Kanishka Dhaundiyal: https://drive.google.com/file/d/1CcVYY4KhjImTgqOSdisLwnCFMS-HBD67/view?usp=sharing



## 1. User Stories for Sprint 2

### Frontend User Cases (4)

| ID | User Story | Demo Focus |
|----|------------|------------|
| **F-US-6** | **As a** user, **I want to** see clearer login errors and loading states **so that** I understand whether authentication succeeded or failed. | Login UX, API error handling, auth persistence |
| **F-US-7** | **As a** user, **I want to** view dashboard summary cards **so that** I can understand the system at a glance. | Dashboard summary UI |
| **F-US-8** | **As a** user, **I want to** search the catalog and filter by availability **so that** I can find books more quickly. | Search, filter, sort, availability |
| **F-US-9** | **As a** developer, **I want** stable UI selectors for browser automation **so that** frontend behavior can be tested reliably. | Cypress selectors, automated login flow |

### Backend User Cases (4)

| ID | User Story | Demo Focus |
|----|------------|------------|
| **B-US-6** | **As a** developer, **I want** JWT verification to use the same fallback secret as login token creation **so that** protected routes work consistently in development. | JWT fallback consistency |
| **B-US-7** | **As a** frontend developer, **I want** standardized JSON API responses **so that** integration logic is simpler and more predictable. | Response shape consistency |
| **B-US-8** | **As a** frontend developer, **I want** a dashboard summary endpoint **so that** the frontend can show live system metrics. | Summary endpoint |
| **B-US-9** | **As a** user and developer, **I want** stronger validation for auth, book, and borrowing requests **so that** invalid requests are rejected safely. | Request decoding, ID checks, borrow validation |

---

## 2. Issues Planned for Sprint 2

### Backend Issues

| Issue | Title | User Story |
|-------|--------|------------|
| B-7 | Fix JWT verification fallback mismatch | B-US-6 |
| B-8 | Standardize JSON API responses for integration | B-US-7 |
| B-9 | Add dashboard summary endpoint | B-US-8 |
| B-10 | Strengthen request validation for auth, books, and borrowing | B-US-9 |

### Frontend Issues

| Issue | Title | User Story |
|-------|--------|------------|
| F-7 | Integrate live dashboard summary cards | F-US-7 |
| F-8 | Improve catalog search and availability filtering | F-US-8 |
| F-9 | Harden authentication UX for integrated APIs | F-US-6 |
| F-10 | Prepare the UI for automated browser testing | F-US-9 |

---

## 3. Issues Successfully Completed

Sprint 2 focused on frontend-backend integration, stronger validation, improved authentication UX, and test readiness.

### Backend

| Issue | Title | Status | Notes |
|-------|--------|--------|-------|
| B-7 | Fix JWT verification fallback mismatch | ✅ Completed | Middleware and login token creation now use the same fallback secret. |
| B-8 | Standardize JSON API responses for integration | ✅ Completed | Auth, book, and borrowing endpoints return consistent JSON response shapes. |
| B-9 | Add dashboard summary endpoint | ⏳ Pending | No `/dashboard` or `/summary` route/controller is present in the current codebase. |
| B-10 | Strengthen request validation for auth, books, and borrowing | ✅ Completed | Controllers validate request bodies, IDs, missing records, and borrow availability. |

### Frontend

| Issue | Title | Status | Notes |
|-------|--------|--------|-------|
| F-7 | Integrate live dashboard summary cards | ⏳ Pending | Dashboard UI exists, but live summary-card API integration is not yet visible in the current code. |
| F-8 | Improve catalog search and availability filtering | ✅ Completed | Search, stock filter, genre filter, favorites filter, and sorting are implemented. |
| F-9 | Harden authentication UX for integrated APIs | ✅ Completed | Login flow shows validation, loading state, and backend connectivity errors clearly. |
| F-10 | Prepare the UI for automated browser testing | ✅ Completed | `data-cy` selectors and Cypress login coverage are present. |

---

## 4. Implementation Notes

### Backend

- JWT fallback consistency is implemented in authentication middleware and login token creation.
- JSON responses are standardized across auth, books, and borrowing controllers.
- Validation improvements cover auth request decoding, book ID/body checks, and borrow/return validation.

### Frontend

- Authentication UX improvements are visible in the login component and auth service integration.
- Catalog improvements are implemented in the books list with search, filtering, and sorting.
- Browser automation support is present through Cypress-friendly selectors on the login page and matching Cypress tests.

---

## 5. Sprint 2 Summary

Sprint 2 successfully improved authentication UX, catalog usability, backend validation, and API consistency.  
The remaining dashboard summary work is identified and documented, but it is not yet implemented in the current codebase.

---
