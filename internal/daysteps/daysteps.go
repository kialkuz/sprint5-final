package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataParts := strings.Split(datastring, ",")
	if len(dataParts) != 2 {
		err := fmt.Errorf("invalid number of arguments")
		log.Println(err)
		return err
	}

	countSteps, err := strconv.Atoi(dataParts[0])
	if err != nil {
		return err
	}

	if countSteps <= 0 {
		return fmt.Errorf("invalid count steps")
	}

	walkDuration, err := time.ParseDuration(dataParts[1])
	if err != nil {
		return err
	}

	if walkDuration <= 0 {
		return fmt.Errorf("zero or negative duration")
	}

	ds.Steps = countSteps
	ds.Duration = walkDuration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
			"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
			ds.Steps,
			spentenergy.Distance(ds.Steps, ds.Personal.Height),
			spentCalories,
		),
		nil
}
