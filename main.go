package main

import (
	"github.com/hidenari-yuda/todo_app/app/controllers"
)

func main() {

	//user, _ := models.GetUser()

	/*user.CreateTodo("test deadline", "2019-01-01")
	getTodos, _ := models.GetTodos()
	fmt.Println(getTodos)*/

	/*u := &models.Chat{}
	u.Content = "こんにちは"
	u.UserID = 3
	u.CreatedAt = time.Now()
	u.GroupID = 2

	u.CreateChat()*/

	/*user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	session, _ := user.CreateSession()

	fmt.Println(session)

	valid, _ := session.CheckSession()
	fmt.Println(valid)*/

	/*fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)*/
	/*user, _ := models.GetUser(2)

	user.CreateTodo("First Todo")*/

	/*user, _ := models.GetUser(3)

	user.CreateTodo("Third Todo")

	/*todos, _ := models.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	}*/

	/*user2, _ := models.GetUser(2)
	todos, _ := user2.()
	for _, v := range todos {
		fmt.Println(v)
	}*/

	/*t, _ := models.GetTodo(3)
	t.DeleteTodo()*/

	controllers.StartMainServer()

}
