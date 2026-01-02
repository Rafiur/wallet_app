# Wallet App

A full-stack wallet application with Go backend and React frontend.

## Backend (Go)

### Setup

1. Install Go (version 1.21 or higher) from https://golang.org/dl/

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Set up PostgreSQL database and update `config.env` with connection details.

4. Run the server:
   ```
   go run main.go
   ```

The API will be available at http://localhost:8000

## Frontend (React)

### Setup

1. Install Node.js (version 18 or higher) from https://nodejs.org/

2. Navigate to frontend directory:
   ```
   cd frontend
   ```

3. Install dependencies:
   ```
   npm install
   ```

4. Start the development server:
   ```
   npm run dev
   ```

5. Open http://localhost:5173 in your browser.

## API Endpoints

All CRUD endpoints are available under `/api/v1`:

- `/accounts` - Account management
- `/users` - User management
- `/transactions` - Transaction management
- `/transfers` - Transfer management
- `/budgets` - Budget management
- `/currencies` - Currency management
- `/expense-categories` - Expense category management
- `/banks` - Bank management
- `/investments` - Investment management
- `/recurring-transactions` - Recurring transaction management
- `/account-currencies` - Account currency management
- `/cash-flow-summaries` - Cash flow summary management

## Notes

- JWT authentication is set up in the backend but not fully implemented in the frontend.
- Ensure database schema is created before running the app.
- This is a basic implementation; add features like authentication UI, better error handling, etc., as needed.