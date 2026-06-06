package calendar

import (
	"fmt"
	"time"

	"github.com/ACaiCat/fuu-calendar-extention/biz/model/calendar"
	"github.com/ACaiCat/fuu-calendar-extention/pkg/errno"
	"github.com/ACaiCat/fuu-calendar-extention/pkg/util"
	ics "github.com/arran4/golang-ical"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/west2-online/jwch"
)

func (s *Service) GetLecturesCalendar(req *calendar.GetLecturesCalendarReq) ([]byte, error) {
	stu := jwch.NewStudent().WithUser(req.SchoolNo, req.Password)
	err := stu.Login()
	if err != nil {
		return nil, errno.JwchLoginFailedErr.WithError(err)
	}

	calendarIcs, err := buildLecturesCalendar(stu)
	if err != nil {
		hlog.CtxErrorf(s.ctx, "buildLecturesCalendar err: %v", err)
		return nil, err
	}

	return calendarIcs, nil
}

func buildLecturesCalendar(stu *jwch.Student) ([]byte, error) {
	// 初始化
	cstSh, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = cstSh

	// 转换为 ics 格式
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	cal.SetXWRCalName(fmt.Sprintf("福州大学讲座 [%s]", stu.ID))
	cal.SetTimezoneId("Asia/Shanghai")
	cal.SetXWRTimezone("Asia/Shanghai")

	lectures, err := stu.GetLectures()

	if err != nil {
		return nil, err
	}

	for _, lecture := range lectures {
		eventIdBase := fmt.Sprintf("%s__%d_%s_%s",
			lecture.Category, lecture.IssueNumber, lecture.Title, lecture.Speaker)
		description := fmt.Sprintf(
			"主讲人：%s\n"+
				"类别：%s\n"+
				"听取情况：%s\n",
			lecture.Speaker, lecture.Category, lecture.AttendanceStatus)

		event := cal.AddEvent(util.Md5Str(eventIdBase))
		event.SetCreatedTime(time.Now())
		event.SetDtStampTime(time.Now())
		event.SetModifiedAt(time.Now())
		event.SetSummary(lecture.Title)
		event.SetDescription(description)

		// 位置信息
		lat, lon := util.FindGeoLocation(lecture.Location)
		if lat != 0 && lon != 0 {
			event.SetGeo(lat, lon)
		}
		event.SetLocation("[讲座] " + lecture.Title)

		// 开始时间
		startTime := time.UnixMilli(lecture.Timestamp)
		event.SetStartAt(startTime)
		event.SetEndAt(startTime.Add(1 * time.Hour))

		// 提醒
		alarmDescription := "地点: " + lecture.Location + "\n"
		alarm := event.AddAlarm()
		alarm.SetAction(ics.ActionDisplay)
		alarm.SetSummary("[讲座] " + lecture.Title)
		alarm.SetTrigger("-PT15M")
		alarm.SetDescription(alarmDescription)
	}

	calendarContent := cal.Serialize()

	return []byte(calendarContent), nil
}
