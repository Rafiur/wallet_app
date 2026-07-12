import { useState, useEffect, useMemo } from 'react';
import axios from '../api';

// Normalizes a field entry: plain strings become basic text fields, objects pass through.
function normalizeField(field) {
  if (typeof field === 'string') {
    return { name: field, label: field, type: 'text' };
  }
  return { label: field.name, type: 'text', required: true, ...field };
}

function CrudPage({
  resource, fields, listFields, idField = 'id', filterFields = [], defaultForm = {},
  transformPayload = (p) => p, transformEditItem = (i) => i,
}) {
  const normalizedFields = useMemo(() => fields.map(normalizeField), [fields]);

  const [items, setItems] = useState([]);
  const [form, setForm] = useState(defaultForm);
  const [editing, setEditing] = useState(null);
  const [error, setError] = useState('');
  const [optionsByField, setOptionsByField] = useState({});
  const [filters, setFilters] = useState(() => {
    const initial = {};
    filterFields.forEach((f) => { initial[f.name] = f.defaultValue || ''; });
    return initial;
  });

  const filtersReady = filterFields.every((f) => filters[f.name]);

  // Fetch dropdown options for any field or filter backed by another resource, once.
  useEffect(() => {
    const sourced = [...normalizedFields, ...filterFields].filter((f) => f.source);
    sourced.forEach(async (f) => {
      try {
        const res = await axios.get(`/${f.source.resource}`);
        setOptionsByField((prev) => ({ ...prev, [f.name]: res.data || [] }));
      } catch {
        // silently ignore; the select just renders empty
      }
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [resource]);

  useEffect(() => {
    if (filtersReady) {
      fetchItems();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [JSON.stringify(filters)]);

  const fetchItems = async () => {
    try {
      setError('');
      const res = await axios.get(`/${resource}`, { params: filters });
      setItems(res.data || []);
    } catch (err) {
      setError(err.response?.data?.error || 'Failed to load data');
    }
  };

  // HTML inputs only ever produce strings; coerce to the types the API expects
  // (numbers as numbers, dates as RFC3339) and drop blank optional fields
  // rather than sending "" for a UUID/number column.
  const buildPayload = () => {
    const payload = { ...form };
    normalizedFields.forEach((f) => {
      const val = payload[f.name];
      const isEmpty = val === '' || val === undefined || val === null;
      if (isEmpty) {
        if (!f.required) delete payload[f.name];
        return;
      }
      if (f.type === 'number') {
        payload[f.name] = Number(val);
      } else if (f.type === 'date') {
        payload[f.name] = `${val}T00:00:00Z`;
      }
    });
    return transformPayload(payload);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      setError('');
      const payload = buildPayload();
      if (editing) {
        await axios.put(`/${resource}/${editing[idField]}`, payload);
      } else {
        await axios.post(`/${resource}`, payload);
      }
      setForm(defaultForm);
      setEditing(null);
      fetchItems();
    } catch (err) {
      setError(err.response?.data?.error || 'Save failed');
    }
  };

  const handleEdit = (item) => {
    const next = { ...item };
    normalizedFields.forEach((f) => {
      // <input type="date"> needs YYYY-MM-DD; the API returns RFC3339 timestamps.
      if (f.type === 'date' && typeof next[f.name] === 'string') {
        next[f.name] = next[f.name].slice(0, 10);
      }
    });
    setForm(transformEditItem(next));
    setEditing(item);
  };

  const handleCancel = () => {
    setForm(defaultForm);
    setEditing(null);
  };

  const handleDelete = async (id) => {
    try {
      setError('');
      await axios.delete(`/${resource}/${id}`);
      fetchItems();
    } catch (err) {
      setError(err.response?.data?.error || 'Delete failed');
    }
  };

  // Resolves a raw value (e.g. a foreign-key id) to a human label when the field
  // has select options, so the list shows "Checking Account" instead of a UUID.
  const displayValue = (fieldName, rawValue) => {
    const fieldDef = normalizedFields.find((f) => f.name === fieldName);
    if (!fieldDef || rawValue === undefined || rawValue === null || rawValue === '') return rawValue ?? '';
    const options = fieldDef.source ? optionsByField[fieldDef.name] : fieldDef.options;
    if (!options) return rawValue;
    const valueField = fieldDef.source?.valueField || 'value';
    const labelField = fieldDef.source?.labelField || 'label';
    const match = options.find((o) => (fieldDef.source ? o[valueField] : o.value) === rawValue);
    if (!match) return rawValue;
    return fieldDef.source ? match[labelField] : match.label;
  };

  const renderInput = (f) => {
    const value = form[f.name] ?? '';
    if (f.type === 'select') {
      const options = f.source ? (optionsByField[f.name] || []) : (f.options || []);
      const valueField = f.source?.valueField || 'value';
      const labelField = f.source?.labelField || 'label';
      return (
        <select
          className="field-input"
          value={value}
          required={f.required}
          onChange={(e) => setForm({ ...form, [f.name]: e.target.value })}
        >
          <option value="" disabled>Select {f.label}</option>
          {options.map((o) => (
            <option key={o[valueField]} value={o[valueField]}>{o[labelField]}</option>
          ))}
        </select>
      );
    }
    if (f.type === 'textarea') {
      return (
        <textarea
          className="field-input"
          rows={3}
          value={value}
          required={f.required}
          onChange={(e) => setForm({ ...form, [f.name]: e.target.value })}
        />
      );
    }
    const inputType = ['number', 'date', 'password', 'email'].includes(f.type) ? f.type : 'text';
    return (
      <input
        className="field-input"
        type={inputType}
        step={f.type === 'number' ? 'any' : undefined}
        value={value}
        required={f.required}
        onChange={(e) => setForm({ ...form, [f.name]: e.target.value })}
      />
    );
  };

  const visibleFields = normalizedFields.filter((f) => f.type !== 'hidden');

  return (
    <div>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-semibold text-slate-800">
          {resource.split('-').map((w) => w[0].toUpperCase() + w.slice(1)).join(' ')}
        </h1>
      </div>

      {filterFields.length > 0 && (
        <div className="card p-4 mb-6 flex flex-wrap items-end gap-4">
          {filterFields.map((f) => {
            const options = f.source ? (optionsByField[f.name] || []) : f.options;
            const valueField = f.source?.valueField || 'value';
            const labelField = f.source?.labelField || 'label';
            return (
              <div key={f.name}>
                <label className="field-label">{f.label || f.name}</label>
                {options ? (
                  <select
                    className="field-input"
                    value={filters[f.name] || ''}
                    onChange={(e) => setFilters({ ...filters, [f.name]: e.target.value })}
                  >
                    <option value="" disabled>Select {f.label}</option>
                    {options.map((o) => (
                      <option key={o[valueField]} value={o[valueField]}>{o[labelField]}</option>
                    ))}
                  </select>
                ) : (
                  <input
                    className="field-input"
                    type="text"
                    value={filters[f.name] || ''}
                    onChange={(e) => setFilters({ ...filters, [f.name]: e.target.value })}
                  />
                )}
              </div>
            );
          })}
          {!filtersReady && <span className="text-sm text-slate-500">Select the filter(s) above to load data.</span>}
        </div>
      )}

      <form onSubmit={handleSubmit} className="card p-6 mb-6">
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
          {visibleFields.map((f) => (
            <div key={f.name}>
              <label className="field-label">{f.label}</label>
              {renderInput(f)}
            </div>
          ))}
        </div>
        <div className="mt-4 flex gap-2">
          <button type="submit" className="btn-primary">{editing ? 'Update' : 'Create'}</button>
          {editing && <button type="button" className="btn-secondary" onClick={handleCancel}>Cancel</button>}
        </div>
      </form>

      {error && (
        <div className="mb-4 rounded-md bg-red-50 border border-red-200 text-red-700 px-4 py-3 text-sm">{error}</div>
      )}

      <div className="card overflow-x-auto">
        <table className="min-w-full divide-y divide-slate-200 text-sm">
          <thead className="bg-slate-50">
            <tr>
              {listFields.map((f) => (
                <th key={f} className="px-4 py-2 text-left font-medium text-slate-500">{f}</th>
              ))}
              <th className="px-4 py-2" />
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {items.map((item) => (
              <tr key={item[idField]} className="hover:bg-slate-50">
                {listFields.map((f) => (
                  <td key={f} className="px-4 py-2 whitespace-nowrap">{String(displayValue(f, item[f]))}</td>
                ))}
                <td className="px-4 py-2 whitespace-nowrap text-right space-x-2">
                  <button className="btn-secondary" onClick={() => handleEdit(item)}>Edit</button>
                  <button className="btn-danger" onClick={() => handleDelete(item[idField])}>Delete</button>
                </td>
              </tr>
            ))}
            {items.length === 0 && (
              <tr>
                <td colSpan={listFields.length + 1} className="px-4 py-6 text-center text-slate-400">
                  No records yet.
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default CrudPage;
