package main

import (
	"context"
	"fmt"
	"github.com/jmjp/go-asaas/asaas"
	"os"
)

var paymentLinkAsaas asaas.PaymentLink

func main() {
	paymentLinkAsaas = asaas.NewPaymentLink(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	createPaymentLink()
	updatePaymentLinkById()
	getPaymentLinkById()
	getAllPaymentLink()
	deletePaymentLinkById()
}

func createPaymentLink() {
	resp, err := paymentLinkAsaas.Create(context.TODO(), asaas.CreatePaymentLinkRequest{
		Name:                "",
		Description:         "",
		BillingType:         "",
		ChargeType:          "",
		EndDate:             asaas.Date{},
		Value:               0,
		DueDateLimitDays:    0,
		SubscriptionCycle:   "",
		MaxInstallmentCount: 0,
		NotificationEnabled: false,
		Callback:            nil,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func updatePaymentLinkById() {
	resp, err := paymentLinkAsaas.UpdateById(context.TODO(), "", asaas.UpdatePaymentLinkRequest{
		Name:                "",
		Description:         nil,
		BillingType:         "",
		ChargeType:          "",
		EndDate:             asaas.Date{},
		Value:               nil,
		DueDateLimitDays:    0,
		SubscriptionCycle:   nil,
		MaxInstallmentCount: 0,
		NotificationEnabled: nil,
		Callback:            nil,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else if resp.IsNoContent() {
		fmt.Println("no content:", resp.Errors)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func deletePaymentLinkById() {
	resp, err := paymentLinkAsaas.DeleteById(context.TODO(), "")
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else if resp.IsNoContent() {
		fmt.Println("no content:", resp.Errors)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func getPaymentLinkById() {
	resp, err := paymentLinkAsaas.GetById(context.TODO(), "")
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else if resp.IsNoContent() {
		fmt.Println("no content:", resp)
	} else {
		fmt.Println("no content:", resp)
	}
}

func getAllPaymentLink() {
	resp, err := paymentLinkAsaas.GetAll(context.TODO(), asaas.GetAllPaymentLinksRequest{
		Name:           "",
		Active:         nil,
		IncludeDeleted: nil,
		Offset:         0,
		Limit:          0,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else if resp.IsNoContent() {
		fmt.Println("no content:", resp)
	} else {
		fmt.Println("no content:", resp)
	}
}
