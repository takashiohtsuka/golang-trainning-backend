package querymodel

type StoreQuery interface {
	IsNil() bool
	GetID() uint
	GetCompanyID() uint
	GetName() string
	GetIsActive() bool
	GetOpenStatus() string
	GetBusinessTypeCode() string
	GetContractPlanCode() string
	GetDistrictID() uint
	GetDistrictName() string
	GetPrefectureID() uint
	GetPrefectureName() string
	GetRegionID() uint
	GetRegionName() string
}

type StoreQueryModel struct {
	ID               uint
	CompanyID        uint
	Name             string
	IsActive         bool
	OpenStatus       string
	BusinessTypeCode string
	ContractPlanCode string
	DistrictID       uint
	DistrictName     string
	PrefectureID     uint
	PrefectureName   string
	RegionID         uint
	RegionName       string
}

func (s *StoreQueryModel) IsNil() bool              { return s.ID == 0 }
func (s *StoreQueryModel) GetID() uint               { return s.ID }
func (s *StoreQueryModel) GetCompanyID() uint        { return s.CompanyID }
func (s *StoreQueryModel) GetName() string           { return s.Name }
func (s *StoreQueryModel) GetIsActive() bool         { return s.IsActive }
func (s *StoreQueryModel) GetOpenStatus() string     { return s.OpenStatus }
func (s *StoreQueryModel) GetBusinessTypeCode() string { return s.BusinessTypeCode }
func (s *StoreQueryModel) GetContractPlanCode() string { return s.ContractPlanCode }
func (s *StoreQueryModel) GetDistrictID() uint       { return s.DistrictID }
func (s *StoreQueryModel) GetDistrictName() string   { return s.DistrictName }
func (s *StoreQueryModel) GetPrefectureID() uint     { return s.PrefectureID }
func (s *StoreQueryModel) GetPrefectureName() string { return s.PrefectureName }
func (s *StoreQueryModel) GetRegionID() uint         { return s.RegionID }
func (s *StoreQueryModel) GetRegionName() string     { return s.RegionName }

type NilStoreQueryModel struct{}

func (s *NilStoreQueryModel) IsNil() bool              { return true }
func (s *NilStoreQueryModel) GetID() uint               { return 0 }
func (s *NilStoreQueryModel) GetCompanyID() uint        { return 0 }
func (s *NilStoreQueryModel) GetName() string           { return "" }
func (s *NilStoreQueryModel) GetIsActive() bool         { return false }
func (s *NilStoreQueryModel) GetOpenStatus() string     { return "" }
func (s *NilStoreQueryModel) GetBusinessTypeCode() string { return "" }
func (s *NilStoreQueryModel) GetContractPlanCode() string { return "" }
func (s *NilStoreQueryModel) GetDistrictID() uint       { return 0 }
func (s *NilStoreQueryModel) GetDistrictName() string   { return "" }
func (s *NilStoreQueryModel) GetPrefectureID() uint     { return 0 }
func (s *NilStoreQueryModel) GetPrefectureName() string { return "" }
func (s *NilStoreQueryModel) GetRegionID() uint         { return 0 }
func (s *NilStoreQueryModel) GetRegionName() string     { return "" }
