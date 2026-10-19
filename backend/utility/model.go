package utility

type ResourceList struct {
	Types     []string   `json:"types"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Example     string   `json:"example"`
	Types       []string `json:"types"`
}
