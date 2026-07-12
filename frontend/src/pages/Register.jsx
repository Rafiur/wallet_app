import { useState } from 'react';
import { Link } from 'react-router-dom';
import axios, { storeSession } from '../api';

function Register({ onLogin }) {
  const [fullName, setFullName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setLoading(true);
    try {
      await axios.post('/register', { full_name: fullName, email, password });
      const res = await axios.post('/login', { email, password });
      storeSession(res.data);
      onLogin();
    } catch (err) {
      setError(err.response?.data?.error || 'Registration failed');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-slate-50 px-4">
      <div className="w-full max-w-sm">
        <h1 className="text-2xl font-semibold text-center text-slate-800 mb-1">Wallet</h1>
        <p className="text-center text-slate-500 mb-6">Create your account</p>
        <form onSubmit={handleSubmit} className="card p-6 space-y-4">
          <div>
            <label className="field-label">Full Name</label>
            <input className="field-input" type="text" value={fullName} onChange={(e) => setFullName(e.target.value)} required autoFocus />
          </div>
          <div>
            <label className="field-label">Email</label>
            <input className="field-input" type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
          </div>
          <div>
            <label className="field-label">Password</label>
            <input className="field-input" type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
          </div>
          {error && <p className="text-sm text-red-600">{error}</p>}
          <button type="submit" className="btn-primary w-full" disabled={loading}>
            {loading ? 'Creating account...' : 'Register'}
          </button>
        </form>
        <p className="text-center text-sm text-slate-500 mt-4">
          Already have an account? <Link to="/" className="text-brand-600 hover:underline">Login</Link>
        </p>
      </div>
    </div>
  );
}

export default Register;
