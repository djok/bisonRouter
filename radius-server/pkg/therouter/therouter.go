// Code generated by radius-dict-gen. DO NOT EDIT.

package therouter

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_TheRouter_VendorID = 12345
)

func _TheRouter_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_TheRouter_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _TheRouter_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _TheRouter_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				values = append(values, vsa[2:int(vsaLen)])
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _TheRouter_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _TheRouter_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				return vsa[2:int(vsaLen)], true
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _TheRouter_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _TheRouter_VendorID {
			i++
			continue
		}
		for j := 0; len(vsa[j:]) >= 3; {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa[j:]) || vsaLen < 3 {
				i++
				break
			}
			if vsaTyp == typ {
				vsa = append(vsa[:j], vsa[j+int(vsaLen):]...)
			}
			j += int(vsaLen)
		}
		if len(vsa) > 0 {
			copy(avp.Attribute[4:], vsa)
			i++
		} else {
			p.Attributes = append(p.Attributes[:i], p.Attributes[i+i:]...)
		}
	}
	return _TheRouter_AddVendor(p, typ, attr)
}

func _TheRouter_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _TheRouter_VendorID {
			i++
			continue
		}
		offset := 0
		for len(vsa[offset:]) >= 3 {
			vsaTyp, vsaLen := vsa[offset], vsa[offset+1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				continue vsaLoop
			}
			if vsaTyp == typ {
				copy(vsa[offset:], vsa[offset+int(vsaLen):])
				vsa = vsa[:len(vsa)-int(vsaLen)]
			} else {
				offset += int(vsaLen)
			}
		}
		if offset == 0 {
			p.Attributes = append(p.Attributes[:i], p.Attributes[i+1:]...)
		} else {
			i++
		}
	}
	return
}

type TherouterIngressCir uint32

var TherouterIngressCir_Strings = map[TherouterIngressCir]string{}

func (a TherouterIngressCir) String() string {
	if str, ok := TherouterIngressCir_Strings[a]; ok {
		return str
	}
	return "TherouterIngressCir(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterIngressCir_Add(p *radius.Packet, value TherouterIngressCir) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 1, a)
}

func TherouterIngressCir_Get(p *radius.Packet) (value TherouterIngressCir) {
	value, _ = TherouterIngressCir_Lookup(p)
	return
}

func TherouterIngressCir_Gets(p *radius.Packet) (values []TherouterIngressCir, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterIngressCir(i))
	}
	return
}

func TherouterIngressCir_Lookup(p *radius.Packet) (value TherouterIngressCir, err error) {
	a, ok := _TheRouter_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterIngressCir(i)
	return
}

func TherouterIngressCir_Set(p *radius.Packet, value TherouterIngressCir) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 1, a)
}

func TherouterIngressCir_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 1)
}

type TherouterEngressCir uint32

var TherouterEngressCir_Strings = map[TherouterEngressCir]string{}

func (a TherouterEngressCir) String() string {
	if str, ok := TherouterEngressCir_Strings[a]; ok {
		return str
	}
	return "TherouterEngressCir(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterEngressCir_Add(p *radius.Packet, value TherouterEngressCir) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 2, a)
}

func TherouterEngressCir_Get(p *radius.Packet) (value TherouterEngressCir) {
	value, _ = TherouterEngressCir_Lookup(p)
	return
}

func TherouterEngressCir_Gets(p *radius.Packet) (values []TherouterEngressCir, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 2) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterEngressCir(i))
	}
	return
}

func TherouterEngressCir_Lookup(p *radius.Packet) (value TherouterEngressCir, err error) {
	a, ok := _TheRouter_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterEngressCir(i)
	return
}

func TherouterEngressCir_Set(p *radius.Packet, value TherouterEngressCir) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 2, a)
}

func TherouterEngressCir_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 2)
}

type TherouterIpv4Addr uint32

var TherouterIpv4Addr_Strings = map[TherouterIpv4Addr]string{}

func (a TherouterIpv4Addr) String() string {
	if str, ok := TherouterIpv4Addr_Strings[a]; ok {
		return str
	}
	return "TherouterIpv4Addr(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterIpv4Addr_Add(p *radius.Packet, value TherouterIpv4Addr) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 3, a)
}

func TherouterIpv4Addr_Get(p *radius.Packet) (value TherouterIpv4Addr) {
	value, _ = TherouterIpv4Addr_Lookup(p)
	return
}

