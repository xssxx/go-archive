package client

import (
	"fmt"
	"go-desing-pattern/adapter/computer"
)

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com computer.Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
