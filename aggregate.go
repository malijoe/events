package events

type Aggregate interface {
	ID() string
	Type() string
	Events() []Event
	Load([]Event) error
	HasUncommittedEvents() bool
	CommitEvents()
}
