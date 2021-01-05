package server
import ("github.com/gin-gonic/gin"
		// dao "example/server/user_service/daos"
		client "github.com/micro/micro/v3/service/client"
		"log"
		)
//declare struct apigateway
type ApiGateWayServer struct{
	//web service
	* gin.Engine
	//authorization service
	Client client.Client

}
//constructor
func NewApiGateWayServer(mysql string) (api *ApiGateWayServer){
	log.Println("call NewApiGateWayServer %v",mysql)
	api = &ApiGateWayServer{
		Engine: gin.Default(),}
	api.createAuthenApi()
	return
}
func (api *ApiGateWayServer) createAuthenApi(){
	v1:=api.Group("/gateway/api")
	v1.POST("/register",api.CreateUserCtrl)
    /*
        192.168.0.105:9091/gateway/api/register
        {
        "Name": "Manh CHu",
        "Age": 27,
        "CardNumber":"2461989"
        }
    */
	v1.GET("/users",api.GetUserCtrl)//192.168.0.105:9091/gateway/api/users/123584
	v1.GET("/users/:id",api.GetUserById)//192.168.0.105:9091/gateway/api/users/123584
    v1.PUT("/delete/:id",api.DeleteById)//192.168.0.105:9091/gateway/api/delete/888888
}