func TherouterIpv4Addr_Gets(p *radius.Packet) (values []TherouterIpv4Addr, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 3) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterIpv4Addr(i))
	}
	return
}

func TherouterIpv4Addr_Lookup(p *radius.Packet) (value TherouterIpv4Addr, err error) {
	a, ok := _TheRouter_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterIpv4Addr(i)
	return
}

func TherouterIpv4Addr_Set(p *radius.Packet, value TherouterIpv4Addr) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 3, a)
}

func TherouterIpv4Addr_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 3)
}

type TherouterIpv4Mask uint32

var TherouterIpv4Mask_Strings = map[TherouterIpv4Mask]string{}

func (a TherouterIpv4Mask) String() string {
	if str, ok := TherouterIpv4Mask_Strings[a]; ok {
		return str
	}
	return "TherouterIpv4Mask(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterIpv4Mask_Add(p *radius.Packet, value TherouterIpv4Mask) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 4, a)
}

func TherouterIpv4Mask_Get(p *radius.Packet) (value TherouterIpv4Mask) {
	value, _ = TherouterIpv4Mask_Lookup(p)
	return
}

func TherouterIpv4Mask_Gets(p *radius.Packet) (values []TherouterIpv4Mask, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 4) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterIpv4Mask(i))
	}
	return
}

func TherouterIpv4Mask_Lookup(p *radius.Packet) (value TherouterIpv4Mask, err error) {
	a, ok := _TheRouter_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterIpv4Mask(i)
	return
}

func TherouterIpv4Mask_Set(p *radius.Packet, value TherouterIpv4Mask) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 4, a)
}

func TherouterIpv4Mask_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 4)
}

type TherouterOuterVid uint32

var TherouterOuterVid_Strings = map[TherouterOuterVid]string{}

func (a TherouterOuterVid) String() string {
	if str, ok := TherouterOuterVid_Strings[a]; ok {
		return str
	}
	return "TherouterOuterVid(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterOuterVid_Add(p *radius.Packet, value TherouterOuterVid) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 5, a)
}

func TherouterOuterVid_Get(p *radius.Packet) (value TherouterOuterVid) {
	value, _ = TherouterOuterVid_Lookup(p)
	return
}

func TherouterOuterVid_Gets(p *radius.Packet) (values []TherouterOuterVid, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 5) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterOuterVid(i))
	}
	return
}

func TherouterOuterVid_Lookup(p *radius.Packet) (value TherouterOuterVid, err error) {
	a, ok := _TheRouter_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterOuterVid(i)
	return
}

func TherouterOuterVid_Set(p *radius.Packet, value TherouterOuterVid) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 5, a)
}

func TherouterOuterVid_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 5)
}

type TherouterInnerVid uint32

var TherouterInnerVid_Strings = map[TherouterInnerVid]string{}

func (a TherouterInnerVid) String() string {
	if str, ok := TherouterInnerVid_Strings[a]; ok {
		return str
	}
	return "TherouterInnerVid(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterInnerVid_Add(p *radius.Packet, value TherouterInnerVid) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 6, a)
}

func TherouterInnerVid_Get(p *radius.Packet) (value TherouterInnerVid) {
	value, _ = TherouterInnerVid_Lookup(p)
	return
}

func TherouterInnerVid_Gets(p *radius.Packet) (values []TherouterInnerVid, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 6) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterInnerVid(i))
	}
	return
}

func TherouterInnerVid_Lookup(p *radius.Packet) (value TherouterInnerVid, err error) {
	a, ok := _TheRouter_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterInnerVid(i)
	return
}

func TherouterInnerVid_Set(p *radius.Packet, value TherouterInnerVid) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 6, a)
}

func TherouterInnerVid_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 6)
}

type TherouterIPUnnumbered uint32

var TherouterIPUnnumbered_Strings = map[TherouterIPUnnumbered]string{}

func (a TherouterIPUnnumbered) String() string {
	if str, ok := TherouterIPUnnumbered_Strings[a]; ok {
		return str
	}
	return "TherouterIPUnnumbered(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterIPUnnumbered_Add(p *radius.Packet, value TherouterIPUnnumbered) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 7, a)
}

func TherouterIPUnnumbered_Get(p *radius.Packet) (value TherouterIPUnnumbered) {
	value, _ = TherouterIPUnnumbered_Lookup(p)
	return
}

