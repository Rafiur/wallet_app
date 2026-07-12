import { BrowserRouter as Router, Routes, Route, NavLink, Navigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import axios, { clearSession, getCurrentUserFullName } from './api';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import Accounts from './pages/Accounts';
import Users from './pages/Users';
import Transactions from './pages/Transactions';
import Transfers from './pages/Transfers';
import Budgets from './pages/Budgets';
import Currencies from './pages/Currencies';
import ExpenseCategories from './pages/ExpenseCategories';
import Banks from './pages/Banks';
import Investments from './pages/Investments';
import RecurringTransactions from './pages/RecurringTransactions';
import AccountCurrencies from './pages/AccountCurrencies';
import CashFlowSummaries from './pages/CashFlowSummaries';

const navSections = [
  {
    title: 'Overview',
    links: [{ to: '/dashboard', label: 'Dashboard' }],
  },
  {
    title: 'Money',
    links: [
      { to: '/accounts', label: 'Accounts' },
      { to: '/transactions', label: 'Transactions' },
      { to: '/transfers', label: 'Transfers' },
      { to: '/expense-categories', label: 'Expense Categories' },
    ],
  },
  {
    title: 'Planning',
    links: [
      { to: '/budgets', label: 'Budgets' },
      { to: '/recurring-transactions', label: 'Recurring' },
      { to: '/investments', label: 'Investments' },
      { to: '/cash-flow-summaries', label: 'Cash Flow' },
    ],
  },
  {
    title: 'Reference',
    links: [
      { to: '/banks', label: 'Banks' },
      { to: '/account-currencies', label: 'Account Currencies' },
      { to: '/currencies', label: 'Currencies' },
      { to: '/users', label: 'Users' },
    ],
  },
];

function AppShell({ onLogout }) {
  const linkClass = ({ isActive }) =>
    `block rounded-md px-3 py-1.5 text-sm transition-colors ${
      isActive ? 'bg-brand-50 text-brand-700 font-medium' : 'text-slate-600 hover:bg-slate-100'
    }`;

  return (
    <div className="min-h-screen flex bg-slate-50">
      <aside className="w-60 shrink-0 border-r border-slate-200 bg-white flex flex-col">
        <div className="px-4 py-5">
          <span className="text-lg font-semibold text-brand-600">Wallet</span>
        </div>
        <nav className="flex-1 overflow-y-auto px-3 space-y-5">
          {navSections.map((section) => (
            <div key={section.title}>
              <p className="px-3 text-xs font-semibold uppercase tracking-wide text-slate-400 mb-1">{section.title}</p>
              <div className="space-y-0.5">
                {section.links.map((link) => (
                  <NavLink key={link.to} to={link.to} className={linkClass}>{link.label}</NavLink>
                ))}
              </div>
            </div>
          ))}
        </nav>
        <div className="border-t border-slate-200 px-4 py-4">
          <p className="text-sm text-slate-600 truncate mb-2">{getCurrentUserFullName()}</p>
          <button onClick={onLogout} className="btn-secondary w-full">Logout</button>
        </div>
      </aside>
      <main className="flex-1 p-6 overflow-x-hidden">
        <div className="max-w-6xl mx-auto">
          <Routes>
            <Route path="/" element={<Navigate to="/dashboard" />} />
            <Route path="/dashboard" element={<Dashboard />} />
            <Route path="/accounts" element={<Accounts />} />
            <Route path="/users" element={<Users />} />
            <Route path="/transactions" element={<Transactions />} />
            <Route path="/transfers" element={<Transfers />} />
            <Route path="/budgets" element={<Budgets />} />
            <Route path="/currencies" element={<Currencies />} />
            <Route path="/expense-categories" element={<ExpenseCategories />} />
            <Route path="/banks" element={<Banks />} />
            <Route path="/investments" element={<Investments />} />
            <Route path="/recurring-transactions" element={<RecurringTransactions />} />
            <Route path="/account-currencies" element={<AccountCurrencies />} />
            <Route path="/cash-flow-summaries" element={<CashFlowSummaries />} />
            <Route path="*" element={<Navigate to="/dashboard" />} />
          </Routes>
        </div>
      </main>
    </div>
  );
}

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(!!localStorage.getItem('accessToken'));

  useEffect(() => {
    const onUnauthenticated = () => setIsLoggedIn(false);
    window.addEventListener('auth:unauthenticated', onUnauthenticated);
    return () => window.removeEventListener('auth:unauthenticated', onUnauthenticated);
  }, []);

  const handleLogin = () => setIsLoggedIn(true);

  const handleLogout = async () => {
    const refreshToken = localStorage.getItem('refreshToken');
    try {
      if (refreshToken) {
        await axios.post('/logout', { refresh_token: refreshToken });
      }
    } catch {
      // best-effort: log out locally regardless of server response
    }
    clearSession();
    setIsLoggedIn(false);
  };

  if (!isLoggedIn) {
    return (
      <Router>
        <Routes>
          <Route path="/register" element={<Register onLogin={handleLogin} />} />
          <Route path="*" element={<Login onLogin={handleLogin} />} />
        </Routes>
      </Router>
    );
  }

  return (
    <Router>
      <AppShell onLogout={handleLogout} />
    </Router>
  );
}

export default App;
