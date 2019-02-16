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

  const fetchData: () => Promise<void> = async (): Promise<void> => {
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
    if (!store.selectedIngredient) {
      fetchData();
    }

    if (id) {
      store.setSelectedIngredientId(id);
    }
  }, []);

  return { isLoading, isError, error };
};

export default useGetIngredients;