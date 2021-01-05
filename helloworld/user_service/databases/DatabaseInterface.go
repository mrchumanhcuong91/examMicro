package databases
import(
	m "helloworld/user_service/models"
)

type Adapter interface{
	InsertUser(u *m.User) error
	UpdateUser(u *m.User, id string) error
	GetUser(id string) (m.User, error)
	GetUsers() ([]m.User, error)

	FindUser(id string) (m.User, error)
	DeleteUser(id string) error
}