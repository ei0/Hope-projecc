package shandle

import (
	Data "data"
	"fmt"
	"snet"

	"github.com/golang/protobuf/proto"
)

type CreateRoom struct {
}

func (C *CreateRoom) Handle(request *snet.Request) {

	Playinfo := &Data.PlayerInfo{}

	proto.Unmarshal(request.GetData(), Playinfo)
	//获取房间ID
	id := snet.RoomMgr.AddRoom(Playinfo.Roomid)

	fmt.Println("[info]房间创建完毕，房间ID为", id, "创建房间的人为：", Playinfo.Uid)

	temp := JoinRoom{}
	temp.Handle(request)

	FlashRoomMsg := FlashRoom{}
	FlashRoomMsg.Handle(request)

	//request.GetConnection().SendMesg(CREAT_ROOM_ACK,[]byte("Successful"))

}
