package http

// ContextIn describes dependecies needed by this package
type ContextIn struct {
	RedirectTarget string
	Port           int
}

// ContextOut describes dependencies exported by this package
type ContextOut struct {
	Server Server
}

// Bootstrap initializes this module with ContextIn and exports
// resulting ContextOut
func Bootstrap(in *ContextIn) *ContextOut {

	out := &ContextOut{}
	out.Server = &server{
		redirectTarget: in.RedirectTarget,
		port:           in.Port,
	}

	return out
}
