package model

type VisitRecord struct {
  Method string  `gorm:"column=method"`     
  Time int64  `gorm:"column=time"`
  Path string  `gorm:"column=path"`
  CostTime int  `gorm:"column=cost_time"`
  IP string  `gorm:"column=ip"`
}

func (table *VisitRecord) TableName() string {
  return "visit_record"
}
