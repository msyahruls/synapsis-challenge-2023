package helper

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	LevelAdmin             = os.Getenv("ADMIN")
	LevelKoordinator       = os.Getenv("KOORDINATOR")
	LevelRelawan           = os.Getenv("RELAWAN")
	LevelGeneralUser       = os.Getenv("GENERALUSER")
	DefaultGoalRecruitment = os.Getenv("DEFAULT_GOAL_RECRUITMENT")
)
