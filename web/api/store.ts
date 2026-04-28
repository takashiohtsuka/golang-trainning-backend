import type { StoreDetail } from "@/interfaces/store";

export type StoreListParams = {
  regionId?: number;
  prefectureId?: number;
  districtIds?: number[];
  businessTypeIds?: number[];
  contractPlanIds?: number[];
};

// クライアントサイド用（ブラウザからnginx経由）
export async function fetchStores(params?: StoreListParams): Promise<StoreDetail[]> {
  const query = new URLSearchParams();
  if (params?.regionId) query.set("region_id", String(params.regionId));
  if (params?.prefectureId) query.set("prefecture_id", String(params.prefectureId));
  params?.districtIds?.forEach((id) => query.append("district_ids", String(id)));
  params?.businessTypeIds?.forEach((id) => query.append("business_type_ids", String(id)));
  params?.contractPlanIds?.forEach((id) => query.append("contract_plan_ids", String(id)));

  const res = await fetch(`/backend/stores?${query.toString()}`);
  if (!res.ok) {
    throw new Error(`Failed to fetch stores: ${res.status}`);
  }
  return res.json();
}

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
