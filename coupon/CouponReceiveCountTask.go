package coupon

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CouponReceiveCountTaskResult struct {
	app.Result
	Count int `json:"count"`
}

type CouponReceiveCountTask struct {
	app.Task
	Uid      int64  `json:"uid"`
	CouponId int64  `json:"couponId"`
	Status   string `json:"status"`
	Result   CouponReceiveCountTaskResult
}

func (task *CouponReceiveCountTask) GetResult() interface{} {
	return &task.Result
}

func (task *CouponReceiveCountTask) GetInhertType() string {
	return "coupon"
}

func (task *CouponReceiveCountTask) GetClientName() string {
	return "Coupon.ReceiveCount"
}
