package service

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"
	"time"
	//"github.com/go-chi/render"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) PerformCalculateProfit(ctx context.Context, w http.ResponseWriter, r *http.Request, req *DataCalculateRevenue) {

	fmt.Printf("Received financial year:%s on %s", req.Config.FinancialYear, time.Now().String())
	transactionData := req.TransactionData

	/*render.JSON(w, r,
	DataCalculateRevenue{
		transactionData,
		req.Config,
	})

	*/
	processTransactions(transactionData, req.Config)

}

func processTransactions(transactions []Transaction, config Config) {
	formatTransactions(&transactions)
	fmt.Printf("\n %v+", transactions)
	buyShares := getbuyShares(transactions, config)
	fmt.Println(buyShares)
	sellShares := getsellShares(transactions, config)
	fmt.Println(sellShares)
	//return calculatePandL(buyShares, sellShares, config)
}

func formatTransactions(transactions *[]Transaction) {

	for i := 0; i < len(*transactions); i++ {

		unitprice := (*transactions)[i].Cost / float32((*transactions)[i].Quantity)

		(*transactions)[i].UnitPrice = float32(math.Abs(float64(unitprice)))

	}

}

func getbuyShares(transactions []Transaction, config Config) []Transaction {

	buytransactions := make([]Transaction, 0)

	for _, t := range transactions {
		if config.SkipCorporateAction {
			if strings.ToLower(t.Activity) != "trade" {
				continue
			}
		}
		if strings.ToLower(t.Direction) == "buy" {
			buytransactions = append(buytransactions, t)

		}

	}
	return buytransactions
}

func getsellShares(transaction []Transaction, config Config) []Transaction {
	selltransactions := make([]Transaction, 0)
	for _, t := range transaction {
		if config.SkipCorporateAction {
			if strings.ToLower(t.Activity) != "trade" {
				continue
			}
		}
		if strings.ToLower(t.Direction) == "sell" {
			selltransactions = append(selltransactions, t)

		}
	}
	sort.Slice(selltransactions, func(i, j int) bool {
		return selltransactions[i].Date.Before(selltransactions[j].Date)
	})
	return selltransactions
}

