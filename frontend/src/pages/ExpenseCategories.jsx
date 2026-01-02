import CrudPage from './CrudPage';

const fields = ['name', 'parent_category_id'];
const listFields = ['id', 'name'];

function ExpenseCategories() {
  return <CrudPage resource="expense-categories" fields={fields} listFields={listFields} />;
}

export default ExpenseCategories;