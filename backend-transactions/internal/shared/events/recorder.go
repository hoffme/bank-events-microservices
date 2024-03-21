package events

type Recorder struct {
	events []MessageGeneric
}

func NewRecorder() *Recorder {
	return &Recorder{events: []MessageGeneric{}}
}

func (r *Recorder) Record(events ...MessageGeneric) {
	r.events = append(r.events, events...)
}

func (r *Recorder) Clear() {
	r.events = []MessageGeneric{}
}

func (r *Recorder) Pull() []MessageRaw {
	events := r.events
	r.events = []MessageGeneric{}

	result := make([]MessageRaw, len(events))
	for i, event := range events {
		result[i] = event.ToRaw()
	}

	return result
}

func (r *Recorder) Get() []MessageGeneric {
	result := make([]MessageGeneric, len(r.events))
	copy(result, r.events)
	return result
}
