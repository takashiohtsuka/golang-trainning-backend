import type { Prefecture } from "@/interfaces/prefecture";

export async function fetchPrefectures(regionId: number): Promise<Prefecture[]> {
  const res = await fetch(`/backend/prefectures?region_id=${regionId}`);
  if (!res.ok) {
    throw new Error(`Failed to fetch prefectures: ${res.status}`);
  }
  return res.json();
}