func TherouterIPUnnumbered_Gets(p *radius.Packet) (values []TherouterIPUnnumbered, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 7) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterIPUnnumbered(i))
	}
	return
}

func TherouterIPUnnumbered_Lookup(p *radius.Packet) (value TherouterIPUnnumbered, err error) {
	a, ok := _TheRouter_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterIPUnnumbered(i)
	return
}

func TherouterIPUnnumbered_Set(p *radius.Packet, value TherouterIPUnnumbered) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 7, a)
}

func TherouterIPUnnumbered_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 7)
}

type TherouterPortID uint32

var TherouterPortID_Strings = map[TherouterPortID]string{}

func (a TherouterPortID) String() string {
	if str, ok := TherouterPortID_Strings[a]; ok {
		return str
	}
	return "TherouterPortID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterPortID_Add(p *radius.Packet, value TherouterPortID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 8, a)
}

func TherouterPortID_Get(p *radius.Packet) (value TherouterPortID) {
	value, _ = TherouterPortID_Lookup(p)
	return
}

func TherouterPortID_Gets(p *radius.Packet) (values []TherouterPortID, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 8) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterPortID(i))
	}
	return
}

func TherouterPortID_Lookup(p *radius.Packet) (value TherouterPortID, err error) {
	a, ok := _TheRouter_LookupVendor(p, 8)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterPortID(i)
	return
}

func TherouterPortID_Set(p *radius.Packet, value TherouterPortID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 8, a)
}

func TherouterPortID_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 8)
}

type TherouterIpv4Gw uint32

var TherouterIpv4Gw_Strings = map[TherouterIpv4Gw]string{}

func (a TherouterIpv4Gw) String() string {
	if str, ok := TherouterIpv4Gw_Strings[a]; ok {
		return str
	}
	return "TherouterIpv4Gw(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterIpv4Gw_Add(p *radius.Packet, value TherouterIpv4Gw) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 9, a)
}

func TherouterIpv4Gw_Get(p *radius.Packet) (value TherouterIpv4Gw) {
	value, _ = TherouterIpv4Gw_Lookup(p)
	return
}

func TherouterIpv4Gw_Gets(p *radius.Packet) (values []TherouterIpv4Gw, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 9) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterIpv4Gw(i))
	}
	return
}

func TherouterIpv4Gw_Lookup(p *radius.Packet) (value TherouterIpv4Gw, err error) {
	a, ok := _TheRouter_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterIpv4Gw(i)
	return
}

func TherouterIpv4Gw_Set(p *radius.Packet, value TherouterIpv4Gw) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 9, a)
}

func TherouterIpv4Gw_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 9)
}

type TherouterPbr uint32

var TherouterPbr_Strings = map[TherouterPbr]string{}

func (a TherouterPbr) String() string {
	if str, ok := TherouterPbr_Strings[a]; ok {
		return str
	}
	return "TherouterPbr(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterPbr_Add(p *radius.Packet, value TherouterPbr) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 10, a)
}

func TherouterPbr_Get(p *radius.Packet) (value TherouterPbr) {
	value, _ = TherouterPbr_Lookup(p)
	return
}

func TherouterPbr_Gets(p *radius.Packet) (values []TherouterPbr, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 10) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterPbr(i))
	}
	return
}

func TherouterPbr_Lookup(p *radius.Packet) (value TherouterPbr, err error) {
	a, ok := _TheRouter_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterPbr(i)
	return
}

func TherouterPbr_Set(p *radius.Packet, value TherouterPbr) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 10, a)
}

func TherouterPbr_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 10)
}

type TherouterInstallSubscRoute uint32

var TherouterInstallSubscRoute_Strings = map[TherouterInstallSubscRoute]string{}

func (a TherouterInstallSubscRoute) String() string {
	if str, ok := TherouterInstallSubscRoute_Strings[a]; ok {
		return str
	}
	return "TherouterInstallSubscRoute(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterInstallSubscRoute_Add(p *radius.Packet, value TherouterInstallSubscRoute) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 17, a)
}

func TherouterInstallSubscRoute_Get(p *radius.Packet) (value TherouterInstallSubscRoute) {
	value, _ = TherouterInstallSubscRoute_Lookup(p)
	return
}

func TherouterInstallSubscRoute_Gets(p *radius.Packet) (values []TherouterInstallSubscRoute, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 17) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterInstallSubscRoute(i))
	}
	return
}

