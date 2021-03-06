package wscli

import (
	"go-snake/akmessage"
	"go-snake/common"
	"go-snake/common/messageBase"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"

	"github.com/gorilla/websocket"
)

const (
	WS_CLOSE      = uint32(0)
	WS_CONNECTING = uint32(1)
	WS_CONNECTED  = uint32(2)
)

type WsNet struct {
	host     string
	c        *websocket.Conn
	fnCh     chan func()
	sendCh   chan []byte
	status   uint32
	lostData chan []byte
}

func NewClient(host string, modelFns func(*WsNet)) {
	ws := &WsNet{host: host,
		sendCh:   make(chan []byte, 100),
		lostData: make(chan []byte, 10)}

	err := ws.dail()
	if err != nil {
		akLog.Fail("dial fail:", err)
		return
	}

	common.DosafeRoutine(func() { ws.readloop() }, nil)
	common.DosafeRoutine(func() { ws.writeloop(modelFns) }, nil)
	common.DosafeRoutine(func() { ws.checkconnect() }, nil)
	common.DosafeRoutine(func() { ws.heartbeat() }, nil)
}

func (this *WsNet) dail() error {
	this.SetConnecting()

	u := url.URL{Scheme: "ws", Host: this.host, Path: "/echo"}
	akLog.FmtPrintf("connecting to %s", u.String())

	dstc, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	this.c = dstc
	this.SetConnected()
	return nil
}

func (this *WsNet) close() {
	if this.GetStatus() == WS_CLOSE {
		return
	}
	this.c.Close()
	this.SetClose()
}

func (this *WsNet) readloop() {
	for {
		select {
		case fn := <-this.fnCh:
			fn()
		default:
			if this.GetStatus() == WS_CONNECTED {

				this.c.SetReadDeadline(time.Now().Add(40 * time.Second))

				_, data, err := this.c.ReadMessage()
				if err != nil {
					akLog.Error("read:", err)
					this.close()
					continue
				}
				akLog.FmtPrintln("recv: ", string(data))
			}

		}
	}
}

func (this *WsNet) writeloop(modelFns func(*WsNet)) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			akLog.FmtPrintln("send test, status: ", this.GetStatus())
			if this.GetStatus() == WS_CONNECTED {
				modelFns(this)
			}
		case data := <-this.sendCh:
			err := this.c.WriteMessage(websocket.BinaryMessage, data)
			if err != nil {
				akLog.Error("write fail: ", err)
				this.close()
				if len(this.lostData) < cap(this.lostData) {
					this.lostData <- data
				}
			}
		}
	}
}

func (this *WsNet) checkconnect() {
	tick := time.NewTicker(3 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			akLog.FmtPrintln("checkconnect status: ", this.GetStatus())
			if this.GetStatus() == WS_CLOSE || this.GetStatus() == WS_CONNECTING {
				this.dail()
			}
			if this.GetStatus() == WS_CONNECTED {
				if len(this.lostData) > 0 {
					for d := range this.lostData {
						this.sendCh <- d
					}
				}
			}
		}
	}
}

func (this *WsNet) heartbeat() {
	tick := time.NewTicker(5 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			if this.GetStatus() == WS_CONNECTED {
				this.sendHeartBeatMsg()
			}
		}
	}
}

func (this *WsNet) sendHeartBeatMsg() {
	hb := &akmessage.CS_HeartBeat{}
	data := messageBase.CSPackMsg_pb(akmessage.MSG_CS_HEARTBEAT, hb)
	if data == nil {
		akLog.Error("pack heart beat msg fail.")
		return
	}
	akLog.FmtPrintf("heart beat, id: %v.", uint32(akmessage.MSG_CS_HEARTBEAT))
	this.SendMsg(data)
}

func (this *WsNet) SendMsg(data []byte) {
	if len(this.sendCh) >= cap(this.sendCh) {
		this.lostData <- data
	} else {
		this.sendCh <- data
	}
}

func (this *WsNet) SetClose() {
	atomic.StoreUint32(&this.status, WS_CLOSE)
}

func (this *WsNet) SetConnected() {
	atomic.StoreUint32(&this.status, WS_CONNECTED)
}

func (this *WsNet) SetConnecting() {
	atomic.StoreUint32(&this.status, WS_CONNECTING)
}

func (this *WsNet) GetStatus() uint32 {
	return atomic.LoadUint32(&this.status)
}
