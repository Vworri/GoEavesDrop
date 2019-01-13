package packet

type Packet struct {
	// Will house full packets for processing
	Pac       string
	Complete  bool
	Processed bool
}
