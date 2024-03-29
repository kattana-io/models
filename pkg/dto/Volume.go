package dto

import (
	"math/big"

	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type Volume struct {
	VolumeA   *big.Int `json:"volume_a"`
	VolumeB   *big.Int `json:"volume_b"`
	VolumeUSD float64  `json:"volume_usd"`
}

func (v *Volume) pack() *storage.Volume {
	return &storage.Volume{
		VolumeA:   v.VolumeA.Bytes(),
		VolumeB:   v.VolumeB.Bytes(),
		VolumeUSD: v.VolumeUSD,
	}
}

func (v *Volume) unpack(data *storage.Volume) *Volume {
	v.VolumeA = new(big.Int).SetBytes(data.VolumeA)
	v.VolumeB = new(big.Int).SetBytes(data.VolumeB)
	v.VolumeUSD = data.VolumeUSD
	return v
}

func (v *Volume) Marshal() ([]byte, error) {
	return proto.Marshal(v.pack())
}

func (v *Volume) UnMarshal(buf []byte) error {
	data := &storage.Volume{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	v.unpack(data)
	return nil
}
