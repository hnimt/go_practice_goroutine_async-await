package helper

import (
	"fmt"
	"imbd_goroutine_concurency/model"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func Crawl(path string) <-chan []model.CrawledFilm {
	log.Printf("Crawl %v", path)

	chanFilms := make(chan []model.CrawledFilm)
	crFilms := []model.CrawledFilm{}

	go func() {
		defer close(chanFilms)

		col := colly.NewCollector()

		col.OnHTML("tbody.lister-list", func(e *colly.HTMLElement) {
			e.ForEach("tr", func(i int, h *colly.HTMLElement) {
				var crFilm model.CrawledFilm
				crFilm.Name = h.ChildText("td.titleColumn > a")

				year := h.ChildText("td.titleColumn > span.secondaryInfo")
				year = strings.Replace(year, "(", "", -1)
				year = strings.Replace(year, ")", "", -1)
				crFilm.Year = year

				rating := h.ChildText("td.ratingColumn > strong")
				rating = strings.Replace(rating, ",", ".", -1)
				ratingInt, _ := strconv.ParseFloat(rating, 64)
				crFilm.Rating = ratingInt
				crFilms = append(crFilms, crFilm)
			})
		})

		col.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})

		col.Visit(path)
		chanFilms <- crFilms
	}()

	return chanFilms
}
