import CrudPage from './CrudPage';

const accountSource = { resource: 'accounts', valueField: 'id', labelField: 'name' };

const fields = [
  { name: 'account_id', label: 'Account', type: 'select', source: accountSource },
  {
    name: 'currency_code',
    label: 'Currency',
    type: 'select',
    source: { resource: 'currencies', valueField: 'code', labelField: 'name' },
  },
  { name: 'balance_according_to_currency', label: 'Balance', type: 'number' },
];
const listFields = ['account_id', 'currency_code', 'balance_according_to_currency'];
const filterFields = [{ name: 'account_id', label: 'Account', source: accountSource }];

function AccountCurrencies() {
  return (
    <CrudPage
      resource="account-currencies"
      fields={fields}
      listFields={listFields}
      filterFields={filterFields}
    />
  );
}

export default AccountCurrencies;
