import type { BusinessType } from "@/interfaces/business_type";

export async function fetchBusinessTypes(): Promise<BusinessType[]> {
  const res = await fetch(`/backend/business_types`);
  if (!res.ok) {
    throw new Error(`Failed to fetch business_types: ${res.status}`);
  }
  return res.json();
}
