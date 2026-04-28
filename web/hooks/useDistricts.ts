import { useEffect, useState } from "react";
import { fetchDistricts } from "@/api/district";
import type { District } from "@/interfaces/district";

export function useDistricts(
  prefectureId: number | null,
  businessTypeIds: number[] = [],
  contractPlanIds: number[] = []
) {
  const [districts, setDistricts] = useState<District[]>([]);

  useEffect(() => {
    if (prefectureId == null) {
      setDistricts([]);
      return;
    }
    fetchDistricts({ prefectureId, businessTypeIds, contractPlanIds })
      .then(setDistricts)
      .catch(console.error);
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [prefectureId, JSON.stringify(businessTypeIds), JSON.stringify(contractPlanIds)]);

  return { districts };
}
