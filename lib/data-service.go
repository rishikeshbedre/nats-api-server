package lib

import ()

//PermissionJSON stores permissions for each user in authorization configuration
type PermissionJSON struct {
	Publish []string `json:"publish"`
	Subscribe []string `json:"subscribe"`
}

//AuthorizationJSON stores info of each user in authorization configuration
type AuthorizationJSON struct {
	User string `json:"user"`
	Password string `json:"password"`
	Permissions PermissionJSON `json:"permissions"`
}

//PreAuthorizationJSON structure to store AuthorizationJSON
type PreAuthorizationJSON struct {
	Users []AuthorizationJSON `json:"users"`
}

//DownloadConfigurationJSON stores JSON that as to be downloaded to configuration file
type DownloadConfigurationJSON struct {
	Authorization PreAuthorizationJSON `json:"authorization"`
}

//AddUserJSON maps and validate adding a user in authorization configuration
type AddUserJSON struct {
	User string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//AddDeleteTopicJSON maps and validate adding and deleting topics in authorization configuration
type AddDeleteTopicJSON struct {
	User string `json:"user" binding:"required"`
	Permissions PermissionJSON `json:"permissions" binding:"required"`
}

//DeleteUserJSON maps and validate deleting a user in authorization configuration
type DeleteUserJSON struct {
	User string `json:"user" binding:"required"`
}