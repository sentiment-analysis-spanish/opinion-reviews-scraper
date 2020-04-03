package scrapers

import (
	"opinion-reviews-scraper/models"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type FilmaffinityReviewsScraper struct {
	Config models.ScrapingConfig
}

func RateTextToFloatFilmaffinity(rateText string) float64 {
	rate, _ := strconv.ParseFloat(rateText, 64)
	rate = rate / 10
	return rate
}

func (scraper *FilmaffinityReviewsScraper) ScrapPage(urlNew UrlNew) []models.ReviewScraped {
	results := []models.ReviewScraped{}
	var ajaxUrl string
	if strings.Contains(urlNew.url, "/es/film") {
		id := strings.Split(urlNew.url, "/es/film")[1]
		ajaxUrl = "https://www.filmaffinity.com/es/reviews/1/" + id
	}
	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("https://elpais.com/"),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 1 * time.Second,
	})

	c.OnHTML("div", func(e *colly.HTMLElement) {
		if e.Attr("class") == "fa-shadow movie-review-wrapper rw-item" {
			result := models.ReviewScraped{}
			text := ""
			rateText := ""
			username := ""
			date := ""
			title := ""

			e.ForEach("div", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("class")=="review-text1" {
					text = elem.Text
				}
			})

			e.ForEach("div[class]", func(_ int, elem *colly.HTMLElement) {
				if strings.Contains(elem.Attr("class"), "user-reviews-movie-rating") {
					rateText = strings.TrimSpace(elem.Text)
				}
			})

			e.ForEach("div[class]", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("class") == "review-date" {
					date =elem.Text
				}
			})
			//social-member-event-MemberEventOnObjectBlock__member--35-jC
			e.ForEach("a", func(_ int, elem *colly.HTMLElement) {
				if strings.Contains(elem.Attr("href"), "/userreviews/") && !strings.Contains(elem.Text, " críticas") {
					username = elem.Text
				}
			})

			//location-review-review-list-parts-ReviewTitle__reviewTitleText--2tFRT
			e.ForEach("div[class]", func(_ int, elem *colly.HTMLElement) {
				if strings.Contains(elem.Attr("class"), "review-title") {
					title = elem.Text
				}
			})

			result.Content = title + " " + text
			result.Title = title
			result.RateText = rateText
			result.Date = date
			result.Rate = RateTextToFloatFilmaffinity(rateText)
			result.User = username

				result.Source = "filmaffinity"

			results = append(results, result)
			log.Println("obtained new review by user " + username + " rate: " + rateText)

		}
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		url := e.Attr("href") //strings.Replace(e.Attr("href"), "/", "", 1)
		if e.Text == ">>" {
			c.Visit(url)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Info("Visiting\n", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Info("Something went wrong:", err)
	})

	c.Visit(ajaxUrl)
	c.Wait()

	return results

}
