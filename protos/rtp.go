package protos

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/sipcapture/heplify/ownlayers"
)

func NewRTP(raw []byte) (*ownlayers.RTP, error) {
	rtpl := gopacket.NewPacket(raw, ownlayers.LayerTypeRTP, gopacket.DecodeOptions{Lazy: true, NoCopy: true})
	rtp, ok := rtpl.Layers()[0].(*ownlayers.RTP)
	if !ok {
		//return nil
		return nil, fmt.Errorf("not a valid RTP packet")
	}

	return rtp, nil
}
