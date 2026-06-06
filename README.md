# fuu-calendar-extension

福uu 日历 API 扩展，提供考试和讲座的日历订阅服务。

## 接口

### 考试

`/v1/calendar/exams?school_no=教务处学号&password=密码`

示例：
https://api.terraria.ink/fzu/v1/calendar/exams?school_no=教务处学号&password=密码

### 讲座

`/v1/calendar/lectures?school_no=教务处学号&password=密码`

示例：
https://api.terraria.ink/fzu/v1/calendar/lectures?school_no=教务处学号&password=密码

> [!IMPORTANT]
> 若密码含特殊字符请使用[URL编码](https://www.toolhelper.cn/EncodeDecode/Url)

> [!NOTE]
> API 需要传入教务处密码，如对隐私问题有所顾虑，请自行部署。

## 开源许可

项目根据 GPL-3.0 License 的要求进行使用和分发，代码部分来自 https://github.com/renbaoshuo/fzu-ics
