package gplog

const (
	JsonFormat    = "json"
	ConsoleFormat = "console"

	ISO8601TimeEncoder     = "ISO8601"
	RFC3339TimeEncoder     = "RFC3339"
	RFC3339NanoTimeEncoder = "RFC3339Nano"
)

type LogConfig struct {
	ServiceName string
	Level       string
	LogFormat   string
	TimeFormat  string
	Filename    string
	Output      string
}
