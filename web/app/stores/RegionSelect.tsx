"use client";

import { useRegions } from "@/hooks/useRegions";

type Props = {
  value: number | null;
  onChange: (regionId: number | null) => void;
};

export default function RegionSelect({ value, onChange }: Props) {
  const { regions } = useRegions();

  const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const regionId = e.target.value ? Number(e.target.value) : null;
    onChange(regionId);
  };

  return (
    <div>
      <label>地方</label>
      <select value={value ?? ""} onChange={handleChange}>
        <option value="">選択して下さい</option>
        {regions.map((r) => (
          <option key={r.ID} value={r.ID}>
            {r.Name}
          </option>
        ))}
      </select>
    </div>
  );
}
