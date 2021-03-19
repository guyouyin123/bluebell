package sonwflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineId int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineId)
	return
}

// 获取唯一id的方法
func GenID() (id int64) {
	if err := Init("2020-07-01", 1); err != nil {
		fmt.Printf("init failed,err:%V\n", err)
		return
	}
	id = node.Generate().Int64()
	return id
}

//func main() {
//	id, err := GenID()
//	fmt.Println(id)
//	fmt.Println(err)
//}
