package controller

type Router struct {
	Path       string
	Controller Controller
}

type RegexRouter struct {
	Pattern    string
	Controller Controller
}
