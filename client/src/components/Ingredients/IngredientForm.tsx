import React, { useState, useContext } from 'react';
import { Form, FormField, TextInput, Text, TextArea, Box } from 'grommet';
import { observer } from 'mobx-react-lite';
import axios, { AxiosResponse } from 'axios';
import StoreContext from '../../state';
import { FormButton } from '../../form/form.styles';

interface IProps {
  form: any;
  dismissForm: () => void;
}

const IngredientForm: React.FC<IProps> = observer(
  ({ form, dismissForm }: IProps) => {
    const store = useContext(StoreContext);
    const [isSubmitting, setIsSubmitting] = useState(false);

    const handleSubmit = async (e: any) => {
      try {
        setIsSubmitting(true);
        const { isValid } = await form.validate({ showErrors: true });
        if (isValid) {
          const { data }: AxiosResponse = await axios.post(
            'http://localhost:5000/api/v1/ingredients',
            JSON.stringify(form.values())
          );
          setIsSubmitting(false);
          store.setIngredients(data);
          form.reset();
          form.showErrors(false);
          dismissForm();
        }
      } catch (err) {
        console.log(err);
      }
    };

    const handleCancel = () => {
      form.reset();
      setIsSubmitting(false);
      dismissForm();
    };

    return (
      <Form style={{ width: '100%' }}>
        <FormField label={form.$('name').label}>
          <TextInput {...form.$('name').bind()} />
        </FormField>
        <p>{form.$('name').error}</p>
        <FormField label={form.$('notes').label}>
          <Box height="small">
            <TextArea fill {...form.$('notes').bind()} />
          </Box>
        </FormField>
        <Box direction="row" justify="end">
          {!isSubmitting && (
            <>
              <FormButton primary label="Submit" onClick={handleSubmit} />
              <FormButton
                color="accent-4"
                label="Cancel"
                onClick={handleCancel}
              />
            </>
          )}
          {isSubmitting && <Text>Submitting...</Text>}
        </Box>
      </Form>
    );
  }
);

export default IngredientForm;
