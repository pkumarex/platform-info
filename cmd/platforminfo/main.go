package main

import (
	"lib-go-platform-info/pkg/platforminfo"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		log.Fatal("Error while executing command line utility. Usage :  ./main platforminfoMethodName")
	}
	var methodName string = os.Args[1]
	p := platforminfo.New()
	switch methodName {
		case "BiosName" :
			log.Printf("Bios Name method called")
			biosName,err := p.BiosName()
			if err != nil {
				log.Fatal("Error during retrieval of Bios Name: ", err)
			}
			log.Printf("Bios Name:%s\n", biosName)

		case "BiosVersion" :
			log.Printf("Bios Version method called")
			biosVersion,err := p.BiosVersion()
			if err != nil {
				log.Fatal("Error during retrieval of Bios Version: ", err)
			}
			log.Printf("Bios Version:%s\n", biosVersion)

		case "HardwareUUID" :
			log.Printf("Hardware UUID method called")
			hardwareUUID,err := p.HardwareUUID()
			if err != nil {
				log.Fatal("Error during retrieval of Hardware UUID: ", err)
			}
			log.Printf("Hardware UUID:%s\n", hardwareUUID)

		case "OSName" :
			log.Printf("OS Name method called")
			osName,err := p.OSName()
			if err != nil {
				log.Fatal("Error during retrieval of OS Name: ", err)
			}
			log.Printf("OS Name:%s\n", osName)

		case "OSVersion" :
			log.Printf("OS Version method called")
			osVersion,err := p.OSVersion()
			if err != nil {
				log.Fatal("Error during retrieval of OS Version: ", err)
			}
			log.Printf("OS Version:%s\n", osVersion)

		case "ProcessorFlags" :
			log.Printf("Processor Flags method called")
			processorFlags,err := p.ProcessorFlags()
			if err != nil {
				log.Fatal("Error during retrieval of Processor Flags: ", err)
			}
			log.Printf("Processor Flags:%s\n", processorFlags)

		case "ProcessorInfo" :
			log.Printf("Processor Info method called")
			processorInfo,err := p.ProcessorInfo()
			if err != nil {
				log.Fatal("Error during retrieval of Processor Info: ", err)
			}
			log.Printf("Processor Info:%s\n", processorInfo)

		case "VMMName" :
			log.Printf("VMM Name method called")
			vmmName,err := p.VMMName()
			if err != nil {
				log.Fatal("Error during retrieval of VMM Name: ", err)
			}
			log.Printf("VMM Name:%s\n", vmmName)

		case "VMMVersion" :
			log.Printf("VMM Version method called")
			vmmVersion,err := p.VMMVersion()
			if err != nil {
				log.Fatal("Error during retrieval of VMM Version: ", err)
			}
			log.Printf("VMM Version:%s\n", vmmVersion)

		case "TPMVersion" :
			log.Printf("TPM Version method called")
			tpmVersion,err := p.TPMVersion()
			if err != nil {
				log.Fatal("Error during retrieval of TPM Version: ", err)
			}
			log.Printf("TPM Version:%s\n", tpmVersion)

		case "HostName" :
			log.Printf("Host Name method called")
			hostName,err := p.HostName()
			if err != nil {
				log.Fatal("Error during retrieval of Host Name: ", err)
			}
			log.Printf("Host Name:%s\n", hostName)

		case "NoOfSockets" :
			log.Printf("Number of Sockets method called")
			noOfSockets,err := p.NoOfSockets()
			if err != nil {
				log.Fatal("Error during retrieval of Number of Sockets: ", err)
			}
			log.Printf("Number of sockets:%s\n", noOfSockets)

		case "TPMEnabled" :
			log.Printf("TPM Enabled method called")
			tpmEnabled,err := p.TPMEnabled()
			if err != nil {
				log.Fatal("Error during retrieval of status of TPM: ", err)
			}
			log.Printf("TPM Enabled:%s\n", tpmEnabled)

		case "TXTEnabled" :
			log.Printf("TXT Enabled method called")
			txtEnabled,err := p.TXTEnabled()
			if err != nil {
				log.Fatal("Error during retrieval of status of TXT: ", err)
			}
			log.Printf("TXT Enabled:%s\n", txtEnabled)

		default :
				log.Printf("Invalid method name mentioned.\nExpected values: BiosName, BiosVersion, HardwareUUID, " +
					"OSName, OSVersion, ProcessorFlags, ProcessorInfo, VMMName, VMMVersion, TPMVersion, HostName, " +
					"NoOfSockets, TPMEnabled, TXTEnabled" )
	}
}
