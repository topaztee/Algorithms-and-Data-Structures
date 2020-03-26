

// return new context
function bind(fn, context) {
	return function() {
		fn.call(context, ...arguments)
	}
}
