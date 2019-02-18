import React, { useContext, useState } from 'react';
import { observer } from 'mobx-react-lite';
import { Heading, Box, Button, Grommet } from 'grommet';
import { useGetIngredients } from '../../hooks';
import { IngredientsTable, IngredientForm } from '.';
import { IngredientHeader } from './ingredients.styles';
import { Add } from 'grommet-icons';
import StoreContext from '../../state';
import { ingredientForm } from '../../form';

const customTheme = {
  text: {
    medium: {
      size: '14px',
      height: '18px;',
    },
  },
  button: {
    padding: {
      vertical: '2px',
      horizontal: '12px',
    },
  },
};

const Ingredients: React.FC = observer(() => {
  const store = useContext(StoreContext);
  const [showNewIngredientForm, setShowNewIngredientForm] = useState(false);
  const { isLoading, isError, error } = useGetIngredients();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error: {error}</div>;
  }

  return (
    <Grommet theme={customTheme}>
      <Box align="start" pad="large" style={{ position: 'relative' }}>
        <IngredientHeader>
          {!showNewIngredientForm && (
            <>
              <Heading level="3" margin="none">
                Ingredients
              </Heading>
              <Button
                style={{ margin: '5px 0' }}
                primary
                icon={<Add size="small" />}
                label="Add"
                onClick={() => setShowNewIngredientForm(true)}
              />
            </>
          )}
          {showNewIngredientForm && (
            <>
              <Heading level="3" margin="none">
                Add New Ingredient
              </Heading>
            </>
          )}
        </IngredientHeader>
        {!showNewIngredientForm && (
          <IngredientsTable ingredients={store.ingredients} />
        )}
        {showNewIngredientForm && (
          <Box direction="row" width="medium">
            <IngredientForm
              form={ingredientForm}
              dismissForm={() => setShowNewIngredientForm(false)}
            />
          </Box>
        )}
      </Box>
    </Grommet>
  );
});

export default Ingredients;
