package main

import (
	"log"
	"sync"

	"github.com/aerosystems/nix-trainee-4/internal/resource"
	"github.com/aerosystems/nix-trainee-4/pkg/client/gorm"
	"github.com/aerosystems/nix-trainee-4/pkg/client/mysql"
)

func main() {
	clientMySQL := mysql.NewClient()
	clientGORM := gorm.NewClient()
	repository := resource.NewRepository(clientMySQL, clientGORM)

	userId := 7
	posts, err := resource.GetPosts(userId)
	if err != nil {
		log.Println(err)
	}

	chanComment := make(chan resource.Comment)
	wg := new(sync.WaitGroup)
	defer wg.Wait()

	for _, post := range posts {
		wg.Add(1)
		go func(post resource.Post) {
			defer wg.Done()
			comments, err := resource.GetComments(post.Id)
			if err != nil {
				log.Println(err)
			}

			for _, comment := range comments {
				wg.Add(1)
				go func(comment resource.Comment) {
					defer wg.Done()
					chanComment <- comment
				}(comment)

			}
		}(post)
	}

	go func() {
		for {
			repository.CreateWithGORM(<-chanComment)
		}
	}()
}
