package daos
import (
	"log"
	d "helloworld/user_service/databases"
	m "helloworld/user_service/models"
	logs "github.com/micro/micro/v3/service/logger"
)


type UserDao struct {
	a d.Adapter
}
func InitUserDao(config string)(dao * UserDao) {
	log.Printf("init user dao %v",config)
	instanse := d.NewMysql(config)
	return &UserDao{a :  instanse}
}

func(u* UserDao) AddUser(user* m.User) error{
	//call global database
	logs.Info("Received UserDao.AddUser request")
	return u.a.InsertUser(user)
}
func(u* UserDao) FindUser(id string) (m.User,error){
	return u.a.FindUser(id)
}
func(u* UserDao) UpdateUser(user* m.User) error{
	return nil
}
func(u* UserDao) DeleteUser(id string) error{
    return u.a.DeleteUser(id)
}
func(u* UserDao) GetUser(id string) (m.User, error){
	return m.User{}, nil
}
func(u* UserDao) GetUsers() ([]m.User, error){
	return u.a.GetUsers()
}
func(u* UserDao) GetUserByName(name string) (m.User, error){
	return u.a.GetUserByName(name)
}
