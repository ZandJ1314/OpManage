package lib

import (
	"opManage/models"
	"time"
)

func DeleteLogs(logintime time.Time){
	_,err := models.LogsDelete(logintime)
	if err != nil{
		NewLog().Error("failed",err)
	}
}

func StartTimer(logintime time.Time) {
	go func() {
		for {
			now := logintime
			next := now.Add(time.Hour*168)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0,0,0,0,next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
			DeleteLogs(logintime)
		}
	}()
}
