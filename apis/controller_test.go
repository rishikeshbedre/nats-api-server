package apis

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/rishikeshbedre/nats-api-server/lib"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func TestShowUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	ShowUsers(c)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status, http.StatusOK, rr.Body.String())
	}
}

func TestAddUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//---------------positive case-------------------------------------------------
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	requestBody, _ := json.Marshal(map[string]string{
		"user":"xyz",
		"password":"123",
	})
	c.Request, _ = http.NewRequest("POST","/user", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")
	AddUser(c)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status, http.StatusOK, rr.Body.String())
	}

	//---------------negative case : 1---------------------------------------------------
	rr2 := httptest.NewRecorder()
	c2, _ := gin.CreateTestContext(rr2)
	requestBody2, _ := json.Marshal(map[string]string{
		"user":"xyz",
		"password":"123",
	})
	c2.Request, _ = http.NewRequest("POST","/user", bytes.NewBuffer(requestBody2))
	c2.Request.Header.Set("Content-Type", "application/json")
	AddUser(c2)

	if status2 := rr2.Code; status2 != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status2, http.StatusBadRequest, rr2.Body.String())
	}

	//---------------negative case : 2---------------------------------------------------
	rr3 := httptest.NewRecorder()
	c3, _ := gin.CreateTestContext(rr3)
	requestBody3, _ := json.Marshal(map[string]string{
		"user":"abc",
	})
	c3.Request, _ = http.NewRequest("POST","/user", bytes.NewBuffer(requestBody3))
	c3.Request.Header.Set("Content-Type", "application/json")
	AddUser(c3)

	if status3 := rr3.Code; status3 != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status3, http.StatusBadRequest, rr3.Body.String())
	}
}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//---------------------------add user before deleting-----------------------------------------
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	requestBody, _ := json.Marshal(map[string]string{
		"user":"abc",
		"password":"123",
	})
	c.Request, _ = http.NewRequest("POST","/user", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")
	AddUser(c)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status, http.StatusOK, rr.Body.String())
	}

	//--------------------------delete user positive case-----------------------------------------
	rr2 := httptest.NewRecorder()
	c2, _ := gin.CreateTestContext(rr2)
	requestBody2, _ := json.Marshal(map[string]string{
		"user":"abc",
	})
	c2.Request, _ = http.NewRequest("DELETE","/user", bytes.NewBuffer(requestBody2))
	c2.Request.Header.Set("Content-Type", "application/json")
	DeleteUser(c2)

	if status2 := rr2.Code; status2 != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status2, http.StatusOK, rr2.Body.String())
	}

	//--------------------------delete user negative case 1-----------------------------------------
	rr3 := httptest.NewRecorder()
	c3, _ := gin.CreateTestContext(rr3)
	requestBody3, _ := json.Marshal(map[string]string{
		"password":"abc",
	})
	c3.Request, _ = http.NewRequest("DELETE","/user", bytes.NewBuffer(requestBody3))
	c3.Request.Header.Set("Content-Type", "application/json")
	DeleteUser(c3)

	if status3 := rr3.Code; status3 != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status3, http.StatusBadRequest, rr3.Body.String())
	}

	//--------------------------delete user negative case 2-----------------------------------------
	rr4 := httptest.NewRecorder()
	c4, _ := gin.CreateTestContext(rr4)
	requestBody4, _ := json.Marshal(map[string]string{
		"user":"abc",
	})
	c4.Request, _ = http.NewRequest("DELETE","/user", bytes.NewBuffer(requestBody4))
	c4.Request.Header.Set("Content-Type", "application/json")
	DeleteUser(c4)

	if status4 := rr4.Code; status4 != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status4, http.StatusBadRequest, rr4.Body.String())
	}
}