func TherouterInstallSubscRoute_Lookup(p *radius.Packet) (value TherouterInstallSubscRoute, err error) {
	a, ok := _TheRouter_LookupVendor(p, 17)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterInstallSubscRoute(i)
	return
}

func TherouterInstallSubscRoute_Set(p *radius.Packet, value TherouterInstallSubscRoute) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 17, a)
}

func TherouterInstallSubscRoute_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 17)
}

type TherouterSubscTTL uint32

var TherouterSubscTTL_Strings = map[TherouterSubscTTL]string{}

func (a TherouterSubscTTL) String() string {
	if str, ok := TherouterSubscTTL_Strings[a]; ok {
		return str
	}
	return "TherouterSubscTTL(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterSubscTTL_Add(p *radius.Packet, value TherouterSubscTTL) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 18, a)
}

func TherouterSubscTTL_Get(p *radius.Packet) (value TherouterSubscTTL) {
	value, _ = TherouterSubscTTL_Lookup(p)
	return
}

func TherouterSubscTTL_Gets(p *radius.Packet) (values []TherouterSubscTTL, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 18) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterSubscTTL(i))
	}
	return
}

func TherouterSubscTTL_Lookup(p *radius.Packet) (value TherouterSubscTTL, err error) {
	a, ok := _TheRouter_LookupVendor(p, 18)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterSubscTTL(i)
	return
}

func TherouterSubscTTL_Set(p *radius.Packet, value TherouterSubscTTL) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 18, a)
}

func TherouterSubscTTL_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 18)
}

type TherouterSubscStaticArp uint32

var TherouterSubscStaticArp_Strings = map[TherouterSubscStaticArp]string{}

func (a TherouterSubscStaticArp) String() string {
	if str, ok := TherouterSubscStaticArp_Strings[a]; ok {
		return str
	}
	return "TherouterSubscStaticArp(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterSubscStaticArp_Add(p *radius.Packet, value TherouterSubscStaticArp) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 19, a)
}

func TherouterSubscStaticArp_Get(p *radius.Packet) (value TherouterSubscStaticArp) {
	value, _ = TherouterSubscStaticArp_Lookup(p)
	return
}

func TherouterSubscStaticArp_Gets(p *radius.Packet) (values []TherouterSubscStaticArp, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 19) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterSubscStaticArp(i))
	}
	return
}

func TherouterSubscStaticArp_Lookup(p *radius.Packet) (value TherouterSubscStaticArp, err error) {
	a, ok := _TheRouter_LookupVendor(p, 19)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterSubscStaticArp(i)
	return
}

func TherouterSubscStaticArp_Set(p *radius.Packet, value TherouterSubscStaticArp) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 19, a)
}

func TherouterSubscStaticArp_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 19)
}

type TherouterSubscProxyArp uint32

var TherouterSubscProxyArp_Strings = map[TherouterSubscProxyArp]string{}

func (a TherouterSubscProxyArp) String() string {
	if str, ok := TherouterSubscProxyArp_Strings[a]; ok {
		return str
	}
	return "TherouterSubscProxyArp(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterSubscProxyArp_Add(p *radius.Packet, value TherouterSubscProxyArp) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 20, a)
}

func TherouterSubscProxyArp_Get(p *radius.Packet) (value TherouterSubscProxyArp) {
	value, _ = TherouterSubscProxyArp_Lookup(p)
	return
}

func TherouterSubscProxyArp_Gets(p *radius.Packet) (values []TherouterSubscProxyArp, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 20) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterSubscProxyArp(i))
	}
	return
}

func TherouterSubscProxyArp_Lookup(p *radius.Packet) (value TherouterSubscProxyArp, err error) {
	a, ok := _TheRouter_LookupVendor(p, 20)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterSubscProxyArp(i)
	return
}

func TherouterSubscProxyArp_Set(p *radius.Packet, value TherouterSubscProxyArp) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 20, a)
}

func TherouterSubscProxyArp_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 20)
}

type TherouterSubscRpFilter uint32

var TherouterSubscRpFilter_Strings = map[TherouterSubscRpFilter]string{}

func (a TherouterSubscRpFilter) String() string {
	if str, ok := TherouterSubscRpFilter_Strings[a]; ok {
		return str
	}
	return "TherouterSubscRpFilter(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterSubscRpFilter_Add(p *radius.Packet, value TherouterSubscRpFilter) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 21, a)
}

