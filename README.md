# Wallet App

A full-stack wallet application built with Go backend and React frontend. This application provides comprehensive financial management features including account tracking, transactions, budgets, investments, and more.

## Project Structure

```
wallet_app/
├── backend/          # Go backend application
│   ├── config.env    # Environment configuration
│   ├── go.mod        # Go module file
│   ├── main.go       # Application entry point
│   ├── internal/     # Internal packages
│   ├── pkg/          # Shared packages
│   └── utils/        # Utility functions
├── frontend/         # React frontend application
│   ├── src/          # Source code
│   ├── package.json  # Node dependencies
│   └── vite.config.js # Vite configuration
├── run.bat           # Script to run both frontend and backend
└── README.md         # This file
```

## Features

- **User Management**: Registration and authentication with JWT tokens
- **Account Management**: Multiple accounts with different currencies
- **Transaction Tracking**: Record income and expenses
- **Transfer Management**: Move money between accounts
- **Budget Planning**: Set and track budgets
- **Investment Tracking**: Monitor investment portfolios
- **Recurring Transactions**: Automate regular payments
- **Cash Flow Analysis**: Generate financial summaries
- **Multi-Currency Support**: Handle different currencies
- **Bank Integration**: Connect with financial institutions

## Database Setup

1. **Install PostgreSQL** and create a database:
   ```sql
   CREATE DATABASE wallet_db;
   ```

2. **Run the setup script**:
   ```bash
   cd backend
   psql -U postgres -d wallet_db -f setup.sql
   ```

3. **Seed default admin user**:
   ```bash
   cd backend
   go run seeder.go
   ```

   Default admin credentials:
   - Email: `admin@walletapp.com`
   - Password: `admin123`

## Quick Start

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd wallet_app
   ```

2. **Set up database** (see Database Setup section above)

3. **Run the application**
   ```bash
   # Double-click run.bat or run from command line
   ./run.bat
   ```

   This will start both the backend (http://localhost:8000) and frontend (http://localhost:5173).

## Manual Setup

### Backend Setup

1. **Prerequisites**
   - Go (version 1.21 or higher)
   - PostgreSQL database

2. **Setup**
   ```bash
   cd backend
   go mod tidy
   ```

3. **Configuration**
   - Update `backend/config.env` with your database credentials
   - Ensure PostgreSQL is running and database is created

4. **Run**
   ```bash
   go run main.go
   ```

### Frontend Setup

1. **Prerequisites**
   - Node.js (version 18 or higher)

2. **Setup**
   ```bash
   cd frontend
   npm install
   ```

3. **Run**
   ```bash
   npm run dev
   ```

## API Documentation

### Authentication

- `POST /api/v1/login` - User login
- `POST /api/v1/refresh` - Refresh access token

### Resources

All CRUD operations are available for the following resources under `/api/v1`:

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

### Authentication

All API endpoints except login and refresh require JWT authentication. Include the access token in the Authorization header:

```
Authorization: Bearer <access_token>
```

## Technology Stack

- **Backend**: Go, Echo framework, PostgreSQL, Bun ORM, JWT authentication
- **Frontend**: React, Vite, Axios, React Router
- **Database**: PostgreSQL
- **Authentication**: JWT with access and refresh tokens

## Development

- Backend uses clean architecture with separate layers for handlers, services, and repositories
- Frontend uses reusable components for CRUD operations
- JWT tokens are automatically refreshed on expiration
- CORS is configured for local development

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test both frontend and backend
5. Submit a pull request

## License

This project is licensed under the MIT License.