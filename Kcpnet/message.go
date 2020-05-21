package Kcpnet

// add by stefan

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/Peakchen/xgameCommon/ado/dbStatistics"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_HeartBeat"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Login"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Server"
	"github.com/xtaci/kcp-go"
)

type TMessageProc struct {
	proc       reflect.Value
	paramTypes []reflect.Type
}

var (
	_MessageTab      = map[uint32]*TMessageProc{}
	_specialLoginMsg = map[uint16]bool{}
)

func RegisterMessage(mainID, subID uint16, proc interface{}) {
	_cmd := EncodeCmd(mainID, subID)
	_, ok := _MessageTab[_cmd]
	if ok {
		return
	}

	cbref := reflect.TypeOf(proc)
	if cbref.Kind() != reflect.Func {
		akLog.FmtPrintln("proc type not is func, but is: %v.", cbref.Kind())
		return
	}

	if cbref.NumIn() != 2 {
		akLog.FmtPrintln("proc num input is not 2, but is: %v.", cbref.NumIn())
		return
	}

	if cbref.NumOut() != 2 {
		akLog.FmtPrintln("proc num output is not 2, but is: %v.", cbref.NumOut())
		return
	}

	if cbref.Out(0) != reflect.TypeOf(bool(false)) {
		akLog.FmtPrintln("proc num out 1 is not string, but is: %v.", cbref.Out(0))
		return
	}

	if cbref.Out(1).Name() != "error" {
		akLog.FmtPrintln("proc num out 2 is not *proto.Message, but is: %v.", cbref.Out(1), reflect.TypeOf(error(nil)), errors.New("0"), fmt.Errorf("0"))
		return
	}

	paramtypes := []reflect.Type{}
	for i := 0; i < cbref.NumIn(); i++ {
		t := cbref.In(i)
		// if t.Kind() == reflect.String ||
		// 	t.Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
		// 	paramtypes = append(paramtypes, t)
		// }
		paramtypes = append(paramtypes, t)
	}

	_MessageTab[_cmd] = &TMessageProc{
		proc:       reflect.ValueOf(proc),
		paramTypes: paramtypes,
	}

	return
}

func GetMessageInfo(mainID, subID uint16) (proc *TMessageProc, finded bool) {
	_cmd := EncodeCmd(mainID, subID)
	proc, finded = _MessageTab[_cmd]
	return
}

func GetAllMessageIDs() (msgs []uint32) {
	msgs = []uint32{}
	for msgid, _ := range _MessageTab {
		msgs = append(msgs, uint32(msgid))
	}
	return
}

func RegisterMessageRet(session TcpSession) (succ bool, err error) {
	rsp := &MSG_Server.SC_ServerRegister_Rsp{}
	rsp.Ret = MSG_Server.ErrorCode_Success
	rsp.Identify = session.GetModuleName()
	return session.SendSvrClientMsg(uint16(MSG_MainModule.MAINMSG_SERVER), uint16(MSG_Server.SUBMSG_SC_ServerRegister), rsp)
}

func SpecialLoginMsgFilter(main, sub uint16) (ok bool) {
	if main != uint16(MSG_MainModule.MAINMSG_LOGIN) {
		return
	}

	if sub == uint16(MSG_Login.SUBMSG_CS_UserRegister) ||
		sub == uint16(MSG_Login.SUBMSG_CS_Login) {
		ok = true
	}

	return
}

func sendHeartBeat(session TcpSession) (succ bool, err error) {
	if !session.Alive() {
		err = fmt.Errorf("session heartbeat disconnection, can not send.")
		return
	}
	rsp := &MSG_HeartBeat.CS_HeartBeat_Req{}
	return session.SendInnerSvrMsg(uint16(MSG_MainModule.MAINMSG_HEARTBEAT), uint16(MSG_HeartBeat.SUBMSG_CS_HeartBeat), rsp)
}

func ResponseHeartBeat(session TcpSession) (succ bool, err error) {
	if !session.Alive() {
		err = fmt.Errorf("session heartbeat disconnection, can not response.")
		return
	}
	rsp := &MSG_HeartBeat.SC_HeartBeat_Rsp{}
	return session.SendInnerSvrMsg(uint16(MSG_MainModule.MAINMSG_HEARTBEAT), uint16(MSG_HeartBeat.SUBMSG_SC_HeartBeat), rsp)
}

