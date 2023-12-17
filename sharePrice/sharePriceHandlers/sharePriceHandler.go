package sharePriceHandlers

import (
	"github.com/guatom999/sharePrice/config"
	"github.com/guatom999/sharePrice/sharePrice/sharePriceUseCases"
	"github.com/labstack/echo/v4"
)

type sharePriceHandler struct {
	cfg               *config.Config
	sharePriceUseCase sharePriceUseCases.SharePriceUseCase
}

func NewSharePriceHandler(cfg *config.Config, sharePriceUseCase sharePriceUseCases.SharePriceUseCase) SharePriceHandler {
	return &sharePriceHandler{
		cfg:               cfg,
		sharePriceUseCase: sharePriceUseCase,
	}
}

func (h *sharePriceHandler) Test(e echo.Context) error {
	return SuccessResponse(e, 200, "test")
}
