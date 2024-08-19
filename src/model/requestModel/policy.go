package requestModel

type CreatePolicyRequest struct {
	Name       string `json:"name,omitempty"`
	Version    string `json:"version,omitempty"`
	Statements []struct {
		Sid       string                       `json:"sid,omitempty"`
		Action    string                       `json:"action,omitempty"`
		Effect    string                       `json:"effect,omitempty"`
		Resource  string                       `json:"resource,omitempty"`
		Condition map[string]map[string]string `json:"condition,omitempty"`
	} `json:"statements,omitempty"`
}

type UpdatePolicyRequest struct {
	ID         string  `json:"ID,omitempty"`
	Name       *string `json:"name,omitempty"`
	Version    *string `json:"version,omitempty"`
	Statements []struct {
		Sid       string                       `json:"sid,omitempty"`
		Action    string                       `json:"action,omitempty"`
		Effect    string                       `json:"effect,omitempty"`
		Resource  string                       `json:"resource,omitempty"`
		Condition map[string]map[string]string `json:"condition,omitempty"`
	} `json:"statements,omitempty"`
}
