package model

type VisitRecord struct {
  Time int64
  Path string
  CostTime int
}

func (table *VisitRecord) TableName() string {
  return "visit_record"
}
