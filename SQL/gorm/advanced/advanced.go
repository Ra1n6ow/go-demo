package advanced

import (
	"fmt"

	"github.com/ra1n6ow/go-demo/SQL/gorm/model"
	"gorm.io/gorm"
)

func string2Ptr(email string) *string {
	return &email
}

type User struct {
	Name string
	Age  int
}

func InitData(db *gorm.DB) {
	var sl []model.Student
	db.Find(&sl).Delete(&sl)
	sl = []model.Student{
		{
			Email:  string2Ptr("duyi@qq.com"),
			ID:     1,
			Name:   "杜一",
			Age:    18,
			Gender: 1,
		},
		{
			Email:  string2Ptr("zhouer@qq.com"),
			ID:     2,
			Name:   "周二",
			Age:    19,
			Gender: 0,
		},
		{
			Email:  string2Ptr("zhangsan@qq.com"),
			ID:     3,
			Name:   "张三",
			Age:    18,
			Gender: 1,
		},
		{
			Email:  string2Ptr("lisi@qq.com"),
			ID:     4,
			Name:   "李四",
			Age:    19,
			Gender: 0,
		},
		{
			Email:  string2Ptr("wangwu@qq.com"),
			ID:     5,
			Name:   "王五",
			Age:    20,
			Gender: 1,
		},
		{
			Email:  string2Ptr("zhaoliu@qq.com"),
			ID:     6,
			Name:   "赵六",
			Age:    21,
			Gender: 1,
		},
		{
			Email:  string2Ptr("liuqi@qq.com"),
			ID:     7,
			Name:   "刘七",
			Age:    19,
			Gender: 1,
		},
		{
			Email:  string2Ptr("yuanba@qq.com"),
			ID:     8,
			Name:   "袁八",
			Age:    20,
			Gender: 0,
		},
		{
			Email:  string2Ptr("huangjiu@qq.com"),
			ID:     9,
			Name:   "黄九",
			Age:    21,
			Gender: 0,
		},
	}
	db.Create(&sl)
}

func Query(db *gorm.DB) {
	var sl []model.Student
	db.Where("name = ?", "杜一").Find(&sl) // SELECT * FROM `student` WHERE name = '杜一'

	db.Where("name <> ?", "杜一").Find(&sl) // SELECT * FROM `student` WHERE name <> '杜一'

	db.Where("name in ?", []string{"杜一", "周二"}).Find(&sl) // SELECT * FROM `student` WHERE name in ('杜一','周二')

	db.Where("name like ?", "杜%").Find(&sl) // SELECT * FROM `student` WHERE name like '杜%'

	db.Where("age > ? and email like ?", "20", "%qq.com").Find(&sl) // SELECT * FROM `student` WHERE age > '20' and email like '%qq.com
	fmt.Println(sl)

	// 结构体查询会过滤0值的字段，意味着 Age 为 0 时，不会加入到查询条件中
	db.Where(&model.Student{Name: "杜一", Age: 0}).Find(&sl) // SELECT * FROM `student` WHERE `student`.`name` = '杜一'

	// map 查询不会过滤0值
	db.Where(map[string]any{"name": "杜一", "age": "0"}).Find(&sl) // SELECT * FROM `student` WHERE `age` = '0' AND `name` = '杜一'

	db.Not("age > ?", 20).Find(&sl) // SELECT * FROM `student` WHERE NOT age > 20

	db.Or("name = ?", "杜一").Or("age > ?", 20).Find(&sl) // SELECT * FROM `student` WHERE name = '杜一' OR age > 20
	fmt.Println(sl)
}

// select 选择字段
func Select(db *gorm.DB) {
	var sl []model.Student
	db.Select("name", "age").Find(&sl) // SELECT `name`,`age` FROM `student`
	fmt.Println(sl)                    // 非选择字段为零值

	// scan, 将选择的字段存入另一个结构体中，可以避免非选择字段为0值
	sl = []model.Student{}
	var ul []User

	// db.Select("name", "age").Find(&sl).Scan(&ul)  // 这样会查询两次
	// db.Model(&model.Student{}).Select("name", "age").Scan(&ul) // 这样只会查询一次
	db.Table("student").Select("name", "age").Scan(&ul) // 同上
	fmt.Println(ul)
}

func Other(db *gorm.DB) {
	var sl []model.Student

	// 排序
	db.Order("age desc").Find(&sl)
	fmt.Println(sl)

	// 分页
	// 一页多少条
	limit := 2
	// 第几页
	page := 2
	offset := (page - 1) * limit
	db.Limit(limit).Offset(offset).Find(&sl) // SELECT * FROM `student` LIMIT 2 OFFSET 2
	fmt.Println(sl)

	// 去重
	var ageList []int
	// db.Table("student").Select("age").Distinct("age").Scan(&ageList)
	// 或者
	db.Table("student").Select("distinct age").Scan(&ageList) // SELECT distinct age FROM `student`
	fmt.Println(ageList)

	// 分组
}
