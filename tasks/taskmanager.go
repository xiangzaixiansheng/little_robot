package tasks

import (
	"fmt"
	util "little_robot/pkg/utils"
	holiday_util "little_robot/plugins/holiday"
	send_util "little_robot/plugins/send"
	"time"
)

type Job1 struct {
}

func (t Job1) Run() {
	// 假期预告
	isWeekend, days := holiday_util.NewHolidayUtil().GetWeekend(time.Now())

	fmt.Printf("isWeekend %v 距离days %d", isWeekend, days)

	holiday, days2 := holiday_util.NewHolidayUtil().GetNextHoliday()

	fmt.Printf("holiday %v 距离days %d", holiday, days2)

	content := ""
	if isWeekend {
		content += fmt.Sprintf("今天是周末啦! 假期愉快 好好休息哦 ♥️\n")
	} else {
		content += fmt.Sprintf("又是新的一天❤️, 距离周末还有 %d 天 \n", days)
	}
	if holiday != "" {
		content += fmt.Sprintf("距离最近的法定假期是 %v, 还有 %d 天😊 \n", holiday, days2)
	}

	send_util.SendMsg(content)
	util.LogrusObj.Infoln("[INFO]task 假期提醒: %s", content)
}

type Job2 struct {
}

func (t Job2) Run() {
	isWeekend, _ := holiday_util.NewHolidayUtil().GetWeekend(time.Now())
	if isWeekend {
		util.LogrusObj.Infoln("[INFO]task 不进行喝水提醒，因为是周末: ")
		return
	}

	content := fmt.Sprintf("小猪猪 喝水提醒 💧💧💧 \n %v", time.Now().Format("2006-01-02 15:04:05"))
	send_util.SendMsg(content)
	util.LogrusObj.Infoln("[INFO]task 喝水提醒: %s", content)
}

func init() {
	util.LogrusObj.Infoln("[INFO]开始加载定时任务")
	tm := NewTimerTask()
	// 每天8点 假期提醒
	tm.AddTaskByJob("job", "0 8 * * *", Job1{})
	tm.AddTaskByJob("job", "0 13-18 * * *", Job2{})

}
