package database

type database struct {
}

func (d *database) Start() {
	println("> [database] Starting...")
	println("> [database] Starting Done! Database running!")
}
 
func (d *database) Stop() {
	println("> [database] Stopping...")
	println("> [database] Stopping Done! Database stopped!")
}

func New() *database {
	return &database{}
}
