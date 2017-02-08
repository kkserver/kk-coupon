package coupon

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CouponUseQueryCounter struct {
	PageIndex int `json:"p"`
	PageSize  int `json:"size"`
	PageCount int `json:"count"`
	RowCount  int `json:"rowCount"`
}

type CouponUse struct {
	CouponReceive
	OfferValue int64 `json:"offerValue"` //优惠金额
}

type CouponUseQueryTaskResult struct {
	app.Result
	Counter *CouponQueryCounter `json:"counter,omitempty"`
	Coupons []CouponUse         `json:"coupons,omitempty"`
}

type CouponUseQueryTask struct {
	app.Task

	Uid        int64  `json:"uid"`
	ReceiveIds string `json:"receiveIds"`

	Value int64 `json:"value"` //总计金额
	Count int64 `json:"count"` //总计数量

	PageIndex int  `json:"p"`
	PageSize  int  `json:"size"`
	Counter   bool `json:"counter"`
	Result    CouponUseQueryTaskResult
}

func (task *CouponUseQueryTask) GetResult() interface{} {
	return &task.Result
}

func (task *CouponUseQueryTask) GetInhertType() string {
	return "coupon"
}

func (task *CouponUseQueryTask) GetClientName() string {
	return "Coupon.UseQuery"
}
