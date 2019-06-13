package main

type TWAMPTestRequest struct {
	Sequence  uint32
	Timestamp Timestamp
	ErrorEst  ErrorEstimate
}

type TWAMPTestResponse struct {
	Sequence        uint32
	Timestamp       Timestamp
	ErrorEst        ErrorEstimate
	MBZ             [2]byte
	RcvTimestamp    Timestamp
	SenderSequence  uint32
	SenderTimestamp Timestamp
	SenderErrorEst  ErrorEstimate
	MBZ2            [2]byte
	SenderTTL       byte
}

type Timestamp struct {
	Seconds  uint32
	Fraction uint32
}

type ErrorEstimate struct {
	SZScale    byte
	Multiplier byte
}
