package pt

type OnStatusDataPacket struct {
}

func (pkt *OnStatusDataPacket) Decode(data []uint8) (err error) {
	return
}
func (pkt *OnStatusDataPacket) Encode() (data []uint8) {
	return
}
func (pkt *OnStatusDataPacket) GetMessageType() uint8 {
	return RTMP_MSG_AMF0DataMessage
}
func (pkt *OnStatusDataPacket) GetPreferCsId() uint32 {
	return RTMP_CID_OverStream
}