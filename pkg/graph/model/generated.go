// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AnalyticsEvent struct {
	AppID                string `json:"appID"`
	SessionID            string `json:"sessionID"`
	PayloadType          string `json:"payloadType"`
	PayloadSchemaVersion int    `json:"payloadSchemaVersion"`
	Payload              string `json:"payload"`
	URI                  string `json:"uri"`
	Timestamp            string `json:"timestamp"`
	InstallationID       string `json:"installationID"`
	EnvironmentClass     string `json:"environmentClass"`
}

type AnalyticsEventInput struct {
	AppID                string `json:"appID"`
	SessionID            string `json:"sessionID"`
	PayloadType          string `json:"payloadType"`
	PayloadSchemaVersion int    `json:"payloadSchemaVersion"`
	Payload              string `json:"payload"`
	URI                  string `json:"uri"`
}

type Identity struct {
	Provider string `json:"provider"`
	Codename string `json:"codename"`
}

type Kubernetes struct {
	APIURL  string `json:"apiUrl"`
	AuthURL string `json:"authUrl"`
	CaCert  string `json:"caCert"`
}
