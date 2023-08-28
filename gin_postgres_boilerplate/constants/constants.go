package constants

const (
	DefaultPageSize   = 10
	MaxPageSize       = 100
	DefaultAPIVersion = "v1"
)

var ErrorCodes = map[string]struct {
	Code    string
	Message string
}{
	"EA001": {
		Code:    "EA001",
		Message: "Invalid login credentials",
	},
	"EA002": {
		Code:    "EA002",
		Message: "Something wrong in token generation",
	},
	"EA003": {
		Code:    "EA003",
		Message: "Password mismatch",
	},
	"EA004": {
		Code:    "EA004",
		Message: "Authorization header is missing",
	},
	"EA005": {
		Code:    "EA005",
		Message: "Invalid authorization header format",
	},
	"EA006": {
		Code:    "EA006",
		Message: "Invalid login token",
	},
	"EA007": {
		Code:    "EA007",
		Message: "Unable to parse claims from token",
	},
	"EU001": {
		Code:    "EU001",
		Message: "User creation unsuccessful",
	},
	"EU002": {
		Code:    "EU002",
		Message: "Failed to retrieve users",
	},
	"EU003": {
		Code:    "EU003",
		Message: "User not found",
	},
}
