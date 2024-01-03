package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Sample struct{
		Name string
		Desc  string
	}
	type Datas struct {
		Host_name            string
		Host_uuid            string
		Host_last_logon      int
		Host_logon_count     int
		Host_id              string
		Host_os              string
		Host_os_version      string
		Host_dns_name        string
		Host_system_critical bool
		Timestamp            int
		Host_watcHed         string
		Host_critical        string
		Host_score           int
		Host_sites           int
		Query_incident       string
		Domain               string
		Host_browser         []string
		Sample Sample
	}

	datadJson := `{"host_name": "UBBOHDC", "host_uuid": "ubbohdc$", "host_last_logon": 1690062235560, 
	"host_logon_count": "135", "host_id": "ubbohdc$", "host_os": "Windows Server 2016 Standard", 
	"host_os_version": "10.0 (14393)", "host_dns_name": "UBBOHDC.ubb.com", "host_system_critical": "true", 
	"timestamp": 1693200002867, "host_watched": "no", "host_critical": "no", "host_score": 0, 
	"host_sites": 0, "query_incident": "","sample":{"name":"macbook","desc":"testing"},"host_browser":["Chrome","Firefox"], "domain": "ubb.com"}`

	var data Datas
	json.Unmarshal([]byte(datadJson), &data)
	fmt.Printf("host_dns_name:%s\n", data.Sample.Desc)
}
