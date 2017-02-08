package coupon

import (
	"bytes"
	"fmt"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/dynamic"
	"strings"
	"time"
)

type CouponService struct {
	app.Service

	Get    *CouponTask
	Set    *CouponSetTask
	Remove *CouponRemoveTask
	Create *CouponCreateTask
	Query  *CouponQueryTask

	Send     *CouponSendTask
	Use      *CouponUseTask
	UseQuery *CouponUseTask

	ReceiveQuery *CouponReceiveQueryTask
}

func (S *CouponService) Handle(a app.IApp, task app.ITask) error {
	return app.ServiceReflectHandle(a, task, S)
}

func (S *CouponService) HandleCouponCreateTask(a ICouponApp, task *CouponCreateTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Coupon{}

	v.Type = task.Type
	v.MaxCount = task.MaxCount
	v.UMaxCount = task.UMaxCount
	v.UseMaxCount = task.UseMaxCount
	v.UseMinCount = task.UseMinCount
	v.UseMaxValue = task.UseMaxValue
	v.UseMinValue = task.UseMinValue
	v.MinValue = task.MinValue
	v.MaxValue = task.MaxValue
	v.MinRebate = task.MinRebate
	v.MaxRebate = task.MaxRebate
	v.StartTimeType = task.StartTimeType
	v.StartTime = task.StartTime
	v.EndTimeType = task.EndTimeType
	v.EndTime = task.EndTime
	v.Title = task.Title
	v.Summary = task.Summary
	v.Remark = task.Remark
	v.Ctime = time.Now().Unix()

	_, err = kk.DBInsert(db, a.GetCouponTable(), a.GetPrefix(), &v)

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Coupon = &v

	return nil
}

func (S *CouponService) HandleCouponSetTask(a ICouponApp, task *CouponSetTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Coupon{}

	rows, err := kk.DBQuery(db, a.GetCouponTable(), a.GetPrefix(), " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {
		scanner := kk.NewDBScaner(&v)
		err = scanner.Scan(rows)
		if err != nil {
			task.Result.Errno = ERROR_COUPON
			task.Result.Errmsg = err.Error()
			return nil
		}
	} else {
		task.Result.Errno = ERROR_COUPON_NOT_FOUND
		task.Result.Errmsg = "Not Found coupon"
		return nil
	}

	keys := map[string]bool{}

	if task.MaxCount != nil {
		v.MaxCount = int(dynamic.IntValue(task.MaxCount, int64(v.MaxCount)))
		keys["maxcount"] = true
	}

	if task.UMaxCount != nil {
		v.UMaxCount = int(dynamic.IntValue(task.UMaxCount, int64(v.UMaxCount)))
		keys["umaxcount"] = true
	}

	if task.UseMaxCount != nil {
		v.UseMaxCount = int(dynamic.IntValue(task.UseMaxCount, int64(v.UseMaxCount)))
		keys["usemaxcount"] = true
	}

	if task.UseMinCount != nil {
		v.UseMinCount = int(dynamic.IntValue(task.UseMinCount, int64(v.UseMinCount)))
		keys["usemincount"] = true
	}

	if task.UseMaxValue != nil {
		v.UseMaxValue = dynamic.IntValue(task.UseMaxValue, v.UseMaxValue)
		keys["usemaxvalue"] = true
	}

	if task.UseMinValue != nil {
		v.UseMinValue = dynamic.IntValue(task.UseMinValue, v.UseMinValue)
		keys["useminvalue"] = true
	}

	if task.MinValue != nil {
		v.MinValue = dynamic.IntValue(task.MinValue, v.MinValue)
		keys["minvalue"] = true
	}

	if task.MaxValue != nil {
		v.MaxValue = dynamic.IntValue(task.MaxValue, v.MaxValue)
		keys["maxvalue"] = true
	}

	if task.MinRebate != nil {
		v.MinRebate = dynamic.IntValue(task.MinRebate, v.MinRebate)
		keys["minrebate"] = true
	}

	if task.MaxRebate != nil {
		v.MaxRebate = dynamic.IntValue(task.MaxRebate, v.MaxRebate)
		keys["maxrebate"] = true
	}

	if task.StartTimeType != nil {
		v.StartTimeType = int(dynamic.IntValue(task.StartTimeType, int64(v.StartTimeType)))
		keys["starttimetype"] = true
	}

	if task.StartTime != nil {
		v.StartTime = dynamic.IntValue(task.StartTime, v.StartTime)
		keys["starttime"] = true
	}

	if task.EndTimeType != nil {
		v.EndTimeType = int(dynamic.IntValue(task.EndTimeType, int64(v.EndTimeType)))
		keys["endtimetype"] = true
	}

	if task.EndTime != nil {
		v.EndTime = dynamic.IntValue(task.EndTime, v.EndTime)
		keys["endtime"] = true
	}

	if task.Title != nil {
		v.Title = dynamic.StringValue(task.Title, v.Title)
		keys["title"] = true
	}

	if task.Summary != nil {
		v.Summary = dynamic.StringValue(task.Summary, v.Summary)
		keys["summary"] = true
	}

	if task.Remark != nil {
		v.Remark = dynamic.StringValue(task.Remark, v.Remark)
		keys["remark"] = true
	}

	if task.Status != nil {
		v.Status = int(dynamic.IntValue(task.Status, int64(v.Status)))
		keys["status"] = true
	}

	_, err = kk.DBUpdateWithKeys(db, a.GetCouponTable(), a.GetPrefix(), &v, keys)

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Coupon = &v

	return nil
}

