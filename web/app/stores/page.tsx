"use client";

import { useEffect, useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import StoreList from "./StoreList";
import AreaFilter, { type AreaFilterValue } from "./AreaFilter";
import { useStores } from "./useStores";

function parseAreaFilter(searchParams: URLSearchParams): AreaFilterValue {
  const regionId = searchParams.get("region_id");
  const prefectureId = searchParams.get("prefecture_id");
  const districtIds = searchParams.getAll("district_ids");
  const businessTypeIds = searchParams.getAll("business_type_ids");
  const contractPlanIds = searchParams.getAll("contract_plan_ids");

  return {
    regionId: regionId ? Number(regionId) : null,
    prefectureId: prefectureId ? Number(prefectureId) : null,
    districtIds: districtIds.map(Number),
    businessTypeIds: businessTypeIds.map(Number),
    contractPlanIds: contractPlanIds.map(Number),
  };
}

export default function Page() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const { stores, loading, error, search } = useStores();
  const [areaFilter, setAreaFilter] = useState<AreaFilterValue>(() =>
    parseAreaFilter(searchParams)
  );

  // URLパラメータが変化したら自動検索
  useEffect(() => {
    const filter = parseAreaFilter(searchParams);
    setAreaFilter(filter);
    search(filter);
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [searchParams.toString()]);

  const handleSearch = () => {
    const params = new URLSearchParams();
    if (areaFilter.regionId) params.set("region_id", String(areaFilter.regionId));
    if (areaFilter.prefectureId) params.set("prefecture_id", String(areaFilter.prefectureId));
    areaFilter.districtIds.forEach((id) => params.append("district_ids", String(id)));
    areaFilter.businessTypeIds.forEach((id) => params.append("business_type_ids", String(id)));
    areaFilter.contractPlanIds.forEach((id) => params.append("contract_plan_ids", String(id)));
    router.push(`?${params.toString()}`);
  };

  return (
    <main>
      <h1>店舗一覧</h1>

      <AreaFilter value={areaFilter} onChange={setAreaFilter} />

      <button onClick={handleSearch}>検索</button>

      {loading && <p>Loading...</p>}
      {error && <p>Error: {error}</p>}
      {!loading && <StoreList stores={stores} />}
    </main>
  );
}
