package Config

var Config Configuration

type Configuration struct {
	DbAddr     string `json:"dbAddr"`
	DbPassWord string `json:"dbPassWord"`
}
