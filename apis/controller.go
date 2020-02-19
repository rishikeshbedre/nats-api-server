package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishikeshbedre/nats-api-server/lib"
	"github.com/rishikeshbedre/nats-api-server/util"
)

//ShowUsers function returns the current authorization configuration
func ShowUsers(c *gin.Context) {
	result, showuserserr := lib.OpShowUsers()
	if showuserserr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": showuserserr})
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}

//AddUser function adds new user to the authorization configuration
func AddUser(c *gin.Context) {
	var adduser lib.AddUserJSON
	if jsonbinderr := c.ShouldBindJSON(&adduser); jsonbinderr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": jsonbinderr.Error()})
		return
	}
	passwd, passhasherr := util.GenHashPassword(adduser.Password)
	if passhasherr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": passhasherr})
	} else {
		addusererr := lib.OpAddUser(adduser.User, passwd)
		if addusererr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": addusererr.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "User:" + adduser.User + " added"})
		}
	}
}

//DeleteUser function deletes the user from authorization configuration
func DeleteUser(c *gin.Context) {
	var deleteuser lib.DeleteUserJSON
	if jsonbinderr := c.ShouldBindJSON(&deleteuser); jsonbinderr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": jsonbinderr.Error()})
		return
	}
	deleteusererr := lib.OpDeleteUser(deleteuser.User)
	if deleteusererr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": deleteusererr.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User:" + deleteuser.User + " deleted"})
	}
}

//AddTopic function adds the topics to the particular user in authorization configuration
//If any of the topics are present in the request JSON are available in the authorization configuration for that particular user, this functions returns a error message
func AddTopic(c *gin.Context) {
	var addtopic lib.AddDeleteTopicJSON
	if jsonbinderr := c.ShouldBindJSON(&addtopic); jsonbinderr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": jsonbinderr.Error()})
		return
	}
	addtopicerr := lib.OpAddTopic(addtopic)
	if addtopicerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": addtopicerr.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Topics Added for the user:" + addtopic.User})
	}
}

//DeleteTopic function deletes the topics from the particular user in authorization configuration
//If any of the topics are present in the request JSON are not available in the authorization configuration for that particular user, this functions returns a error message
func DeleteTopic(c *gin.Context) {
	var deletetopic lib.AddDeleteTopicJSON
	if jsonbinderr := c.ShouldBindJSON(&deletetopic); jsonbinderr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": jsonbinderr.Error()})
		return
	}
	deletetopicerr := lib.OpDeleteTopic(deletetopic)
	if deletetopicerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": deletetopicerr.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Topics deleted for the user:" + deletetopic.User})
	}
}

//DownloadConfiguration function stores the authorization configuration to the file and reload the nats server
func DownloadConfiguration(c *gin.Context) {
	result, err := lib.OpDownloadConfiguration()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": result})
	}
}
