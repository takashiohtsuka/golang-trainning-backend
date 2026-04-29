import { useEffect, useState } from "react";
import { fetchPrefectures } from "@/api/prefecture";
import type { Prefecture } from "@/interfaces/prefecture";

export function usePrefectures(regionId: number | null) {
  const [prefectures, setPrefectures] = useState<Prefecture[]>([]);

  useEffect(() => {
    if (regionId == null) {
      setPrefectures([]);
      return;
    }
    fetchPrefectures(regionId).then(setPrefectures).catch(console.error);
  }, [regionId]);

  return { prefectures };
}
