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

func (s *Service) GetExamsCalendar(req *calendar.GetExamsCalendarReq) ([]byte, error) {
	stu := jwch.NewStudent().WithUser(req.SchoolNo, req.Password)
	err := stu.Login()
	if err != nil {
		return nil, errno.JwchLoginFailedErr.WithError(err)
	}

	calendarIcs, err := buildExamsCalendar(stu)
	if err != nil {
		hlog.CtxErrorf(s.ctx, "buildExamsCalendar err: %v", err)
		return nil, err
	}

	return calendarIcs, nil
}

func buildExamsCalendar(stu *jwch.Student) ([]byte, error) {
	// 初始化
	cstSh, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = cstSh

	// 转换为 ics 格式
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	cal.SetXWRCalName(fmt.Sprintf("福州大学考试 [%s]", stu.ID))
	cal.SetTimezoneId("Asia/Shanghai")
	cal.SetXWRTimezone("Asia/Shanghai")

	schoolCalendar, err := stu.GetSchoolCalendar()
	if err != nil {
		return nil, err
	}

	exams, err := stu.GetExamRoom(jwch.ExamRoomReq{Term: schoolCalendar.CurrentTerm})

	if err != nil {
		return nil, err
	}

	for _, exam := range exams {
		if exam.Time == "" {
			continue
		}

		eventIdBase := fmt.Sprintf("%s__%s_%s",
			exam.CourseName, exam.Teacher, exam.Time)
		description := fmt.Sprintf(
			"任课教师：%s\n"+
				"学分：%d"+
				exam.Teacher, exam.Credit)

		event := cal.AddEvent(util.Md5Str(eventIdBase))
		event.SetCreatedTime(time.Now())
		event.SetDtStampTime(time.Now())
		event.SetModifiedAt(time.Now())
		event.SetSummary("[考试] " + exam.CourseName)
		event.SetDescription(description)

		// 位置信息
		lat, lon := util.FindGeoLocation(exam.Location)
		if lat != 0 && lon != 0 {
			event.SetGeo(lat, lon)
		}
		event.SetLocation(exam.Location)

		startTime, endTime, err := util.ParseExamTime(exam.Time)

		if err != nil {
			return nil, err
		}

		// 开始时间
		event.SetStartAt(startTime)
		event.SetEndAt(endTime)

		// 提醒
		alarmDescription := "地点: " + exam.Location + "\n"
		alarm := event.AddAlarm()
		alarm.SetAction(ics.ActionDisplay)
		alarm.SetSummary("[考试] " + exam.CourseName)
		alarm.SetTrigger("-PT60M")
		alarm.SetDescription(alarmDescription)
	}

	calendarContent := cal.Serialize()

	return []byte(calendarContent), nil
}
