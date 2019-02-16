import React from 'react';
import { IIngredient, Ingredient } from './domain/Ingredient';

interface IStore {
  // Values
  ingredients: IIngredient[];
  selectedIngredientId: number | null;

  // Setters
  setIngredients(i: IIngredient[]): void;
  setSelectedIngredientId(id: number): void;

  // Computed
  selectedIngredient: IIngredient | undefined;
}

const store: IStore = {
  ingredients: [] as IIngredient[],
  selectedIngredientId: null,

  setIngredients(i: IIngredient[]): void {
    store.ingredients = i;
  },

  setSelectedIngredientId(id: number): void {
    store.selectedIngredientId = id;
  },

  get selectedIngredient(): IIngredient | undefined {
    return store.ingredients.find(
      (i: IIngredient) => i.id === store.selectedIngredientId
    );
  },
};

const StoreContext = React.createContext(store);

export default StoreContext;
