package andoncloud

import "gin_server/lib"

type Bloodpressureinfo struct {
	BPId           int64  `xorm:"'BPId'"`
	GUId           string `xorm:"'GUId'"`
	UserId         int64  `xorm:"'UserId'"`
	APPSDataID     string `xorm:"'APPSDataID'"`
	DataSourceType int64  `xorm:"'DataSourceType'"`
}

func (Bloodpressureinfo) FindOne(BPId int64) (Bloodpressureinfo, error) {
	var bloodpressureinfo Bloodpressureinfo
	engine := lib.GetConn()
	_, err := engine.Where("BPId = ?", BPId).Get(&bloodpressureinfo)
	if err != nil {
		return bloodpressureinfo, err
	}
	return bloodpressureinfo, nil
}
