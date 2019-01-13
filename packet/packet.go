package packet

type Packet struct {
	// will house full packets for processing
	Pac       string
	Complete  bool
	Processed bool
}
