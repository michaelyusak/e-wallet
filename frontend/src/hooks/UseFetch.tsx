import Cookies from 'js-cookie';
import { useState, useEffect } from 'react';

type FetchState<T> = {
  data: T | null;
  isLoading: boolean;
  error: Error | null;
  refetch: () => void;
};

const useFetch = <T,>(url: string): FetchState<T> => {
  const [data, setData] = useState<T | null>(null);
  const [isLoading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<Error | null>(null);

  const options = {
    headers: {
      'Content-Type': 'application/json',
      Authorization: Cookies.get('token') || '',
    },
  };

  const fetchData = async (url: string) => {
    try {
      const response = await fetch(url, options);
      if (!response.ok) {
        throw new Error(response.statusText);
      }
      const jsonData = await response.json();
      setData(jsonData);
    } catch (err) {
      setError(err as Error);
    } finally {
      setLoading(false);
    }
  };

  const refetch = () => fetchData(url);

  useEffect(() => {
    fetchData(url);
  }, [url]);

  return { data, isLoading, error, refetch };
};

export default useFetch;
