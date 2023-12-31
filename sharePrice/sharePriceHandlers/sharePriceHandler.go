package sharePriceHandlers

import (
	"context"

	"github.com/guatom999/sharePrice/config"
	"github.com/guatom999/sharePrice/pkg/response"
	"github.com/guatom999/sharePrice/sharePrice/sharePriceEntity"
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

func (h *sharePriceHandler) AddShareTracker(c echo.Context) error {

	ctx := context.Background()

	req := new(sharePriceEntity.ShareTrackReq)
	if err := c.Bind(req); err != nil {
		return response.ErrResponse(c, 400, err.Error())
	}

	result, err := h.sharePriceUseCase.AddTrackShare(ctx, req)
	if err != nil {
		return response.ErrResponse(c, 400, err.Error())
	}

	return response.SuccessResponse(c, 200, result)
}
