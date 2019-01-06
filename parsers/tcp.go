package parsers

type tcp struct {
	Bytes_on_wire    int
	Source_ip        int
	SourcePort       int
	DEstinationPort  int
	content_type     string
	delta_tcp_stream float64
	delta_tcp_packet float64
	payload interface
}
