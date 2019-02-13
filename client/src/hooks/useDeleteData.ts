import { useState, useEffect } from 'react';
import axios, { AxiosResponse } from 'axios';

const useDeleteData = (
  resource: string,
  id: number
): {
  success: boolean;
  isDeleting: boolean;
  isError: boolean;
  error: string;
} => {
  const [success, setSuccess] = useState(true);
  const [isDeleting, setIsDeleting] = useState(false);
  const [isError, setIsError] = useState(false);
  const [error, setError] = useState('');

  const deleteData: (resource: string, id: number) => Promise<void> = async (
    resource: string,
    id: number
  ) => {
    setIsDeleting(true);
    setIsError(false);
    setError('');

    try {
      const response: AxiosResponse = await axios.delete(
        `http://localhost:5000/api/v1/${resource}/${id}`
      );
      if (response.status === 200) {
        setSuccess(true);
      } else {
        setSuccess(false);
        setIsError(true);
        setError(response.statusText);
      }
    } catch (e) {
      setIsError(true);
      setError(e.message);
    }

    setIsDeleting(false);
  };

  useEffect(() => {
    deleteData(resource, id);
  }, []);

  return { success, isDeleting, isError, error };
};

export default useDeleteData;
