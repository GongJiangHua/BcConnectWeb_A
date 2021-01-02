package models

import (
	"BcConnectWeb_A/dbmysql"
	"fmt"
)

type SmsRecord struct {
	BizId     string `form:"biz_id"`
	Phone     string `form:"phone"`
	Code      string `form:"code"`
	Status    string `form:"status"`
	Message   string `form:"message"`
	TimeStamp int64  `form:"timestamp"`
}

/**
查询登录的验证码和手机号是否正确
*/
func QuerySmsRecord( phone string, code string) (*SmsRecord, error) {
	var sms SmsRecord
	row := dbmysql.Db.QueryRow("select biz_id,timestamp from sms_record where phone = ? and code = ?",  phone, code)

	err := row.Scan(&sms.BizId, &sms.TimeStamp)
	if err != nil {
		return nil, err
	}
	return &sms, nil

	//row, err := dbmysql.Db.Query("select phone,password from user where phone = ? and password = ?",
	//	u.Phone, u.Password)
	//if err != nil {
	//	log.Fatal(err)
	//}
	////defer row.Close()
	//for row.Next() {
	//	err := row.Scan(&u.Phone, &u.Password)
	//	if err != nil {
	//		log.Fatal(err)
	//		return nil,err
	//	}
	//	log.Fatal(u.Phone, u.Password)
	//}
	//return &u, nil
}

/**
将验证记录存到数据库中
*/
func (s SmsRecord) SaveSmsRecord() (int64, error) {
	rs, err := dbmysql.Db.Exec("insert into sms_record(biz_id,phone,code,status,message,timestamp)"+
		"value(?,?,?,?,?,?)", s.BizId, s.Phone, s.Code, s.Status, s.Message, s.TimeStamp)
	if err != nil {
		fmt.Println("验证保存错误：", err.Error())
		return -1, err
	}
	return rs.RowsAffected()
}
