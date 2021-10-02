package progress

import "fmt"

type DataUnit int64
type SpeedUnit int64

const (
	B   DataUnit = 1
	KiB          = 1024 * B
	MiB          = 1024 * KiB
	GiB          = 1024 * MiB
	TiB          = 1024 * GiB
	PiB          = 1024 * TiB
)

func (d DataUnit) Value() (float64, string) {
	switch {
	case d < KiB:
		return float64(d), "B"
	case d < MiB:
		return float64(d) / float64(KiB), "KiB"
	case d < GiB:
		return float64(d) / float64(MiB), "MiB"
	case d < TiB:
		return float64(d) / float64(GiB), "GiB"
	case d < PiB:
		return float64(d) / float64(TiB), "TiB"
	default:
		return float64(d), "B"
	}
}

func (d DataUnit) String() string {
	v, u := d.Value()
	return fmt.Sprintf("%6.3f %s", v, u)
}

func (s SpeedUnit) Value() (float64, string) {
	switch {
	case s < SpeedUnit(KiB):
		return float64(s), "B/s"
	case s < SpeedUnit(MiB):
		return float64(s) / float64(KiB), "KiB/s"
	case s < SpeedUnit(GiB):
		return float64(s) / float64(MiB), "MiB/s"
	case s < SpeedUnit(TiB):
		return float64(s) / float64(GiB), "GiB/s"
	case s < SpeedUnit(PiB):
		return float64(s) / float64(TiB), "TiB/s"
	default:
		return float64(s), "B/s"
	}
}

func (s SpeedUnit) String() string {
	v, u := s.Value()
	return fmt.Sprintf("%6.3f %s", v, u)
}
