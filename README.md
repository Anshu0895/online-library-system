# Online Library Management System

## Objective
The objective of this task is to build REST APIs for a Library Management System.

## Description
The Online Library System allows anyone to create their own library. Each library has two types of user profiles:
- **LibraryAdmin**
- **Readers**

The **Library Owner** (creator) can onboard a **LibraryAdmin** directly and share a sign-up portal with **Readers**. The system currently supports only one LibraryAdmin per library.

Each book is identified uniquely using an ISBN number. The system maintains details of books, including versions, authors, and publishers.

## Profiles and Roles

### ✅ Admin:
- ➕ Add or Remove Books
- ✏️ Update book details (e.g., number of copies)
- 📜 List issue requests
- ✅ Approve or ❌ Disapprove issue requests
- 📖 Issue info for a Reader

### ✅ Reader:
- 📩 Raise an issue request for a book
- 🔍 Search books by Title, Author, or Publisher

### ✅ Owner:
- 🏛️ Create and manage libraries
- 👤 Create and manage Library Admins

## Database Info

### Tables:
- **Library**: (ID, Name)
- **Users**: (ID, Name, Email, ContactNumber, Role, LibID)
- **BookInventory**: (ISBN, LibID, Title, Authors, Publisher, Version, TotalCopies, AvailableCopies)
- **RequestEvents**: (ReqID, BookID, ReaderID, RequestDate, ApprovalDate, ApproverID, RequestType)
- **IssueRegistery**: (IssueID, ISBN, ReaderID, IssueApproverID, IssueStatus, IssueDate, ExpectedReturnDate, ReturnDate, ReturnApproverID)

## ER Daigram
![Daigram](./Library_er_daigram.png)


### User Authentication & Authorization
- Every user must pass their email address for authentication.
- Basic authentication ensures valid users exist in the database.
- Authorization is based on user roles.

### ✅ Library Owner Flows:
- 🏛️ **Library Creation** – Registers a new library with an Owner. Duplicate names are not allowed.
- 👤 **Create Admin** – The Owner can add a LibraryAdmin.

### ✅ Reader Flows:
- 🔍 **Search Book** – Readers can search by title, author, or publisher and check availability.
- 📩 **Raise Issue Request** – Readers can request books using the book ID and email. If unavailable, the request is rejected immediately.

### ✅ Library Admin Flows:
- 📚 **Add Books** – Adds books to inventory. If the book exists, the system increments the count.
- 🗑️ **Remove Books** – Decrements book count until 0. Issued books cannot be removed.
- ✏️ **Update Books** – Updates book details using ISBN.
- 📜 **List Issue Requests** – Admin lists issue requests.
- ✅ **Approve or ❌ Reject Issue Requests** – Sets approver ID, approval date, and updates issue registry accordingly.

## API Routes

### 📌 User Authentication
- `POST /signup` – User registration
- `POST /login` – User login

### 📌 Book Management (LibraryAdmin)
- `POST /books` – Add books
- `PUT /books/:isbn` – Update book details
- `DELETE /books/:isbn` – Remove a book
- `GET /books` – Get all books
- `GET /books/:isbn` – Get a book by ISBN
- `GET /books/search` – Search books

### 📌 Library Management (Owner)
- `POST /libraries` – Create a library
- `GET /libraries` – Get libraries
- `DELETE /libraries/:id` – Delete a library

### 📌 Issue Requests
- `POST /raise-request` – Raise an issue request
- `GET /requests` – Get all issue requests
- `GET /pending-requests` – Get pending requests
- `GET /requests/:id` – Get request by ID
- `PUT /requests/:id/approve` – Approve request
- `PUT /requests/:id/reject` – Reject request

### 📌 User Management
- `POST /users` – Create an Admin
- `GET /users` – Get all users
- `GET /admins` – Get all admins
- `GET /users/:id` – Get user by ID
- `PUT /users/:id` – Update user
- `DELETE /users/:id` – Delete user


Online Library System - Test Report

📌 Overview

This document provides a test report for the Online Library System, detailing the test execution results, observations, and areas for improvement.

📅 Date: February 25, 2025

⏱ Total Execution Time: 0.244s

📂 Module: online-library-system/controllers

✅ Passed Tests

1️⃣ User Authentication

TestSignup – Successfully created a new user (John Doe)

TestLogin – Successfully authenticated user and generated a token

2️⃣ Books Management

TestAddBook – Successfully added a new book

TestGetBooks – Successfully retrieved all books

TestGetBook – Successfully retrieved a specific book (ISBN: 3333333333)

TestRemoveBook – Successfully removed a book (ISBN: 4444444444)

3️⃣ Library Management

TestGetLibraries – Successfully retrieved all libraries

TestCreateLibrary – Successfully created a new library (New Library)

TestCreateDuplicateLibrary – Prevented duplicate library creation (Conflict 409)

TestDeleteLibrary – Successfully deleted library (ID: 1)

TestDeleteNonExistentLibrary – Handled non-existent library deletion (ID: 999)

4️⃣ Requests & Users Management

TestCreateIssueRequest – Successfully created a book issue request (Reader ID: 1)

TestCreateAdmin – Successfully created an admin user

TestGetUser – Successfully retrieved user details (ID: 1)

TestGetUsers – Successfully retrieved all users

TestGetAdmins – Successfully retrieved all admins

TestUpdateUser – Successfully updated user information (ID: 1)

TestDeleteUser – Successfully deleted user (ID: 1)

📊 Summary

✔ Total Tests Run: 18✅ Passed: 18❌ Failed: 0⚠️ Warnings: Some records not found but did not affect test results

📜 How to Run Tests

To execute the test suite, use the following command: