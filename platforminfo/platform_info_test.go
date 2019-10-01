/*
 * Copyright (C) 2019 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
package platforminfo

import (
	"fmt"
	"strings"
	"testing"
)

func TestBiosName(t *testing.T) {
	biosName, err := BiosName()
	fmt.Println("Bios Name:")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bios Name:\n", biosName)
	}
}

func TestBiosVersion(t *testing.T) {
	biosVersion, err := BiosVersion()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bios Version:\n", biosVersion)
	}
}

func TestHardwareUUID(t *testing.T) {
	hardwareUUID, err := HardwareUUID()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hardware UUID:\n", hardwareUUID)
	}
}

func TestOSName(t *testing.T) {
	osName, err := OSName()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OS Name:\n", osName)
	}
}

func TestOSVersion(t *testing.T) {
	osVersion, err := OSVersion()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OS Version:\n", osVersion)
	}
}

func TestProcessorFlags(t *testing.T) {
	processorFlags, err := ProcessorFlags()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Flags for processor 0:\n", strings.Join(processorFlags, ", "))
	}
}

func TestProcessorID(t *testing.T) {
	processorID, err := ProcessorID()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ID for processor 0:\n", processorID)
	}
}

func TestVMMName(t *testing.T) {
	vmmName, err := VMMName()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hypervisor name:\n", vmmName)
	}
}

func TestVMMVersion(t *testing.T) {
	vmmVersion, err := VMMVersion()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hypervisor version:\n", vmmVersion)
	}
}

func TestTPMVersion(t *testing.T) {
	tpmVersion, err := TPMVersion()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("TPM version:\n", tpmVersion)
	}
}

func TestHostName(t *testing.T) {
	hostName, err := HostName()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Host name:\n", hostName)
	}
}

func TestNoOfSockets(t *testing.T) {
	numSockets, err := NoOfSockets()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Number of sockets:\n", string(numSockets))
	}
}

func TestTPMEnabled(t *testing.T) {
	tpmStatus, err := TPMEnabled()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("TPM Enabled:\n", tpmStatus)
	}
}

func TestTXTEnabled(t *testing.T) {
	txtStatus, err := TXTEnabled()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("TXT Enabled:\n", txtStatus)
	}
}
