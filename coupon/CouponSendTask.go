package coupon

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CouponSendTaskResult struct {
	app.Result
	Coupon *CouponReceive `json:"coupon,omitempty"`
}

type CouponSendTask struct {
	app.Task

	Id  int64 `json:"id"`
	Uid int64 `json:"uid"`

	Result CouponSendTaskResult
}

func (task *CouponSendTask) GetResult() interface{} {
	return &task.Result
}

func (task *CouponSendTask) GetInhertType() string {
	return "coupon"
}

func (task *CouponSendTask) GetClientName() string {
	return "Coupon.Send"
}
