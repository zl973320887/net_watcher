package mynet

type watcher struct {
	n Net
}

func NewWatcher(n Net) watcher {
	return watcher{n: n}
}

func (w *watcher) Watch() {
	w.n.Ping()
}
