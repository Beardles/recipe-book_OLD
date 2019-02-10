import { Home } from './components/Home';
import { Ingredients } from './components/Ingredients';

export interface IRouteConfig {
  path: string;
  component: React.FC;
  exact?: boolean;
}

export const routes: IRouteConfig[] = [
  {
    path: '/',
    component: Home,
  },
  {
    path: '/ingredients',
    component: Ingredients,
  },
];
