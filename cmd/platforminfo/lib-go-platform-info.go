package main

import (
	"lib-go-platform-info/pkg/platforminfo"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		log.Fatal("Error while executing integration test. Usage :  ./lib-go-platform-info platforminfoMethodName")
	}
	var methodName string = os.Args[1]
	switch methodName {
		case "BiosName" :
			log.Printf("Bios Name method called")
			p := platforminfo.New()
			biosName,err := p.BiosName()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("Bios Name:%s\n", biosName)

		case "BiosVersion" :
			log.Printf("Bios Version method called")
			p := platforminfo.New()
			biosVersion,err := p.BiosVersion()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("Bios Version:%s\n", biosVersion)

		case "HardwareUUID" :
			log.Printf("Hardware UUID method called")
			p := platforminfo.New()
			hardwareUUID,err := p.HardwareUUID()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("Hardware UUID:%s\n", hardwareUUID)

		case "OSName" :
			log.Printf("OS Name method called")
			p := platforminfo.New()
			osName,err := p.OSName()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("OS Name:%s\n", osName)

		case "OSVersion" :
			log.Printf("OS Version method called")
			p := platforminfo.New()
			osVersion,err := p.OSVersion()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("OS Version:%s\n", osVersion)

		case "ProcessorFlags" :
			log.Printf("Processor Flags method called")
			p := platforminfo.New()
			processorFlags,err := p.ProcessorFlags()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("Processor Flags:%s\n", processorFlags)

		case "ProcessorInfo" :
			log.Printf("Processor Info method called")
			p := platforminfo.New()
			processorInfo,err := p.ProcessorInfo()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("Processor Info:%s\n", processorInfo)

		case "VMMName" :
			log.Printf("VMM Name method called")
			p := platforminfo.New()
			vmmName,err := p.VMMName()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("VMM Name:%s\n", vmmName)

		case "VMMVersion" :
			log.Printf("VMM Version method called")
			p := platforminfo.New()
			vmmVersion,err := p.VMMVersion()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("VMM Version:%s\n", vmmVersion)

		case "TPMVersion" :
			log.Printf("TPM Version method called")
			p := platforminfo.New()
			tpmVersion,err := p.TPMVersion()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("TPM Version:%s\n", tpmVersion)

		case "HostName" :
			log.Printf("Host Name method called")
			p := platforminfo.New()
			hostName,err := p.HostName()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("Host Name:%s\n", hostName)

		case "NoOfSockets" :
			log.Printf("Number of Sockets method called")
			p := platforminfo.New()
			noOfSockets,err := p.NoOfSockets()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("Number of sockets:%s\n", noOfSockets)

		case "TPMEnabled" :
			log.Printf("TPM Enabled method called")
			p := platforminfo.New()
			tpmEnabled,err := p.TPMEnabled()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("TPM Enabled:%s\n", tpmEnabled)

		case "TXTEnabled" :
			log.Printf("TXT Enabled method called")
			p := platforminfo.New()
			txtEnabled,err := p.TXTEnabled()
			if err != nil {
				log.Fatal("Error occurred: ", err.Error())
			}
			log.Printf("TXT Enabled:%s\n", txtEnabled)

		default :
				log.Printf("Invalid method name mentioned.\nExpected values: BiosName, BiosVersion, HardwareUUID, " +
					"OSName, OSVersion, ProcessorFlags, ProcessorInfo, VMMName, VMMVersion, TPMVersion, HostName, " +
					"NoOfSockets, TPMEnabled, TXTEnabled" )
	}
}
