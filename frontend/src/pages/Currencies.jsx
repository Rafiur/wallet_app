import CrudPage from './CrudPage';

const fields = ['code', 'name', 'symbol'];
const listFields = ['code', 'name', 'symbol'];

function Currencies() {
  return <CrudPage resource="currencies" fields={fields} listFields={listFields} idField="code" />;
}

export default Currencies;