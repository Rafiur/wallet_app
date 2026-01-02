import { BrowserRouter as Router, Routes, Route, Link, Navigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import axios from 'axios';
import Login from './pages/Login';
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

const API_BASE = 'http://localhost:8000/api/v1';

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('accessToken');
    if (token) {
      setIsLoggedIn(true);
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
    }
  }, []);

  const handleLogin = () => {
    setIsLoggedIn(true);
  };

  const handleLogout = () => {
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
    delete axios.defaults.headers.common['Authorization'];
    setIsLoggedIn(false);
  };

  // Axios interceptor to refresh token
  axios.interceptors.response.use(
    (response) => response,
    async (error) => {
      if (error.response?.status === 401) {
        const refreshToken = localStorage.getItem('refreshToken');
        if (refreshToken) {
          try {
            // Assume endpoint to refresh
            const res = await axios.post(`${API_BASE}/refresh`, { refresh_token: refreshToken }, {
              headers: { Authorization: undefined } // Remove auth header for refresh
            });
            localStorage.setItem('accessToken', res.data.access_token);
            axios.defaults.headers.common['Authorization'] = `Bearer ${res.data.access_token}`;
            // Retry original request
            return axios(error.config);
          } catch (err) {
            handleLogout();
          }
        } else {
          handleLogout();
        }
      }
      return Promise.reject(error);
    }
  );

  if (!isLoggedIn) {
    return <Login onLogin={handleLogin} />;
  }

  return (
    <Router>
      <div className="app">
        <nav>
          <ul>
            <li><Link to="/accounts">Accounts</Link></li>
            <li><Link to="/users">Users</Link></li>
            <li><Link to="/transactions">Transactions</Link></li>
            <li><Link to="/transfers">Transfers</Link></li>
            <li><Link to="/budgets">Budgets</Link></li>
            <li><Link to="/currencies">Currencies</Link></li>
            <li><Link to="/expense-categories">Expense Categories</Link></li>
            <li><Link to="/banks">Banks</Link></li>
            <li><Link to="/investments">Investments</Link></li>
            <li><Link to="/recurring-transactions">Recurring Transactions</Link></li>
            <li><Link to="/account-currencies">Account Currencies</Link></li>
            <li><Link to="/cash-flow-summaries">Cash Flow Summaries</Link></li>
            <li><button onClick={handleLogout} style={{ background: 'none', border: 'none', color: 'white', cursor: 'pointer' }}>Logout</button></li>
          </ul>
        </nav>
        <main>
          <Routes>
            <Route path="/" element={<Navigate to="/accounts" />} />
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
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;