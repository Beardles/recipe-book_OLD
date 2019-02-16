import React, { useContext } from 'react';
import { Edit, Trash } from 'grommet-icons';
import { observer } from 'mobx-react-lite';
import { IIngredient } from '../../domain/Ingredient';
import {
  Text,
  Table,
  TableHeader,
  TableRow,
  TableCell,
  TableBody,
} from 'grommet';
import { IngredientTableActions, IngredientLink } from './ingredients.styles';
import StoreContext from '../../state';

interface IProps {
  ingredients: IIngredient[];
}

const IngredientsTable: React.FC<IProps> = observer(
  ({ ingredients }: IProps) => {
    const store = useContext(StoreContext);

    return (
      <Table>
        <TableHeader>
          <TableRow>
            <TableCell size="small">
              <Text>Name</Text>
            </TableCell>
            <TableCell size="small">
              <Text>Notes</Text>
            </TableCell>
            <TableCell size="small">
              <Text alignSelf="end">Actions</Text>
            </TableCell>
          </TableRow>
        </TableHeader>
        <TableBody>
          {ingredients.map((ingredient: IIngredient, idx: number) => (
            <TableRow key={idx}>
              <TableCell>
                <IngredientLink
                  onClick={() => store.setSelectedIngredientId(ingredient.id)}
                  to={`/ingredients/${ingredients[idx].id}`}
                >
                  {ingredient.name}
                </IngredientLink>
              </TableCell>
              <TableCell>{ingredient.notes}</TableCell>
              <TableCell style={{ textAlign: 'right' }}>
                <IngredientTableActions>
                  <Edit color="dark-2" onClick={() => alert('Edit')} />
                  <Trash color="dark-2" onClick={() => alert('Delete')} />
                </IngredientTableActions>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    );
  }
);

export default IngredientsTable;
