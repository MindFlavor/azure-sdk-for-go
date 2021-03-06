package logic

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// WorkflowVersions Client
type WorkflowVersionsClient struct {
	LogicManagementClient
}

func NewWorkflowVersionsClient(subscriptionId string) WorkflowVersionsClient {
	return NewWorkflowVersionsClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewWorkflowVersionsClientWithBaseUri(baseUri string, subscriptionId string) WorkflowVersionsClient {
	return WorkflowVersionsClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// Get gets a workflow version.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. versionId is the workflow versionId.
func (client WorkflowVersionsClient) Get(resourceGroupName string, workflowName string, versionId string) (result WorkflowVersion, ae autorest.Error) {
	req, err := client.NewGetRequest(resourceGroupName, workflowName, versionId)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowVersionsClient", "Get", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowVersionsClient", "Get", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "logic.WorkflowVersionsClient", "Get", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "logic.WorkflowVersionsClient", "Get", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the Get request.
func (client WorkflowVersionsClient) NewGetRequest(resourceGroupName string, workflowName string, versionId string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
		"versionId":         url.QueryEscape(versionId),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.GetRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Get request.
func (client WorkflowVersionsClient) GetRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/versions/{versionId}"))
}
