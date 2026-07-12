import CrudPage from './CrudPage';
import { getCurrentUserId } from '../api';

const fields = [
  { name: 'user_id', type: 'hidden' },
  { name: 'name', label: 'Bank Name' },
  { name: 'branch', label: 'Branch', required: false },
  { name: 'account_number', label: 'Account Number', required: false },
  {
    name: 'account_id',
    label: 'Linked Account',
    type: 'select',
    source: { resource: 'accounts', valueField: 'id', labelField: 'name' },
  },
  { name: 'annual_charge', label: 'Annual Charge', type: 'number', required: false },
];
const listFields = ['name', 'branch', 'account_number', 'annual_charge'];

function Banks() {
  return <CrudPage resource="banks" fields={fields} listFields={listFields} defaultForm={{ user_id: getCurrentUserId() }} />;
}

export default Banks;
