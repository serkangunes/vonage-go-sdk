/*
 * Voice API
 *
 * The Voice API lets you create outbound calls, control in-progress calls and get information about historical calls. More information about the Voice API can be found at <https://developer.nexmo.com/voice/voice-api/overview>.
 *
 * API version: 1.2.9
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package voice
// StartStreamResponse struct for StartStreamResponse
type StartStreamResponse struct {
	// Description of the action taken
	Message string `json:"message,omitempty"`
	// The unique identifier for this call leg. The UUID is created when your call request is accepted by Nexmo. You use the UUID in all requests for individual live calls
	Uuid string `json:"uuid,omitempty"`
}
