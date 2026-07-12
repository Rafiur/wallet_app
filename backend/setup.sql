-- Database setup for Wallet App
-- Run this script to create the database and tables

-- Create database (run as postgres superuser)
-- CREATE DATABASE wallet_db;

-- Connect to wallet_db and run the following:

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Currencies table
CREATE TABLE currencies (
    code TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    symbol TEXT NOT NULL
);

-- Accounts table
CREATE TABLE accounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    type TEXT NOT NULL, -- 'bank', 'wallet', 'cash', 'investment'
    balance DOUBLE PRECISION DEFAULT 0,
    currency TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Banks table
CREATE TABLE banks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    branch TEXT,
    account_number TEXT UNIQUE,
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    annual_charge DOUBLE PRECISION DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Expense Categories table
CREATE TABLE expense_categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    type TEXT DEFAULT 'expense', -- 'expense', 'income'
    user_id UUID REFERENCES users(id), -- null = global category
    parent_category_id UUID REFERENCES expense_categories(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Recurring Transactions table
CREATE TABLE recurring_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    name TEXT NOT NULL, -- e.g. "Bi-monthly utility", "Annual bank fee"
    amount DOUBLE PRECISION NOT NULL,
    frequency TEXT NOT NULL, -- 'monthly', 'bi-monthly', 'quarterly', 'semi-annual', 'annual'
    frequency_interval INTEGER DEFAULT 1, -- e.g. 2 for every 2 months
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP,
    next_due_date TIMESTAMP NOT NULL,
    transaction_type TEXT NOT NULL, -- 'income', 'expense'
    expense_category_id UUID REFERENCES expense_categories(id),
    note TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Transactions table
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    transaction_name TEXT NOT NULL,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    amount DOUBLE PRECISION NOT NULL, -- positive for income, negative for expense
    expense_category_id UUID REFERENCES expense_categories(id),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    note TEXT,
    tags TEXT[],
    transaction_type TEXT NOT NULL, -- 'income' or 'expense'
    is_billable BOOLEAN DEFAULT false,
    recurring_id UUID REFERENCES recurring_transactions(id),
    deleted_at TIMESTAMP
);

-- Transfers table
CREATE TABLE transfers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    from_account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    to_account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    amount DOUBLE PRECISION NOT NULL,
    currency TEXT NOT NULL,
    exchange_rate DOUBLE PRECISION,
    transfer_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    note TEXT,
    status TEXT DEFAULT 'pending', -- 'pending', 'completed', 'failed'
    deleted_at TIMESTAMP
);

-- Budgets table
CREATE TABLE budgets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    expense_category_id UUID REFERENCES expense_categories(id), -- per category or global (null)
    period TEXT NOT NULL, -- 'monthly', 'quarterly', 'yearly'
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP,
    amount DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Investments table
CREATE TABLE investments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_id UUID REFERENCES accounts(id),
    type TEXT NOT NULL, -- 'fdr', 'stocks', 'bonds'
    amount DOUBLE PRECISION NOT NULL,
    interest_rate DOUBLE PRECISION,
    start_date TIMESTAMP NOT NULL,
    maturity_date TIMESTAMP,
    status TEXT DEFAULT 'active', -- 'active', 'matured'
    note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Account Currencies table
CREATE TABLE account_currencies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    currency_code TEXT NOT NULL,
    balance_according_to_currency REAL DEFAULT 0
);

-- Cash Flow Summary table
CREATE TABLE cash_flow_summaries (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_id UUID REFERENCES accounts(id), -- per account or global
    period TEXT NOT NULL, -- 'daily', 'monthly', 'yearly'
    start_date TIMESTAMP NOT NULL,
    inflow DOUBLE PRECISION DEFAULT 0,
    outflow DOUBLE PRECISION DEFAULT 0,
    net_flow DOUBLE PRECISION DEFAULT 0,
    billable_outflow DOUBLE PRECISION DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Sessions table
CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    refresh_token TEXT NOT NULL UNIQUE,
    user_agent TEXT,
    ip_address TEXT,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_accounts_user_id ON accounts(user_id);
CREATE INDEX idx_transactions_account_id ON transactions(account_id);
CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transfers_from_account ON transfers(from_account_id);
CREATE INDEX idx_transfers_to_account ON transfers(to_account_id);
CREATE INDEX idx_budgets_user_id ON budgets(user_id);
CREATE INDEX idx_investments_user_id ON investments(user_id);
CREATE INDEX idx_recurring_transactions_user_id ON recurring_transactions(user_id);
CREATE INDEX idx_account_currencies_account_id ON account_currencies(account_id);
CREATE INDEX idx_cash_flow_summaries_user_id ON cash_flow_summaries(user_id);
CREATE INDEX idx_sessions_user_id ON sessions(user_id);
CREATE INDEX idx_sessions_refresh_token ON sessions(refresh_token);

-- Default currencies
INSERT INTO currencies (code, name, symbol) VALUES
    ('USD', 'US Dollar', '$'),
    ('EUR', 'Euro', '€'),
    ('GBP', 'British Pound', '£'),
    ('BDT', 'Bangladeshi Taka', '৳'),
    ('INR', 'Indian Rupee', '₹');

-- Default global expense/income categories (user_id NULL = visible to every user)
INSERT INTO expense_categories (name, type) VALUES
    ('Food & Dining', 'expense'),
    ('Groceries', 'expense'),
    ('Transport', 'expense'),
    ('Utilities', 'expense'),
    ('Rent/Housing', 'expense'),
    ('Healthcare', 'expense'),
    ('Entertainment', 'expense'),
    ('Shopping', 'expense'),
    ('Education', 'expense'),
    ('Other', 'expense'),
    ('Salary', 'income'),
    ('Interest', 'income'),
    ('Bonus', 'income'),
    ('Business Income', 'income'),
    ('Other Income', 'income');
