import { useState, useEffect } from 'react';
import axios, { AxiosResponse } from 'axios';

const useGetData = <T extends any>(
  url: string,
  initialData: T
): { data: T; isLoading: boolean; isError: boolean; error: string } => {
  const [data, setData] = useState(initialData);
  const [isLoading, setIsLoading] = useState(false);
  const [isError, setIsError] = useState(false);
  const [error, setError] = useState('');

  const fetchData = async (resource: string) => {
    setIsLoading(true);
    setIsError(false);
    setError('');

    try {
      const response: AxiosResponse = await axios.get(
        `http://localhost:5000/api/v1/${resource}`
      );
      setData(response.data);
    } catch (e) {
      console.log(e);
      setIsError(true);
      setError(e.message);
    }

    setIsLoading(false);
  };

  useEffect(() => {
    fetchData(url);
  }, []);

  return { data, isLoading, isError, error };
};

export default useGetData;
