import CrudPage from './CrudPage';

const fields = ['full_name', 'email', 'password'];
const listFields = ['id', 'full_name', 'email'];

function Users() {
  return <CrudPage resource="users" fields={fields} listFields={listFields} />;
}

export default Users;