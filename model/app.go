package model

// App App info
type App struct {
	Identifier  string `json:"identifier" bson:"identifier"`
	Title       string `json:"title" bson:"title"`
	URL         string `json:"URL" bson:"URL"`
	IconURL     string `json:"icon_url" bson:"icon_url"`
	DisplayName string `json:"display_name" bson:"display_name"`
}
