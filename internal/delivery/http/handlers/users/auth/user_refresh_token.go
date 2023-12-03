package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/config"
	_ "github.com/legocy-co/legocy/docs"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	resources "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/pkg/auth/jwt"
	"net/http"
)

// RefreshToken
//
//	@Summary	refresh jwt tokens
//	@Tags		authentication
//	@ID			refresh-jwt
//	@Produce	json
//	@Param		data	body		schemas.RefreshTokenRequest	true	"jwt request"
//	@Success	200		{object}	schemas.AccessTokenResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/auth/refresh [post]
func (th *TokenHandler) RefreshToken(c *gin.Context) {

	var jwtRequest resources.RefreshTokenRequest
	if err := c.ShouldBindJSON(&jwtRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := jwt.ValidateRefreshToken(
		jwtRequest.RefreshToken,
		config.GetAppConfig().JwtConf.SecretKey,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload, _ := jwt.ParseTokenClaims(
		jwtRequest.RefreshToken,
		config.GetAppConfig().JwtConf.SecretKey,
	)

	user, appErr := th.service.GetUserByID(c, payload.ID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	accessToken, err := jwt.GenerateAccessToken(
		user.Email,
		user.ID,
		user.Role,
		config.GetAppConfig().JwtConf.SecretKey,
		config.GetAppConfig().JwtConf.AccessTokenLifeTime,
	)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, resources.AccessTokenResponse{
		AccessToken: accessToken,
	},
	)
}
