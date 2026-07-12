import CrudPage from './CrudPage';
import { getCurrentUserId } from '../api';

const fields = [
  { name: 'user_id', type: 'hidden' },
  {
    name: 'account_id',
    label: 'Linked Account',
    type: 'select',
    required: false,
    source: { resource: 'accounts', valueField: 'id', labelField: 'name' },
  },
  {
    name: 'type',
    label: 'Type',
    type: 'select',
    options: [
      { value: 'fdr', label: 'Fixed Deposit (FDR)' },
      { value: 'stocks', label: 'Stocks' },
      { value: 'bonds', label: 'Bonds' },
    ],
  },
  { name: 'amount', label: 'Amount', type: 'number' },
  { name: 'interest_rate', label: 'Interest Rate (%)', type: 'number', required: false },
  { name: 'start_date', label: 'Start Date', type: 'date' },
  { name: 'maturity_date', label: 'Maturity Date', type: 'date', required: false },
  {
    name: 'status',
    label: 'Status',
    type: 'select',
    required: false,
    options: [
      { value: 'active', label: 'Active' },
      { value: 'matured', label: 'Matured' },
    ],
  },
  { name: 'note', label: 'Note', type: 'textarea', required: false },
];
const listFields = ['type', 'amount', 'interest_rate', 'status'];

function Investments() {
  return <CrudPage resource="investments" fields={fields} listFields={listFields} defaultForm={{ user_id: getCurrentUserId() }} />;
}

export default Investments;
