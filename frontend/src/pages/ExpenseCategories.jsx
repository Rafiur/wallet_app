import { useState, useEffect, useMemo } from 'react';
import axios from '../api';

const emptyForm = { name: '', type: 'expense', parent_category_id: '' };

function ExpenseCategories() {
  const [categories, setCategories] = useState([]);
  const [form, setForm] = useState(emptyForm);
  const [editing, setEditing] = useState(null);
  const [error, setError] = useState('');

  useEffect(() => { fetchCategories(); }, []);

  const fetchCategories = async () => {
    try {
      setError('');
      const res = await axios.get('/expense-categories');
      setCategories(res.data || []);
    } catch (err) {
      setError(err.response?.data?.error || 'Failed to load categories');
    }
  };

  // Only top-level categories of the same type can be a parent.
  const parentOptions = useMemo(
    () => categories.filter((c) => c.type === form.type && !c.parent_category_id && c.id !== editing?.id),
    [categories, form.type, editing]
  );

  const grouped = useMemo(() => {
    const topLevel = categories.filter((c) => !c.parent_category_id);
    const childrenOf = (id) => categories.filter((c) => c.parent_category_id === id);
    return topLevel.map((parent) => ({ parent, children: childrenOf(parent.id) }));
  }, [categories]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      setError('');
      const payload = { ...form, parent_category_id: form.parent_category_id || null };
      if (editing) {
        await axios.put(`/expense-categories/${editing.id}`, payload);
      } else {
        await axios.post('/expense-categories', payload);
      }
      setForm(emptyForm);
      setEditing(null);
      fetchCategories();
    } catch (err) {
      setError(err.response?.data?.error || 'Save failed');
    }
  };

  const handleEdit = (cat) => {
    setForm({ name: cat.name, type: cat.type, parent_category_id: cat.parent_category_id || '' });
    setEditing(cat);
  };

  const handleCancel = () => {
    setForm(emptyForm);
    setEditing(null);
  };

  const handleDelete = async (id) => {
    try {
      setError('');
      await axios.delete(`/expense-categories/${id}`);
      fetchCategories();
    } catch (err) {
      setError(err.response?.data?.error || 'Delete failed');
    }
  };

  return (
    <div>
      <h1 className="text-2xl font-semibold text-slate-800 mb-6">Expense Categories</h1>

      <form onSubmit={handleSubmit} className="card p-6 mb-6">
        <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
          <div>
            <label className="field-label">Name</label>
            <input
              className="field-input"
              value={form.name}
              onChange={(e) => setForm({ ...form, name: e.target.value })}
              required
            />
          </div>
          <div>
            <label className="field-label">Type</label>
            <select
              className="field-input"
              value={form.type}
              onChange={(e) => setForm({ ...form, type: e.target.value, parent_category_id: '' })}
            >
              <option value="expense">Expense</option>
              <option value="income">Income</option>
            </select>
          </div>
          <div>
            <label className="field-label">Parent Category</label>
            <select
              className="field-input"
              value={form.parent_category_id}
              onChange={(e) => setForm({ ...form, parent_category_id: e.target.value })}
            >
              <option value="">None (top-level)</option>
              {parentOptions.map((c) => (
                <option key={c.id} value={c.id}>{c.name}</option>
              ))}
            </select>
          </div>
        </div>
        <div className="mt-4 flex gap-2">
          <button type="submit" className="btn-primary">{editing ? 'Update' : 'Create'}</button>
          {editing && <button type="button" className="btn-secondary" onClick={handleCancel}>Cancel</button>}
        </div>
      </form>

      {error && (
        <div className="mb-4 rounded-md bg-red-50 border border-red-200 text-red-700 px-4 py-3 text-sm">{error}</div>
      )}

      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {['expense', 'income'].map((type) => (
          <div key={type} className="card p-4">
            <h2 className="text-sm font-semibold uppercase tracking-wide text-slate-500 mb-3">{type}</h2>
            <ul className="space-y-2">
              {grouped.filter((g) => g.parent.type === type).map(({ parent, children }) => (
                <li key={parent.id}>
                  <div className="flex items-center justify-between rounded-md px-2 py-1 hover:bg-slate-50">
                    <span className="font-medium text-slate-700">
                      {parent.name}
                      {!parent.user_id && <span className="ml-2 text-xs text-slate-400">(global)</span>}
                    </span>
                    {parent.user_id && (
                      <span className="space-x-2">
                        <button className="text-xs text-brand-600 hover:underline" onClick={() => handleEdit(parent)}>Edit</button>
                        <button className="text-xs text-red-600 hover:underline" onClick={() => handleDelete(parent.id)}>Delete</button>
                      </span>
                    )}
                  </div>
                  {children.length > 0 && (
                    <ul className="ml-4 mt-1 space-y-1 border-l border-slate-200 pl-3">
                      {children.map((child) => (
                        <li key={child.id} className="flex items-center justify-between rounded-md px-2 py-1 hover:bg-slate-50">
                          <span className="text-slate-600">
                            {child.name}
                            {!child.user_id && <span className="ml-2 text-xs text-slate-400">(global)</span>}
                          </span>
                          {child.user_id && (
                            <span className="space-x-2">
                              <button className="text-xs text-brand-600 hover:underline" onClick={() => handleEdit(child)}>Edit</button>
                              <button className="text-xs text-red-600 hover:underline" onClick={() => handleDelete(child.id)}>Delete</button>
                            </span>
                          )}
                        </li>
                      ))}
                    </ul>
                  )}
                </li>
              ))}
            </ul>
          </div>
        ))}
      </div>
    </div>
  );
}

export default ExpenseCategories;
