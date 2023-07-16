package requests

type Automations struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			AutomationName string `json:"automation_name" gorm:"automation_name"`
		} `json:"attributes"`
	} `json:"data"`
}
