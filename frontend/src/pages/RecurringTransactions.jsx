import CrudPage from './CrudPage';

const fields = ['user_id', 'account_id', 'name', 'frequency', 'amount'];
const listFields = ['id', 'name', 'frequency', 'amount'];

function RecurringTransactions() {
  return <CrudPage resource="recurring-transactions" fields={fields} listFields={listFields} />;
}

export default RecurringTransactions;