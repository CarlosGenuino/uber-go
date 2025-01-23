package domain

type Driver struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	LicenseID string    `db:"license_id"`
	Available bool      `db:"available"`
	Location  *Location `db:"-"`
	Car       Car       `db:"-"`
}
