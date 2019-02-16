import React, { useContext } from 'react';
import { observer } from 'mobx-react-lite';
import { Heading, Box, Button, Grommet } from 'grommet';
import { useGetData, useGetIngredients } from '../../hooks';
import { IIngredient } from '../../domain/Ingredient';
import { IngredientsTable } from '.';
import { IngredientHeader } from './ingredients.styles';
import { Add } from 'grommet-icons';
import StoreContext from '../../state';

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
          <Heading level="3" margin="none">
            Ingredients
          </Heading>
          <Button
            style={{ margin: '5px 0' }}
            primary
            icon={<Add size="small" />}
            label="Add"
            onClick={() => alert('Add')}
          />
        </IngredientHeader>
        <IngredientsTable ingredients={store.ingredients} />
      </Box>
    </Grommet>
  );
});

export default Ingredients;
