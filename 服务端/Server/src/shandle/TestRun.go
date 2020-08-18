package shandle

import (
	"fmt"
	"snet"
)

type TestRun struct {
}

func (S *TestRun) Handle(request *snet.Request) {
	fmt.Println("删除房间")
	ALLROOM := snet.RoomMgr.AllRoom

	for k, v := range ALLROOM {
		v.Lock()
		delete(ALLROOM, k)
		v.Unlock()
	}
}
