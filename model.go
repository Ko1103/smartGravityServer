package model

import "time"
import "strings"
// define weight model
// 

type Weight struct {
	id String //このweightID
	scaleId String //はかりのID
	weight Int //重さ
	createdAt time.Time //作成された時間
	UpdatedAt time.Time //更新された時間
}