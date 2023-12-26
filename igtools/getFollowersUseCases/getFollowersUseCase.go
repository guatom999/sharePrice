package getFollowersUseCases

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/guatom999/sharePrice/igtools"
	"github.com/guatom999/sharePrice/igtools/getFollowersRepositories"
)

type (
	GetFollowersUsecaseService interface {
		GetFollowers(name string) ([]*igtools.Followers, error)
	}

	getFollowersUseCase struct {
		// cron cron.N
		getFollowersRepo getFollowersRepositories.GetFollowersRepositoriesService
	}
)

func NewGetFollowersUseCase(getFollowersRepo getFollowersRepositories.GetFollowersRepositoriesService) GetFollowersUsecaseService {
	return &getFollowersUseCase{
		getFollowersRepo: getFollowersRepo,
	}
}

func (u *getFollowersUseCase) GetFollowers(name string) ([]*igtools.Followers, error) {

	log.Println("GetFollowers function is working +==============================")
	// Create a new collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.instagram.com"),
		// colly.CacheDir("./_instagram_cache/"),
	)

	// xpathExpr := "//div[2]/div/div/div[2]/div/div/div[1]/div[1]/div[2]/div[2]/section/main/div/header/section/ul/li[2]/a"

	c.OnXML("//div[2]/div/div/div[2]/div/div/div[1]/div[1]/div[2]/div[2]/section/main/div/header/section/ul/li[2]/a", func(e *colly.XMLElement) {
		Name := e.ChildText("span._ac2a")
		// Use XPath to select elements
		// nodes, _ := xmlquery.QueryAll(&xmlquery.Node{}, xpathExpr)
		fmt.Println(Name)
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	// Replace 'INSTAGRAM_PROFILE_URL' with the actual Instagram profile URL
	err := c.Visit("https://www.instagram.com/al__pha____jt/followers/")
	if err != nil {
		log.Println(err)
	}
	return nil, nil

}