func checkHeartBeatRet(pack IMessagePack) (exist bool) {
	mainID, subID := pack.GetMessageID()
	if mainID == uint16(MSG_MainModule.MAINMSG_HEARTBEAT) &&
		uint16(MSG_HeartBeat.SUBMSG_SC_HeartBeat) == subID {
		//akLog.FmtPrintf("<heart beat> RemoteAddr: %v.", pack.GetRemoteAddr())
		exist = true
	}
	return
}

//receive logic message call.
func msgCallBack(sessionobj TcpSession) (succ bool) {
	protocolPack := sessionobj.GetPack()
	msg, cb, unpackerr, exist := protocolPack.UnPackData()
	if unpackerr != nil || !exist {
		akLog.FmtPrintf("unpack data, ModuleName: %v, reg point: %v, err: %v.", sessionobj.GetModuleName(), sessionobj.GetRegPoint(), unpackerr)
		return
	}

	// record db operation stack log.
	mainid, subid := protocolPack.GetMessageID()
	sessionobj.RefreshHeartBeat(mainid, subid)
	identify := protocolPack.GetIdentify()
	dbStatistics.DBMsgStatistics(identify, mainid, subid)

	params := []reflect.Value{
		reflect.ValueOf(sessionobj),
		reflect.ValueOf(msg),
	}

	ret := cb.Call(params)
	succ = ret[0].Interface().(bool)
	reterr := ret[1].Interface()
	if reterr != nil || !succ {
		akLog.FmtPrintln("[client] message return err: ", reterr.(error).Error())
	}

	return
}

/*
@func: UnPackExternalMsg 解服务器外部消息（客户端，clientsession 注册消息）
@parma1: 连接对象 c *kcp.UDPSession
@param2: 解包对象 pack IMessagePack
*/
func UnPackExternalMsg(c *kcp.UDPSession, pack IMessagePack) (succ bool) {
	packLenBuf := make([]byte, EnMessage_NoDataLen)
	readn, err := io.ReadFull(c, packLenBuf)
	if err != nil || readn < EnMessage_NoDataLen {
		if err.Error() == "EOF" {
			succ = true
		} else {
			akLog.FmtPrintln("pack External msg read data fail, err: ", err, readn)
		}
		return
	}

	akLog.FmtPrintln("identify is empty, read data: ", len(packLenBuf))
	packlen := binary.LittleEndian.Uint32(packLenBuf[EnMessage_DataPackLen:EnMessage_NoDataLen])
	if packlen > maxMessageSize {
		akLog.FmtPrintln("error receiving packLen:", packlen)
		return
	}

	data := make([]byte, EnMessage_NoDataLen+packlen)
	readn, err = io.ReadFull(c, data[EnMessage_NoDataLen:])
	if err != nil || readn < int(packlen) {
		akLog.FmtPrintln("error receiving msg, readn:", readn, "packLen:", packlen, "reason:", err)
		return
	}

	//todo: unpack message then read real date.
	copy(data[:EnMessage_NoDataLen], packLenBuf[:])
	_, err = pack.UnPackMsg4Client(data)
	if err != nil {
		akLog.FmtPrintln("unpack action err: ", err)
		return
	}

	succ = true
	return
}

/*
@func: UnPackInnerMsg 解服务器内部消息（server 间客户端发来的请求或者其他rpc消息传递）
@parma1: 连接对象 c *kcp.UDPSession
@param2: 解包对象 pack IMessagePack
*/
func UnPackInnerMsg(c *kcp.UDPSession, pack IMessagePack) (succ bool) {
	packLenBuf := make([]byte, EnMessage_SvrNoDataLen)
	readn, err := io.ReadFull(c, packLenBuf)
	if err != nil || readn < EnMessage_SvrNoDataLen {
		if err.Error() == "EOF" {
			succ = true
		} else {
			akLog.FmtPrintln("pack Inner message read data fail, err: ", err, readn)
		}
		return
	}

	//akLog.FmtPrintln("identify not empty, read data: ", len(packLenBuf))
	packlen := binary.LittleEndian.Uint32(packLenBuf[EnMessage_SvrDataPackLen:EnMessage_SvrNoDataLen])
	if packlen > maxMessageSize {
		akLog.FmtPrintln("error receiving packLen:", packlen)
		return
	}

	data := make([]byte, EnMessage_SvrNoDataLen+packlen)
	readn, err = io.ReadFull(c, data[EnMessage_SvrNoDataLen:])
	if err != nil || readn < int(packlen) {
		akLog.FmtPrintln("error receiving msg, readn:", readn, "packLen:", packlen, "reason:", err)
		return
	}

	//todo: unpack message then read real date.
	copy(data[:EnMessage_SvrNoDataLen], packLenBuf[:])
	_, err = pack.UnPackMsg4Svr(data)
	if err != nil {
		akLog.FmtPrintln("unpack action err: ", err)
		return
	}
	succ = true
	return
}

