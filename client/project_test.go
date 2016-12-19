package client

import (
	"github.com/stretchr/testify/assert"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
	"testing"
)

var testProjectName = "Example project"

func TestSakuraAPIClient_ProjectCRUD(t *testing.T) {

	SakuraAPI.SetAPIKey(token, secret)
	defer SakuraAPI.ClearAPIKey()

	projectName := testProjectName
	projectUpdate := &models.ProjectUpdate{
		Name: &projectName,
	}

	// create project
	createParam := SakuraAPI.NewCreateProjectParam()
	createParam.SetProject(projectUpdate)

	project, err := SakuraAPI.CreateProject(createParam)

	assert.NoError(t, err)
	assert.NotEqual(t, project.ID, 0) // will set ID after create
	assert.Equal(t, project.Name, projectName)

	// find project
	findParam := SakuraAPI.NewFindProjectsParam()
	findParam.SetName(&projectName)

	projects, err := SakuraAPI.FindProjects(findParam)
	assert.Len(t, projects, 1)

	// read project
	readParam := SakuraAPI.NewReadProjectParam(project.ID)
	project, err = SakuraAPI.ReadProject(readParam)
	assert.NoError(t, err)
	assert.NotNil(t, project)

	// update project
	updProjectName := projectName + "UPD"
	projectUpdate.Name = &updProjectName

	updateParam := SakuraAPI.NewUpdateProjectParam(project.ID)
	updateParam.SetProject(projectUpdate)

	project, err = SakuraAPI.UpdateProject(updateParam)
	assert.NoError(t, err)
	assert.Equal(t, project.Name, updProjectName)

	// delete project
	deleteParam := SakuraAPI.NewDeleteProjectParam(project.ID)

	err = SakuraAPI.DeleteProject(deleteParam)
	assert.NoError(t, err)
}
