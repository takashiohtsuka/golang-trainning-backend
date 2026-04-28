import type { Region } from "@/interfaces/region";

export async function fetchRegions(): Promise<Region[]> {
  const res = await fetch("/backend/regions");
  if (!res.ok) {
    throw new Error(`Failed to fetch regions: ${res.status}`);
  }
  return res.json();
}
