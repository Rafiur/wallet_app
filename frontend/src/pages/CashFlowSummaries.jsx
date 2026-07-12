import CrudPage from './CrudPage';
import { getCurrentUserId } from '../api';

const periodOptions = [
  { value: 'daily', label: 'Daily' },
  { value: 'monthly', label: 'Monthly' },
  { value: 'yearly', label: 'Yearly' },
];

const fields = [
  { name: 'user_id', type: 'hidden' },
  {
    name: 'account_id',
    label: 'Account (leave blank for all)',
    type: 'select',
    required: false,
    source: { resource: 'accounts', valueField: 'id', labelField: 'name' },
  },
  { name: 'period', label: 'Period', type: 'select', options: periodOptions },
  { name: 'start_date', label: 'Start Date', type: 'date' },
  { name: 'inflow', label: 'Inflow', type: 'number', required: false },
  { name: 'outflow', label: 'Outflow', type: 'number', required: false },
];
const listFields = ['period', 'start_date', 'inflow', 'outflow', 'net_flow'];
const filterFields = [{ name: 'period', label: 'Period (daily/monthly/yearly)', defaultValue: 'monthly' }];

function CashFlowSummaries() {
  return (
    <CrudPage
      resource="cash-flow-summaries"
      fields={fields}
      listFields={listFields}
      filterFields={filterFields}
      defaultForm={{ user_id: getCurrentUserId() }}
    />
  );
}

export default CashFlowSummaries;
