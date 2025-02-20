package middlewares

import (
	"context"

	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/repositories/repositories_auth"
)

func ValidateToken(ctx context.Context, repositoriesAuth repositories_auth.RepositoryAuthInterface) string {
	defer func() {
		validateFail := recover()
		if validateFail != nil {
			panic("authorization invalid")
		}
	}()
	var tokenString string
	token := ctx.Value(helpers.ContextKey("token"))

	tokenString, ok := token.(string)
	if !ok || tokenString == "" {
		panic("authorization required")
	}

	username, isValid := repositoriesAuth.Validate(tokenString)
	if !isValid {
		panic("authorization invalid")
	}
	return username
}
