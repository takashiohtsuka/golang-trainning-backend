import type { ContractPlan } from "@/interfaces/contract_plan";

export async function fetchContractPlans(): Promise<ContractPlan[]> {
  const res = await fetch(`/backend/contract_plans`);
  if (!res.ok) {
    throw new Error(`Failed to fetch contract_plans: ${res.status}`);
  }
  return res.json();
}
