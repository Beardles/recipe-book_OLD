export interface IIngredient {
  id: number;
  name: string;
}

export class Ingredient implements IIngredient {
  id: number;
  name: string;

  constructor(ingredient: IIngredient) {
    this.id = ingredient.id;
    this.name = ingredient.name;
  }
}
