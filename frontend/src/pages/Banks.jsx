import CrudPage from './CrudPage';

const fields = ['name', 'account_id', 'account_number'];
const listFields = ['id', 'name', 'account_number'];

function Banks() {
  return <CrudPage resource="banks" fields={fields} listFields={listFields} />;
}

export default Banks;