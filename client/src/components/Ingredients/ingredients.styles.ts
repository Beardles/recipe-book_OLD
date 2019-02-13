import styled from 'styled-components';
import { Link } from 'react-router-dom';

export const IngredientHeader = styled.div`
  margin-bottom: 25px;
`;

export const IngredientLink = styled(Link)`
  color: #7d4cdb;
  text-decoration: none;

  &:hover {
    color: #7d4cdb;
    text-decoration: none;
  }
`;

export const IngredientTableActions = styled.div`
  display: flex;
  justify-content: flex-end;
  width: 100%;

  svg {
    margin: 0 7px;

    &:hover {
      cursor: pointer;
    }
  }
`;
