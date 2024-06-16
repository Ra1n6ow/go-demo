package one

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ra1n6ow/go-demo/SQL/gorm/model"
	"gorm.io/gorm"
)

func InsertOne(db *gorm.DB) {
	s := &model.Student{
		Age:      18,
		Gender:   0,
		Birthday: time.Now(),
	}

	// INSERT INTO `student` (`name`,`age`,`gender`,`email`,`birthday`) VALUES ('',18,0,NULL,'2024-06-15 07:03:06.816') RETURNING `id`
	// 这就是空字符串和空指针的区别
	err := db.Create(&s).Error
	if err != nil {
		fmt.Println(err)
	}
}

func InsertMany(db *gorm.DB) {
	var studentList []model.Student
	for i := 0; i < 5; i++ {
		studentList = append(studentList, model.Student{
			Name:     fmt.Sprintf("du%d", i),
			Age:      18,
			Gender:   0,
			Birthday: time.Now(),
		})
	}
	err := db.Create(&studentList).Error
	fmt.Println(err)
	fmt.Println(studentList)
}

func QueryOne(db *gorm.DB) {
	var s model.Student
	db.Take(&s) // SELECT * FROM `student` LIMIT 1
	fmt.Println("take:", s)

	s = model.Student{}
	db.First(&s) // SELECT * FROM `student` ORDER BY `student`.`id` LIMIT 1
	fmt.Println("first:", s)

	s = model.Student{}
	db.Last(&s) // SELECT * FROM `student` ORDER BY `student`.`id` DESC LIMIT 1
	fmt.Println("last:", s)

	s = model.Student{}
	// 使用 ? 拼接能防止 sql 注入, 原理就是将参数转义
	db.Take(&s, "name = ?", "du") // SELECT * FROM `student` WHERE name = 'du' LIMIT 1
	fmt.Println("take:", s)

	s = model.Student{ID: 2}
	// s = model.Student{Name: "du"} , 无效，只能使用主键
	db.Take(&s) // SELECT * FROM `student` WHERE `student`.`id` = 2 LIMIT 1
	fmt.Println("take:", s)

	//err: not found
	s = model.Student{}
	err := db.Take(&s, "name = ?", "xx").Error
	switch err {
	case gorm.ErrRecordNotFound:
		fmt.Println("record not found")
	default:
		fmt.Println("other error")
	}

	// 查询所有记录，并json 序列化
	var sl []model.Student
	rows := db.Find(&sl).RowsAffected // SELECT * FROM `student`
	for _, student := range sl {
		fmt.Println(student)
	}
	fmt.Println(rows)
	data, _ := json.Marshal(sl)
	fmt.Println(string(data))

	fmt.Println("--------------------------------------------------")

	// 根据主键列表查询
	// rows = db.Find(&sl, []int{2, 3, 4}).RowsAffected // SELECT * FROM `student` WHERE `student`.`id` IN (2,3,4)
	rows = db.Find(&sl, "name in ?", []string{"du", "du0"}).RowsAffected // SELECT * FROM `student` WHERE name in ('du','du0')
	fmt.Println(rows)
}

func UpdateOne(db *gorm.DB) {
	var s model.Student

	// Save 保存所有字段
	// 先找
	db.Take(&s, "id=?", 3)
	fmt.Println(s)
	// 再更新, 全字段更新
	s.Name = "linda"
	db.Save(&s) // UPDATE `student` SET `name`='linda',`age`=18,`gender`=0,`email`=NULL,`birthday`='2024-06-15 07:03:06' WHERE `id` = 3
	fmt.Println(s)

	// Select 更新指定字段
	s = model.Student{}
	db.Take(&s, "id=?", 4)
	s.Name = "linda2"
	// 指定 mysql 的 name 字段进行更新
	db.Select("name").Save(&s) // UPDATE `student` SET `name`='linda2' WHERE `id` = 4

	// Update 批量更新
	var sl []model.Student
	db.Model(&sl).Where("id in ?", []int{3, 4, 5}).Update("gender", 1) // UPDATE `student` SET `gender`=1 WHERE id in (3,4,5)

	fmt.Println("--------------------------------------------------")

	// Updates 批量更新
	sl = []model.Student{}
	db.Table("student").Where("id in ?", []int{3, 4, 5}).Updates(map[string]interface{}{"age": "20", "gender": 0}) // UPDATE `student` SET `age`='20',`gender`=0 WHERE id in (3,4,5)
	db.Model(&sl).Where("id in ?", []int{3, 4, 5}).Updates(model.Student{Age: 20, Gender: 0})                      // 同上，但是更新0值会被忽略, UPDATE `student` SET `age`='20' WHERE id in (3,4,5)
}

func DeleteOne(db *gorm.DB) {
	var s model.Student
	db.Delete(&s, "id=?", 2) // DELETE FROM `student` WHERE id=2
	fmt.Println(s)           // 0值

	db.Take(&s, "id=?", 3)
	db.Delete(&s)  // DELETE FROM `student` WHERE `student`.`id` = 3
	fmt.Println(s) // 保存删除的对象

	// 批量删除
	db.Delete(&s, []int{4, 5})       // DELETE FROM `student` WHERE `student`.`id` IN (4,5)
	db.Where("id > ?", 6).Delete(&s) // DELETE FROM `student` WHERE id > 6
}
