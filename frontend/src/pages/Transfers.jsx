import CrudPage from './CrudPage';

const fields = ['from_account_id', 'to_account_id', 'amount', 'currency', 'exchange_rate', 'note'];
const listFields = ['id', 'from_account_id', 'to_account_id', 'amount', 'status'];

function Transfers() {
  return <CrudPage resource="transfers" fields={fields} listFields={listFields} />;
}

export default Transfers;