func TestAddTopic(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//---------------------------add user before adding topic-----------------------------------------
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	requestBody, _ := json.Marshal(map[string]string{
		"user":"desert",
		"password":"lion",
	})
	c.Request, _ = http.NewRequest("POST","/user", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")
	AddUser(c)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status, http.StatusOK, rr.Body.String())
	}

	//--------------------------add topic positive case-----------------------------------------------
	rr2 := httptest.NewRecorder()
	c2, _ := gin.CreateTestContext(rr2)
	var exaddtopic = lib.AddDeleteTopicJSON {
		User: "desert",
		Permissions: lib.PermissionJSON{
			Publish: []string{"test"},
			Subscribe: []string{"test"},
		},
	}
	requestBody2, _ := json.Marshal(exaddtopic)
	c2.Request, _ = http.NewRequest("POST","/topic", bytes.NewBuffer(requestBody2))
	c2.Request.Header.Set("Content-Type", "application/json")
	AddTopic(c2)

	if status2 := rr2.Code; status2 != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status2, http.StatusOK, rr2.Body.String())
	}

	//--------------------------add topic negative case 1-----------------------------------------------
	rr3 := httptest.NewRecorder()
	c3, _ := gin.CreateTestContext(rr3)
	var exaddtopic2 = lib.AddDeleteTopicJSON {
		User: "desert",
		Permissions: lib.PermissionJSON{
			Publish: []string{"demo","test"},
			Subscribe: []string{"demo"},
		},
	}
	requestBody3, _ := json.Marshal(exaddtopic2)
	c3.Request, _ = http.NewRequest("POST","/topic", bytes.NewBuffer(requestBody3))
	c3.Request.Header.Set("Content-Type", "application/json")
	AddTopic(c3)

	if status3 := rr3.Code; status3 != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status3, http.StatusBadRequest, rr3.Body.String())
	}

	//--------------------------add topic negative case 2-----------------------------------------------
	rr4 := httptest.NewRecorder()
	c4, _ := gin.CreateTestContext(rr4)
	var exaddtopic3 = lib.AddDeleteTopicJSON {
		User: "desert",
		Permissions: lib.PermissionJSON{
			Publish: []string{"demo"},
			Subscribe: []string{"demo","test"},
		},
	}
	requestBody4, _ := json.Marshal(exaddtopic3)
	c4.Request, _ = http.NewRequest("POST","/topic", bytes.NewBuffer(requestBody4))
	c4.Request.Header.Set("Content-Type", "application/json")
	AddTopic(c4)

	if status4 := rr4.Code; status4 != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status4, http.StatusBadRequest, rr4.Body.String())
	}

	//--------------------------add topic negative case 3-----------------------------------------------
	rr5 := httptest.NewRecorder()
	c5, _ := gin.CreateTestContext(rr5)
	var exaddtopic4 = lib.AddDeleteTopicJSON {
		User: "moon",
		Permissions: lib.PermissionJSON{
			Publish: []string{"water"},
			Subscribe: []string{"water"},
		},
	}
	requestBody5, _ := json.Marshal(exaddtopic4)
	c5.Request, _ = http.NewRequest("POST","/topic", bytes.NewBuffer(requestBody5))
	c5.Request.Header.Set("Content-Type", "application/json")
	AddTopic(c5)

	if status5 := rr5.Code; status5 != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status5, http.StatusBadRequest, rr5.Body.String())
	}

	//--------------------------add topic negative case 4-----------------------------------------------
	rr6 := httptest.NewRecorder()
	c6, _ := gin.CreateTestContext(rr6)
	requestBody6, _ := json.Marshal(map[string]string{
		"password":"abc",
	})
	c6.Request, _ = http.NewRequest("POST","/topic", bytes.NewBuffer(requestBody6))
	c6.Request.Header.Set("Content-Type", "application/json")
	AddTopic(c6)

	if status6 := rr6.Code; status6 != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status6, http.StatusBadRequest, rr6.Body.String())
	}
}

func TestDeleteTopic(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//--------------------------delete topic positive case-----------------------------------------------
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	var exaddtopic = lib.AddDeleteTopicJSON {
		User: "desert",
		Permissions: lib.PermissionJSON{
			Publish: []string{"test"},
			Subscribe: []string{"test"},
		},
	}
	requestBody, _ := json.Marshal(exaddtopic)
	c.Request, _ = http.NewRequest("DELETE","/topic", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")
	DeleteTopic(c)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status, http.StatusOK, rr.Body.String())
	}

	//--------------------------delete topic negative case 1-----------------------------------------------
	rr2 := httptest.NewRecorder()
	c2, _ := gin.CreateTestContext(rr2)
	var exaddtopic2 = lib.AddDeleteTopicJSON {
		User: "desert",
		Permissions: lib.PermissionJSON{
			Publish: []string{"test"},
			Subscribe: []string{"test"},
		},
	}
	requestBody2, _ := json.Marshal(exaddtopic2)
	c2.Request, _ = http.NewRequest("DELETE","/topic", bytes.NewBuffer(requestBody2))
	c2.Request.Header.Set("Content-Type", "application/json")
	DeleteTopic(c2)

	if status2 := rr2.Code; status2 != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status2, http.StatusBadRequest, rr2.Body.String())
	}

	//--------------------------delete topic negative case 2-----------------------------------------------
	rr3 := httptest.NewRecorder()
	c3, _ := gin.CreateTestContext(rr3)
	requestBody3, _ := json.Marshal(map[string]string{
		"password":"abc",
	})
	c3.Request, _ = http.NewRequest("DELETE","/topic", bytes.NewBuffer(requestBody3))
	c3.Request.Header.Set("Content-Type", "application/json")
	DeleteTopic(c3)

	if status3 := rr3.Code; status3 != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status3, http.StatusBadRequest, rr3.Body.String())
	}
}

func TestDownloadConfiguration(t *testing.T) {
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request, _ = http.NewRequest("POST","/reload", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	DownloadConfiguration(c)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code : got %v want %v, body: %v", status, http.StatusBadRequest, rr.Body.String())
	}
}