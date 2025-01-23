package domain

type Passenger struct {
	ID       string    `db:"id"`
	Name     string    `db:"name"`
	Location *Location `db:"-"`
}
