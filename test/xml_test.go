package main

import (
	"os"
	"testing"
	"fmt"
	"encoding/xml"
)


type Servers struct {
    XMLName xml.Name `xml:"servers"`
    Version string   `xml:"version,attr"`
    Svs     []server `xml:"server"`
}

type server struct {
    ServerName string `xml:"serverName"`
    ServerIP   string `xml:"serverIP"`
}

func TestXmlEncode(t *testing.T) {
	v := &Servers{Version: "1"}
    v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
    v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
    output, err := xml.MarshalIndent(v, "", "    ")
    if err != nil {
        t.Fatal(err.Error())
    }
    
    os.Stdout.Write([]byte(xml.Header))
    os.Stdout.Write(output)
	
	var decodedV Servers
    err = xml.Unmarshal(output, &decodedV)
    if err != nil {
        t.Fatal(err.Error())
    }
    
    fmt.Println()
    fmt.Println("Version: ", decodedV.Version)
}