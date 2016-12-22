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

	// cleanup modules
	SakuraAPI.SetAPIKey(token, secret)
	p1 := SakuraAPI.NewFindModulesParam()
	modules, err := SakuraAPI.FindModules(p1)
	if err != nil {
		panic(err)
	}

	for _, m := range modules {
		deleteParam := SakuraAPI.NewDeleteModuleParam(m.ID)
		SakuraAPI.DeleteModule(deleteParam)
	}

	// cleanup projects
	p2 := SakuraAPI.NewFindProjectsParam()
	projects, err := SakuraAPI.FindProjects(p2)
	if err != nil {
		panic(err)
	}

	for _, p := range projects {
		deleteParam := SakuraAPI.NewDeleteProjectParam(p.ID)
		SakuraAPI.DeleteProject(deleteParam)
	}

	os.Exit(ret)
}
