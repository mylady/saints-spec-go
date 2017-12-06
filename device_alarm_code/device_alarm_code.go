package device_alarm_code

const (
	Unknown           = -1
	AreaAlarm         = 0
	MotionDetect      = 1
	VideoLost         = 2
	VideoShelter      = 3
	SoundCheck        = 4
	DiskFull          = 5
	DiskFault         = 6
	SoundLost         = 7
	DeviceAbnormal    = 8
	AreaIntrusion     = 9
	LimitHeight       = 10
	Wander            = 11
	Raiseup           = 12
	TargetLost        = 13
	AbnormalBehavior  = 14
	ObjectLeft        = 15
	Stay              = 16
	Timeout           = 17
	InvalidOperation  = 18
	ThreatOperation   = 19
	Remove            = 20
	Offline           = 21
	PowerUnstable     = 22
	Disturb           = 23
	DeviceLost        = 24
	Bypass            = 25
	Inactive          = 26
	RecordStop        = 27
	PowerAbnormal     = 28
	VideoQuality      = 29
	DutyCheck         = 30
	FierceMotion      = 31
	DeviceFault       = 32
	AudioLost         = 33
	FanFault          = 34
	TemperatureAlarm  = 35
	HumidityAlarm     = 36
	DoorPointAlarm    = 37
	ForceOperation    = 38
	SoundAlarm        = 39
	LowBattery        = 40
	VideoAnalyseAlarm = 41
	FaceDefenceAlarm  = 42
)

var AlarmCodeMap = make(map[int]string)

func init() {
	AlarmCodeMap[Unknown] = "未知类型"
	AlarmCodeMap[AreaAlarm] = "防区报警"
	AlarmCodeMap[MotionDetect] = "移动侦测"
	AlarmCodeMap[VideoLost] = "视频丢失"
	AlarmCodeMap[VideoShelter] = "视频遮挡"
	AlarmCodeMap[SoundCheck] = "声音检测"
	AlarmCodeMap[DiskFull] = "硬盘满"
	AlarmCodeMap[DiskFault] = "硬盘故障"
	AlarmCodeMap[SoundLost] = "声音丢失"
	AlarmCodeMap[DeviceAbnormal] = "设备异常"
	AlarmCodeMap[AreaIntrusion] = "区域入侵"
	AlarmCodeMap[LimitHeight] = "限高"
	AlarmCodeMap[Wander] = "徘徊"
	AlarmCodeMap[Raiseup] = "起身"
	AlarmCodeMap[TargetLost] = "目标丢失"
	AlarmCodeMap[AbnormalBehavior] = "异常行为"
	AlarmCodeMap[ObjectLeft] = "物品遗留"
	AlarmCodeMap[Stay] = "停留"
	AlarmCodeMap[Timeout] = "超时"
	AlarmCodeMap[InvalidOperation] = "非法操作"
	AlarmCodeMap[ThreatOperation] = "胁迫操作"
	AlarmCodeMap[Remove] = "移除"
	AlarmCodeMap[Offline] = "离线"
	AlarmCodeMap[PowerUnstable] = "电源不稳定"
	AlarmCodeMap[Disturb] = "干扰"
	AlarmCodeMap[DeviceLost] = "丢失"
	AlarmCodeMap[Bypass] = "旁路"
	AlarmCodeMap[Inactive] = "失效"
	AlarmCodeMap[RecordStop] = "停止录制"
	AlarmCodeMap[PowerAbnormal] = "电源异常"
	AlarmCodeMap[VideoQuality] = "视质检测"
	AlarmCodeMap[DutyCheck] = "值岗检测"

	AlarmCodeMap[FierceMotion] = "剧烈运动"
	AlarmCodeMap[DeviceFault] = "设备故障"
	AlarmCodeMap[FanFault] = "风扇故障"
	AlarmCodeMap[TemperatureAlarm] = "温度报警"
	AlarmCodeMap[HumidityAlarm] = "湿度报警"
	AlarmCodeMap[DoorPointAlarm] = "门磁报警"
	AlarmCodeMap[ForceOperation] = "强行闯入"
	AlarmCodeMap[SoundAlarm] = "分贝报警"
	AlarmCodeMap[LowBattery] = "电量低"
	AlarmCodeMap[VideoAnalyseAlarm] = "视频分析报警"
	AlarmCodeMap[FaceDefenceAlarm] = "人脸布防报警"
}
