// routes/routes.go
package routes

import (
	"github.com/disaster_management_backend/internals/controllers"
	"github.com/disaster_management_backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)

	authorized := router.Group("/api")

	authorized.POST("/donate", controllers.AddDonation)
	authorized.GET("/donation", controllers.GetTotalDonations)
	authorized.GET("/carts", controllers.GetFundsAndExpenses)
	authorized.GET("/volunteers", controllers.GetAllVolunteer)
	authorized.GET("/volunteer/:id", controllers.GetVolunteerByID)
	authorized.POST("/crisis", controllers.ReportCrisis)

	// protected routes
	authorized.Use(middleware.AuthMiddleware())
	{
		// Admin-only routes
		adminRoutes := authorized.Group("/admin")
		adminRoutes.Use(middleware.AdminOnly())
		{
			adminRoutes.POST("/approve-crisis", controllers.ApproveCrisis)
			adminRoutes.POST("/assign-task", controllers.AssignTask)
			adminRoutes.POST("/approve-volunteer", controllers.ApproveVolunteer)
			adminReports := adminRoutes.Group("/reports")
			{
				adminReports.GET("/expenses", controllers.GenerateDailyExpenseReport)   // Daily expense report
				adminReports.GET("/donations", controllers.GenerateDailyDonationReport) // Daily donation report
				adminReports.GET("/inventory", controllers.GenerateInventoryReportCSV)  // Daily inventory report
			}
		}

		// Volunteer-only routes
		volunteerRoutes := authorized.Group("/volunteer")
		volunteerRoutes.Use(middleware.VolunteerOnly())
		{
			volunteerRoutes.POST("/respond-crisis", controllers.RespondToCrisis)
			volunteerRoutes.POST("/add", controllers.AddInventory)
			volunteerRoutes.PUT("/update/:id", controllers.UpdateInventory)
			volunteerRoutes.DELETE("/delete/:id", controllers.DeleteInventory)
			volunteerRoutes.GET("/", controllers.ListInventory)
		}
	}

	return router
}
