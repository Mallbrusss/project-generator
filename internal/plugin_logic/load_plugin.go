package pluginlogic

//import (
//	"log"
//	"plugin"
//)
//
//func LoadPlugin(pluginPath string) {
//	p, err := plugin.Open(pluginPath)
//	if err != nil {
//		log.Println("Error loading plugin:", err)
//		return
//	}
//
//	plugSymbol, err := p.Lookup("RegisterPlugin")
//	if err != nil {
//		log.Printf("Error looking up route function: %v\n", err)
//	}
//
//	regPlugin := plugSymbol.(func())
//	regPlugin()
//	log.Printf("Plugin loaded from %s\n", pluginPath)
//}
