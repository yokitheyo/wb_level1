package main

import "fmt"

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("client inserts lightning connector into computer.")
	com.InsertIntoLightningPort()
}

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("lightning connector is plugged into mac machine")
}

type Windows struct{}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("adapter converts lightning signal to USB")
	w.windowsMachine.InsertIntoUSBPort()
}

func main() {
	client := &Client{}
	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{windowsMachine}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
