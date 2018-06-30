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
	updatedAt time.Time //更新された時間
}

type Minimum struct {
	id String //このオブジェクトのID
	minimumWeight Int //はかりの下限値、これを下回った場合は購入を促す。
	createdAt time.Time //作成された時間
	updatedAt time.Time //更新された時間
}