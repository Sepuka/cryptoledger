package structs

type WatchEntity struct {
	Address string
	MinAlertValue int64
}

type Configuration struct {
	ApiTokenEtherscan string
	CheckFrequencyMinutes uint8
	SlackApiToken string
	SlackChannel string
	Ethereum []WatchEntity
	Bitcoin []WatchEntity
}
