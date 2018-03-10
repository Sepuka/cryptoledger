package structs

type WatchEntity struct {
	Address string
	MinAlertValue int64
}

type Configuration struct {
	ApiTokenEtherscan string
	Ethereum []WatchEntity
	Bitcoin []WatchEntity
}
