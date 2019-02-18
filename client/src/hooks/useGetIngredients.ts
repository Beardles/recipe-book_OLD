import StoreContext from '../state';
import { useContext, useState, useEffect } from 'react';
import Axios, { AxiosResponse } from 'axios';

const useGetIngredients = (
  id?: number
): {
  isLoading: boolean;
  isError: boolean;
  error: string;
} => {
  const store = useContext(StoreContext);
  const [isLoading, setIsLoading] = useState(false);
  const [isError, setIsError] = useState(false);
  const [error, setError] = useState('');

  const fetchIngredients: () => Promise<void> = async (): Promise<void> => {
    setIsLoading(true);
    setIsError(false);
    setError('');

    try {
      const response: AxiosResponse = await Axios.get(
        'http://localhost:5000/api/v1/ingredients'
      );
      store.setIngredients(response.data);
    } catch (e) {
      setIsError(true);
      setError(e.message);
    }

    setIsLoading(false);
  };

  useEffect(() => {
    // TODO: This could be cleaner
    // If we are working with a single ingredient
    if (id) {
      if (!store.selectedIngredient) {
        fetchIngredients();
      }
      store.setSelectedIngredientId(id);
    } else {
      // If we are working with the list
      if (store.ingredients.length === 0) {
        fetchIngredients();
      }
    }
  }, []);

  return { isLoading, isError, error };
};

export default useGetIngredients;
