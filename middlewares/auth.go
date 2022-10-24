package middlewares

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"strings"
)

func verifyFirebaseIDToken(ctx echo.Context, auth *auth.Client) (*auth.Token, error) {
	// リクエストのヘッダからトークンを取り出す
	headerAuth := ctx.Request().Header.Get("Authorization")
	token := strings.Replace(headerAuth, "Bearer ", "", 1)
	// 取り出したトークンを渡して検証する
	jwtToken, err := auth.VerifyIDToken(context.Background(), token)

	return jwtToken, err
}

// ログインした状態でなければ利用できないAPIに対して使用する
// FIXME: トークンの確認に失敗する。
// ID token has invalid 'aud' (audience) claim; expected "nuxtgovideo" but got "manager-d606f"; make sure the ID token comes from the same Firebase project as the credential used to authenticate this SDK; see https://firebase.google.com/docs/auth/admin/verify-id-tokens for details on how to retrieve a valid ID token
func FirebaseGuard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authClient := c.Get("firebase").(*auth.Client)
			jwtToken, err := verifyFirebaseIDToken(c, authClient)

			if err != nil {
				return c.JSON(fasthttp.StatusUnauthorized, "Not Authenticated")
			}

			c.Set("auth", jwtToken) // コンテキストに*auth.Tokenを保存する。*auth.TokenはUIDを持つ構造体。このUIDを各機能で使用する

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}

// ログインしていなくても利用できるAPIに使用する
func FirebaseAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authClient := c.Get("firebase").(*auth.Client)
			jwtToken, _ := verifyFirebaseIDToken(c, authClient)

			c.Set("auth", jwtToken)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}
