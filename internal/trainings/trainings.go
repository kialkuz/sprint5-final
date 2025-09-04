package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	dataParts := strings.Split(datastring, ",")
	if len(dataParts) != 3 {
		return fmt.Errorf("invalid number of arguments")
	}

	countSteps, err := strconv.Atoi(dataParts[0])
	if err != nil {
		return err
	}

	if countSteps <= 0 {
		return fmt.Errorf("invalid count steps")
	}

	activityDuration, err := time.ParseDuration(dataParts[2])
	if err != nil {
		return err
	}

	if activityDuration <= 0 {
		return fmt.Errorf("zero or negative duration")
	}

	t.Steps = countSteps
	t.TrainingType = dataParts[1]
	t.Duration = activityDuration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var spentCalories float64

	var err error
	switch t.TrainingType {
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	// с конкатенацией сделал для читабельности, чтобы не растягивать строку
	return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\n",
			t.TrainingType,
			t.Duration.Hours(),
			spentenergy.Distance(t.Steps, t.Personal.Height),
		) +
			fmt.Sprintf("Скорость: %.2f км/ч\n", spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)) +
			fmt.Sprintf("Сожгли калорий: %.2f\n", spentCalories),
		nil
}
