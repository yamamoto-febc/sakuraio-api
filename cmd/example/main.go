// sakura-iot - A command line client for the Sakura-IOT platform.
package main

import (
	"fmt"
	"github.com/yamamoto-febc/sakuraio-api/client"
	"github.com/yamamoto-febc/sakuraio-api/cmd"
	"github.com/yamamoto-febc/sakuraio-api/version"
	"gopkg.in/urfave/cli.v2"
	"os"
	"strings"
)

var (
	appName      = "sakura-iot"
	appUsage     = "A command line client for the Sakura-IOT platform."
	appCopyright = "Copyright (C) 2016 Kazumichi Yamamoto."
)

func main() {

	app := &cli.App{}
	option := &cmd.Option{}

	app.Name = appName
	app.Usage = appUsage
	app.HelpName = appName
	app.Copyright = appCopyright

	app.Flags = cliFlags(option)
	app.Action = cliCommand(option)
	app.Version = version.FullVersion()

	app.Run(os.Args)
}

func cliFlags(option *cmd.Option) []cli.Flag {

	return []cli.Flag{
		&cli.StringFlag{
			Name:        "auth-token",
			Aliases:     []string{"sakuraiot-token"},
			Usage:       "API Token of Sakura-IOT",
			EnvVars:     []string{"SAKURAIOT_AUTH_TOKEN"},
			DefaultText: "",
			Destination: &option.AuthToken,
		},
		&cli.StringFlag{
			Name:        "auth-secret",
			Aliases:     []string{"sakuraiot-token-secret"},
			Usage:       "API Secret of Sakura-IOT",
			EnvVars:     []string{"SAKURAIOT_AUTH_SECRET"},
			DefaultText: "",
			Destination: &option.AuthSecret,
		},
	}

}

func cliCommand(option *cmd.Option) func(c *cli.Context) error {
	return func(c *cli.Context) error {

		errors := option.Validate()
		if len(errors) != 0 {
			return flattenErrors(errors)
		}

		// API client
		sakuraAPI := client.SakuraAPI

		// for Auth by API token and secret
		sakuraAPI.SetAPIKey(option.AuthToken, option.AuthSecret)

		// ------------------------------------------------------------
		// GET /v1/auth/
		// ------------------------------------------------------------
		auth, err := sakuraAPI.Auth()
		if err != nil {
			return err
		}

		fmt.Printf("==================================================\n")
		fmt.Printf("GET /v1/auth/\n")
		fmt.Printf("==================================================\n")
		fmt.Printf("UserName : %s\n", auth.Username)
		fmt.Printf("Level    : %d\n", auth.Level)
		fmt.Printf("JWT      : %s\n", auth.JWT)
		fmt.Printf("==================================================\n\n")

		// ------------------------------------------------------------
		// GET /v1/projects/
		// ------------------------------------------------------------
		projectFindParam := sakuraAPI.NewFindProjectsParam()
		//p1.SetName("cond")
		//p1.SetSort("col")

		projects, err := sakuraAPI.FindProjects(projectFindParam)
		if err != nil {
			return err
		}
		fmt.Printf("==================================================\n")
		fmt.Printf("GET /v1/projects/\n")
		fmt.Printf("==================================================\n")

		for i, p := range projects {
			fmt.Printf("    ----------------------------------------------\n")
			fmt.Printf("    Index    : %d\n", i+1)
			fmt.Printf("    ID       : %d\n", p.ID)
			fmt.Printf("    Name     : %s\n", p.Name)
			fmt.Printf("    ----------------------------------------------\n")
		}
		fmt.Println("")

		fmt.Printf("\n**************************************************\n\n")

		// ------------------------------------------------------------
		// GET /v1/modules/
		// ------------------------------------------------------------
		moduleFindParam := sakuraAPI.NewFindModulesParam()

		modules, err := sakuraAPI.FindModules(moduleFindParam)
		if err != nil {
			return err
		}
		fmt.Printf("==================================================\n")
		fmt.Printf("GET /v1/modules/\n")
		fmt.Printf("==================================================\n")

		for i, m := range modules {
			fmt.Printf("    ----------------------------------------------\n")
			fmt.Printf("    Index    : %d\n", i+1)
			fmt.Printf("    ID       : %d\n", m.ID)
			fmt.Printf("    Name     : %s\n", m.Name)
			fmt.Printf("    Project  : %d\n", m.Project)
			fmt.Printf("    IsOnline : %b\n", m.IsOnline)
			fmt.Printf("    ----------------------------------------------\n")
		}
		fmt.Println("")

		fmt.Printf("\n**************************************************\n\n")

		// ------------------------------------------------------------
		// GET /v1/services/
		// ------------------------------------------------------------
		serviceFindParam := sakuraAPI.NewFindServicesParam()

		services, err := sakuraAPI.FindServices(serviceFindParam)
		if err != nil {
			return err
		}
		fmt.Printf("==================================================\n")
		fmt.Printf("GET /v1/services/\n")
		fmt.Printf("==================================================\n")

		for i, s := range services {
			fmt.Printf("    ----------------------------------------------\n")
			fmt.Printf("    Index    : %d\n", i+1)
			fmt.Printf("    ID       : %d\n", s.ID)
			fmt.Printf("    Name     : %s\n", s.Name)
			fmt.Printf("    Project  : %d\n", s.Project)
			fmt.Printf("    IsOnline : %s\n", s.Type)
			fmt.Printf("    ----------------------------------------------\n")
		}
		fmt.Println("")

		fmt.Printf("\n**************************************************\n\n")

		return nil

	}
}

func flattenErrors(errors []error) error {
	var list = make([]string, 0)
	for _, str := range errors {
		list = append(list, str.Error())
	}
	return fmt.Errorf(strings.Join(list, "\n"))
}
