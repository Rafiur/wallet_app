import CrudPage from './CrudPage';
import { getCurrentUserId } from '../api';

const fields = [
  { name: 'user_id', type: 'hidden' },
  {
    name: 'account_id',
    label: 'Account',
    type: 'select',
    source: { resource: 'accounts', valueField: 'id', labelField: 'name' },
  },
  { name: 'name', label: 'Name' },
  { name: 'amount', label: 'Amount', type: 'number' },
  {
    name: 'transaction_type',
    label: 'Type',
    type: 'select',
    options: [
      { value: 'income', label: 'Income' },
      { value: 'expense', label: 'Expense' },
    ],
  },
  {
    name: 'expense_category_id',
    label: 'Category',
    type: 'select',
    required: false,
    source: { resource: 'expense-categories', valueField: 'id', labelField: 'name' },
  },
  {
    name: 'frequency',
    label: 'Frequency',
    type: 'select',
    options: [
      { value: 'monthly', label: 'Monthly' },
      { value: 'bi-monthly', label: 'Bi-monthly' },
      { value: 'quarterly', label: 'Quarterly' },
      { value: 'semi-annual', label: 'Semi-annual' },
      { value: 'annual', label: 'Annual' },
    ],
  },
  { name: 'start_date', label: 'Start Date', type: 'date' },
  { name: 'next_due_date', label: 'Next Due Date', type: 'date' },
  { name: 'end_date', label: 'End Date', type: 'date', required: false },
  { name: 'note', label: 'Note', type: 'textarea', required: false },
];
const listFields = ['name', 'amount', 'transaction_type', 'frequency', 'next_due_date'];

function RecurringTransactions() {
  return (
    <CrudPage
      resource="recurring-transactions"
      fields={fields}
      listFields={listFields}
      defaultForm={{ user_id: getCurrentUserId() }}
    />
  );
}

export default RecurringTransactions;
