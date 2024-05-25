package model

import (
	"fmt"
	"strconv"
)

func init() {
	fmt.Println("Struct Todo initialized")
}

type Todo struct {
	Id   int64
	Item string
}

func (t Todo) PrintInfo() string {
	//PARSE INT64 TO STRING
	sId := strconv.FormatInt(t.Id, 10)

	return "\n Id: " + sId + " Item: " + t.Item
}
