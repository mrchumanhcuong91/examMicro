package helper
import(
	"fmt"
)

func NewToken(params map[string]interface{})(string,string,string, error){
	secretKey := NewRandomKey()
	pubKey := NewRandomKey()
	//encrypty key
	secretEnCrypt, err := Encrypt(secretKey)
	if err != nil{
		fmt.Printf("error encrypt key %v",err)
		return "","","",err
	}
	
	//make token
	claims := NewClaims(params)

	token, err := EncodeJWT(string(secretKey[:]), claims)

	return token,string(pubKey[:]),string(secretEnCrypt[:]),err

}