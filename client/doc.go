// Package client is library to call the Sakura-IoT APIs
//
// Basic usage
//
// The basic functions of this package are provided by the SakuraAPIClient struct.
//
// An instance of the SakuraAPIClient struct can be referenced from the following fields.
//
//    // Instance of SakuraAPIClient
//    client.SakuraAPI
//
// In order to use the SakuraIoT APIs , API-key must be set for this instance as follows.
// If the API-key is not set, or if an invalid API-key is specified, it will results in a runtime error.
//
//    // Set API-key to instance of SakuraAPIClient
//    client.SakuraAPI.SetAPIKey(token , secret)
//
// CRUD operation for each resource
//
// There is a structs that to set values of parameters for each resource and each operation.
// In operation, please set parameters through these structs.
//
// For more detail, please see follow Example
//
//
package client
