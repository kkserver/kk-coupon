package coupon

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CouponUseTaskResult struct {
	app.Result
	Receive *CouponReceive `json:"receive,omitempty"`
}

type CouponUseTask struct {
	app.Task

	ReceiveId int64 `json:"receiveId"`

	Value int64 `json:"value"` //总计金额
	Count int64 `json:"count"` //总计数量

	UseType    string `json:"useType"`    //使用类型
	UseTradeNo string `json:"useTradeNo"` //使用流水号

	Result CouponUseTaskResult
}

func (task *CouponUseTask) GetResult() interface{} {
	return &task.Result
}

func (task *CouponUseTask) GetInhertType() string {
	return "coupon"
}

func (task *CouponUseTask) GetClientName() string {
	return "Coupon.Use"
}
