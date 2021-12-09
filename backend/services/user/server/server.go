package server

// Init
func Init() {
	router := NewRouter()
	router.Run(":5000")
}
