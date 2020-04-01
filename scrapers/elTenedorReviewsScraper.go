package scrapers

import (
	"opinion-reviews-scraper/models"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type ElTenedorReviewsScraper struct {
	Config models.ScrapingConfig
}

func RateTextToFloatElTenedor(rateText string) float64 {
	rateText = strings.ReplaceAll(rateText, ",", ".")
	rate, _ := strconv.ParseFloat(rateText, 64)
	rate = rate / 10
	return rate
}

func (scraper *ElTenedorReviewsScraper) ScrapPage(urlNew UrlNew) []models.ReviewScraped {
	results := []models.ReviewScraped{}

	ajaxUrl := urlNew.url + "/opiniones"
	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("https://elpais.com/"),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 1 * time.Second,
	})

	c.OnHTML("li[data-test]", func(e *colly.HTMLElement) {
		if e.Attr("data-test") == "restaurant-page-review-item" {
			result := models.ReviewScraped{}
			text := ""
			rateText := ""
			username := ""
			date := ""
			title := ""

			e.ForEach("p", func(_ int, elem *colly.HTMLElement) {
				if elem.Attr("class")=="eveXu" {
					text = elem.Text
				}
			})

			e.ForEach("strong", func(_ int, elem *colly.HTMLElement) {
				if strings.Contains(elem.Attr("data-test"), "rating-value") {
					rateText = elem.Text
				}
			})

			e.ForEach("time", func(_ int, elem *colly.HTMLElement) {
					date = elem.Attr("datetime")
			})
			//social-member-event-MemberEventOnObjectBlock__member--35-jC
			e.ForEach("cite", func(_ int, elem *colly.HTMLElement) {
				if strings.Contains(elem.Attr("class"), "_1rTLO") {
					username = elem.Text
				}
			})


			result.Content = text
			result.Title = title
			result.RateText = rateText
			result.Date = date
			result.Rate = RateTextToFloatElTenedor(rateText)
			result.User = username
			result.Source = "eltenedor"
			results = append(results, result)
			log.Println("obtained new review by user " + username + " rate: " + rateText)

		}
	})

	c.OnHTML("ul", func(e *colly.HTMLElement) {
		if e.Attr("class")=="_1fOA6" {
			e.ForEach("li", func(_ int, elem *colly.HTMLElement) {
				elem.ForEach("a", func(_ int, elem2 *colly.HTMLElement) {
					url := "https://www.eltenedor.es" + elem2.Attr("href") //strings.Replace(e.Attr("href"), "/", "", 1)
					c.Visit(url)
				})
			})
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
