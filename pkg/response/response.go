package response

import "github.com/labstack/echo/v4"

type (
	MsgResponse struct {
		Message string `json:"message"`
	}
)

func ErrResponse(c echo.Context, responseCode int, message string) error {
	return c.JSON(responseCode, &MsgResponse{Message: message})
}

func SuccessResponse(c echo.Context, responseCode int, data any) error {
	return c.JSON(responseCode, data)
}
