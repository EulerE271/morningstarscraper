package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	isin := "SE0012193019"

	client := &http.Client{}
	url := fmt.Sprintf("https://tools.morningstar.co.uk/api/rest.svc/klr5zyak8x/security/screener?page=1&pageSize=20&sortOrder=ReturnM60%%20DESC&outputType=json&version=1&languageId=sv-SE&currencyId=SEK&universeIds=FOSWE%%24%%24ALL_5498&securityDataPoints=SecId%%7CName%%7CPriceCurrency%%7CTenforeId%%7CReturnM0%%7CReturnM60%%7CStandardDeviationM60%%7COngoingCharge%%7CFundTNAV%%7CStarRatingM255%%7CMedalist_RatingNumber%%7CSustainabilityRank%%7CTrailingDate%%7CClosePrice%%7CReturnD1%%7CReturnW1%%7CReturnM1%%7CReturnM3%%7CReturnM6%%7CReturnM12%%7CReturnM36%%7CReturnM120%%7CReturnM180%%7CMaxFrontEndLoad%%7CMaximumExitCostAcquired%%7CPerformanceFeeActual%%7CFeeLevel%%7CInitialPurchase%%7CEquityStyleBox%%7CBondStyleBox%%7CAverageCreditQualityCode%%7CEffectiveDuration%%7CPortfolioDate%%7CCollectedSRRI%%7CMorningstarRiskM255%%7CStandardDeviationM12%%7CStandardDeviationM36%%7CstandardDeviationM120%%7CIsin%%7CppmCode%%7CTrackRecordExtension&filters=&term=%s&subUniverseId=", isin)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "tools.morningstar.co.uk")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("origin", "https://www.morningstar.se")
	req.Header.Set("referer", "https://www.morningstar.se/se/screener/fund.aspx")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="117", "Not;A=Brand";v="8", "Chromium";v="117"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
