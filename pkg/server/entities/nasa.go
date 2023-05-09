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
	Copyright      string `json:"copyright"`
}

type MarsRoverPhotosRequest struct {
	Sol    int    `json:"sol" url:"sol"`
	Camera string `json:"camera" url:"camera"`
	ApiKey string `json:"api_key" url:"api_key"`
}

type MarsRoverPhotots struct {
	Photos []MarsRoverPhoto `json:"photos"`
}

type MarsRoverPhoto struct {
	ID        int             `json:"id"`
	Sol       int             `json:"sol"`
	Camera    MarsRoverCamera `json:"camera"`
	ImgSrc    string          `json:"img_src"`
	EarthDate string          `json:"string"`
	Rover     MarsRover       `json:"rover"`
}

type MarsRoverCamera struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RoverID  int    `json:"rover_id"`
	FullName string `json:"full_name"`
}

type MarsRover struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	LandingDate string `json:"landing_date"`
	LaunchDate  string `json:"launch_date"`
	Status      string `json:"status"`
}
