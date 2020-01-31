package main

func (s *server) routes() {
	s.router.HandleFunc("/health", s.log(s.handleHealth()))
	s.router.HandleFunc("/", s.log(s.handleNotFound()))
}
