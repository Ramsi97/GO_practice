
# Library Management System (Go)

This is a simple console-based Library Management System built in Go to demonstrate:
- Structs and interfaces
- Maps and slices
- Methods and error handling
- Modular project structure

## Features
- Add and remove books
- Borrow and return books
- View available and borrowed books
- Command-line menu interface

## Packages
- `models/`: Data structures (Book, Member)
- `services/`: Business logic (LibraryManager implementation)
- `controllers/`: Console input/output handling
- `main.go`: Entry point

## How to Run
1. Initialize module:
   ```bash
   go mod init library_management
