import { Home } from './components/Home';
import { Ingredients } from './components/Ingredients';
import { Recipes } from './components/Recipes';

export interface IRouteConfig {
  path: string;
  component: React.FC;
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
    component: Recipes,
  },
  {
    path: '/ingredients',
    component: Ingredients,
  },
];
