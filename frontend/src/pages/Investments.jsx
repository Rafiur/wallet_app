import CrudPage from './CrudPage';

const fields = ['user_id', 'type', 'amount'];
const listFields = ['id', 'type', 'amount'];

function Investments() {
  return <CrudPage resource="investments" fields={fields} listFields={listFields} />;
}

export default Investments;