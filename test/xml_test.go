package test

import (
	"testing"
    "fmt"
	"encoding/xml"
    "../src"
)

func TestXmlEncode(t *testing.T) {
    var s = supersocket.ServerConfig { Name: "TestServer" }
    
    s.Listeners = append(s.Listeners, supersocket.ListenerConfig { Address: "Any:1012" })
    s.Listeners = append(s.Listeners, supersocket.ListenerConfig { Address: "Any:1013" })
    
    output, err := xml.MarshalIndent(&s, "", "    ")
    if err != nil {
        t.Fatal(err.Error())
    }
    
    fmt.Print(xml.Header)
    fmt.Print(string(output))
    fmt.Println()

    var ds supersocket.ServerConfig
    err = xml.Unmarshal(output, &ds)
    
    if err != nil {
        t.Fatal(err.Error())
    }
    
    if (ds.Name != "TestServer") {
        t.Error("Unexpected server name")
    }
    
    if (len(ds.Listeners) != 2) {
        t.Error("Unexpected listener count")
    }
    
    if (ds.Listeners[0].Address != "Any:1012") {
        t.Error("Unexpected listener 0 address")
    }
    
    if (ds.Listeners[1].Address != "Any:1013") {
        t.Error("Unexpected listener 1 address")
    }
}