/*
Slurm Rest API

Testing SlurmApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_SlurmApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SlurmApiService SlurmctldCancelJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		httpRes, err := apiClient.SlurmApi.SlurmctldCancelJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmctldDiag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmctldDiag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmctldGetJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmApi.SlurmctldGetJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmctldGetJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmctldGetJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmctldGetNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmctldGetNode(context.Background(), nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmctldGetNodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmctldGetNodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmctldGetPartition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var partitionName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmctldGetPartition(context.Background(), partitionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmctldGetPartitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmctldGetPartitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmctldGetReservation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var reservationName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmctldGetReservation(context.Background(), reservationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmctldGetReservations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmctldGetReservations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmctldPing", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmctldPing(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmctldSubmitJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmctldSubmitJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmctldUpdateJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		httpRes, err := apiClient.SlurmApi.SlurmctldUpdateJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
