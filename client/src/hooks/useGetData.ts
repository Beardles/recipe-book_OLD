import { useState, useEffect, useContext } from 'react';
import axios, { AxiosResponse } from 'axios';

function useGetData<T>(
  resource: string,
  initialState: T | {}
): { data: T; isLoading: boolean; isError: boolean; error: string };

function useGetData<T>(
  resource: string,
  initialState: T
): { data: T; isLoading: boolean; isError: boolean; error: string };

function useGetData(resource: string, initialState: any): any {
  const [data, setData] = useState(initialState);
  const [isLoading, setIsLoading] = useState(false);
  const [isError, setIsError] = useState(false);
  const [error, setError] = useState('');

  const fetchData: (resource: string) => Promise<void> = async (
    resource: string
  ): Promise<void> => {
    setIsLoading(true);
    setIsError(false);
    setError('');

    try {
      const response: AxiosResponse = await axios.get(
        `http://localhost:5000/api/v1/${resource}`
      );
      setData(response.data);
    } catch (e) {
      setIsError(true);
      setError(e.message);
    }

    setIsLoading(false);
  };

  useEffect(() => {
    fetchData(resource);
  }, []);

  return { data, isLoading, isError, error };
}

export default useGetData;
