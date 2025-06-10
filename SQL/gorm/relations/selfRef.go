package relations

import (
	"fmt"

	"github.com/ra1n6ow/go-demo/SQL/gorm/model"
	"gorm.io/gorm"
)

func CreateChildren(db *gorm.DB) {

	// menu1 := model.Menu{
	// 	Name: "menu1",
	// }
	// db.Create(&menu1)

	// menu2 := model.Menu{
	// 	Name:     "menu2",
	// 	ParentID: &menu1.ID,
	// }
	// db.Create(&menu2)

	// menu3 := model.Menu{
	// 	Name:     "menu3",
	// 	ParentID: &menu2.ID,
	// }
	// db.Create(&menu3)

	// var menus []model.Menu
	// menus = append(menus, model.Menu{
	// 	Name: "menu1",
	// 	Children: []*model.Menu{
	// 		{
	// 			Name: "menu2",
	// 			Children: []*model.Menu{
	// 				{
	// 					Name: "menu3",
	// 				},
	// 			},
	// 		},
	// 	},
	// })
	// menus = append(menus, model.Menu{
	// 	Name: "menu4",
	// 	Children: []*model.Menu{
	// 		{
	// 			Name: "menu5",
	// 		},
	// 	},
	// })

	// db.Create(&menu)
}

func QueryChildren(db *gorm.DB) {
	var menu model.Menu
	db.Preload("Children").Find(&menu, "parent_id = ?", 1)
	// data, _ := json.Marshal(menu)
	// fmt.Println(string(data))
	fmt.Println(menu)
}

func QueryParent(db *gorm.DB) {
	var menu model.Menu
	db.Find(&menu, "id = ?", 2)
	fmt.Println(menu)
}
