import CrudPage from './CrudPage';

const fields = ['user_id', 'account_id', 'transaction_name', 'amount', 'transaction_type', 'expense_category_id', 'note'];
const listFields = ['id', 'transaction_name', 'amount', 'transaction_type'];

function Transactions() {
  return <CrudPage resource="transactions" fields={fields} listFields={listFields} />;
}

export default Transactions;