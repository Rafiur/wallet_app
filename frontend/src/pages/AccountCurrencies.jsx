import CrudPage from './CrudPage';

const fields = ['currency_code', 'account_id'];
const listFields = ['id', 'currency_code', 'account_id'];

function AccountCurrencies() {
  return <CrudPage resource="account-currencies" fields={fields} listFields={listFields} />;
}

export default AccountCurrencies;