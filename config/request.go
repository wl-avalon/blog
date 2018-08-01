package config

type requestConf struct {
	Domain string
	InterfaceMap map[string]string
}

var RequestConfMap = map[string]requestConf{
	"idgent" : {
		Domain: "123.56.156.172:4021",
		InterfaceMap: map[string]string{
			"nextID": "/nextId",
			"batch": "/batch",
		},
	},
}
