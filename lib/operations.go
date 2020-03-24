package lib

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"sync"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type concurrentAuthInfo struct {
	sync.RWMutex
	authorizationInfo []AuthorizationJSON
}

var cai = concurrentAuthInfo{}

//init function reads configuration from file and stores in the concurrentAuthInfo
func init() {
	byteconf, readerr := ioutil.ReadFile("./configuration/authorization/auth.conf")
	if readerr != nil {
		panic(readerr)
	}
	var downloadconf DownloadConfigurationJSON
	jsonbinderr := json.Unmarshal(byteconf, &downloadconf)
	if jsonbinderr != nil {
		panic(jsonbinderr)
	}
	cai.Lock()
	cai.authorizationInfo = downloadconf.Authorization.Users
	cai.Unlock()
}

//OpShowUsers function returns current authorization configuration
func OpShowUsers() ([]AddDeleteTopicJSON, error) {
	cai.RLock()
	temp := cai.authorizationInfo
	cai.RUnlock()
	tempconf, jsonbinderr := json.Marshal(temp)
	if jsonbinderr != nil {
		return nil, jsonbinderr
	}
	var adddeletetopic = []AddDeleteTopicJSON{}
	jsonunbinderr := json.Unmarshal(tempconf, &adddeletetopic)
	if jsonunbinderr != nil {
		return nil, jsonunbinderr
	}
	return adddeletetopic, nil
}

//OpAddUser function puts new user to authorization configuration
func OpAddUser(user, password string) error {
	cai.RLock()
	for _, tempinfo := range cai.authorizationInfo {
		if tempinfo.User == user {
			cai.RUnlock()
			return errors.New("User:" + user + " already present")
		}
	}
	cai.RUnlock()
	addinfo := AuthorizationJSON{User: user, Password: password}
	cai.Lock()
	cai.authorizationInfo = append(cai.authorizationInfo, addinfo)
	cai.Unlock()
	return nil
}

//OpDeleteUser function deletes the user from authorization configuration
func OpDeleteUser(user string) error {
	cai.Lock()
	for index, tempinfo := range cai.authorizationInfo {
		if tempinfo.User == user {
			cai.authorizationInfo = append(cai.authorizationInfo[:index], cai.authorizationInfo[index+1:]...)
			cai.Unlock()
			return nil
		}
	}
	cai.Unlock()
	return errors.New("User:" + user + " cannot be deleted")
}

//OpShowTopics function returns all the topics present in the authorization configuration
func OpShowTopics() map[string]int {
	topicmap := make(map[string]int)
	cai.RLock()
	temp := cai.authorizationInfo
	cai.RUnlock()
	for _, tempinfo := range temp {
		for _, topicpublish := range tempinfo.Permissions.Publish {
			pvalue, _ := topicmap[topicpublish]
			topicmap[topicpublish] = pvalue + 1
		}
		for _, topicsubscribe := range tempinfo.Permissions.Subscribe {
			svalue, _ := topicmap[topicsubscribe]
			topicmap[topicsubscribe] = svalue + 1
		}
	}
	return topicmap
}

//OpAddTopic function adds topics to the particular user in the authorization configuration
//If any of the topics are present in the request JSON are available in the authorization configuration for that particular user, this functions returns false
func OpAddTopic(info AddDeleteTopicJSON) error {
	cai.Lock()
	for indexuser, authinfo := range cai.authorizationInfo {
		if authinfo.User == info.User {
			for _, infotopicpublish := range info.Permissions.Publish {
				for _, authtopicpublish := range authinfo.Permissions.Publish {
					if infotopicpublish == authtopicpublish {
						cai.Unlock()
						return errors.New("" + infotopicpublish + " topic is already present for the user:" + info.User)
					}
				}
			}
			for _, infotopicsubscribe := range info.Permissions.Subscribe {
				for _, authtopicsubscribe := range authinfo.Permissions.Subscribe {
					if infotopicsubscribe == authtopicsubscribe {
						cai.Unlock()
						return errors.New("" + infotopicsubscribe + " topic is already present for the user:" + info.User)
					}
				}
			}
			temppublish := append(authinfo.Permissions.Publish, info.Permissions.Publish...)
			tempsubscribe := append(authinfo.Permissions.Subscribe, info.Permissions.Subscribe...)
			cai.authorizationInfo[indexuser].Permissions.Publish = temppublish
			cai.authorizationInfo[indexuser].Permissions.Subscribe = tempsubscribe
			cai.Unlock()
			return nil
		}
	}
	cai.Unlock()
	return errors.New("Cannot add topics for the user:" + info.User)
}

