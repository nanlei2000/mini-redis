package db

type DbDropGuard struct {
	db Db
}

type Db struct {
	shared Shared
}

type Shared struct {
	state          State
	backgroundTask interface{}
}

type State struct {
	entries map[string]Entry
	pubSub  map[string]interface{}
	// TODO:
	expirations map[string]interface{}
	nextID      uint64
	shutdown    bool
}

type Entry struct {
	id   uint64
	data byte
	// TODO:
	expiresAt interface{}
}

func New() Db {
	shared := Shared{
		state: State{
			entries:     make(map[string]Entry),
			pubSub:      make(map[string]interface{}),
			expirations: make(map[string]interface{}),
			nextID:      0,
			shutdown:    false,
		},
		backgroundTask: "TODO:",
	}

	//  // Start the background task.
	//  tokio::spawn(purge_expired_tasks(shared.clone()));

	return Db{shared}
}

func (d *Db) Get(key *string) interface{} {
	s := d.shared.state

	return s.entries[*key]
}
