package router

import (
	"golang-trainning-backend/pkg/adapter/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	g := e.Group("/backend")

	g.POST("/companies", func(ctx echo.Context) error { return c.Company.CreateCompany(ctx) })
	g.GET("/stores", func(ctx echo.Context) error { return c.Store.GetStores(ctx) })
	g.GET("/stores/:id", func(ctx echo.Context) error { return c.Store.GetStore(ctx) })
	g.POST("/stores", func(ctx echo.Context) error { return c.Store.CreateStore(ctx) })
	g.PUT("/stores/:id", func(ctx echo.Context) error { return c.Store.UpdateStore(ctx) })
	g.POST("/management_staffs", func(ctx echo.Context) error { return c.ManagementStaff.CreateManagementStaff(ctx) })
	g.POST("/women", func(ctx echo.Context) error { return c.Woman.CreateWoman(ctx) })
	g.PUT("/women/:id", func(ctx echo.Context) error { return c.Woman.UpdateWoman(ctx) })

	g.GET("/regions", func(ctx echo.Context) error { return c.Region.GetRegions(ctx) })
	g.GET("/prefectures", func(ctx echo.Context) error { return c.Prefecture.GetPrefectures(ctx) })
	g.GET("/districts", func(ctx echo.Context) error { return c.District.GetDistricts(ctx) })
	g.GET("/business_types", func(ctx echo.Context) error { return c.BusinessType.GetBusinessTypes(ctx) })
	g.GET("/contract_plans", func(ctx echo.Context) error { return c.ContractPlan.GetContractPlans(ctx) })
	g.POST("/stores/:id/immediate_available_women", func(ctx echo.Context) error {
		return c.ImmediateAvailableWoman.Create(ctx)
	})
	g.DELETE("/stores/:id/immediate_available_women/:woman_id", func(ctx echo.Context) error {
		return c.ImmediateAvailableWoman.Delete(ctx)
	})

	return e
}
