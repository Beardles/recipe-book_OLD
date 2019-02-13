import React from 'react';
import { IIngredient } from '../../domain/Ingredient';
import {
  Text,
  Table,
  TableHeader,
  TableRow,
  TableCell,
  TableBody,
  Box,
} from 'grommet';
import { Link } from 'react-router-dom';
import { Edit, Trash } from 'grommet-icons';
import { IngredientTableActions, IngredientLink } from './ingredients.styles';

interface IProps {
  ingredients: IIngredient[];
}

const IngredientsTable: React.FC<IProps> = ({ ingredients }) => (
  <Table>
    <TableHeader>
      <TableRow>
        <TableCell size="small">
          <Text>Name</Text>
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
            <IngredientLink to={`/ingredients/${ingredients[idx].id}`}>
              {ingredient.name}
            </IngredientLink>
          </TableCell>
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

export default IngredientsTable;
