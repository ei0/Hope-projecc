package shandle

import (
	Data "data"
	"snet"

	"github.com/golang/protobuf/proto"
)

//准备游戏的readygame

type ReadyGame struct {
}

func (R *ReadyGame) Handle(request *snet.Request) {
	//获得玩家信息，如果如果该玩家准备的话，那么Username处给所有房间内的人广播Ready，客户端进行一次切换状态即可。
	PlayerInfoMsg := &Data.PlayerInfo{
		Roomid:   0,
		Uid:      0,
		Username: "",
	}
	proto.Unmarshal(request.GetData(), PlayerInfoMsg)

	NowRoom := snet.RoomMgr.GetRoom(PlayerInfoMsg.Roomid)

	PlayerInfoMsg.Username = "Ready"
	buf, _ := proto.Marshal(PlayerInfoMsg)
	Msg := snet.NewMsgPackage(PLAYER_READY_ACK, buf)
	NowRoom.BroadRoom(Msg)
}
