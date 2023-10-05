package utils

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ExtractLeaderDocumentFromToken(userToken interface{}) (string, error) {
	token, ok := userToken.(*jwt.Token)
	if !ok || token == nil {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Token de usuario no válido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Error al procesar el token")
	}

	leaderDocument, ok := claims["document"].(string)
	if !ok {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Este usuario no tiene ningún documento de líder asignado")
	}

	return leaderDocument, nil
}
