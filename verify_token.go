package gfa

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Go handles auth with credentials from its own api.
// We just have to check for a valid token on secure routes.
func VerifyToken(app *Handler, nonRestrictivePolicy bool) gin.HandlerFunc {
	return func (c *gin.Context) {
		// Google automatically insert the auth token in the Authorization header.
		token := c.Request.Header.Get("authorization")

		// Set request data.
		c.Set("logged", false)
		c.Set("credentials", nil)

		// An empty token was provided.
		if token == "" || token == "null" {
			// We sometimes want to allow non logged users to access a resource in a limited way.
			if nonRestrictivePolicy == true {
				c.Next()
				return
			} else {
				// Abort request for non authorized users (resource will not be accessible at all).
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
				return
			}
		}

		// Decode auth token to retrieve some credentials
		credentials, err := app.Auth.VerifyIDToken(context.Background(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
			return
		}

		// Credentials are of type *firebase.Token (for type assertion).
		c.Set("credentials", credentials)

		// Go to next middleware.
		c.Next()
		return
	}
}
