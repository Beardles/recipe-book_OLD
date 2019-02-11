export interface IIngredient {
  name: string;
}

export class Ingredient implements IIngredient {
  name: string;

  constructor(ingredient: IIngredient) {
    this.name = ingredient.name;
  }
}
