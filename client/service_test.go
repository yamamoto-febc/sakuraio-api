package client

import (
	"github.com/stretchr/testify/assert"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
	"testing"
)

var testServiceName = "Example Module"

func TestSakuraAPIClient_WebSocketServiceCRUD(t *testing.T) {
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

	// create WebSocket service
	webSocketCreateParam := SakuraAPI.NewCreateWebSocketServiceParam()
	webSocketCreateParam.SetWebSocketService(&models.WebSocketServiceUpdate{
		Project: &project.ID,
		Name:    &testServiceName,
	})

	webSocket, err := SakuraAPI.CreateWebSocketService(webSocketCreateParam)

	assert.NoError(t, err)
	assert.NotNil(t, webSocket)
	assert.NotEmpty(t, webSocket.URL)

	// read WebSocket service
	webSocketReadParam := SakuraAPI.NewReadWebSocketServiceParam(webSocket.ID)
	webSocket, err = SakuraAPI.ReadWebSocketService(webSocketReadParam)

	assert.NoError(t, err)
	assert.NotNil(t, webSocket)
	assert.NotEmpty(t, webSocket.URL)

	// update WebSocket service
	testServiceNameUpd := testServiceName + "UPD"
	webSocketUpdateParam := SakuraAPI.NewUpdateWebSocketServiceParam(webSocket.ID)
	webSocketUpdateParam.SetWebSocketService(&models.WebSocketServiceUpdate{
		Project: &project.ID,
		Name:    &testServiceNameUpd,
	})
	webSocket, err = SakuraAPI.UpdateWebSocketService(webSocketUpdateParam)

	assert.NoError(t, err)
	assert.NotNil(t, webSocket)
	assert.Equal(t, webSocket.Name, testServiceNameUpd)

	// delete WebSocket Service
	webSocketDeleteParam := SakuraAPI.NewDeleteServiceParam(webSocket.ID)
	err = SakuraAPI.DeleteService(webSocketDeleteParam)

	assert.NoError(t, err)

	// delete project
	deleteParam := SakuraAPI.NewDeleteProjectParam(project.ID)

	err = SakuraAPI.DeleteProject(deleteParam)
	assert.NoError(t, err)

}

func TestSakuraAPIClient_OutgoingWebhookServiceCRUD(t *testing.T) {
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

	// create OutgoingWebhook service
	tagetURL := "https://8.8.8.8/"
	webhookCreateParam := SakuraAPI.NewCreateOutgoingWebhookServiceParam()
	webhookCreateParam.SetOutgoingWebhookService(&models.OutgoingWebhookServiceUpdate{
		Project: &project.ID,
		Name:    &testServiceName,
		URL:     &tagetURL,
	})

	webhook, err := SakuraAPI.CreateOutgoingWebhookService(webhookCreateParam)

	assert.NoError(t, err)
	assert.NotNil(t, webhook)
	assert.NotEmpty(t, webhook.URL)

	// read Outgoing Webhook service
	webhookReadParam := SakuraAPI.NewReadOutgoingWebhookServiceParam(webhook.ID)
	webhook, err = SakuraAPI.ReadOutgoingWebhookService(webhookReadParam)

	assert.NoError(t, err)
	assert.NotNil(t, webhook)
	assert.NotEmpty(t, webhook.URL)

	// update Outgoing Webhook service
	testServiceNameUpd := testServiceName + "UPD"
	webhookUpdateParam := SakuraAPI.NewUpdateOutgoingWebhookServiceParam(webhook.ID)
	webhookUpdateParam.SetOutgoingWebhookService(&models.OutgoingWebhookServiceUpdate{
		Project: &project.ID,
		Name:    &testServiceNameUpd,
		URL:     &tagetURL,
	})
	webhook, err = SakuraAPI.UpdateOutgoingWebhookService(webhookUpdateParam)

	assert.NoError(t, err)
	assert.NotNil(t, webhook)
	assert.Equal(t, webhook.Name, testServiceNameUpd)

	// delete Outgoing Webhook Service
	webhookDeleteParam := SakuraAPI.NewDeleteServiceParam(webhook.ID)
	err = SakuraAPI.DeleteService(webhookDeleteParam)

	assert.NoError(t, err)

	// delete project
	deleteParam := SakuraAPI.NewDeleteProjectParam(project.ID)

	err = SakuraAPI.DeleteProject(deleteParam)
	assert.NoError(t, err)

}

func TestSakuraAPIClient_IncomingWebhookServiceCRUD(t *testing.T) {
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

	// create Incoming Webhook service
	webhookCreateParam := SakuraAPI.NewCreateIncomingWebhookServiceParam()
	webhookCreateParam.SetIncomingWebhookService(&models.IncomingWebhookServiceUpdate{
		Project: &project.ID,
		Name:    &testServiceName,
	})

	webhook, err := SakuraAPI.CreateIncomingWebhookService(webhookCreateParam)

	assert.NoError(t, err)
	assert.NotNil(t, webhook)
	assert.NotEmpty(t, webhook.URL)

	// read Incoming Webhook service
	webhookReadParam := SakuraAPI.NewReadIncomingWebhookServiceParam(webhook.ID)
	webhook, err = SakuraAPI.ReadIncomingWebhookService(webhookReadParam)

	assert.NoError(t, err)
	assert.NotNil(t, webhook)
	assert.NotEmpty(t, webhook.URL)

	// update Incoming Webhook service
	testServiceNameUpd := testServiceName + "UPD"
	webhookUpdateParam := SakuraAPI.NewUpdateIncomingWebhookServiceParam(webhook.ID)
	webhookUpdateParam.SetIncomingWebhookService(&models.IncomingWebhookServiceUpdate{
		Project: &project.ID,
		Name:    &testServiceNameUpd,
	})
	webhook, err = SakuraAPI.UpdateIncomingWebhookService(webhookUpdateParam)

	assert.NoError(t, err)
	assert.NotNil(t, webhook)
	assert.Equal(t, webhook.Name, testServiceNameUpd)

	// delete Outgoing Webhook Service
	webhookDeleteParam := SakuraAPI.NewDeleteServiceParam(webhook.ID)
	err = SakuraAPI.DeleteService(webhookDeleteParam)

	assert.NoError(t, err)

	// delete project
	deleteParam := SakuraAPI.NewDeleteProjectParam(project.ID)

	err = SakuraAPI.DeleteProject(deleteParam)
	assert.NoError(t, err)
}
