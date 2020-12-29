package models

import "BcConnectWeb_A/dbmysql"

type SmsRecord struct {
	BizId string
	Phone string
	Code string
	Status string
	Message string
	TimeStamp int64
}

/**
查询登录的验证码和手机号是否正确
 */
func QuerySmsRecord(bizId string,phone string,code string) (*SmsRecord,error) {
	var sms SmsRecord
	row := dbmysql.Db.QueryRow("select biz_id,timestamp from sms_record where biz_id = ? and phone = ? and code = ?",bizId,phone,code)
	err := row.Scan(&sms.BizId,&sms.TimeStamp)
	if err != nil {
		return nil,err
	}
	return &sms,nil
}

/**
将验证记录存到数据库中
 */
func (s SmsRecord) SaveSmsRecord() (int64,error) {
	rs,err := dbmysql.Db.Exec("insert into sms_record(biz_id,phone,code,status,message,timestamp)" +
		"value(?,?,?,?,?,?",s.BizId,s.Phone,s.Status,s.Code,s.Message,s.TimeStamp)
	if err != nil {
		return -1,err
	}
	return rs.RowsAffected()
}