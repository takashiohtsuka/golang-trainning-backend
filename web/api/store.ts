import type { StoreDetail } from "@/interfaces/store";

// クライアントサイド用（ブラウザからnginx経由）
export async function fetchStore(id: string): Promise<StoreDetail> {
  const res = await fetch(`/backend/stores/${id}`);
  if (!res.ok) {
    throw new Error(`Failed to fetch store: ${res.status}`);
  }
  return res.json();
}

// サーバーサイド用（Next.jsサーバーからGoバックエンドへ直接）
export async function fetchStoreServer(id: string): Promise<StoreDetail> {
  const res = await fetch(`http://localhost:8080/backend/stores/${id}`);
  if (!res.ok) {
    throw new Error(`Failed to fetch store: ${res.status}`);
  }
  return res.json();
}
