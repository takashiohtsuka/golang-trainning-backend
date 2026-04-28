"use client";

import { useBusinessTypes } from "@/hooks/useBusinessTypes";

type Props = {
  value: number[];
  onChange: (businessTypeIds: number[]) => void;
};

export default function BusinessTypeCheckboxes({ value, onChange }: Props) {
  const { businessTypes } = useBusinessTypes();

  const handleChange = (id: number, checked: boolean) => {
    const businessTypeIds = checked
      ? [...value, id]
      : value.filter((v) => v !== id);
    onChange(businessTypeIds);
  };

  if (businessTypes.length === 0) return null;

  return (
    <div>
      <label>業種</label>
      {businessTypes.map((bt) => (
        <label key={bt.ID}>
          <input
            type="checkbox"
            value={bt.ID}
            checked={value.includes(bt.ID)}
            onChange={(e) => handleChange(bt.ID, e.target.checked)}
          />
          {bt.Code}
        </label>
      ))}
    </div>
  );
}
