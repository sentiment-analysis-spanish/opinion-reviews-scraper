package scrapers

import (
	"opinion-reviews-scraper/models"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type TripAdvisorReviewsScraperInProduct struct {
	Config models.ScrapingConfig
}

func RateTextToFloatTripAdvisor(rateText string) float64 {
	splitted := strings.Split(rateText, "_")
	rate, _ := strconv.ParseFloat(splitted[1], 64)
	rate = rate / 50
	return rate
}

func (scraper *TripAdvisorReviewsScraperInProduct) ScrapPage(urlNew UrlNew) []models.ReviewScraped {
	results := []models.ReviewScraped{}

	ajaxUrl := urlNew.url
	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("https://elpais.com/"),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 1 * time.Second,
	})

	c.OnHTML("div[data-test-target]", func(e *colly.HTMLElement) {
		if e.Attr("data-test-target") == "HR_CC_CARD" {
			result := models.ReviewScraped{}
			text := ""
			rateText := ""
			username := ""
			date := ""
			title := ""

			e.ForEach("q", func(_ int, elem *colly.HTMLElement) {
				text = elem.Text
			})

			e.ForEach("span[class]", func(_ int, elem *colly.HTMLElement) {
				if strings.Contains(elem.Attr("class"), "ui_bubble_rating") {
					rateText = strings.Split(elem.Attr("class"), " ")[1]
				}
			})

			e.ForEach("span[class]", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("class") == "location-review-review-list-parts-EventDate__event_date--1epHa" {
					date = strings.ReplaceAll(elem.Text, "Fecha de la estancia:", "")
				}
			})
			//social-member-event-MemberEventOnObjectBlock__member--35-jC
			e.ForEach("a[class]", func(_ int, elem *colly.HTMLElement) {
				if strings.Contains(elem.Attr("class"), "social-member-event-MemberEventOnObjectBlock__member--35-jC") {
					username = elem.Text
				}
			})

			//location-review-review-list-parts-ReviewTitle__reviewTitleText--2tFRT
			e.ForEach("a[class]", func(_ int, elem *colly.HTMLElement) {
				if strings.Contains(elem.Attr("class"), "location-review-review-list-parts-ReviewTitle__reviewTitleText--2tFRT") {
					title = elem.Text
				}
			})

			result.Content = title + " " + text
			result.Title = title
			result.RateText = rateText
			result.Date = date
			result.Rate = RateTextToFloatTripAdvisor(rateText)
			result.User = username
			result.Source = "tripadvisor"
			results = append(results, result)
			log.Println("obtained new review by user " + username + " rate: " + rateText)

		}
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		url := "https://www.tripadvisor.es/" + e.Attr("href") //strings.Replace(e.Attr("href"), "/", "", 1)
		if e.Text == "Siguiente" {
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
