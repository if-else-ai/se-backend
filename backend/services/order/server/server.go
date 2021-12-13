package server

func Init() {
	router := NewRouter()
	router.Run(":5000")
}
