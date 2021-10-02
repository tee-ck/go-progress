package progress

import (
	"time"
)

const (
	BlocksFilledDefault = "▏▎▍▌▋▊▉█"
	BlocksFilledArrow   = ">="
	BlockEmptyDefault   = ' '
	BlockEmptyArrow     = '-'
)

type Config struct {
	BlocksFilling string
	BlocksEmpty   rune
	Length        uint
}

type Speed struct {
	Peak    SpeedUnit
	Current SpeedUnit
	Average SpeedUnit
}

type Block struct {
	Size     int
	Index    int
	Fillable int // length of Filled
	Cached   []rune
	Filled   []rune
	Empty    rune
}

type Value struct {
	Now DataUnit
	Max DataUnit
}

type Progress struct {
	started   time.Time
	updated   time.Time
	progress  float64
	completed bool

	speed Speed
	block Block
	value Value
}

func (p *Progress) PeakSpeed() SpeedUnit {
	return p.speed.Peak
}

func (p *Progress) CurrentSpeed() SpeedUnit {
	return p.speed.Current
}

func (p *Progress) AverageSpeed() SpeedUnit {
	return p.speed.Average
}

func (p *Progress) Eta(speed SpeedUnit) time.Time {
	eta := float64(p.value.Max - p.value.Now) / float64(speed)

	return time.Now().Add(time.Duration(eta*1000) * time.Millisecond)
}

func (p *Progress) Progress() float64 {
	return p.progress
}

func (p *Progress) Bar() string {
	return string(p.block.Cached)
}

func (p *Progress) IsDone() bool {
	return p.completed
}

func (p *Progress) NowValue() DataUnit {
	return p.value.Now
}

func (p *Progress) SetMaxValue(value DataUnit) {
	p.value.Max = value
}

func (p *Progress) IncMaxValue(increment DataUnit) DataUnit {
	p.value.Max += increment
	return p.value.Max
}

func (p *Progress) MaxValue() DataUnit {
	return p.value.Max
}

func (p *Progress) Second() int {
	return time.Now().Second() - p.started.Second()
}

func (p *Progress) Increase(value DataUnit) {
	now := time.Now()

	p.speed.Current = SpeedUnit(float64(value) / (float64(now.UnixMilli()-p.updated.UnixMilli()) / 1000))
	if p.speed.Current > p.speed.Peak {
		p.speed.Peak = p.speed.Current
	}
	p.updated = now

	p.value.Now += value
	if p.value.Now >= p.value.Max {
		p.value.Now = p.value.Max
		p.progress = 100
		p.completed = true
	} else {
		p.progress = (float64(p.value.Now) / float64(p.value.Max)) * 100
	}
	p.speed.Average = SpeedUnit(float64(p.value.Now) / (float64(now.UnixMilli()-p.started.UnixMilli()) / 1000))

	width := p.block.Size * p.block.Fillable
	blocks := int(float64(width) * (p.progress / 100))
	filled := blocks / p.block.Fillable

	for i := p.block.Index; i < filled; i++ {
		p.block.Cached[i] = p.block.Filled[p.block.Fillable-1]

		p.block.Index = i + 1
	}

	if p.block.Index < p.block.Size {
		n := blocks % p.block.Fillable
		if n > 0 {
			n = n - 1
		}
		p.block.Cached[p.block.Index] = p.block.Filled[n]
	}
}

func New(total DataUnit, conf ...Config) *Progress {
	var (
		config Config
	)

	if len(conf) > 0 {
		config = conf[0]
	}

	if config.BlocksFilling == "" {
		config.BlocksFilling = BlocksFilledDefault
	}

	if config.BlocksEmpty == 0 {
		config.BlocksEmpty = BlockEmptyDefault
	}

	if config.Length == 0 {
		config.Length = 20
	}

	filled := []rune(config.BlocksFilling)
	cached := make([]rune, config.Length)
	for i := uint(0); i < config.Length; i++ {
		cached[i] = config.BlocksEmpty
	}

	return &Progress{
		started: time.Now().Add(-time.Second),
		updated: time.Now().Add(-time.Second),
		value: Value{
			Max: total,
		},
		speed: Speed{
			Peak:    0,
			Current: 0,
			Average: 0,
		},
		block: Block{
			Size:     int(config.Length),
			Index:    0,
			Fillable: len(filled),
			Filled:   filled,
			Cached:   cached,
			Empty:    config.BlocksEmpty,
		},
	}
}
