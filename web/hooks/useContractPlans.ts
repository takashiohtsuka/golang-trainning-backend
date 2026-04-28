import { useEffect, useState } from "react";
import { fetchContractPlans } from "@/api/contract_plan";
import type { ContractPlan } from "@/interfaces/contract_plan";

export function useContractPlans() {
  const [contractPlans, setContractPlans] = useState<ContractPlan[]>([]);

  useEffect(() => {
    fetchContractPlans().then(setContractPlans).catch(console.error);
  }, []);

  return { contractPlans };
}
