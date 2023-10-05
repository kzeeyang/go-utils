package util

import (
	"fmt"
	"syncCompanyInfo/common"
	"syncCompanyInfo/models"
	"time"

	//"github.com/astaxie/beego"

	"github.com/dgrijalva/jwt-go"
	"github.com/xiya-team/helpers"
)

func CreateToken(user models.User) string {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(72)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = user.Id
	claims["verification"] = helpers.Md5(user.UserName)
	claims["user_name"] = user.UserName
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//token.Claims=claims
	tokenString, _ := token.SignedString([]byte("key"))

	return tokenString
}

func CheckToken(tokenString string) (b bool, t string, code int) {
	//kv := strings.Split(tokenString, " ")
	//if len(kv) != 2 || kv[0] != "Bearer" {
	//	fmt.Println("AuthString invalid:", tokenString)
	//	return false, nil
	//}

	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte("key"), nil
	})

	fmt.Println(err)

	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError: // something was wrong during the validation
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				//ctx.Output.SetStatus(401)
				//resBody, err := json.Marshal(controllers.OutResponse(401, nil, "登录已过期，请重新登录"))
				//ctx.Output.Body(resBody)
				//if err != nil {
				//	panic(err)
				//}
				return false, "登录已过期，请重新登录", 50014
			default:
				//ctx.Output.SetStatus(401)
				//resBytes, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求，请重新登录"))
				//ctx.Output.Body(resBytes)
				//if err != nil {
				//	panic(err)
				//}
				return false, "非法请求，请重新登录", 50008
			}
		default: // something else went wrong
			//ctx.Output.SetStatus(401)
			//resBytes, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求，请重新登录"))
			//ctx.Output.Body(resBytes)
			//if err != nil {
			//	panic(err)
			//}
			return false, "非法请求，请重新登录", 50008
		}

		fmt.Println("转换为jwt claims失败.", err)
		return false, "非法请求，请重新登录", 50008
	}

	if !token.Valid {
		// but may still be invalid
		//ctx.Output.SetStatus(401)
		//resBytes, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求，请重新登录"))
		//ctx.Output.Body(resBytes)
		//if err != nil {
		//	panic(err)
		//}
		return false, "非法请求，请重新登录", 50008
	}
	//GetUserNameByToken(kv[1])

	// username := GetUserNameByToken(tokenString)

	// verification := GetVerificationByToken(tokenString)

	// if verification != helpers.Md5(username) {
	// 	return false, "非法请求，请重新登录", 50008
	// }

	userId := GetUserIdByToken(tokenString)
	user_name := GetUserNameByToken(tokenString)
	common.UserName = user_name
	return true, "验证通过", userId
}

func GetUserIdByToken(tokenString string) int {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("key"), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	return int(id)
}

func GetUserNameByToken(tokenString string) string {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("key"), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	user_name := claims["user_name"].(string)
	return user_name
}

func GetVerificationByToken(tokenString string) string {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("key"), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	user_name := claims["verification"].(string)
	return user_name
}
