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
  { name: 'transaction_name', label: 'Description' },
  { name: 'amount', label: 'Amount', type: 'number' }, // always entered as a positive number; sign is derived from type on submit
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
  { name: 'transaction_date', label: 'Date', type: 'date', required: false },
  { name: 'note', label: 'Note', type: 'textarea', required: false },
];
const listFields = ['transaction_name', 'amount', 'transaction_type', 'transaction_date'];

// The API stores signed amounts (negative = expense, positive = income) so account
// balances can be updated with a plain addition; the form always collects a positive number.
const transformPayload = (payload) => ({
  ...payload,
  amount: payload.transaction_type === 'expense' ? -Math.abs(payload.amount) : Math.abs(payload.amount),
});
const transformEditItem = (item) => ({ ...item, amount: Math.abs(item.amount) });

function Transactions() {
  return (
    <CrudPage
      resource="transactions"
      fields={fields}
      listFields={listFields}
      defaultForm={{ user_id: getCurrentUserId() }}
      transformPayload={transformPayload}
      transformEditItem={transformEditItem}
    />
  );
}

export default Transactions;
