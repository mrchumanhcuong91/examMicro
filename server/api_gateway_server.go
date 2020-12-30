package server
import ("github.com/gin-gonic/gin"
		dao "example/server/user_service/daos"
		client "github.com/micro/micro/v3/service/client"
		"log"
		)
//declare struct apigateway
type ApiGateWayServer struct{
	//web service
	* gin.Engine
	//authorization service
	uDao *dao.UserDao
	Client client.Client

}
//constructor
func NewApiGateWayServer(mysql string) (api *ApiGateWayServer){
	log.Println("call NewApiGateWayServer %v",mysql)
	api = &ApiGateWayServer{
		Engine: gin.Default()}
	api.createAuthenApi()
	return
}
func (api *ApiGateWayServer) createAuthenApi(){
	v1:=api.Group("/gateway/api")
	v1.POST("/register",api.CreateUserCtrl)
}