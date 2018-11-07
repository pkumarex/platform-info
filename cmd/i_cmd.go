package cmd

type ICommand interface {

	retrieveBiosName() string
	retrieveHardwareUUID() string
	retrieveOSVersion() string
	retrieveProcessorFlags() string
	retrieveProcessorInfo() string
	retrieveVMMName() string
	retrieveVMMVersion() string
	retrieveTpmVersion() string
	retrieveHostName() string
	retrieveNoOfSockets() string
	retrieveTPMEnabled() string
	retrieveTXTEnabled() string
}
