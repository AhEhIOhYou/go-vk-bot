package entities

type APODRequset struct {
	ApiKey string `json:"api_key" url:"api_key"`
	Count  int    `json:"count,omitempty" url:"count,omitempty"`
}

type APOD struct {
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	HDUrl          string `json:"hdurl"`
	Url            string `json:"url"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
}
