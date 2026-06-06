struct ErrorResp {
    1: required i32 code;
    2: required string msg;
}

struct GetExamsCalendarReq {
    1: required string school_no
    2: required string password
}

struct GetLecturesCalendarReq {
    1: required string school_no
    2: required string password
}

service CalendarService {
    binary GetExamsCalendar(1: GetExamsCalendarReq req)(api.get="/v1/calendar/exams")
    binary GetLecturesCalendar(1: GetLecturesCalendarReq req)(api.get="/v1/calendar/lectures")
}