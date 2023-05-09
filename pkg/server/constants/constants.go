package constants

const (
	// Error
	RequestFailed              = "request failed: %v\n"
	RequestCreationError       = "Request creation error: %v\n"
	QueryCreationError         = "Query creation error: %v\n"
	DecodingJSONError          = "Error when decoding json: %v\n"
	ServiceInitializationError = "service initialization error: %v\n"
	ServerErrorOccurred        = "Server error occurred >_<"
	SecretWordMissmatch        = "secret word missmatch"

	// Other

	// Front Hazard Avoidance Camera
	RoverCameraFHAZ = "FHAZ"
	// Rear Hazard Avoidance Camera
	RoverCameraRHAZ = "RHAZ"
	// Mast Camera
	RoverCameraMAST = "MAST"
	// Chemistry and Camera Complex
	RoverCameraCHEMCAM = "CHEMCAM"

	// Nasa API
	NasaApiUrl             = "https://api.nasa.gov/"
	NasaApiMethodApod      = "planetary/apod"
	NasaApiMethodMarsPhoto = "mars-photos/api/v1/rovers/curiosity/photos"

	// VK API
	VkApiVersion           = "5.131"
	VkApiUrl               = "https://api.vk.com/method/"
	VkApiMethodMessageSend = "messages.send"
)

// Bot answers
var BotUnknownCommandsMsg = []string{
	"I'm sorry, I wasn't programmed for free-form conversation.",
	"I'm afraid I don't understand what you're asking. Please choose from the available options.",
	"That's not something I'm capable of doing. Is there anything else I can help you with?",
	"My apologies, but I don't have the ability to engage in casual conversation.",
	"I'm designed to assist with specific tasks and questions. Is there something specific you need?",
}
