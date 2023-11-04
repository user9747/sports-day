package entity

type Event struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Category  string `db:"category" json:"category"`
	StartTime string `db:"start_time" json:"startTime"`
	EndTime   string `db:"end_time" json:"endTime"`
	CreatedAt string `db:"created_at" json:"createdAt"`
	UpdatedAt string `db:"updated_at" json:"updatedAt"`
}
