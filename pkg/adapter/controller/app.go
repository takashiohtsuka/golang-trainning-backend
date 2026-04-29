package controller

type AppController struct {
	Company                  interface{ Company }
	Store                    interface{ Store }
	ManagementStaff          interface{ ManagementStaff }
	Woman                    interface{ Woman }
	Region                   interface{ Region }
	Prefecture               interface{ Prefecture }
	District                 interface{ District }
	BusinessType             interface{ BusinessType }
	ContractPlan             interface{ ContractPlan }
	ImmediateAvailableWoman  interface{ ImmediateAvailableWoman }
}
