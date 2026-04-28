import { useEffect, useState } from "react";
import { fetchBusinessTypes } from "@/api/business_type";
import type { BusinessType } from "@/interfaces/business_type";

export function useBusinessTypes() {
  const [businessTypes, setBusinessTypes] = useState<BusinessType[]>([]);

  useEffect(() => {
    fetchBusinessTypes().then(setBusinessTypes).catch(console.error);
  }, []);

  return { businessTypes };
}
