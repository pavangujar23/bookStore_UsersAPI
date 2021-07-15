package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pavangujar23/bookStore_UsersAPI/domain/users"
	"github.com/pavangujar23/bookStore_UsersAPI/services"
	"github.com/pavangujar23/bookStore_UsersAPI/utils/errorsUtil"
)

//every request is handelled by controller
func CreateUser(c *gin.Context) {
	var user users.User

	//these logic should be musttt

	// fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)

	// if err != nil {
	// 	//TODO: Handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//TODO: Handle json error
	// 	fmt.Println(err.Error())
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO: Handle json error
		restErr := errorsUtil.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: Handle userCreation error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errorsUtil.NewBadRequestError("User Id Should be number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		//TODO: Handle userCreation error
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

	// c.String(http.StatusNotImplemented, "implement me...!")
}
