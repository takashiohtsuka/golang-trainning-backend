"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import { fetchStore } from "@/api/store";
import type { StoreDetail } from "@/interfaces/store";
import StoreDetailView from "./StoreDetail";

export default function Page() {
  const { id } = useParams<{ id: string }>();
  const [store, setStore] = useState<StoreDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setLoading(true);
    fetchStore(id)
      .then(setStore)
      .catch((e) => setError(e.message))
      .finally(() => setLoading(false));
  }, [id]);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;
  if (!store) return <p>店舗が見つかりませんでした。</p>;

  return (
    <main>
      <StoreDetailView store={store} />
    </main>
  );
}
