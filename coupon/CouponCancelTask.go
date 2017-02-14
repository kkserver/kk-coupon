package coupon

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CouponCancelTaskResult struct {
	app.Result
	Coupon *CouponReceive `json:"coupon,omitempty"`
}

type CouponCancelTask struct {
	app.Task

	Id int64 `json:"id"`

	Result CouponCancelTaskResult
}

func (task *CouponCancelTask) GetResult() interface{} {
	return &task.Result
}

func (task *CouponCancelTask) GetInhertType() string {
	return "coupon"
}

func (task *CouponCancelTask) GetClientName() string {
	return "Coupon.Cancel"
}
