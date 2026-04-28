"use client";

import { usePrefectures } from "@/hooks/usePrefectures";

type Props = {
  regionId: number | null;
  businessTypeIds: number[];
  contractPlanIds: number[];
  value: number | null;
  onChange: (prefectureId: number | null) => void;
};

export default function PrefectureSelect({ regionId, businessTypeIds, contractPlanIds, value, onChange }: Props) {
  const { prefectures } = usePrefectures(regionId, businessTypeIds, contractPlanIds);

  const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const prefectureId = e.target.value ? Number(e.target.value) : null;
    onChange(prefectureId);
  };

  return (
    <div>
      <label>都道府県</label>
      <select value={value ?? ""} onChange={handleChange} disabled={prefectures.length === 0}>
        <option value="">選択して下さい</option>
        {prefectures.map((p) => (
          <option key={p.ID} value={p.ID}>
            {p.Name}
          </option>
        ))}
      </select>
    </div>
  );
}
