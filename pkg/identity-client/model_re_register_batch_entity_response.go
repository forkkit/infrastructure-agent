// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
/*
 * Identity API specifications
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package identity
// ReRegisterBatchEntityResponse struct for ReRegisterBatchEntityResponse
type ReRegisterBatchEntityResponse struct {
	EntityId int64 `json:"entityId,omitempty"`
	EntityName string `json:"entityName,omitempty"`
	Guid string `json:"guid,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	Error string `json:"error,omitempty"`
}