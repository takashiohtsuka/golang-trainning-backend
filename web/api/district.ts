import type { District } from "@/interfaces/district";

export type DistrictListParams = {
  prefectureId: number;
  businessTypeIds?: number[];
  contractPlanIds?: number[];
};

export async function fetchDistricts(params: DistrictListParams): Promise<District[]> {
  const query = new URLSearchParams();
  query.set("prefecture_id", String(params.prefectureId));
  params.businessTypeIds?.forEach((id) => query.append("business_type_ids", String(id)));
  params.contractPlanIds?.forEach((id) => query.append("contract_plan_ids", String(id)));

  const res = await fetch(`/backend/districts?${query.toString()}`);
  if (!res.ok) {
    throw new Error(`Failed to fetch districts: ${res.status}`);
  }
  return res.json();
}
