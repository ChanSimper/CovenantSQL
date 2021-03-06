package types

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	hsp "github.com/CovenantSQL/HashStablePack/marshalhash"
)

// MarshalHash marshals for hash
func (z *Block) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 2
	o = append(o, 0x82, 0x82)
	if oTemp, err := z.SignedHeader.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x82)
	o = hsp.AppendArrayHeader(o, uint32(len(z.Queries)))
	for za0001 := range z.Queries {
		if z.Queries[za0001] == nil {
			o = hsp.AppendNil(o)
		} else {
			if oTemp, err := z.Queries[za0001].MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Block) Msgsize() (s int) {
	s = 1 + 13 + z.SignedHeader.Msgsize() + 8 + hsp.ArrayHeaderSize
	for za0001 := range z.Queries {
		if z.Queries[za0001] == nil {
			s += hsp.NilSize
		} else {
			s += z.Queries[za0001].Msgsize()
		}
	}
	return
}

// MarshalHash marshals for hash
func (z Blocks) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	o = hsp.AppendArrayHeader(o, uint32(len(z)))
	for za0001 := range z {
		if z[za0001] == nil {
			o = hsp.AppendNil(o)
		} else {
			if oTemp, err := z[za0001].MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Blocks) Msgsize() (s int) {
	s = hsp.ArrayHeaderSize
	for za0001 := range z {
		if z[za0001] == nil {
			s += hsp.NilSize
		} else {
			s += z[za0001].Msgsize()
		}
	}
	return
}

// MarshalHash marshals for hash
func (z *Header) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 6
	o = append(o, 0x86, 0x86)
	if oTemp, err := z.GenesisHash.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x86)
	if oTemp, err := z.ParentHash.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x86)
	if oTemp, err := z.MerkleRoot.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x86)
	o = hsp.AppendInt32(o, z.Version)
	o = append(o, 0x86)
	if oTemp, err := z.Producer.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x86)
	o = hsp.AppendTime(o, z.Timestamp)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Header) Msgsize() (s int) {
	s = 1 + 12 + z.GenesisHash.Msgsize() + 11 + z.ParentHash.Msgsize() + 11 + z.MerkleRoot.Msgsize() + 8 + hsp.Int32Size + 9 + z.Producer.Msgsize() + 10 + hsp.TimeSize
	return
}

// MarshalHash marshals for hash
func (z *SignedHeader) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 4
	o = append(o, 0x84, 0x84)
	if z.Signee == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.Signee.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x84)
	if z.Signature == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.Signature.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x84)
	if oTemp, err := z.Header.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x84)
	if oTemp, err := z.BlockHash.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SignedHeader) Msgsize() (s int) {
	s = 1 + 7
	if z.Signee == nil {
		s += hsp.NilSize
	} else {
		s += z.Signee.Msgsize()
	}
	s += 10
	if z.Signature == nil {
		s += hsp.NilSize
	} else {
		s += z.Signature.Msgsize()
	}
	s += 7 + z.Header.Msgsize() + 10 + z.BlockHash.Msgsize()
	return
}
