package databases
import (
	// "authservice/db/models"
	// "github.com/jinzhu/gorm"
	"fmt"
	"database/sql"
    _ "github.com/lib/pq"
)
type MyPostGres struct{
	Uri string
}

func NewAuthenDb(host,dbName,pg_username, pg_password string) (*MyPostGres){
	//postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",host, "cuongcm", "authenDb", "123456")
	db, err := sql.Open("postgres", dbUri)
	if err != nil{
		fmt.Printf("open db failed  %v",err)
		return nil
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("create sql connection failed %v",err)
	}

	return &MyPostGres{Uri : dbUri}

}
func createConnection(uri string) (*sql.DB){
	db, err := sql.Open("postgres", uri)
	if err != nil{
		fmt.Printf("create sql connection failed %v",err)
		return nil
	}
	return db
}
func (pg* MyPostGres) InsertKey(id string, value string) (error){
	fmt.Printf("InsertKey sql %v",pg.Uri)	
	db := createConnection(pg.Uri)
	defer db.Close()
	//insert key to db
	sqlStr := "INSERT INTO keys(pubkey,encryptkey) VALUES($1,$2)"
	db.Exec(sqlStr, id, value)
	return nil
}
func (pg* MyPostGres) GetKey(id string) (string, error){
	db := createConnection(pg.Uri)
	defer db.Close()
	var key string
	err := db.QueryRow("select encryptkey from keys where pubkey=$1",id).Scan(&key)
	if err != nil{
		fmt.Printf("GetKey error  %v",err)
		return "",nil
	}
	return key,err
}
func (pg* MyPostGres) DeleteKey(id string) error{
	db := createConnection(pg.Uri)
	defer db.Close()
	_,err := db.Exec("DELETE from keys where pubkey=$1",id)
	if err != nil{
		fmt.Printf("DeleteKey error  %v",err)
		return nil
	}
	return err
}
