package sharePriceHandlers

import "github.com/labstack/echo/v4"

type SharePriceHandler interface {
	Test(e echo.Context) error
}
