# News scraper go

This repository contains the code of the workers that extract data from newspapers **el mundo**, **el pais**, **la vanguardia** and **la razón**. 

This code asumes that you are running the [backend rest service](https://github.com/news-scrapers/backend-rest) so go there, configure it and run it.

## Installation

You will also need **golang** (version 1.12 at least). Follow [this](https://golang.org/doc/install) to install it.

## Configuration
* Clone this repository to a directory:

     git clone https://github.com/news-scrapers/news-scraper-workers-go.git

* Move to the cloned directory and create a file named **scrapingConfig.json** . Inside this file you will need to specify the url to the backend rest instance (by default localhost:8000 if you are running locally) and the newspapers that you want to scrap. Also an id for your scraper. Here is an example that will work if you are running the backend locally and if you want to scrap all newspapers:
  
        {
        "url_base": "http://localhost:8000",
        "scraper_id": "scraperTest",
        "device_id": "deviceTest",
        "app_id": "app_id_test",
        "newspaper": ["elpais","elmundo", "lavanguardia", "abc"]
        }
* Run the golang code. Inside the project folder run:
  
        go run main.go
    you should see the logs with the scraped news of the different newspapers.