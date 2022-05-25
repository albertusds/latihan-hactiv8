package params

type Response struct {
	Status         int         `json:"status"`
	Message        string      `json:"message"`
	AdditionalInfo string      `json:"additional_info,omitempty"`
	Payload        interface{} `json:"payload,omitempty"`
}
