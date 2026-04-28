package store

import (
	"golang-trainning-backend/pkg/helper"
	"golang-trainning-backend/pkg/querymodel"
)

func ToQueryModel(row map[string]any) *querymodel.StoreQueryModel {
	return &querymodel.StoreQueryModel{
		ID:               helper.ToUint(row["id"]),
		CompanyID:        helper.ToUint(row["company_id"]),
		Name:             helper.ToString(row["name"]),
		IsActive:         helper.ToBool(helper.ToInt(row["is_active"])),
		OpenStatus:       helper.ToString(row["open_status"]),
		BusinessTypeCode: helper.ToString(row["business_type_code"]),
		ContractPlanCode: helper.ToString(row["contract_plan_code"]),
		DistrictID:       helper.ToUint(row["district_id"]),
		DistrictName:     helper.ToString(row["district_name"]),
		PrefectureID:     helper.ToUint(row["prefecture_id"]),
		PrefectureName:   helper.ToString(row["prefecture_name"]),
		RegionID:         helper.ToUint(row["region_id"]),
		RegionName:       helper.ToString(row["region_name"]),
	}
}
