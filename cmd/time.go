package cmd

import (
	"github.com/spf13/cobra"
	"go-tools/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	calculateTime string
	duration      string
)

var (
	timeCmd = &cobra.Command{
		Use:   "time",
		Short: "时间格式处理",
		Long:  "时间格式处理",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
)

var (
	nowTimeCmd = &cobra.Command{
		Use:   "now",
		Short: "获取当前时间",
		Long:  "获取当前时间",
		Run: func(cmd *cobra.Command, args []string) {
			nowTime := timer.GetNowTime()
			log.Printf("当前时间: %s, 当前时间戳: %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
		},
	}
)

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "时间推算",
	Long:  "时间推算,如果-c参数为空,则默认为当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("时间: %s, 时间戳: %d", t.Format(layout), t.Unix())
	},
}
