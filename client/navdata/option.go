package navdata

type OptionHeader struct {
	Tag    OptionTag
	Length uint16
}

type OptionTag uint16

const (
	DEMO OptionTag = iota
	TIME
	RAW_MEASURES
	PHY_MEASURES
	GYROS_OFFSETS
	EULER_ANGLES
	REFERENCES
	TRIMS
	RC_REFERENCES
	PWM
	ALTITUDE
	VISION_RAW
	VISION_OF
	VISION
	VISION_PERF
	TRACKERS_SEND
	VISION_DETECT
	WATCHDOG
	ADC_DATA_FRAME
	VIDEO_STREAM
	GAMES
	PRESSURE_RAW
	MAGNETO
	WIND_SPEED
	KALMAN_PRESSURE
	HDVIDEO_STREAM
	WIFI
	ZIMMU3000
	CHECKSUM = 0xffff
)
