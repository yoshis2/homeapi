package util

import "time"

// JapaneseNowTime は日本の現在時間を取得
func JapaneseNowTime() (time.Time, error) {
	// Time Zone を日本時間に設定
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := time.Now().UTC().In(jst)

	// 時間フォーマット yyyy-mm-dd
	japaneseTime := nowJST.Format("2006-01-02 15:04:05")
	return ParseTime(japaneseTime)
}

//ParseTime 日付文字列からTimeを返す
func ParseTime(datetime string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", datetime)
}
