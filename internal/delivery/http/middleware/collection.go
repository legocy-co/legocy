package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	r "legocy-go/internal/domain/collections/repository"
	models "legocy-go/internal/domain/users/models"
	"legocy-go/pkg/auth/jwt/middleware"
	"net/http"
	"strconv"
)

func CollectionSetOwnerOrAdmin(
	lookUpParam string, repo r.UserCollectionRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenPayload, err := middleware.GetUserPayload(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if tokenPayload.Role == models.ADMIN {
			logrus.Printf("Current User is Admin. Access Allowed")
			ctx.Next()
			return
		}

		setID, err := strconv.Atoi(ctx.Param(lookUpParam))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "cannot extract set ID from URL"})
			return
		}

		setOwnerID, err := repo.GetCollectionSetOwner(ctx, setID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if tokenPayload.ID != setOwnerID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access Denied"})
			return
		}

		ctx.Next()
	}
}
