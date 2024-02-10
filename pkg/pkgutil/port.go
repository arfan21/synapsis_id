package pkgutil

import "github.com/arfan21/synapsis_id/config"

func GetPort() string {
	port := config.GetConfig().HttpPort
	if port != "" {
		return ":" + port
	}
	return ":8888"
}
