import { Home } from './components/Home';
import { IngredientList, IngredientView } from './components/Ingredients';
import { Recipes } from './components/Recipes';

export interface IRouteConfig {
  path: string;
  component: React.FC<any>;
  exact?: boolean;
}

export const routes: IRouteConfig[] = [
  {
    path: '/',
    component: Home,
    exact: true,
  },
  {
    path: '/recipes',
    exact: true,
    component: Recipes,
  },
  {
    path: '/ingredients',
    exact: true,
    component: IngredientList,
  },
  {
    path: '/ingredients/:id',
    exact: true,
    component: IngredientView,
  },
];
