package client

import (
	"github.com/stretchr/testify/assert"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
	"testing"
)

var testModuleName = "Example Module"

func TestSakuraAPIClient_ModuleCRUD(t *testing.T) {

	SakuraAPI.SetAPIKey(token, secret)
	defer SakuraAPI.ClearAPIKey()

	projectName := testProjectName
	projectUpdate := &models.ProjectUpdate{
		Name: &projectName,
	}

	// create project
	projectParam := SakuraAPI.NewCreateProjectParam()
	projectParam.SetProject(projectUpdate)
	project, err := SakuraAPI.CreateProject(projectParam)
	assert.NoError(t, err)

	// create module
	moduleRegister := &models.ModuleRegister{
		Name:         &testModuleName,
		Project:      &project.ID,
		RegisterID:   &moduleID,
		RegisterPass: &modulePass,
	}

	createParam := SakuraAPI.NewCreateModuleParam()
	createParam.SetRegisterInformation(moduleRegister)

	module, err := SakuraAPI.CreateModule(createParam)
	assert.NoError(t, err)
	assert.NotNil(t, module)
	assert.Equal(t, module.Name, testModuleName)

	// find module
	findParam := SakuraAPI.NewFindModulesParam()
	findParam.SetSerialNumber(&moduleSerial)

	modules, err := SakuraAPI.FindModules(findParam)
	assert.Len(t, modules, 1)

	// read module
	readParam := SakuraAPI.NewReadModuleParam(module.ID)
	module, err = SakuraAPI.ReadModule(readParam)
	assert.NoError(t, err)
	assert.NotNil(t, project)

	// update module
	updModuleName := testModuleName + "UPD"
	moduleUpdate := models.ModuleUpdate{}
	moduleUpdate.Name = &updModuleName
	moduleUpdate.Project = &project.ID

	updateParam := SakuraAPI.NewUpdateModuleParam(module.ID)
	updateParam.SetBody(&moduleUpdate)

	module, err = SakuraAPI.UpdateModule(updateParam)
	assert.NoError(t, err)
	assert.Equal(t, module.Name, updModuleName)

	// delete module
	deleteParam := SakuraAPI.NewDeleteModuleParam(module.ID)

	err = SakuraAPI.DeleteModule(deleteParam)
	assert.NoError(t, err)

	// delete project
	deleteProjectParam := SakuraAPI.NewDeleteProjectParam(project.ID)

	err = SakuraAPI.DeleteProject(deleteProjectParam)
	assert.NoError(t, err)
}
