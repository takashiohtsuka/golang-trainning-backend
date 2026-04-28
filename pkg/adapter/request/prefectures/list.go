package prefectures

type List struct {
	RegionID        uint   `query:"region_id" validate:"required,min=1"`
	BusinessTypeIDs []uint `query:"business_type_ids"`
	ContractPlanIDs []uint `query:"contract_plan_ids"`
}
