package main

import (
	"fmt"
	"imbd_goroutine_concurency/config"
	"imbd_goroutine_concurency/db"
	"imbd_goroutine_concurency/entity"
	"imbd_goroutine_concurency/helper"
	"imbd_goroutine_concurency/model"

	"github.com/google/uuid"
)

var (
	Config = config.Config()
	DB     = db.Connect(Config)
)

func crawl() <-chan bool {

	crawled := make(chan bool)

	go func() {
		films := []model.CrawledFilm{}
		f1 := helper.Crawl("https://www.imdb.com/chart/moviemeter/?ref_=nv_mv_mpm")
		f2 := helper.Crawl("https://www.imdb.com/chart/top/?ref_=nv_mv_250")

		if films1 := <-f1; films1 != nil {
			fmt.Println("Finish f1")
			films = append(films, films1...)
		}
		if films2 := <-f2; films2 != nil {
			fmt.Println("Finish f2")
			films = append(films, films2...)
		}

		for _, v := range films {
			filmID := uuid.NewString()
			film := entity.Film{
				FilmID: filmID,
				Name:   v.Name,
				Year:   v.Year,
				Rating: v.Rating,
			}

			DB.MustExec("BEGIN;")
			DB.MustExec("INSERT INTO films VALUES (?, ?, ?, ?)", filmID, film.Name, film.Year, film.Rating)
			DB.MustExec("COMMIT;")
		}

		crawled <- true
	}()

	return crawled
}

func main() {
	defer DB.Close()

	crawled := crawl()

	select {
	case <-crawled:
		fmt.Println("Crawled successfully")
	}
}
