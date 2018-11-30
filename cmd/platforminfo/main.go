package main

import (
	p "intel/isecl/lib/platform-info"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		log.Fatal("Error while executing command line utility. Usage :  ./lib-platform-info platforminfoMethodName")
	}
	var methodName = os.Args[1]
	switch methodName {
	case "BiosName":
		log.Println("Bios Name method called")
		biosName, err := p.BiosName()
		if err != nil {
			log.Fatal("Error during retrieval of Bios Name: ", err)
		}
		log.Println("Bios Name:\n", biosName)

	case "BiosVersion":
		log.Println("Bios Version method called")
		biosVersion, err := p.BiosVersion()
		if err != nil {
			log.Fatal("Error during retrieval of Bios Version: ", err)
		}
		log.Println("Bios Version:\n", biosVersion)

	case "HardwareUUID":
		log.Println("Hardware UUID method called")
		hardwareUUID, err := p.HardwareUUID()
		if err != nil {
			log.Fatal("Error during retrieval of Hardware UUID: ", err)
		}
		log.Println("Hardware UUID:\n", hardwareUUID)

	case "OSName":
		log.Println("OS Name method called")
		osName, err := p.OSName()
		if err != nil {
			log.Fatal("Error during retrieval of OS Name: ", err)
		}
		log.Println("OS Name:\n", osName)

	case "OSVersion":
		log.Println("OS Version method called")
		osVersion, err := p.OSVersion()
		if err != nil {
			log.Fatal("Error during retrieval of OS Version: ", err)
		}
		log.Println("OS Version:\n", osVersion)

	case "ProcessorFlags":
		log.Println("Processor Flags method called")
		processorFlags, err := p.ProcessorFlags()
		if err != nil {
			log.Fatal("Error during retrieval of Processor Flags: ", err)
		}
		log.Println("Processor Flags:\n", processorFlags)

	case "ProcessorInfo":
		log.Println("Processor Info method called")
		processorInfo, err := p.ProcessorInfo()
		if err != nil {
			log.Fatal("Error during retrieval of Processor Info: ", err)
		}
		log.Println("Processor Info:\n", processorInfo)

	case "VMMName":
		log.Println("VMM Name method called")
		vmmName, err := p.VMMName()
		if err != nil {
			log.Fatal("Error during retrieval of VMM Name: ", err)
		}
		log.Println("VMM Name:\n", vmmName)

	case "VMMVersion":
		log.Println("VMM Version method called")
		vmmVersion, err := p.VMMVersion()
		if err != nil {
			log.Fatal("Error during retrieval of VMM Version: ", err)
		}
		log.Println("VMM Version:\n", vmmVersion)

	case "TPMVersion":
		log.Println("TPM Version method called")
		tpmVersion, err := p.TPMVersion()
		if err != nil {
			log.Fatal("Error during retrieval of TPM Version: ", err)
		}
		log.Println("TPM Version:\n", tpmVersion)

	case "HostName":
		log.Println("Host Name method called")
		hostName, err := p.HostName()
		if err != nil {
			log.Fatal("Error during retrieval of Host Name: ", err)
		}
		log.Println("Host Name:\n", hostName)

	case "NoOfSockets":
		log.Println("Number of Sockets method called")
		noOfSockets, err := p.NoOfSockets()
		if err != nil {
			log.Fatal("Error during retrieval of Number of Sockets: ", err)
		}
		log.Println("Number of sockets:\n", noOfSockets)

	case "TPMEnabled":
		log.Println("TPM Enabled method called")
		tpmEnabled, err := p.TPMEnabled()
		if err != nil {
			log.Fatal("Error during retrieval of status of TPM: ", err)
		}
		log.Println("TPM Enabled:\n", tpmEnabled)

	case "TXTEnabled":
		log.Println("TXT Enabled method called")
		txtEnabled, err := p.TXTEnabled()
		if err != nil {
			log.Fatal("Error during retrieval of status of TXT: ", err)
		}
		log.Println("TXT Enabled:\n", txtEnabled)

	default:
		log.Println("Invalid method name mentioned.\nExpected values: BiosName, BiosVersion, HardwareUUID, " +
			"OSName, OSVersion, ProcessorFlags, ProcessorInfo, VMMName, VMMVersion, TPMVersion, HostName, " +
			"NoOfSockets, TPMEnabled, TXTEnabled")
	}
}
