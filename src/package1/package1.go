package package1

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Site struct {
	ID         int64
	Sitename   string `gorm:"column:sitename"`
	Uri        string `gorm:"column:uri"`
	Createtime string `gorm:"column:create_time"`
	VisitNum   int64  `gorm:"column:visit_num"`
}

type Food struct {
	Id         int  // 表字段名为：id
	Name       string // 表字段名为：name
	Price      float64 // 表字段名为：price
	TypeId     int  // 表字段名为：type_id
	CreateTime int64 `gorm:"column:createtime"`  // 表字段名为：createtime
}

func (v Food) TableName() string {
	return "food"
}

func testGormCreate() {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	Dbname := "orangen"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	s1 := Site{
		Sitename:   "vscode指南",
		Uri:        "https://geek-docs.com/vscode",
		Createtime: time.Now().Format("2006-01-02 15:04:05"),
		VisitNum:   7,
	}
	if err := db.Create(&s1).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
}

func testGormSelect() {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	Dbname := "orangen"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	s2 := Site{}
	result := db.Where("visit_num = ?", 7).First(&s2)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("找不到记录")
		return
	}
	fmt.Println(s2)
	fmt.Printf("type s2.Createtime %T\n", s2.Createtime)
}

func ConnectDB() *gorm.DB {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	Dbname := "orangen"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	return db
}

func testSelectSites(db *gorm.DB) []Site {
	var sites []Site
	db.Find(&sites)
	for _, a := range sites {
		fmt.Printf("| %d | %s | %s | %d \n", a.ID, a.Sitename, a.Createtime, a.VisitNum)
	}
	return sites
}

func main() {
	// testGormCreate()
	testGormSelect()

	db := ConnectDB()
	testSelectSites(db)
}

func GetAllSites() []Site {
	db := ConnectDB()
	return testSelectSites(db)
}
