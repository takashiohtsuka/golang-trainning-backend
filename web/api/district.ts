import type { District } from "@/interfaces/district";

export async function fetchDistricts(prefectureId: number): Promise<District[]> {
  const res = await fetch(`/backend/districts?prefecture_id=${prefectureId}`);
  if (!res.ok) {
    throw new Error(`Failed to fetch districts: ${res.status}`);
  }
  return res.json();
}
