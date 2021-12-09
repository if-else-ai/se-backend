package server

func Start(){
	router := NewRouter()
	router.Run(":5000")
}