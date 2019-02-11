import React from 'react';
import { toJS } from 'mobx';
import { observer, useObservable } from 'mobx-react-lite';
import { useGetData } from '../../hooks';
import { IIngredient } from '../../domain/Ingredient';

const Ingredients: React.FC = observer(() => {
  const store = useObservable({
    ingredients: [] as IIngredient[],

    setIngredients(i: IIngredient[]): void {
      store.ingredients = i;
    },

    get ingredientCount(): number {
      return store.ingredients.length;
    },
  });
  const { data, isLoading, isError, error } = useGetData('ingredients', []);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error: {error}</div>;
  }

  store.setIngredients(data);

  return (
    <>
      <div>Total Ingredients: {store.ingredientCount}</div>
      {store.ingredients.map((ingredient: IIngredient, idx: number) => (
        <div key={idx}>{ingredient.name}</div>
      ))}
    </>
  );
});

export default Ingredients;
