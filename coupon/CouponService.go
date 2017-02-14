package coupon

import (
	"bytes"
	"fmt"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/dynamic"
	"math/rand"
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
	UseQuery *CouponUseQueryTask
	Cancel   *CouponCancelTask

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

func (S *CouponService) HandleCouponSendTask(a ICouponApp, task *CouponSendTask) error {

	if task.Id == 0 {
		task.Result.Errno = ERROR_COUPON_NOT_FOUND_ID
		task.Result.Errmsg = "Not Found id"
		return nil
	}

	if task.Uid == 0 {
		task.Result.Errno = ERROR_COUPON_NOT_FOUND_UID
		task.Result.Errmsg = "Not Found uid"
		return nil
	}

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Coupon{}
	vv := CouponReceive{}

	tx, err := db.Begin()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	err = func() error {

		rows, err := kk.DBQuery(tx, a.GetCouponTable(), a.GetPrefix(), " WHERE id=? FOR UPDATE", task.Id)

		if err != nil {
			return err
		}

		if rows.Next() {

			scanner := kk.NewDBScaner(&v)

			err = scanner.Scan(rows)

			rows.Close()

			if err != nil {
				return err
			}

			if v.Status != CouponStatusIn {
				return app.NewError(ERROR_COUPON_STATUS, "Coupon is not released and can not be distributed")
			}

			if v.Count >= v.MaxCount {
				return app.NewError(ERROR_COUPON_COUNT, "Not enough quantity")
			}

			count, err := kk.DBQueryCount(tx, a.GetCouponReceiveTable(), a.GetPrefix(), " WHERE couponid=? AND uid=?", v.Id, task.Uid)

			if err != nil {
				return err
			}

			if count >= v.UMaxCount {
				return app.NewError(ERROR_COUPON_UCOUNT, "Exceeds the user receiving limit")
			}

			vv.CouponId = v.Id
			vv.Uid = task.Uid
			vv.Type = v.Type

			vv.UseMaxCount = v.UseMaxCount
			vv.UseMinCount = v.UseMinCount
			vv.UseMaxValue = v.UseMaxValue
			vv.UseMinValue = v.UseMinValue

			now := time.Now()

			r := rand.New(rand.NewSource(now.UnixNano()))

			if v.MaxValue > v.MinValue {
				vv.Value = v.MinValue + r.Int63n(v.MaxValue-v.MinValue)
			} else {
				vv.Value = v.MinValue
			}

			if v.MaxRebate > v.MinRebate {
				vv.Rebate = v.MinRebate + r.Int63n(v.MaxRebate-v.MinRebate)
			} else {
				vv.Rebate = v.MinRebate
			}

			vv.Ctime = now.Unix()

			switch v.StartTimeType {
			case CouponTimeTypeRelative:
				vv.StartTime = vv.Ctime + v.StartTime
			case CouponTimeTypeRelativeDay:
				vv.StartTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix() + v.StartTime*24*3600
			default:
				vv.StartTime = v.StartTime
			}

			switch v.EndTimeType {
			case CouponTimeTypeRelative:
				vv.EndTime = vv.Ctime + v.EndTime
			case CouponTimeTypeRelativeDay:
				vv.EndTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix() + v.EndTime*24*3600
			default:
				vv.EndTime = v.EndTime
			}

			_, err = kk.DBInsert(tx, a.GetCouponReceiveTable(), a.GetPrefix(), &vv)

			if err != nil {
				return err
			}

			v.Count = v.Count + 1

			_, err = kk.DBUpdateWithKeys(tx, a.GetCouponTable(), a.GetPrefix(), &v, map[string]bool{"count": true})

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

	task.Result.Coupon = &vv

	return nil
}

func (S *CouponService) HandleCouponUseTask(a ICouponApp, task *CouponUseTask) error {

	if task.Id == 0 {
		task.Result.Errno = ERROR_COUPON_NOT_FOUND_ID
		task.Result.Errmsg = "Not Found id"
		return nil
	}

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := CouponReceive{}

	tx, err := db.Begin()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	err = func() error {

		rows, err := kk.DBQuery(tx, a.GetCouponReceiveTable(), a.GetPrefix(), " WHERE id=? FOR UPDATE", task.Id)

		if err != nil {
			return err
		}

		if rows.Next() {

			scanner := kk.NewDBScaner(&v)

			err = scanner.Scan(rows)

			rows.Close()

			if err != nil {
				return err
			}

			if v.Status != CouponReceiveStatusNone {
				return app.NewError(ERROR_COUPON_STATUS, "Coupon can not be used")
			}

			if time.Now().Unix() < v.StartTime {
				return app.NewError(ERROR_COUPON_TIME, "Coupon can not start")
			}

			if time.Now().Unix() > v.EndTime {
				return app.NewError(ERROR_COUPON_TIME, "Coupon has expired")
			}

			if task.Count < v.UseMinCount {
				return app.NewError(ERROR_COUPON_COUNT, "Coupon can not be used")
			}

			if v.UseMaxCount != -1 && task.Count > v.UseMaxCount {
				return app.NewError(ERROR_COUPON_COUNT, "Coupon can not be used")
			}

			if task.Value < v.UseMinValue {
				return app.NewError(ERROR_COUPON_VALUE, "Coupon can not be used")
			}

			if v.UseMaxValue != -1 && task.Value > v.UseMaxValue {
				return app.NewError(ERROR_COUPON_VALUE, "Coupon can not be used")
			}

			rows, err = kk.DBQuery(tx, a.GetCouponTable(), a.GetPrefix(), " WHERE id=? FOR UPDATE", v.CouponId)

			if err != nil {
				return err
			}

			coupon := Coupon{}

			if rows.Next() {

				scanner = kk.NewDBScaner(&coupon)

				err = scanner.Scan(rows)

				rows.Close()

				if err != nil {
					return err
				}

				v.Status = CouponReceiveStatusUse
				v.UseTime = time.Now().Unix()
				v.UseType = task.UseType
				v.UseTradeNo = task.UseTradeNo

				switch coupon.Type {
				case CouponTypeCash:
					v.UseValue = v.Value
				case CouponTypeRebate:
					v.UseValue = (100 - v.Rebate) * task.Value / 100
				}

				coupon.UseCount = coupon.UseCount + 1
				coupon.UseValue = coupon.UseValue + v.UseValue

				_, err = kk.DBUpdateWithKeys(tx, a.GetCouponTable(), a.GetPrefix(), &coupon, map[string]bool{"usecount": true, "usevalue": true})

				if err != nil {
					return err
				}

				_, err = kk.DBUpdateWithKeys(tx, a.GetCouponReceiveTable(), a.GetPrefix(), &v, map[string]bool{"status": true, "usetime": true, "usetype": true, "usetradeno": true, "usevalue": true})

				if err != nil {
					return err
				}

			} else {
				rows.Close()
				return app.NewError(ERROR_COUPON_NOT_FOUND, "Not Found Coupon")
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

func (S *CouponService) HandleCouponCancelTask(a ICouponApp, task *CouponCancelTask) error {

	if task.Id == 0 {
		task.Result.Errno = ERROR_COUPON_NOT_FOUND_ID
		task.Result.Errmsg = "Not Found id"
		return nil
	}

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := CouponReceive{}

	tx, err := db.Begin()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	err = func() error {

		rows, err := kk.DBQuery(tx, a.GetCouponReceiveTable(), a.GetPrefix(), " WHERE id=? FOR UPDATE", task.Id)

		if err != nil {
			return err
		}

		if rows.Next() {

			scanner := kk.NewDBScaner(&v)

			err = scanner.Scan(rows)

			rows.Close()

			if err != nil {
				return err
			}

			if v.Status != CouponReceiveStatusUse {
				return app.NewError(ERROR_COUPON_STATUS, "Coupon can not be canceled")
			}

			rows, err = kk.DBQuery(tx, a.GetCouponTable(), a.GetPrefix(), " WHERE id=? FOR UPDATE", v.CouponId)

			if err != nil {
				return err
			}

			coupon := Coupon{}

			if rows.Next() {

				scanner = kk.NewDBScaner(&coupon)

				err = scanner.Scan(rows)

				rows.Close()

				if err != nil {
					return err
				}

				v.Status = CouponReceiveStatusNone

				coupon.UseCount = coupon.UseCount - 1
				coupon.UseValue = coupon.UseValue - v.UseValue

				_, err = kk.DBUpdateWithKeys(tx, a.GetCouponTable(), a.GetPrefix(), &coupon, map[string]bool{"usecount": true, "usevalue": true})

				if err != nil {
					return err
				}

				_, err = kk.DBUpdateWithKeys(tx, a.GetCouponReceiveTable(), a.GetPrefix(), &v, map[string]bool{"status": true})

				if err != nil {
					return err
				}

			} else {
				rows.Close()
				return app.NewError(ERROR_COUPON_NOT_FOUND, "Not Found Coupon")
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

func (S *CouponService) HandleCouponReceiveQueryTask(a ICouponApp, task *CouponReceiveQueryTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	var coupons = []CouponReceive{}

	var args = []interface{}{}

	var sql = bytes.NewBuffer(nil)

	sql.WriteString(" WHERE 1")

	if task.CouponId != 0 {
		sql.WriteString(" AND couponid=?")
		args = append(args, task.CouponId)
	}

	if task.Uid != 0 {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
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
		var counter = CouponReceiveQueryCounter{}
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

	var v = CouponReceive{}
	var scanner = kk.NewDBScaner(&v)

	rows, err := kk.DBQuery(db, a.GetCouponReceiveTable(), a.GetPrefix(), sql.String(), args...)

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

func (S *CouponService) HandleCouponUseQueryTask(a ICouponApp, task *CouponUseQueryTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_COUPON
		task.Result.Errmsg = err.Error()
		return nil
	}

	var coupons = []CouponUse{}

	var args = []interface{}{}

	var sql = bytes.NewBuffer(nil)

	sql.WriteString(" WHERE uid=?")

	args = append(args, task.Uid)

	sql.WriteString(" AND status=?")

	args = append(args, CouponReceiveStatusNone)

	sql.WriteString(" AND usemincount <= ? AND (usemaxcount = -1 OR usemaxcount >= ?)")

	args = append(args, task.Count, task.Count)

	sql.WriteString(" AND useminvalue <= ? AND (usemaxvalue = -1 OR usemaxvalue >= ?)")

	args = append(args, task.Value, task.Value)

	var pageIndex = task.PageIndex
	var pageSize = task.PageSize

	if pageIndex < 1 {
		pageIndex = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	if task.Counter {
		var counter = CouponUseQueryCounter{}
		counter.PageIndex = pageIndex
		counter.PageSize = pageSize
		counter.RowCount, err = kk.DBQueryCount(db, a.GetCouponReceiveTable(), a.GetPrefix(), sql.String(), args...)
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

	sql.WriteString(" ORDER BY offer DESC,id ASC")

	sql.WriteString(fmt.Sprintf(" LIMIT %d,%d", (pageIndex-1)*pageSize, pageSize))

	var v = CouponUse{}
	var scanner = kk.NewDBScaner(&v)

	rows, err := db.Query(fmt.Sprintf("SELECT *,(CASE type WHEN 0 THEN value ELSE FLOOR((100 - rebate) * %d * 0.01) END) as offer FROM %s%s %s", task.Value, a.GetPrefix(), a.GetCouponReceiveTable().Name, sql.String()), args...)

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
