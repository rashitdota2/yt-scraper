package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains("youtube.com", "www.youtube.com"))

	c.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36"
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "YSC=qbcuSsPIOnY; SOCS=CAISNQgDEitib3FfaWRlbnRpdHlmcm9udGVuZHVpc2VydmVyXzIwMjUwOTEwLjA2X3AwGgJlbiACGgYIgK6dxgY; VISITOR_INFO1_LIVE=v36GU1Zz_Sg; GPS=1; VISITOR_INFO1_LIVE=v36GU1Zz_Sg; VISITOR_PRIVACY_METADATA=CgJCRxIhEh0SGwsMDg8QERITFBUWFxgZGhscHR4fICEiIyQlJiBX; __Secure-ROLLOUT_TOKEN=CIagw6H7hdnN4gEQxcPBhePajwMY_tiQhuPajwM%3D; PREF=f4=4000000&f6=40000000&tz=Asia.Ashgabat")
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("VISITED")

	})

	c.OnHTML("a", func(h *colly.HTMLElement) {
		href := h.Attr("href")
		if href != "" {
			fmt.Println("Нашли ссылку:", href)
		}
	})

	err := c.Visit("https://www.youtube.com/@JohnWatsonRooney/videos")
	if err != nil {
		fmt.Println(err)
	}
}
