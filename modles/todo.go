package modles

import (
	"bubble/dao"
)

//TODO模型
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

//Todo 增删改查
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error//只有一个返回值，不用判断error
	return
}

func GetAllTodo( ) (todoList []*Todo, err error) {
	if err := dao.DB.Find(&todoList).Error; err != nil{
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}