func TherouterSubscRpFilter_Get(p *radius.Packet) (value TherouterSubscRpFilter) {
	value, _ = TherouterSubscRpFilter_Lookup(p)
	return
}

func TherouterSubscRpFilter_Gets(p *radius.Packet) (values []TherouterSubscRpFilter, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 21) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterSubscRpFilter(i))
	}
	return
}

func TherouterSubscRpFilter_Lookup(p *radius.Packet) (value TherouterSubscRpFilter, err error) {
	a, ok := _TheRouter_LookupVendor(p, 21)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterSubscRpFilter(i)
	return
}

func TherouterSubscRpFilter_Set(p *radius.Packet, value TherouterSubscRpFilter) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 21, a)
}

func TherouterSubscRpFilter_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 21)
}

type TherouterShaperType uint32

var TherouterShaperType_Strings = map[TherouterShaperType]string{}

func (a TherouterShaperType) String() string {
	if str, ok := TherouterShaperType_Strings[a]; ok {
		return str
	}
	return "TherouterShaperType(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterShaperType_Add(p *radius.Packet, value TherouterShaperType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 22, a)
}

func TherouterShaperType_Get(p *radius.Packet) (value TherouterShaperType) {
	value, _ = TherouterShaperType_Lookup(p)
	return
}

func TherouterShaperType_Gets(p *radius.Packet) (values []TherouterShaperType, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 22) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterShaperType(i))
	}
	return
}

func TherouterShaperType_Lookup(p *radius.Packet) (value TherouterShaperType, err error) {
	a, ok := _TheRouter_LookupVendor(p, 22)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterShaperType(i)
	return
}

func TherouterShaperType_Set(p *radius.Packet, value TherouterShaperType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 22, a)
}

func TherouterShaperType_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 22)
}

func TherouterShaperIngressParams_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _TheRouter_AddVendor(p, 23, a)
}

func TherouterShaperIngressParams_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _TheRouter_AddVendor(p, 23, a)
}

func TherouterShaperIngressParams_Get(p *radius.Packet) (value []byte) {
	value, _ = TherouterShaperIngressParams_Lookup(p)
	return
}

func TherouterShaperIngressParams_GetString(p *radius.Packet) (value string) {
	value, _ = TherouterShaperIngressParams_LookupString(p)
	return
}

func TherouterShaperIngressParams_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _TheRouter_GetsVendor(p, 23) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TherouterShaperIngressParams_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _TheRouter_GetsVendor(p, 23) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TherouterShaperIngressParams_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _TheRouter_LookupVendor(p, 23)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TherouterShaperIngressParams_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _TheRouter_LookupVendor(p, 23)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TherouterShaperIngressParams_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _TheRouter_SetVendor(p, 23, a)
}

func TherouterShaperIngressParams_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _TheRouter_SetVendor(p, 23, a)
}

func TherouterShaperIngressParams_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 23)
}

func TherouterShaperEgressParams_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _TheRouter_AddVendor(p, 24, a)
}

func TherouterShaperEgressParams_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _TheRouter_AddVendor(p, 24, a)
}

func TherouterShaperEgressParams_Get(p *radius.Packet) (value []byte) {
	value, _ = TherouterShaperEgressParams_Lookup(p)
	return
}

func TherouterShaperEgressParams_GetString(p *radius.Packet) (value string) {
	value, _ = TherouterShaperEgressParams_LookupString(p)
	return
}

func TherouterShaperEgressParams_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _TheRouter_GetsVendor(p, 24) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TherouterShaperEgressParams_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _TheRouter_GetsVendor(p, 24) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TherouterShaperEgressParams_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _TheRouter_LookupVendor(p, 24)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TherouterShaperEgressParams_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _TheRouter_LookupVendor(p, 24)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TherouterShaperEgressParams_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _TheRouter_SetVendor(p, 24, a)
}

func TherouterShaperEgressParams_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _TheRouter_SetVendor(p, 24, a)
}

func TherouterShaperEgressParams_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 24)
}

type TherouterSubscAddrPrefixMapID uint32

var TherouterSubscAddrPrefixMapID_Strings = map[TherouterSubscAddrPrefixMapID]string{}

