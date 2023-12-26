package getFollowersHandlers

import (
	"github.com/guatom999/sharePrice/config"
	"github.com/guatom999/sharePrice/igtools/getFollowersUseCases"
	"github.com/guatom999/sharePrice/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	GetFollowersHandlerService interface {
		GetFollowers(e echo.Context) error
	}

	getFollowersHandler struct {
		cfg                 *config.Config
		getFollowersUseCase getFollowersUseCases.GetFollowersUsecaseService
	}
)

func NewGetFollowersHandler(cfg *config.Config, getFollowersUseCase getFollowersUseCases.GetFollowersUsecaseService) GetFollowersHandlerService {
	return &getFollowersHandler{cfg, getFollowersUseCase}
}

func (h *getFollowersHandler) GetFollowers(e echo.Context) error {

	result, err := h.getFollowersUseCase.GetFollowers("mimienpi")
	if err != nil {
		return response.ErrResponse(e, 400, "erro: resutl is null or something")
	}

	_ = result

	return response.SuccessResponse(e, 200, "Test")
}
