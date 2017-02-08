package coupon

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CouponRemoveTaskResult struct {
	app.Result
	Coupon *Coupon `json:"coupon,omitempty"`
}

type CouponRemoveTask struct {
	app.Task

	Id int64 `json:"id"`

	Result CouponRemoveTaskResult
}

func (task *CouponRemoveTask) GetResult() interface{} {
	return &task.Result
}

func (task *CouponRemoveTask) GetInhertType() string {
	return "coupon"
}

func (task *CouponRemoveTask) GetClientName() string {
	return "Coupon.Remove"
}
