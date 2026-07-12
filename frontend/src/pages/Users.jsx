import CrudPage from './CrudPage';

const fields = [
  { name: 'full_name', label: 'Full Name' },
  { name: 'email', label: 'Email' },
  { name: 'password', label: 'Password', type: 'password' },
];
const listFields = ['full_name', 'email'];

function Users() {
  return <CrudPage resource="users" fields={fields} listFields={listFields} />;
}

export default Users;
