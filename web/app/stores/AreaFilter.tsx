"use client";

import RegionSelect from "./RegionSelect";
import PrefectureSelect from "./PrefectureSelect";
import DistrictCheckboxes from "./DistrictCheckboxes";
import BusinessTypeCheckboxes from "./BusinessTypeCheckboxes";
import ContractPlanCheckboxes from "./ContractPlanCheckboxes";

export type AreaFilterValue = {
  regionId: number | null;
  prefectureId: number | null;
  districtIds: number[];
  businessTypeIds: number[];
  contractPlanIds: number[];
};

type Props = {
  value: AreaFilterValue;
  onChange: (value: AreaFilterValue) => void;
};

export default function AreaFilter({ value, onChange }: Props) {
  const handleRegionChange = (regionId: number | null) => {
    onChange({ ...value, regionId, prefectureId: null, districtIds: [] });
  };

  const handlePrefectureChange = (prefectureId: number | null) => {
    onChange({ ...value, prefectureId, districtIds: [] });
  };

  const handleDistrictChange = (districtIds: number[]) => {
    onChange({ ...value, districtIds });
  };

  const handleBusinessTypeChange = (businessTypeIds: number[]) => {
    onChange({ ...value, businessTypeIds, prefectureId: null, districtIds: [] });
  };

  const handleContractPlanChange = (contractPlanIds: number[]) => {
    onChange({ ...value, contractPlanIds, prefectureId: null, districtIds: [] });
  };

  return (
    <div>
      <BusinessTypeCheckboxes
        value={value.businessTypeIds}
        onChange={handleBusinessTypeChange}
      />
      <ContractPlanCheckboxes
        value={value.contractPlanIds}
        onChange={handleContractPlanChange}
      />
      <RegionSelect value={value.regionId} onChange={handleRegionChange} />
      <PrefectureSelect
        regionId={value.regionId}
        businessTypeIds={value.businessTypeIds}
        contractPlanIds={value.contractPlanIds}
        value={value.prefectureId}
        onChange={handlePrefectureChange}
      />
      <DistrictCheckboxes
        prefectureId={value.prefectureId}
        businessTypeIds={value.businessTypeIds}
        contractPlanIds={value.contractPlanIds}
        value={value.districtIds}
        onChange={handleDistrictChange}
      />
    </div>
  );
}
