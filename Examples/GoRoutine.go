package Examples

func GoRoutine() {
	go Count("Dock")
	Count("Ship") // fi you add go here too it will exit directly because the main will not wait the routines that is why we need to use waitGroup
}