//OpDeleteTopic function deletes topics to the particular user in the authorization configuration
//If any of the topics are present in the request JSON are not available in the authorization configuration for that particular user, this functions returns false
func OpDeleteTopic(info AddDeleteTopicJSON) error {
	cai.Lock()
	for indexuser, authinfo := range cai.authorizationInfo {
		if authinfo.User == info.User {
			countpublish := 0
			countsubscribe := 0
			for _, infopublish := range info.Permissions.Publish {
				for _, authpublish := range authinfo.Permissions.Publish {
					if infopublish == authpublish {
						countpublish = countpublish + 1
					}
				}
			}
			for _, infosubscribe := range info.Permissions.Subscribe {
				for _, authsubscribe := range authinfo.Permissions.Subscribe {
					if infosubscribe == authsubscribe {
						countsubscribe = countsubscribe + 1
					}
				}
			}
			if countpublish == len(info.Permissions.Publish) && countsubscribe == len(info.Permissions.Subscribe) {
				temppublish := authinfo.Permissions.Publish
				tempsubscribe := authinfo.Permissions.Subscribe
				for _, infotopicpublish := range info.Permissions.Publish {
					for i, temptopicpublish := range temppublish {
						if infotopicpublish == temptopicpublish {
							temppublish = append(temppublish[:i], temppublish[i+1:]...)
						}
					}
				}
				for _, infotopicsubscribe := range info.Permissions.Subscribe {
					for j, temptopicsubscribe := range tempsubscribe {
						if infotopicsubscribe == temptopicsubscribe {
							tempsubscribe = append(tempsubscribe[:j], tempsubscribe[j+1:]...)
						}
					}
				}
				cai.authorizationInfo[indexuser].Permissions.Publish = temppublish
				cai.authorizationInfo[indexuser].Permissions.Subscribe = tempsubscribe
				cai.Unlock()
				return nil
			}
		}
	}
	cai.Unlock()
	return errors.New("Cannot delete topics for the user:" + info.User)
}

//OpDownloadConfiguration function stores the authorization configuration to the file and reload the nats server
func OpDownloadConfiguration() (string, error) {
	writeerr := writeConfiguration()
	if writeerr != nil {
		return "", writeerr
	}
	reloaderr := reloadConfiguration()
	if reloaderr != nil {
		return "", reloaderr
	}
	return "Download and reload of Configuration Successful", nil
}

//**************************************************************************************************************************

//writeConfiguration function stores the authorization configuration to the file
func writeConfiguration() error {
	var downloadconf = DownloadConfigurationJSON{}
	cai.RLock()
	downloadconf.Authorization.Users = cai.authorizationInfo
	cai.RUnlock()
	byteconf, jsonbinderr := json.Marshal(downloadconf)
	if jsonbinderr != nil {
		return jsonbinderr
	}
	filewriteerr := ioutil.WriteFile("./configuration/authorization/auth.conf", byteconf, 0777)
	if filewriteerr != nil {
		return filewriteerr
	}
	time.Sleep(1 * time.Second)
	return nil
}

//reloadConfiguration function reloads the nats server
func reloadConfiguration() error {
	args := []string{"--signal", "reload"}
	cmd := exec.Command("./nats-server", args...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
