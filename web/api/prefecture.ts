import type { Prefecture } from "@/interfaces/prefecture";

export type PrefectureListParams = {
  regionId: number;
  businessTypeIds?: number[];
  contractPlanIds?: number[];
};

export async function fetchPrefectures(params: PrefectureListParams): Promise<Prefecture[]> {
  const query = new URLSearchParams();
  query.set("region_id", String(params.regionId));
  params.businessTypeIds?.forEach((id) => query.append("business_type_ids", String(id)));
  params.contractPlanIds?.forEach((id) => query.append("contract_plan_ids", String(id)));

  const res = await fetch(`/backend/prefectures?${query.toString()}`);
  if (!res.ok) {
    throw new Error(`Failed to fetch prefectures: ${res.status}`);
  }
  return res.json();
}
