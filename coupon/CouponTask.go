package coupon

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CouponTaskResult struct {
	app.Result
	Coupon *Coupon `json:"coupon,omitempty"`
}

type CouponTask struct {
	app.Task

	Id int64 `json:"id"`

	Result CouponTaskResult
}

func (task *CouponTask) GetResult() interface{} {
	return &task.Result
}

func (task *CouponTask) GetInhertType() string {
	return "coupon"
}

func (task *CouponTask) GetClientName() string {
	return "Coupon.Get"
}
