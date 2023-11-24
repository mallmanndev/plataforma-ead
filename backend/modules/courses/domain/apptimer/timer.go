package apptimer

import "time"

type Timer interface {
	Now() time.Time
}

type AppTimer struct{}

func NewAppTimer() Timer {
	return AppTimer{}
}

func (t AppTimer) Now() time.Time {
	return time.Now()
}

type FakeTimer struct {
	t time.Time
}

func NewFakeTimer(t time.Time) Timer {
	return FakeTimer{t}
}

func (ft FakeTimer) Now() time.Time {
	return ft.t
}
