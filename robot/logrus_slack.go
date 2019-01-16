package robot

import (
	"github.com/johntdyer/slackrus"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initLogrus() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//file, _ := os.Create(time.Now().Format("2006_01_02.log"))
	//logrus.SetOutput(file)
	
	logrus.SetLevel(logrus.DebugLevel)
	
	logrus.AddHook(&slackrus.SlackrusHook{
		HookURL:        viper.GetString("slack_hook"),
		AcceptedLevels: slackrus.LevelThreshold(logrus.DebugLevel),
		Channel:        viper.GetString("slack_channel"),
		IconEmoji:      ":shark:",
		Username:       "Eric",
	})
}

type SlackHook struct {
	Url   string
	Name  string
	Level logrus.Level
}

func (h *SlackHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (h *SlackHook) Fire(e *logrus.Entry) error {
	
	return nil
}
