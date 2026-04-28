"use client";

import { useContractPlans } from "@/hooks/useContractPlans";

type Props = {
  value: number[];
  onChange: (contractPlanIds: number[]) => void;
};

export default function ContractPlanCheckboxes({ value, onChange }: Props) {
  const { contractPlans } = useContractPlans();

  const handleChange = (id: number, checked: boolean) => {
    const contractPlanIds = checked
      ? [...value, id]
      : value.filter((v) => v !== id);
    onChange(contractPlanIds);
  };

  if (contractPlans.length === 0) return null;

  return (
    <div>
      <label>契約プラン</label>
      {contractPlans.map((cp) => (
        <label key={cp.ID}>
          <input
            type="checkbox"
            value={cp.ID}
            checked={value.includes(cp.ID)}
            onChange={(e) => handleChange(cp.ID, e.target.checked)}
          />
          {cp.Code}
        </label>
      ))}
    </div>
  );
}
