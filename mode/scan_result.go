package mode

type ScanResult struct {
	Host  string            `json:"host"`
	Port  string            `json:"port,omitempty"`
	Proto string            `json:"proto,omitempty"`
	Probe string            `json:"probe,omitempty"`
	Name  string            `json:"name,omitempty"`
	Nets  []string          `json:"nets,omitempty"`
	Info  map[string]string `json:"info"`
}
