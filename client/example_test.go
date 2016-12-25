package client_test

import (
	"fmt"
	"github.com/yamamoto-febc/sakuraio-api/client"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

func Example_find() {
	token := "[put-your-token]"
	secret := "[put-your-secret]"

	projectName := "project-name"
	sortCol := client.ProjectSortIDAsc // Sort by ID asc
	//sortCol := client.ProjectSortIDDesc	// Sort by ID desc
	//sortCol := client.ProjectSortNameAsc	// Sort by Name asc
	//sortCol := client.ProjectSortNameDesc	// Sort by Name desc

	// Instance of SakuraAPIClient
	c := client.SakuraAPI

	// Set API-key to instance of SakuraAPIClient
	c.SetAPIKey(token, secret)

	// Create a parameter to find a project
	p := c.NewFindProjectsParam()
	p.SetName(&projectName)
	p.SetSort(&sortCol)

	// Call find API [GET /project/]
	projects, err := c.FindProjects(p)
	if err != nil {
		panic(err)
	}
	for i, project := range projects {
		fmt.Printf("Finded Project[%d]: ID=%d , Name=%s", i, project.ID, project.Name)

	}
}

func Example_create() {

	token := "[put-your-token]"
	secret := "[put-your-secret]"
	projectName := "project-name"

	// Instance of SakuraAPIClient
	c := client.SakuraAPI

	// Set API-key to instance of SakuraAPIClient
	c.SetAPIKey(token, secret)

	// Create a parameter to create a project
	p := c.NewCreateProjectParam()
	p.SetProject(&models.ProjectUpdate{
		Name: &projectName,
	})

	// Call create API [POST /project/]
	project, err := c.CreateProject(p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created Project: ID=%d , Name=%s", project.ID, project.Name)
}

func Example_read() {
	token := "[put-your-token]"
	secret := "[put-your-secret]"
	targetProjectID := int64(999999999)

	// Instance of SakuraAPIClient
	c := client.SakuraAPI

	// Set API-key to instance of SakuraAPIClient
	c.SetAPIKey(token, secret)

	// Create a parameter to read a project
	p := c.NewReadProjectParam(targetProjectID)

	// Call read API [GET /project/{ProjectID}/]
	project, err := c.ReadProject(p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Readed Project: ID=%d , Name=%s", project.ID, project.Name)
}

func Example_update() {
	token := "[put-your-token]"
	secret := "[put-your-secret]"
	targetProjectID := int64(999999999)
	projectNameUpdate := "project-name-upd"

	// Instance of SakuraAPIClient
	c := client.SakuraAPI

	// Set API-key to instance of SakuraAPIClient
	c.SetAPIKey(token, secret)

	// Create a parameter to update a project
	p := c.NewUpdateProjectParam(targetProjectID)
	p.SetProject(&models.ProjectUpdate{
		Name: &projectNameUpdate,
	})

	// Call update API [PUT /project/{ProjectID}/]
	project, err := c.UpdateProject(p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Updated Project: ID=%d , Name=%s", project.ID, project.Name)
}

func Example_delete() {
	token := "[put-your-token]"
	secret := "[put-your-secret]"
	targetProjectID := int64(999999999)

	// Instance of SakuraAPIClient
	c := client.SakuraAPI

	// Set API-key to instance of SakuraAPIClient
	c.SetAPIKey(token, secret)

	// Create a parameter to delete a project
	p := c.NewDeleteProjectParam(targetProjectID)

	// Call read API [DELETE /project/{ProjectID}/]
	err := c.DeleteProject(p)
	if err != nil {
		panic(err)
	}
}