func (S *CouponService) HandleCouponTask(a ICouponApp, task *CouponTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Coupon{}

	rows, err := kk.DBQuery(db, a.GetCouponTable(), a.GetPrefix(), " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {
		scanner := kk.NewDBScaner(&v)
		err = scanner.Scan(rows)
		if err != nil {
			task.Result.Errno = ERROR_COUPON
			task.Result.Errmsg = err.Error()
			return nil
		}
	} else {
		task.Result.Errno = ERROR_COUPON_NOT_FOUND
		task.Result.Errmsg = "Not Found coupon"
		return nil
	}

	task.Result.Coupon = &v

	return nil
}

func (S *CouponService) HandleCouponRemoveTask(a ICouponApp, task *CouponRemoveTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Coupon{}

	tx, err := db.Begin()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	err = func() error {

		rows, err := kk.DBQuery(tx, a.GetCouponTable(), a.GetPrefix(), " WHERE id=? FOR UPDATE", task.Id)

		if err != nil {
			task.Result.Errno = ERROR_COUPON
			task.Result.Errmsg = err.Error()
			return nil
		}

		if rows.Next() {

			scanner := kk.NewDBScaner(&v)

			err = scanner.Scan(rows)

			rows.Close()

			if err != nil {
				return err
			}

			if v.Count > 0 {
				return app.NewError(ERROR_COUPON_COUNT, "Coupon sended can not be deleted")
			}

			_, err = kk.DBDelete(tx, a.GetCouponTable(), a.GetPrefix(), " WHERE id=?", task.Id)

			if err != nil {
				return err
			}

		} else {
			rows.Close()
			return app.NewError(ERROR_COUPON_NOT_FOUND, "Not Found coupon")
		}

		return nil
	}()

	if err == nil {
		err = tx.Commit()
	}

	if err != nil {
		tx.Rollback()
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Coupon = &v

	return nil
}

func (S *CouponService) HandleCouponQueryTask(a ICouponApp, task *CouponQueryTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	var coupons = []Coupon{}

	var args = []interface{}{}

	var sql = bytes.NewBuffer(nil)

	sql.WriteString(" WHERE 1")

	if task.Id != 0 {
		sql.WriteString(" AND id=?")
		args = append(args, task.Id)
	}

	if task.Status != "" {

		sql.WriteString(" AND status IN (")

		for i, v := range strings.Split(task.Status, ",") {
			if i != 0 {
				sql.WriteString(",")
			}
			sql.WriteString("?")
			args = append(args, v)
		}

		sql.WriteString(")")
	}

	if task.Keyword != "" {
		q := "%" + task.Keyword + "%"
		sql.WriteString(" AND (title LIKE ? OR summary LIKE ? OR remark LIKE ?)")
		args = append(args, q, q, q)
	}

	if task.OrderBy == "asc" {
		sql.WriteString(" ORDER BY id ASC")
	} else {
		sql.WriteString(" ORDER BY id DESC")
	}

	var pageIndex = task.PageIndex
	var pageSize = task.PageSize

	if pageIndex < 1 {
		pageIndex = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	if task.Counter {
		var counter = CouponQueryCounter{}
		counter.PageIndex = pageIndex
		counter.PageSize = pageSize
		counter.RowCount, err = kk.DBQueryCount(db, a.GetCouponTable(), a.GetPrefix(), sql.String(), args...)
		if err != nil {
			task.Result.Errno = ERROR_COUPON
			task.Result.Errmsg = err.Error()
			return nil
		}
		if counter.RowCount%pageSize == 0 {
			counter.PageCount = counter.RowCount / pageSize
		} else {
			counter.PageCount = counter.RowCount/pageSize + 1
		}
		task.Result.Counter = &counter
	}

	sql.WriteString(fmt.Sprintf(" LIMIT %d,%d", (pageIndex-1)*pageSize, pageSize))

	var v = Coupon{}
	var scanner = kk.NewDBScaner(&v)

	rows, err := kk.DBQuery(db, a.GetCouponTable(), a.GetPrefix(), sql.String(), args...)

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	for rows.Next() {

		err = scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_COUPON
			task.Result.Errmsg = err.Error()
			return nil
		}

		coupons = append(coupons, v)
	}

	task.Result.Coupons = coupons

	return nil
}
