package client

import (
	"log"
	"os"
	"testing"
)

var (
	token        string
	secret       string
	moduleID     string
	modulePass   string
	moduleSerial string
)

func TestMain(m *testing.M) {

	//Tested only when a API token/secret is set in environment variable
	token = os.Getenv("SAKURAIOT_AUTH_TOKEN")
	secret = os.Getenv("SAKURAIOT_AUTH_SECRET")
	moduleID = os.Getenv("SAKURAIOT_MODULE_ID")
	modulePass = os.Getenv("SAKURAIOT_MODULE_PASS")
	moduleSerial = os.Getenv("SAKURAIOT_MODULE_SERIAL")

	if token == "" || secret == "" {
		log.Println("Please Set ENV 'SAKURAIOT_AUTH_TOKEN' and 'SAKURAIOT_AUTH_SECRET'")
		os.Exit(0) // exit normal
	}
	if moduleID == "" || modulePass == "" {
		log.Println("Please Set ENV 'SAKURAIOT_MODULE_ID' and 'SAKURAIOT_MODULE_PASS'")
		os.Exit(0) // exit normal
	}

	ret := m.Run()

	// cleanup projects
	SakuraAPI.SetAPIKey(token, secret)
	param := SakuraAPI.NewFindProjectsParam()
	projects, err := SakuraAPI.FindProjects(param)
	if err != nil {
		panic(err)
	}

	for _, p := range projects {
		deleteParam := SakuraAPI.NewDeleteProjectParam(p.ID)
		SakuraAPI.DeleteProject(deleteParam)
	}

	os.Exit(ret)
}
