import { useEffect, useState } from "react";
import { fetchDistricts } from "@/api/district";
import type { District } from "@/interfaces/district";

export function useDistricts(prefectureId: number | null) {
  const [districts, setDistricts] = useState<District[]>([]);

  useEffect(() => {
    if (prefectureId == null) {
      setDistricts([]);
      return;
    }
    fetchDistricts(prefectureId).then(setDistricts).catch(console.error);
  }, [prefectureId]);

  return { districts };
}
