// tests contains end to end tests for links manager server.
package tests

import (
	"github.com/viktor-br/links-manager-server/app/controllers"
	"github.com/viktor-br/links-manager-server/client"
	"net/http"
	"os"
	"strings"
	"testing"
)

func setUp() *client.DefaultApi {
	baseURL := os.Getenv("LMS_TEST_SERVER_BASE_URL")
	return client.NewDefaultApiWithBasePath(baseURL)
}

func setUpWithAccessToken() (api *client.DefaultApi, token string, err error) {
	api = setUp()

	userAuth := client.UserAuth{
		Username: "admin",
		Password: "admin",
	}

	response, err := api.UserLoginPost(userAuth)
	if err != nil {
		return nil, "", err
	}
	token = response.Header.Get(controllers.XAuthToken)

	return api, token, nil
}

func testUserLogin(username, password string) (*client.APIResponse, error) {
	api := setUp()

	userAuth := client.UserAuth{
		Username: username,
		Password: password,
	}

	return api.UserLoginPost(userAuth)
}

func TestUserAuthenticatedSuccessfully(t *testing.T) {
	response, err := testUserLogin("admin", "admin")
	if err != nil {
		t.Errorf("Faile to send login request: %s", err.Error())
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected HTTP status code: %d, actual: %d", http.StatusOK, response.StatusCode)
	}
}

func TestUserAuthenticateFailed(t *testing.T) {
	response, err := testUserLogin("non-existing-user", "admin")
	if err != nil {
		t.Errorf("Faile to send login request: %s", err.Error())
	}

	if response.StatusCode != http.StatusForbidden {
		t.Errorf("Expected HTTP status code: %d, actual: %d", http.StatusForbidden, response.StatusCode)
	}
}

func TestUserAuthenticateEmptyPassword(t *testing.T) {
	response, err := testUserLogin("admin", "")
	if err != nil {
		t.Errorf("Faile to send login request: %s", err.Error())
	}

	if response.StatusCode != http.StatusForbidden {
		t.Errorf("Expected HTTP status code: %d, actual: %d", http.StatusForbidden, response.StatusCode)
	}
}

func TestUserAuthenticateCorruptedJSON(t *testing.T) {
	json := strings.NewReader("{corrupted json")

	url := os.Getenv("LMS_TEST_SERVER_BASE_URL")
	url = url + "/user/login"

	resp, err := http.Post(url, "application/json", json)
	if err != nil {
		t.Errorf("Expected err nil, %s given", err.Error())
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected HTTP status code: %d, actual: %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestUserCreateSuccess(t *testing.T) {
	api, token, err := setUpWithAccessToken()
	if err != nil {
		t.Errorf("Expect err is nil, actual: %s", err.Error())
	}

	user := client.User{
		Username: "test",
		Password: "test",
	}

	resp, err := api.UserPut(user, token)
	if err != nil {
		t.Errorf("Expect err is nil, actual %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected HTTP status code: %d, actual: %d", http.StatusOK, resp.StatusCode)
	}

	resp, err = api.UserPut(user, token)
	if err != nil {
		t.Errorf("Expect err is nil, actual %s", err.Error())
	}

	if resp.StatusCode != http.StatusConflict {
		t.Errorf("Expected HTTP status code: %d, actual: %d", http.StatusConflict, resp.StatusCode)
	}
}

func TestUserCreateUnauthorizedAccess(t *testing.T) {
	api := setUp()

	user := client.User{
		Username: "test",
		Password: "test",
	}

	resp, err := api.UserPut(user, "wrong token")
	if err != nil {
		t.Errorf("Expect err is nil, actual %s", err.Error())
	}

	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("Expected HTTP status code: %d, actual: %d", http.StatusForbidden, resp.StatusCode)
	}
}
