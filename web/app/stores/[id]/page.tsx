"use client";

import { useParams } from "next/navigation";
import { useStore } from "./useStore";
import StoreDetailView from "./StoreDetail";

export default function Page() {
  const { id } = useParams<{ id: string }>();
  const { store, loading, error } = useStore(id);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;
  if (!store) return <p>店舗が見つかりませんでした。</p>;

  return (
    <main>
      <StoreDetailView store={store} />
    </main>
  );
}
