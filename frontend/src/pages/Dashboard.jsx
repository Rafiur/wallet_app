import { useState, useEffect, useMemo } from 'react';
import { Link } from 'react-router-dom';
import axios from '../api';

function formatMoney(amount, currency) {
  const n = Number(amount) || 0;
  return `${n < 0 ? '-' : ''}${currency ? currency + ' ' : ''}${Math.abs(n).toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })}`;
}

function shiftPeriod(period, delta) {
  const [y, m] = period.split('-').map(Number);
  const date = new Date(Date.UTC(y, m - 1 + delta, 1));
  return `${date.getUTCFullYear()}-${String(date.getUTCMonth() + 1).padStart(2, '0')}`;
}

function StatTile({ label, value, tone = 'default' }) {
  const toneClass = {
    default: 'text-slate-800',
    good: 'text-good',
    critical: 'text-critical',
  }[tone];
  return (
    <div className="card p-5">
      <p className="text-sm text-slate-500">{label}</p>
      <p className={`mt-1 text-2xl font-semibold tabular-nums ${toneClass}`}>{value}</p>
    </div>
  );
}

function RankedBars({ title, rows, emptyText }) {
  const max = Math.max(1, ...rows.map((r) => r.amount));
  return (
    <div className="card p-5">
      <h3 className="text-sm font-semibold text-slate-600 mb-4">{title}</h3>
      {rows.length === 0 && <p className="text-sm text-slate-400">{emptyText}</p>}
      <ul className="space-y-3">
        {rows.map((r) => (
          <li key={r.category}>
            <div className="flex items-center justify-between text-sm mb-1">
              <span className="text-slate-700">{r.category}</span>
              <span className="tabular-nums text-slate-600">{formatMoney(r.amount)}</span>
            </div>
            <div className="h-2 rounded-full bg-slate-100">
              <div
                className="h-2 rounded-full bg-brand-500"
                style={{ width: `${Math.max(4, (r.amount / max) * 100)}%` }}
              />
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
}

function Dashboard() {
  const now = new Date();
  const currentPeriod = `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, '0')}`;
  const [period, setPeriod] = useState(currentPeriod);
  const [data, setData] = useState(null);
  const [categoryNames, setCategoryNames] = useState({});
  const [error, setError] = useState('');

  useEffect(() => {
    axios.get('/expense-categories').then((res) => {
      const map = {};
      (res.data || []).forEach((c) => { map[c.id] = c.name; });
      setCategoryNames(map);
    }).catch(() => {});
  }, []);

  useEffect(() => {
    let cancelled = false;
    setError('');
    axios.get('/dashboard', { params: { period } })
      .then((res) => { if (!cancelled) setData(res.data); })
      .catch((err) => { if (!cancelled) setError(err.response?.data?.error || 'Failed to load dashboard'); });
    return () => { cancelled = true; };
  }, [period]);

  const periodLabel = useMemo(() => {
    const [y, m] = period.split('-').map(Number);
    return new Date(Date.UTC(y, m - 1, 1)).toLocaleDateString(undefined, { month: 'long', year: 'numeric', timeZone: 'UTC' });
  }, [period]);

  if (error) {
    return <div className="rounded-md bg-red-50 border border-red-200 text-red-700 px-4 py-3 text-sm">{error}</div>;
  }
  if (!data) {
    return <p className="text-slate-500">Loading...</p>;
  }

  return (
    <div>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-semibold text-slate-800">Dashboard</h1>
        <div className="flex items-center gap-2">
          <button className="btn-secondary px-3" onClick={() => setPeriod((p) => shiftPeriod(p, -1))}>&larr;</button>
          <span className="text-sm font-medium text-slate-600 w-36 text-center">{periodLabel}</span>
          <button className="btn-secondary px-3" onClick={() => setPeriod((p) => shiftPeriod(p, 1))}>&rarr;</button>
        </div>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
        <StatTile label="Total Balance" value={formatMoney(data.total_balance)} />
        <StatTile label="Income" value={formatMoney(data.income_total)} tone="good" />
        <StatTile label="Expenses" value={formatMoney(data.expense_total)} tone="critical" />
        <StatTile label="Net" value={formatMoney(data.net)} tone={data.net >= 0 ? 'good' : 'critical'} />
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-4 mb-6">
        <div className="card p-5 lg:col-span-1">
          <h3 className="text-sm font-semibold text-slate-600 mb-4">Accounts</h3>
          {(data.accounts || []).length === 0 && (
            <p className="text-sm text-slate-400">No accounts yet. <Link to="/accounts" className="text-brand-600 hover:underline">Add one</Link>.</p>
          )}
          <ul className="space-y-3">
            {(data.accounts || []).map((a) => (
              <li key={a.id} className="flex items-center justify-between text-sm">
                <div>
                  <p className="text-slate-700">{a.name}</p>
                  <p className="text-xs text-slate-400">{a.type}</p>
                </div>
                <span className="tabular-nums text-slate-700">{formatMoney(a.balance, a.currency)}</span>
              </li>
            ))}
          </ul>
        </div>

        <RankedBars title="Expenses by Category" rows={data.expense_by_category || []} emptyText="No expenses recorded this period." />
        <RankedBars title="Income by Category" rows={data.income_by_category || []} emptyText="No income recorded this period." />
      </div>

      <div className="card overflow-x-auto">
        <h3 className="text-sm font-semibold text-slate-600 px-5 pt-5">Recent Transactions</h3>
        <table className="min-w-full divide-y divide-slate-200 text-sm mt-3">
          <thead className="bg-slate-50">
            <tr>
              <th className="px-4 py-2 text-left font-medium text-slate-500">Date</th>
              <th className="px-4 py-2 text-left font-medium text-slate-500">Description</th>
              <th className="px-4 py-2 text-left font-medium text-slate-500">Category</th>
              <th className="px-4 py-2 text-right font-medium text-slate-500">Amount</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {(data.recent_transactions || []).map((t) => (
              <tr key={t.id} className="hover:bg-slate-50">
                <td className="px-4 py-2 whitespace-nowrap">{new Date(t.transaction_date).toLocaleDateString()}</td>
                <td className="px-4 py-2 whitespace-nowrap">{t.transaction_name}</td>
                <td className="px-4 py-2 whitespace-nowrap">{categoryNames[t.expense_category_id] || 'Uncategorized'}</td>
                <td className={`px-4 py-2 whitespace-nowrap text-right tabular-nums ${t.transaction_type === 'income' ? 'text-good' : 'text-critical'}`}>
                  {formatMoney(t.amount)}
                </td>
              </tr>
            ))}
            {(data.recent_transactions || []).length === 0 && (
              <tr>
                <td colSpan={4} className="px-4 py-6 text-center text-slate-400">No transactions yet this period.</td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default Dashboard;
