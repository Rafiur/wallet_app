import CrudPage from './CrudPage';

const fields = [
  {
    name: 'from_account_id',
    label: 'From Account',
    type: 'select',
    source: { resource: 'accounts', valueField: 'id', labelField: 'name' },
  },
  {
    name: 'to_account_id',
    label: 'To Account',
    type: 'select',
    source: { resource: 'accounts', valueField: 'id', labelField: 'name' },
  },
  { name: 'amount', label: 'Amount', type: 'number' },
  {
    name: 'currency',
    label: 'Currency',
    type: 'select',
    source: { resource: 'currencies', valueField: 'code', labelField: 'name' },
  },
  { name: 'exchange_rate', label: 'Exchange Rate', type: 'number', required: false },
  { name: 'note', label: 'Note', type: 'textarea', required: false },
];
const listFields = ['from_account_id', 'to_account_id', 'amount', 'status'];

function Transfers() {
  return <CrudPage resource="transfers" fields={fields} listFields={listFields} />;
}

export default Transfers;
