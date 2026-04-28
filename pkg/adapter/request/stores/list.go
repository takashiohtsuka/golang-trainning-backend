package stores

import (
	"golang-trainning-backend/pkg/usecase/query"
)

type List struct {
	RegionID        uint   `query:"region_id"`
	PrefectureID    uint   `query:"prefecture_id"`
	DistrictIDs     []uint `query:"district_ids"`
	BusinessTypeIDs []uint `query:"business_type_ids"`
	ContractPlanIDs []uint `query:"contract_plan_ids"`
}

func (l *List) ToConditions() []query.Condition {
	var conditions []query.Condition

	if l.RegionID != 0 {
		conditions = append(conditions, query.Where("r.id", l.RegionID))
	}
	if l.PrefectureID != 0 {
		conditions = append(conditions, query.Where("p.id", l.PrefectureID))
	}
	if len(l.DistrictIDs) > 0 {
		conditions = append(conditions, query.WhereIn("d.id", l.DistrictIDs))
	}
	if len(l.BusinessTypeIDs) > 0 {
		conditions = append(conditions, query.WhereIn("s.business_type_id", l.BusinessTypeIDs))
	}
	if len(l.ContractPlanIDs) > 0 {
		conditions = append(conditions, query.WhereIn("s.contract_plan_id", l.ContractPlanIDs))
	}

	return conditions
}
