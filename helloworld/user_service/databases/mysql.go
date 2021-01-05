package databases
import(
	"log"
	m "helloworld/user_service/models"
	_ "github.com/go-sql-driver/mysql"
	sql "database/sql"
)
type MySqlAdapter struct{
	db     *sql.DB
}

func NewMysql(sqlPath string) (*MySqlAdapter){
	DB, err := sql.Open("mysql",sqlPath)
	log.Printf("error %v",err)
	log.Printf("error DB  %p",DB)
	_, err = DB.Exec("use microservice")
	log.Printf("error use database %v",err)

	_, err = DB.Exec("create table users(id int auto_increment primary key,cardnumber varchar(32),Name varchar(64),Age int)")
	log.Printf("error create tables %v",err)

	return &MySqlAdapter{db: DB}
}
func (a *MySqlAdapter)InsertUser(u *m.User) error{
	log.Printf("InsertUser mysql %v",u.Name)
	_, err:= a.db.Exec("insert into users(cardnumber, name, age) values(?,?,?)",u.CardNumber,u.Name, u.Age)
	log.Printf("InsertUser error  %v",err)
	return nil
}
func (a *MySqlAdapter)UpdateUser(u *m.User, id string) error{
	a.db.Exec("update users set name=?,age=? where id=?", u.Name, u.Age, id)
	return nil
}
func (a *MySqlAdapter)GetUsers() ([]m.User, error){
	rows, err := a.db.Query("select * from users")
	log.Printf("Get users err %v",err)

	if err == nil{
		results :=  make([]m.User,0)
		for rows.Next(){
			item := m.User{}
			var ii int
			if err := rows.Scan(&ii, &item.CardNumber,&item.Name,&item.Age); err != nil{
				log.Printf("Scan error %v",err)
				continue
			}
			results = append(results, item)

		}
		return results, nil
	}

	return nil, nil
}
func (a *MySqlAdapter)GetUser(id string) (m.User, error){
	return m.User{}, nil
}
func (a *MySqlAdapter)FindUser(id string) (m.User,error){
	rows := a.db.QueryRow("select * from users where cardnumber=?",id)
	result := m.User{}
    var ii int
	if err := rows.Scan(&ii, &result.CardNumber,&result.Name,&result.Age); err != nil{
        log.Printf("FindUser Scan error %v",err)
		return result, err
	}
	return result, nil
}
func (a *MySqlAdapter)DeleteUser(id string) error{
	a.db.Exec("delete from users where cardnumber=?",id)
	return nil
}
