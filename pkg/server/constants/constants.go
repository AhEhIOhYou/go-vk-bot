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

	BotDescription = "Greetings!\n\nI am a chatbot.\nMy purpose is to give modest access to NASA's public database. Read the description of the commands and enjoy the articles and pictures.\n\nMy creator is ahehiohyou."
	BotCommands    = "The following commands will show you photos of Mars from different cameras of the research rovers\n- FHAZ\n- RHAZ\n- MAST\n- CHEMCAM\n\nIf the photo could not be found, try again!\n\nThis command will show you a random picture and its article from NASA's most popular website, Astronomy Picture of the Day.\n- APOD\n\nWell, test :)\n- Test"

	TestSuccess = "Test? Test! test TEST testtttt"

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
	"I'm designed to assist with specific tasks. Is there something specific you need?",
}
