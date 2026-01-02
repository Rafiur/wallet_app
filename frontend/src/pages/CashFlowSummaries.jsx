import CrudPage from './CrudPage';

const fields = ['user_id', 'period', 'start_date', 'end_date'];
const listFields = ['id', 'user_id', 'period', 'start_date', 'end_date'];

function CashFlowSummaries() {
  return <CrudPage resource="cash-flow-summaries" fields={fields} listFields={listFields} />;
}

export default CashFlowSummaries;