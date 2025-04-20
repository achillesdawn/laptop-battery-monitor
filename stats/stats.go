package stats

import (
	"fmt"
	"time"
)

type BatStats struct {
	PowerNow    float32
	EnergyNow   float32
	EnergyTotal float32

	Percent  float32
	TimeLeft *time.Duration
}

func readPowerAndEnergy() (float32, float32, error) {

	powerNow, err := readFloat("/sys/class/power_supply/BAT0/power_now")
	if err != nil {
		return 0, 0, err
	}

	energy, err := readFloat("/sys/class/power_supply/BAT0/energy_now")
	if err != nil {
		return 0, 0, err
	}

	return powerNow, energy, nil
}

func New() (*BatStats, error) {

	energyTotal, err := readFloat(
		"/sys/class/power_supply/BAT0/energy_full",
	)
	if err != nil {
		return nil, err
	}

	powerNow, energy, err := readPowerAndEnergy()

	percent := energy / energyTotal * 100

	b := BatStats{
		PowerNow:    powerNow,
		EnergyNow:   energy,
		EnergyTotal: energyTotal,
		Percent:     percent,
		TimeLeft:    nil,
	}

	return &b, nil
}

func (b *BatStats) CalcTimeLeft() (*time.Duration, error) {

	timeLeft := b.EnergyNow / b.PowerNow

	duration, err := time.ParseDuration(
		fmt.Sprintf("%.2fm", timeLeft*60),
	)
	if err != nil {
		return nil, err
	}

	b.TimeLeft = &duration

	return &duration, nil
}

func (b *BatStats) PrintBatteryStats() {
	fmt.Printf(
		"%.1fw\t%s\t%.0f%%\n",
		b.PowerNow,
		b.TimeLeft.String(),
		b.Percent,
	)

}

func (b *BatStats) String() string {
	return fmt.Sprintf(
		"%.1fw\t%s\t%.0f%%",
		b.PowerNow,
		b.TimeLeft.Round(time.Second*60).String(),
		b.Percent,
	)
}

func (b *BatStats) ReadPowerAndEnergy() error {

	power, energy, err := readPowerAndEnergy()
	if err != nil {
		return err
	}

	b.PowerNow = power

	b.EnergyNow = energy

	percent := energy / b.EnergyTotal * 100

	b.Percent = percent

	return nil
}

func (b *BatStats) Monitor() error {
	for {
		err := b.ReadPowerAndEnergy()
		if err != nil {
			return err
		}

		b.CalcTimeLeft()

		b.PrintBatteryStats()

		time.Sleep(time.Second)
	}
}
