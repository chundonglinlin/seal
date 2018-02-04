package co

import (
	"UtilsTools/identify_panic"
	"log"
	"seal/rtmp/pt"
)

func (rc *RtmpConn) msgAbort(msg *pt.Message) (err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err, ",panic at ", identify_panic.IdentifyPanic())
		}
	}()

	log.Println("MsgAbort")

	if nil == msg {
		return
	}
	
	return
}
