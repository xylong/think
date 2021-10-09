package mode

/*
 * 观察者模式
 */

// observer 观察者
type observer interface {
	handle(interface{})
}

// subject
type subject interface {
	notify()
	add()
}
