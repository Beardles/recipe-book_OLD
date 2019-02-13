import React from 'react';
import { useGetData } from '../../hooks';
import { match } from 'react-router';
import { IIngredient } from '../../domain/Ingredient';

interface IProps {
  match: match<{ id: string }>;
}

const IngredientView: React.FC<IProps> = ({ match }) => {
  const { data, isLoading, isError, error } = useGetData<IIngredient>(
    `ingredients/${match.params.id}`,
    {}
  );

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error: {error}</div>;
  }

  if (data.id) {
    return <div>{data.name}</div>;
  }

  return <div>No Ingredient Found!</div>;
};

export default IngredientView;
