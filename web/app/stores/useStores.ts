import { useState } from "react";
import { fetchStores } from "@/api/store";
import type { StoreDetail } from "@/interfaces/store";
import type { AreaFilterValue } from "./AreaFilter";

type UseStoresReturn = {
  stores: StoreDetail[];
  loading: boolean;
  error: string | null;
  search: (filter: AreaFilterValue) => void;
};

export function useStores(): UseStoresReturn {
  const [stores, setStores] = useState<StoreDetail[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const search = (filter: AreaFilterValue) => {
    setLoading(true);
    setError(null);
    fetchStores({
      regionId: filter.regionId ?? undefined,
      prefectureId: filter.prefectureId ?? undefined,
      districtIds: filter.districtIds.length > 0 ? filter.districtIds : undefined,
      businessTypeIds: filter.businessTypeIds.length > 0 ? filter.businessTypeIds : undefined,
      contractPlanIds: filter.contractPlanIds.length > 0 ? filter.contractPlanIds : undefined,
    })
      .then(setStores)
      .catch((e: Error) => setError(e.message))
      .finally(() => setLoading(false));
  };

  return { stores, loading, error, search };
}
