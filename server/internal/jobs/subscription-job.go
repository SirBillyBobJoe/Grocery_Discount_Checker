package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"learn_go/internal/app"
	"learn_go/internal/data"
	"learn_go/internal/model"
	"learn_go/internal/service"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const woolWorthsEndpoint string = "https://www.woolworths.co.nz/api/v1/products/"

func SubscriptionJob(ctx context.Context, app *app.App) {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Background Job Done!")
			return

		case <-ticker.C:
			fmt.Println("Running Background Job")
			if subscriptions, err := app.SubscriptionRepository.GetAllSubscriptions(ctx); err != nil {
				log.Error("Server error: ", err)
			} else {
				var waitGroup sync.WaitGroup

				for _, sub := range subscriptions {
					waitGroup.Add(1)

					go retrieveProductData(sub, &waitGroup, ctx, app)
				}
				waitGroup.Wait()
			}
		}

	}
}

func retrieveProductData(sub model.SubcriptionModel, waitGroup *sync.WaitGroup, ctx context.Context, app *app.App) {
	defer waitGroup.Done()

	cli := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", woolWorthsEndpoint+sub.ItemId, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Add("user-agent", "some agent")
	req.Header.Add("X-Requested-With", "OnlineShopping.WebApp")

	response, err := cli.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()

	var item data.WoolWorthsItemResponse
	json.NewDecoder(response.Body).Decode(&item)

	checkAndUpdatePrice(item, sub, ctx, app)
}

func checkAndUpdatePrice(item data.WoolWorthsItemResponse, subscription model.SubcriptionModel, ctx context.Context, app *app.App) {
	if subscription.OriginalPrice != item.Price.OriginalPrice || subscription.CurrentPrice != item.Price.SalePrice {
		app.SubscriptionRepository.UpdateSubscription(ctx, subscription, item.Price.OriginalPrice, item.Price.SalePrice)
	}

	if item.Price.SalePrice < subscription.CurrentPrice {
		for _, email := range subscription.Emails {
			go service.SendAsyncEmail(item.Name, subscription.OriginalPrice, item.Price.SalePrice, email)
		}
		return
	}

	if item.Price.OriginalPrice < subscription.OriginalPrice {

		for _, email := range subscription.Emails {
			go service.SendAsyncEmail(item.Name, subscription.OriginalPrice, item.Price.OriginalPrice, email)
		}
		return
	}
}
