package shop

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/davecgh/go-spew/spew"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/smartwalle/paypal"
	"gorm.io/gorm"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/shop"
	"project.io/shop/web/http/response/shop"
	"strconv"
	"strings"
)

func DoPaymentOfPaypal(c context.Context, ctx *app.RequestContext) {

	var request req.DoPaymentOfPaypal

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var payment model.ShpPayment

	fp := facades.Gorm.Scopes(scope.Platform(ctx)).Preload("Order").First(&payment, "`id`=? and `channel`=? and `user_id`=?", request.ID, model.ShpPaymentOfChannelPaypal, auth.ID(ctx))

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "Order not found.")
		return
	} else if fp.Error != nil || payment.Order == nil {
		http.Fail(ctx, "Order query failed. Please try again later.")
		return
	}

	var address model.ShpOrderAddress

	facades.Gorm.Scopes(scope.Platform(ctx)).First(&address, "`order_id`=?", payment.OrderID)

	var details []model.ShpOrderDetail

	facades.Gorm.Scopes(scope.Platform(ctx)).Find(&details, "`order_id`=?", payment.OrderID)

	client := paypal.New(facades.Cfg.GetString("payment.paypal.client_id"), facades.Cfg.GetString("payment.paypal.secret_id"), facades.Cfg.GetBool("payment.paypal.live"))

	information := &paypal.Payment{
		Intent: paypal.PaymentIntentSale,
		Payer: &paypal.Payer{
			PaymentMethod: "paypal",
			PayerInfo: &paypal.PayerInfo{
				Email:     address.Email,
				FirstName: address.FirstName,
				LastName:  address.LastName,
			},
		},
		RedirectURLs: &paypal.RedirectURLs{
			ReturnURL: facades.Cfg.GetString("payment.paypal.url.return"),
			CancelURL: facades.Cfg.GetString("payment.paypal.url.cancel"),
		},
	}

	transaction := &paypal.Transaction{
		ReferenceId: "",
		Amount: &paypal.Amount{
			Total:    strconv.FormatFloat(float64(payment.Money)/100, 'f', 2, 64),
			Currency: "USD",
			Details: &paypal.AmountDetails{
				Subtotal: strconv.FormatFloat(float64(payment.Order.TotalPrice)/100, 'f', 2, 64),
				Shipping: strconv.FormatFloat(float64(payment.Order.CostShipping)/100, 'f', 2, 64),
			},
		},
		InvoiceNumber: payment.ID,
		ItemList: &paypal.ItemList{
			Items: make([]*paypal.Item, len(details)),
		},
	}

	for idx, item := range details {
		transaction.ItemList.Items[idx] = &paypal.Item{
			Name:     item.Name,
			Price:    strconv.FormatFloat(float64(item.Price)/100, 'f', 2, 64),
			Currency: "USD",
			Quantity: fmt.Sprintf("%v", item.Quantity),
			SKU:      strings.Join(item.Specifications, ";"),
		}
	}

	information.Transactions = []*paypal.Transaction{transaction}

	order, err := client.CreatePayment(information)

	if err != nil {
		spew.Dump(err.Error())
		http.Fail(ctx, "Payment initiation failed. Please try again later.")
		return
	}

	if result := facades.Gorm.Model(&payment).Update("no", order.Id); result.Error != nil {
		http.Fail(ctx, "Payment initiation failed. Please try again later.")
		return
	}

	responses := shop.DoPaymentOfPaypal{}

	for _, item := range order.Links {
		if item.Rel == "approval_url" {
			responses.Link = item.Href
			break
		}
	}

	http.Success(ctx, responses)
}
