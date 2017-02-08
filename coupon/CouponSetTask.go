package coupon

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CouponSetTaskResult struct {
	app.Result
	Coupon *Coupon `json:"coupon,omitempty"`
}

type CouponSetTask struct {
	app.Task

	Id int64 `json:"id"`

	MaxCount  interface{} `json:"maxCount"`  //最大派发数量
	UMaxCount interface{} `json:"umaxCount"` //用户领取最大数量

	UseMaxCount interface{} `json:"useMaxCount"` //最大使用数量
	UseMinCount interface{} `json:"useMinCount"` //最小使用数量
	UseMaxValue interface{} `json:"useMaxValue"` //最大使用金额
	UseMinValue interface{} `json:"useMinValue"` //最小使用金额

	MinValue interface{} `json:"minValue"` //最小金额
	MaxValue interface{} `json:"maxValue"` //最大金额

	MinRebate interface{} `json:"minRebate"` //最小折扣 (0~100)
	MaxRebate interface{} `json:"maxRebate"` //最大折扣 (0~100)

	StartTimeType interface{} `json:"startTimeType"` //开始有效时间类型
	StartTime     interface{} `json:"startTime"`     //开始有效时间
	EndTimeType   interface{} `json:"endTimeType"`   //开始有效时间类型
	EndTime       interface{} `json:"endTime"`       //结束有效时间

	Title   interface{} `json:"title"`
	Summary interface{} `json:"summary"`
	Remark  interface{} `json:"remark"`

	Status interface{} `json:"status"`

	Result CouponSetTaskResult
}

func (task *CouponSetTask) GetResult() interface{} {
	return &task.Result
}

func (task *CouponSetTask) GetInhertType() string {
	return "coupon"
}

func (task *CouponSetTask) GetClientName() string {
	return "Coupon.Set"
}
