package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("invalid number of steps")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("invalid weight")
	}

	if height <= 0 {
		return 0, fmt.Errorf("invalid height")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration")
	}

	averageSpeed := MeanSpeed(steps, height, duration)

	return (weight * averageSpeed * duration.Minutes()) * walkingCaloriesCoefficient / 60, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("invalid number of steps")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("invalid weight")
	}

	if height <= 0 {
		return 0, fmt.Errorf("invalid height")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration")
	}

	averageSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * averageSpeed * durationInMinutes) / float64(minInH), nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	return (float64(steps) * stepLength) / mInKm
}
