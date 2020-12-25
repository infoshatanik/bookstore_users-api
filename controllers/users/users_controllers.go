package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/infoshatanik/bookstore_users-api/domain/users"
	//"io/ioutil"
	//"fmt"
	//"encoding/json"
	"github.com/infoshatanik/bookstore_users-api/services"
	"github.com/infoshatanik/bookstore_users-api/utils/errors"
	"strconv"
)

func CreateUser(c *gin.Context){
	var user users.User 
	if err := c.ShouldBindJSON(&user); err != nil {
		// restErr := errors.RestErr{
		// 	Message: "Invalid json body",
		// 	Status: http.StatusBadRequest,
		// 	Error: "bad_request",
		// }
		restErr := errors.NewBadRequstError("invalid json body")
		c.JSON(restErr.Status, restErr)
		//fmt.Println(err)
		//TODO : return bad request to the caller
		return
	}
	//fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//TODO : Handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes,&user); err != nil {
	// 	fmt.Println(err.Error())
	// 	//TODO : Handle json error
	// 	return
	// }
	// fmt.Println(string(bytes))
	// fmt.Println(err)
	result, saveErr := services.CreateUser(user)
	if saveErr !=nil {
		// TODO : handle user creation error
		c.JSON(saveErr.Status, saveErr)
		//fmt.Println(saveErr.Error())
		return
	}
	//fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil {
		err := errors.NewBadRequstError("user id should be a number")
		c.JSON(err.Status,err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
	//c.String(http.StatusNotImplemented,"implement me")
}

func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"implement meeeeeeeeex")
}