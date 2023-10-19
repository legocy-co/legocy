package user_collection

import (
	"github.com/gin-gonic/gin"
	v1 "legocy-go/internal/delievery/http/middleware"
	"legocy-go/internal/delievery/http/resources/collections"
	"net/http"
)

// GetUserCollection
//
//	@Summary	Get User Collection
//	@Tags		user_collections
//	@ID			get_user_collection
//	@Produce	json
//	@Success	200	{object} 	collections.UserLegoSetCollectionResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/collections/ [get]
//
//	@Security	JWT
func (h UserLegoCollectionHandler) GetUserCollection(c *gin.Context) {
	tokenPayload, err := v1.GetUserPayload(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := tokenPayload.ID

	userCollection, err := h.s.GetUserCollection(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userCollectionResponse := collections.GetUserLegoCollectionResponse(*userCollection)
	c.JSON(http.StatusOK, userCollectionResponse)
}
