package user

import "time"

type User struct {
	ID        int       `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"-"`
	NimNip       string    `db:"nim_nip" json:"nim_nip,omitempty"`
	Role      string    `db:"role" json:"role"`
	Jabatan   string    `db:"jabatan" json:"jabatan"`
	ProgramStudi string `db:"program_studi" json:"program_studi,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}