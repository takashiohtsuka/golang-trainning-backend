import { useEffect, useState } from "react";
import { fetchPrefectures } from "@/api/prefecture";
import type { Prefecture } from "@/interfaces/prefecture";

export function usePrefectures(
  regionId: number | null,
  businessTypeIds: number[] = [],
  contractPlanIds: number[] = []
) {
  const [prefectures, setPrefectures] = useState<Prefecture[]>([]);

  useEffect(() => {
    if (regionId == null) {
      setPrefectures([]);
      return;
    }
    fetchPrefectures({ regionId, businessTypeIds, contractPlanIds })
      .then(setPrefectures)
      .catch(console.error);
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [regionId, JSON.stringify(businessTypeIds), JSON.stringify(contractPlanIds)]);

  return { prefectures };
}
