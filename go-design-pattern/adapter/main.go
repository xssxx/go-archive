package main

import (
	"go-desing-pattern/adapter/computer/mac"
	"go-desing-pattern/adapter/computer/windows"

	"go-desing-pattern/adapter/client"
)

func main() {
	client := &client.Client{}
	mac := &mac.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &windows.Windows{}
	windowsMachineAdapter := &windows.WindowsAdapter{
		WindowsMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
