import CrudPage from './CrudPage';

const fields = ['user_id', 'period', 'amount'];
const listFields = ['id', 'user_id', 'period', 'amount'];

function Budgets() {
  return <CrudPage resource="budgets" fields={fields} listFields={listFields} />;
}

export default Budgets;