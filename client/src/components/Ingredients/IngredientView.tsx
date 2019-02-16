import React, { useContext, useEffect, useState } from 'react';
import { observer } from 'mobx-react-lite';
import { match } from 'react-router';
import { useGetIngredients } from '../../hooks';
import StoreContext from '../../state';

interface IProps {
  match: match<{ id: string }>;
}

const IngredientView: React.FC<IProps> = observer(
  ({ match }: { match: match<{ id: string }> }) => {
    const store = useContext(StoreContext);
    const { isLoading, isError, error } = useGetIngredients(
      parseInt(match.params.id)
    );

    if (isLoading) {
      return <div>Loading...</div>;
    }

    if (isError) {
      return <div>Error: {error}</div>;
    }

    if (!store.selectedIngredient) {
      return <div>No Ingredient Selected!</div>;
    }

    return <div>{store.selectedIngredient.name}</div>;
  }
);

export default IngredientView;
