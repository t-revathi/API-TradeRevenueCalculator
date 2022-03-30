package service

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"math"
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
 processTransactions(transactionData,req.Config)


}

func processTransactions(transactions []Transaction,config Config){
	formatTransactions(&transactions)
	fmt.Printf("%v+",transactions)
	/*buyShares := getbuyShares(transactions, config)
	//fmt.Println(buyShares)
	sellShares := getsellShares(transactions, config)
	return calculatePandL(buyShares, sellShares, config)*/
}

func formatTransactions(transactions *[]Transaction){
var date string
var unitprice float32
	for i,v := range *transactions{
		
		date = v.Date.Format("1/2/2006")
		v.Date,_ = time.Parse("1/2/2006",date)
		unitprice = v.Cost/float32(v.Quantity)
		v.UnitPrice = float32(math.Abs(float64(unitprice)))
		fmt.Printf("index: %v , value : %v \n",i,v)
	}
	

}