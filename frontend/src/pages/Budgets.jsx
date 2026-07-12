import CrudPage from './CrudPage';
import { getCurrentUserId } from '../api';

const fields = [
  { name: 'user_id', type: 'hidden' },
  {
    name: 'expense_category_id',
    label: 'Category',
    type: 'select',
    required: false,
    source: { resource: 'expense-categories', valueField: 'id', labelField: 'name' },
  },
  {
    name: 'period',
    label: 'Period',
    type: 'select',
    options: [
      { value: 'monthly', label: 'Monthly' },
      { value: 'quarterly', label: 'Quarterly' },
      { value: 'yearly', label: 'Yearly' },
    ],
  },
  { name: 'amount', label: 'Limit', type: 'number' },
  { name: 'start_date', label: 'Start Date', type: 'date' },
  { name: 'end_date', label: 'End Date', type: 'date', required: false },
];
const listFields = ['period', 'amount', 'start_date', 'end_date'];

function Budgets() {
  return <CrudPage resource="budgets" fields={fields} listFields={listFields} defaultForm={{ user_id: getCurrentUserId() }} />;
}

export default Budgets;
