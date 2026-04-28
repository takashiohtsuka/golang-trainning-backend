import { useEffect, useState } from "react";
import { fetchRegions } from "@/api/region";
import type { Region } from "@/interfaces/region";

export function useRegions() {
  const [regions, setRegions] = useState<Region[]>([]);

  useEffect(() => {
    fetchRegions().then(setRegions).catch(console.error);
  }, []);

  return { regions };
}
