import CrudPage from './CrudPage';

const fields = ['user_id', 'name', 'type', 'currency', 'balance'];
const listFields = ['id', 'name', 'type', 'balance', 'currency'];

function Accounts() {
  return <CrudPage resource="accounts" fields={fields} listFields={listFields} />;
}

export default Accounts;