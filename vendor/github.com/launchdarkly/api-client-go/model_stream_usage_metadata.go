/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.24
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type StreamUsageMetadata struct {
	// The language of the sdk
	Sdk string `json:"sdk,omitempty"`
	// The version of the SDK
	Version string `json:"version,omitempty"`
	Source string `json:"source,omitempty"`
}