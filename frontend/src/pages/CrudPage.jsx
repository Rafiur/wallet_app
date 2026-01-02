import { useState, useEffect } from 'react';
import axios from 'axios';

const API_BASE = 'http://localhost:8000/api/v1'; // Adjust if port is different

function CrudPage({ resource, fields, listFields, idField = 'id' }) {
  const [items, setItems] = useState([]);
  const [form, setForm] = useState({});
  const [editing, setEditing] = useState(null);

  useEffect(() => {
    fetchItems();
  }, []);

  const fetchItems = async () => {
    try {
      const res = await axios.get(`${API_BASE}/${resource}`);
      setItems(res.data);
    } catch (err) {
      console.error(err);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if (editing) {
        await axios.put(`${API_BASE}/${resource}/${editing[idField]}`, form);
      } else {
        await axios.post(`${API_BASE}/${resource}`, form);
      }
      setForm({});
      setEditing(null);
      fetchItems();
    } catch (err) {
      console.error(err);
    }
  };

  const handleEdit = (item) => {
    setForm(item);
    setEditing(item);
  };

  const handleDelete = async (id) => {
    try {
      await axios.delete(`${API_BASE}/${resource}/${id}`);
      fetchItems();
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div>
      <h1>{resource.charAt(0).toUpperCase() + resource.slice(1)}</h1>
      <form onSubmit={handleSubmit}>
        {fields.map(field => (
          <div key={field}>
            <label>{field}: </label>
            <input
              type="text"
              value={form[field] || ''}
              onChange={(e) => setForm({ ...form, [field]: e.target.value })}
              required
            />
          </div>
        ))}
        <button type="submit">{editing ? 'Update' : 'Create'}</button>
        {editing && <button type="button" onClick={() => { setForm({}); setEditing(null); }}>Cancel</button>}
      </form>
      <ul>
        {items.map(item => (
          <li key={item[idField]}>
            {listFields.map(field => `${field}: ${item[field]}`).join(', ')}
            <button onClick={() => handleEdit(item)}>Edit</button>
            <button onClick={() => handleDelete(item[idField])}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default CrudPage;