func (a TherouterSubscAddrPrefixMapID) String() string {
	if str, ok := TherouterSubscAddrPrefixMapID_Strings[a]; ok {
		return str
	}
	return "TherouterSubscAddrPrefixMapID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterSubscAddrPrefixMapID_Add(p *radius.Packet, value TherouterSubscAddrPrefixMapID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 25, a)
}

func TherouterSubscAddrPrefixMapID_Get(p *radius.Packet) (value TherouterSubscAddrPrefixMapID) {
	value, _ = TherouterSubscAddrPrefixMapID_Lookup(p)
	return
}

func TherouterSubscAddrPrefixMapID_Gets(p *radius.Packet) (values []TherouterSubscAddrPrefixMapID, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 25) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterSubscAddrPrefixMapID(i))
	}
	return
}

func TherouterSubscAddrPrefixMapID_Lookup(p *radius.Packet) (value TherouterSubscAddrPrefixMapID, err error) {
	a, ok := _TheRouter_LookupVendor(p, 25)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterSubscAddrPrefixMapID(i)
	return
}

func TherouterSubscAddrPrefixMapID_Set(p *radius.Packet, value TherouterSubscAddrPrefixMapID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 25, a)
}

func TherouterSubscAddrPrefixMapID_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 25)
}

type TherouterSubscAddrPrefixMapValue uint32

var TherouterSubscAddrPrefixMapValue_Strings = map[TherouterSubscAddrPrefixMapValue]string{}

func (a TherouterSubscAddrPrefixMapValue) String() string {
	if str, ok := TherouterSubscAddrPrefixMapValue_Strings[a]; ok {
		return str
	}
	return "TherouterSubscAddrPrefixMapValue(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterSubscAddrPrefixMapValue_Add(p *radius.Packet, value TherouterSubscAddrPrefixMapValue) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 26, a)
}

func TherouterSubscAddrPrefixMapValue_Get(p *radius.Packet) (value TherouterSubscAddrPrefixMapValue) {
	value, _ = TherouterSubscAddrPrefixMapValue_Lookup(p)
	return
}

func TherouterSubscAddrPrefixMapValue_Gets(p *radius.Packet) (values []TherouterSubscAddrPrefixMapValue, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 26) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterSubscAddrPrefixMapValue(i))
	}
	return
}

func TherouterSubscAddrPrefixMapValue_Lookup(p *radius.Packet) (value TherouterSubscAddrPrefixMapValue, err error) {
	a, ok := _TheRouter_LookupVendor(p, 26)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterSubscAddrPrefixMapValue(i)
	return
}

func TherouterSubscAddrPrefixMapValue_Set(p *radius.Packet, value TherouterSubscAddrPrefixMapValue) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 26, a)
}

func TherouterSubscAddrPrefixMapValue_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 26)
}

type TherouterSubscNatConnectionLimit uint32

var TherouterSubscNatConnectionLimit_Strings = map[TherouterSubscNatConnectionLimit]string{}

func (a TherouterSubscNatConnectionLimit) String() string {
	if str, ok := TherouterSubscNatConnectionLimit_Strings[a]; ok {
		return str
	}
	return "TherouterSubscNatConnectionLimit(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func TherouterSubscNatConnectionLimit_Add(p *radius.Packet, value TherouterSubscNatConnectionLimit) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_AddVendor(p, 27, a)
}

func TherouterSubscNatConnectionLimit_Get(p *radius.Packet) (value TherouterSubscNatConnectionLimit) {
	value, _ = TherouterSubscNatConnectionLimit_Lookup(p)
	return
}

func TherouterSubscNatConnectionLimit_Gets(p *radius.Packet) (values []TherouterSubscNatConnectionLimit, err error) {
	var i uint32
	for _, attr := range _TheRouter_GetsVendor(p, 27) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, TherouterSubscNatConnectionLimit(i))
	}
	return
}

func TherouterSubscNatConnectionLimit_Lookup(p *radius.Packet) (value TherouterSubscNatConnectionLimit, err error) {
	a, ok := _TheRouter_LookupVendor(p, 27)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = TherouterSubscNatConnectionLimit(i)
	return
}

func TherouterSubscNatConnectionLimit_Set(p *radius.Packet, value TherouterSubscNatConnectionLimit) (err error) {
	a := radius.NewInteger(uint32(value))
	return _TheRouter_SetVendor(p, 27, a)
}

func TherouterSubscNatConnectionLimit_Del(p *radius.Packet) {
	_TheRouter_DelVendor(p, 27)
}
