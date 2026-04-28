import { useEffect, useState } from "react";
import { fetchStore } from "@/api/store";
import type { StoreDetail } from "@/interfaces/store";

type UseStoreReturn = {
  store: StoreDetail | null;
  loading: boolean;
  error: string | null;
};

export function useStore(id: string): UseStoreReturn {
  const [store, setStore] = useState<StoreDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setLoading(true);
    fetchStore(id)
      .then(setStore)
      .catch((e: Error) => setError(e.message))
      .finally(() => setLoading(false));
  }, [id]);

  return { store, loading, error };
}
