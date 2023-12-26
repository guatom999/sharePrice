package server

import (
	"github.com/guatom999/sharePrice/igtools/getFollowersHandlers"
	"github.com/guatom999/sharePrice/igtools/getFollowersRepositories"
	"github.com/guatom999/sharePrice/igtools/getFollowersUseCases"
)

func (s *server) igToolsServer() {
	getFollowersRepo := getFollowersRepositories.NewGetFollowersRepository(s.db)
	getFollowersUseCase := getFollowersUseCases.NewGetFollowersUseCase(getFollowersRepo)
	getFollowersHandler := getFollowersHandlers.NewGetFollowersHandler(s.cfg, getFollowersUseCase)

	router := s.app.Group("/igtools_v1")

	router.GET("/test", getFollowersHandler.GetFollowers)

}