/*
	内网关路由 inner gateway for message route (request and response).
*/
func innerMsgRouteAct(pointType ESessionType, route Define.ERouteId, mainID uint16, data []byte) (succ bool) {
	var (
		session TcpSession
	)

	if mainID == uint16(MSG_MainModule.MAINMSG_RPC) {
		//game rpc call back.
		akLog.FmtPrintln("inner game rpc route.")
		session = GServer2ServerSession.GetSession(Define.ERouteId_ER_Game)
	} else {
		if route != 0 && pointType == ESessionType_Client {
			//内网转发外网路由请求至xxx服务器 gateway route external message to some one server.
			//akLog.FmtPrintf("inner route requst message, route: %v.", route)
			session = GServer2ServerSession.GetSession(Define.ERouteId(route))
		} else {
			// 内网转发xxx服务器消息至外网 gateway route some one server message to external gateway.
			//akLog.FmtPrintln("inner route respnse message.")
			session = GServer2ServerSession.GetSession(Define.ERouteId_ER_ESG)
		}
	}

	if session == nil {
		akLog.Error("can not find session from inner gateway, mainID: %v.", mainID)
		return
	}

	if !session.Alive() {
		GServer2ServerSession.RemoveSession(session.GetIdentify())
	} else {
		succ = session.WriteMessage(data)
	}
	return
}

// send message for server by inner gateway from external gateway.
func sendInnerSvr(obj TcpSession) (succ bool) {
	session := GServer2ServerSession.GetSession(Define.ERouteId_ER_ISG)
	if session == nil {
		akLog.Error("[request] can not find session inner route from external gateway.")
		return
	}

	if !session.Alive() {
		GServer2ServerSession.RemoveSession(session.GetRemoteAddr())
		akLog.FmtPrintln("s2s session not alive, addr: ", session.GetRemoteAddr())
		return
	}

	out := make([]byte, EnMessage_SvrNoDataLen+int(obj.GetPack().GetDataLen()))
	err := obj.GetPack().PackAction(out)
	if err != nil {
		akLog.Error("unpack action err: ", err)
		return
	}

	succ = session.WriteMessage(out)
	return
}

// send message for user from external gateway.
func sendUserClient(obj TcpSession) (succ bool) {
	out := make([]byte, EnMessage_NoDataLen+int(obj.GetPack().GetDataLen()))
	err := obj.GetPack().PackAction4Client(out)
	if err != nil {
		akLog.Error("[response user client] unpack action err: ", err)
		return
	}

	akLog.FmtPrintln("external response, addr: ", obj.GetPack().GetRemoteAddr(), len(obj.GetPack().GetRemoteAddr()))
	if obj.GetPack().GetPostType() == MsgPostType_Single {
		session := GClient2ServerSession.GetSessionByIdentify(obj.GetPack().GetRemoteAddr())
		if session == nil {
			akLog.Error("[response user client] can not find session route from external gateway.")
			return
		}

		if !session.Alive() {
			GClient2ServerSession.RemoveSession(session.GetRemoteAddr())
			akLog.FmtPrintln("c2s session not alive, addr: ", session.GetRemoteAddr())
			return
		}
		succ = session.WriteMessage(out)
	} else {
		allsession := GClient2ServerSession.GetAllSession()
		allsession.Range(func(key, value interface{}) bool {
			if value != nil {
				sess := value.(TcpSession)
				sess.WriteMessage(out)
			}
			return true
		})
	}

	return
}

/*
	外网关路由 external gateway for message route (request and response).
*/
func externalRouteAct(route Define.ERouteId, obj TcpSession, responseCliented bool) (succ bool) {
	//客户端请求消息 receive user client message.
	if Define.ERouteId(route) != Define.ERouteId_ER_ISG && false == responseCliented {
		akLog.FmtPrintf("external request, route: %v, StrIdentify: %v.", route, obj.GetIdentify())
		// add session.
		GClient2ServerSession.AddSession(obj.GetRemoteAddr(), obj)
		//内网关转发至相关服务器 route message to some one server.
		return sendInnerSvr(obj)
	}

	//外网回复客户端消息 external gateway response user client message.
	return sendUserClient(obj)
}