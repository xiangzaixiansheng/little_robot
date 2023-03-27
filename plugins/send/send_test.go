package send_util

import (
	"fmt"
	holiday_util "little_robot/plugins/holiday"
	"testing"
	"time"
)

func TestHoliday(t *testing.T) {
	// 假期预告
	isWeekend, days := holiday_util.NewHolidayUtil().GetWeekend(time.Now())

	fmt.Printf("isWeekend %v 距离days %d", isWeekend, days)

	holiday, days2 := holiday_util.NewHolidayUtil().GetNextHoliday()

	fmt.Printf("holiday %v 距离days %d", holiday, days2)

	content := ""
	if isWeekend {
		content += fmt.Sprintf("今天是周末啦! 假期愉快 好好休息哦 \n")
	} else {
		content += fmt.Sprintf("又是新的一天, 距离周末还有 %d 天 ❤️ \n", days)
	}
	if holiday != "" {
		content += fmt.Sprintf("距离最近的法定假期是 %v, 还有 %d 天 ❤️\n", holiday, days)
	}
	fmt.Println("====content", content)
	SendMsg(content)
}
