package supersocket

import (
	"encoding/xml"
)

type ServerConfig struct {
    XMLName xml.Name `xml:"server"`
    Name string `xml:"name,attr"`
    Listeners []ListenerConfig `xml:"listeners>listener"`
}

type ListenerConfig struct {
    Address string `xml:"address,attr"`
}