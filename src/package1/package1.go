package package1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Site struct {
	ID         int64  `gorm:"column:id" json:"id"`
	Sitename   string `gorm:"column:sitename" json:"name"`
	Uri        string `gorm:"column:uri" json:"url"`
	Createtime string `gorm:"column:create_time" json:"create_time"`
	VisitNum   int64  `gorm:"column:visit_num" json:"visited"`
}

type Food struct {
	Id         int     // 表字段名为：id
	Name       string  // 表字段名为：name
	Price      float64 // 表字段名为：price
	TypeId     int     // 表字段名为：type_id
	CreateTime int64   `gorm:"column:createtime"` // 表字段名为：createtime
}

func (v Food) TableName() string {
	return "food"
}

func testGormCreate() {
	username := "root"
	password := "123456"
	host := "10.103.164.161"
	port := 3306
	Dbname := "fruit_buckets"
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
	host := "43.139.176.247"
	port := 3000
	Dbname := "fruit_buckets"
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
	return sites
}

func main() {
	// testGormCreate()
	// testGormSelect()

	// db := ConnectDB()
	// testSelectSites(db)
	fmt.Println("==__T__==")
}

func GetAllSites() []Site {
	db := ConnectDB()
	defer func ()  {
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.Close()
	}()
	return testSelectSites(db)
}

func PostSite(c *gin.Context) {
	// 处理post 增加site的请求
	siteObj := Site{}

	if err := c.ShouldBindJSON(&siteObj); err != nil {
		c.String(http.StatusOK, `json绑定为Site结构体失败`)
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	siteObj.Createtime = now
	siteObj.VisitNum = 9

	db := ConnectDB()
	result := db.Create(&siteObj)

	c.JSON(http.StatusOK, gin.H{
		"Error":        result.Error,
		"RowsAffected": result.RowsAffected,
		"id":           siteObj.ID,
	})
}

type UserInfo struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type HeroInfo struct {
	Name        string
	Age         int
	Attack      int
	Wise        int
	DynamicTags []string
}

func TestBindJSON(c *gin.Context) {
	var u UserInfo
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(200, gin.H{
			"msg": "你错了",
		})
	}
	c.JSON(200, u)
}

func DeleteSite(c *gin.Context) {
	id := c.Query("id")

	db := ConnectDB()
	db.Delete(&Site{}, id)
	c.JSON(200, gin.H{
		"id": id,
	})
}

func IncrVisitNum(id string) int64 {
	site := Site{}
	db := ConnectDB()
	defer func ()  {
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.Close()
	}()
	db.Where("id = ?", id).First(&site)
	fmt.Println(site.VisitNum)
	db.Model(&site).Where("id = ?", id).Update("visit_num", site.VisitNum+1)
	return site.VisitNum + 1
}

func UpdateSite(c *gin.Context) {
	id := c.Query("id")
	c.JSON(200, gin.H{
		"visit_num": IncrVisitNum(id),
	})
}

func GetHero(c *gin.Context) {
	c.JSON(200, gin.H{
		"info": []interface{}{
			HeroInfo{"朱元璋", 34, 76, 86, []string{"和尚", "陈友谅", "岳父🐂"}},
			HeroInfo{"宋江", 45, 65, 82, []string{"黑厮", "及时雨"}},
			HeroInfo{"洪秀全", 32, 89, 56, []string{"农民", "太平天国"}},
			HeroInfo{"项羽", 28, 95, 75, []string{"大力举鼎", "学万人敌"}},
			HeroInfo{"张全蛋", 26, 56, 80, []string{"foxxxx"}},
			HeroInfo{"陈二", 28, 99, 65, []string{"打工仔", "小镇做题家"}},
			HeroInfo{"吴佩孚", 46, 81, 79, []string{"军阀", "奉系", "孙传芳"}},
		},
	})
}

func GetSites(c *gin.Context) {
	sites := GetAllSites()
	c.JSON(http.StatusOK, gin.H{
		"info": sites,
	})
}
