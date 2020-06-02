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
// GetCallResponse struct for GetCallResponse
type GetCallResponse struct {
	Links GetCallResponseLinks `json:"_links,omitempty"`
	// The unique identifier for this call leg. The UUID is created when your call request is accepted by Nexmo. You use the UUID in all requests for individual live calls
	Uuid string `json:"uuid,omitempty"`
	// The unique identifier for the conversation this call leg is part of.
	ConversationUuid string `json:"conversation_uuid,omitempty"`
	// The single or mixed collection of endpoint types you connected to
	To []map[string]interface{} `json:"to,omitempty"`
	// The endpoint you called from. Possible values are the same as `to`.
	From []map[string]interface{} `json:"from,omitempty"`
	// The status of the call. [See possible values](https://developer.nexmo.com/voice/voice-api/guides/call-flow#event-objects)
	Status string `json:"status,omitempty"`
	Direction Direction `json:"direction,omitempty"`
	// The price per minute for this call. This is only sent if `status` is `completed`.
	Rate string `json:"rate,omitempty"`
	// The total price charged for this call. This is only sent if `status` is `completed`.
	Price string `json:"price,omitempty"`
	// The time elapsed for the call to take place in seconds. This is only sent if `status` is `completed`.
	Duration string `json:"duration,omitempty"`
	// The time the call started in the following format: `YYYY-MM-DD HH:MM:SS`. For example, `2020-01-01 12:00:00`. This is only sent if `status` is `completed`.
	StartTime string `json:"start_time,omitempty"`
	// The time the call started in the following format: `YYYY-MM-DD HH:MM:SS`. For xample, `2020-01-01 12:00:00`. This is only sent if `status` is `completed`.
	EndTime string `json:"end_time,omitempty"`
	// The Mobile Country Code Mobile Network Code ([MCCMNC](https://en.wikipedia.org/wiki/Mobile_country_code)) for the carrier network used to make this call.
	Network string `json:"network,omitempty"`
}
