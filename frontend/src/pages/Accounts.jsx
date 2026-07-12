import CrudPage from './CrudPage';
import { getCurrentUserId } from '../api';

const fields = [
  { name: 'user_id', type: 'hidden' },
  { name: 'name', label: 'Name' },
  {
    name: 'type',
    label: 'Type',
    type: 'select',
    options: [
      { value: 'bank', label: 'Bank' },
      { value: 'mobile_banking', label: 'Mobile Banking' },
      { value: 'cash', label: 'Cash' },
      { value: 'wallet', label: 'Wallet' },
      { value: 'investment', label: 'Investment' },
      { value: 'other', label: 'Other' },
    ],
  },
  {
    name: 'currency',
    label: 'Currency',
    type: 'select',
    source: { resource: 'currencies', valueField: 'code', labelField: 'name' },
  },
  { name: 'balance', label: 'Opening Balance', type: 'number' },
];
const listFields = ['name', 'type', 'balance', 'currency'];

function Accounts() {
  return <CrudPage resource="accounts" fields={fields} listFields={listFields} defaultForm={{ user_id: getCurrentUserId() }} />;
}

export default Accounts;
