package main

func (s *server) routes() {
	s.router.HandleFunc("/health", s.log(s.handleHealth()))

	s.router.HandleFunc("/stations", s.log(s.handleStations()))
	s.router.HandleFunc("/stations/", s.log(s.stationRouter.handle()))

	s.router.HandleFunc("/", s.log(s.handleNotFound()))
}
