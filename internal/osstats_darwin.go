package internal

// No OS stats implemented for Darwin.

type OSData struct{}

func (s *Server) InitOSData()   {}
func (s *Server) checkOSStats() {}
