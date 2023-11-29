package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/internal/delivery/http/handlers/lego/legoset"
	s "legocy-go/internal/domain/lego/service"
	m "legocy-go/pkg/auth/jwt/middleware"
)

func (r V1router) addLegoSets(rg *gin.RouterGroup, service s.LegoSetService) {
	handler := h.NewLegoSetHandler(service)

	sets := rg.Group("/sets").Use(m.IsAuthenticated())
	{
		sets.GET("/", handler.ListSets)
		sets.GET("/:setID", handler.SetDetail)
	}
	setsAdmin := rg.Group("/admin/sets").Use(m.IsAdmin())
	{
		setsAdmin.POST("/", handler.SetCreate)
		setsAdmin.DELETE("/:setID", handler.SetDelete)
	}
}
