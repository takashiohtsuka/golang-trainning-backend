package querymodel

type ContractPlanQuery interface {
	IsNil() bool
	GetID() uint
	GetCode() string
}

type ContractPlan struct {
	ID   uint
	Code string
}

func (c *ContractPlan) IsNil() bool     { return c.ID == 0 }
func (c *ContractPlan) GetID() uint     { return c.ID }
func (c *ContractPlan) GetCode() string { return c.Code }

type NilContractPlan struct{}

func (c *NilContractPlan) IsNil() bool     { return true }
func (c *NilContractPlan) GetID() uint     { return 0 }
func (c *NilContractPlan) GetCode() string { return "" }
