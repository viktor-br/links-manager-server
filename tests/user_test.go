// tests contains end to end tests for links manager server.
package tests

import (
	"testing"
	"github.com/viktor-br/links-manager-server/client"
	"net/http"
	"os"
)

func setUp() *client.DefaultApi{
	baseURL := os.Getenv("LMS_TEST_SERVER_BASE_URL")
	return client.NewDefaultApiWithBasePath(baseURL)
}

func TestUserAuthenticatedSuccessfully(t *testing.T) {
	api := setUp()

	userAuth := client.UserAuth{
		Username: "admin",
		Password: "admin",
	}

	response, err := api.UserLoginPost(userAuth)
	if err != nil {
		t.Errorf("Faile to send login request: %s", err.Error())
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected HTTP status code: %d, actual: %d", http.StatusOK, response.StatusCode)
	}
}

func TestUserAuthenticateFailed(t *testing.T) {
	api := setUp()

	userAuth := client.UserAuth{
		Username: "non-existing-user",
		Password: "admin",
	}

	response, err := api.UserLoginPost(userAuth)
	if err != nil {
		t.Errorf("Faile to send login request: %s", err.Error())
	}

	if response.StatusCode != http.StatusForbidden {
		t.Errorf("Expected HTTP status code: %d, actual: %d", http.StatusForbidden, response.StatusCode)
	}
}

func TestUserEmptyPasswordAuthenticateFailed(t *testing.T) {
	api := setUp()

	userAuth := client.UserAuth{
		Username: "admin",
		Password: "",
	}

	response, err := api.UserLoginPost(userAuth)
	if err != nil {
		t.Errorf("Faile to send login request: %s", err.Error())
	}

	if response.StatusCode != http.StatusForbidden {
		t.Errorf("Expected HTTP status code: %d, actual: %d", http.StatusForbidden, response.StatusCode)
	}
}
