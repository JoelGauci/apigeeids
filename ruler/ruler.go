package ruler

// Interface Ruler
type Ruler interface {
	GetRules() error   // get rules/ip addreses,... from a source of trust
	SetRules() error   // mediate rules/ip addresses,... into compliant snort rules
	WriteRules() error // write rules into a dedicated file (.rules extension) to be used by snort IDS engine
}
