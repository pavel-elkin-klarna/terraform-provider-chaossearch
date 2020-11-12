package client

// Configuration stores the configuration for a Client
type Configuration struct {
	URL             string
	AccessKeyID     string
	SecretAccessKey string
	AWSServiceName  string
	AWSRegion       string
}

// NewConfiguration creates a default Configuration struct
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		AWSServiceName: "s3",
		AWSRegion:      "eu-west-1",
	}

	return cfg
}
