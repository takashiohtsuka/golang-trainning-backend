"use client";

import { useDistricts } from "@/hooks/useDistricts";

type Props = {
  prefectureId: number | null;
  value: number[];
  onChange: (districtIds: number[]) => void;
};

export default function DistrictCheckboxes({ prefectureId, value, onChange }: Props) {
  const { districts } = useDistricts(prefectureId);

  const handleChange = (districtId: number, checked: boolean) => {
    const districtIds = checked
      ? [...value, districtId]
      : value.filter((id) => id !== districtId);
    onChange(districtIds);
  };

  if (districts.length === 0) return null;

  return (
    <div>
      <label>地区</label>
      {districts.map((d) => (
        <label key={d.ID}>
          <input
            type="checkbox"
            value={d.ID}
            checked={value.includes(d.ID)}
            onChange={(e) => handleChange(d.ID, e.target.checked)}
          />
          {d.Name}
        </label>
      ))}
    </div>
  );
}
