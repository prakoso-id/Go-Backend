package portfolio

import (
	"github.com/gin-gonic/gin"
)

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	portfolios := router.Group("/portfolios")
	{
		// Public routes
		portfolios.GET("", m.Handler.GetAllPortfolios)
		portfolios.GET("/:user_id", m.Handler.GetUserPortfolio)
	}
}
