// @ts-ignore
import dvr from 'mobx-react-form/lib/validators/DVR';
import validatorjs from 'validatorjs';
// @ts-ignore
import MobxReactForm from 'mobx-react-form';

const plugins = {
  dvr: dvr(validatorjs),
};

const fields = [
  {
    name: 'name',
    label: 'Name',
    placeholder: 'Enter Ingredient Name',
    rules: 'required|string',
  },
  {
    name: 'notes',
    label: 'Notes',
    placeholder: 'Enter Notes (Optional)',
    rules: 'string',
  },
];

const form = new MobxReactForm({ fields }, { plugins });

export default form;
