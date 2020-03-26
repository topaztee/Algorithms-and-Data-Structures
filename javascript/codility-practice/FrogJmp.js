// Count minimal number of jumps from position X to Y.

// Complexity:
// 		expected worst-case time complexity is O(1);
//
// expected worst-case space complexity is O(1)

const {performance} = require('perf_hooks');

function solution(X, Y, D) {
	if (Y < X | D < 0) throw new Error;
	let steps = 0;
	// while(X < Y) {
	// 	X = X+D;
	// 	steps++
	// }
	steps = Math.ceil((Y - X)/ D);
	return steps;
}

var t0 = performance.now();
console.log(solution(10,85,30))
var t1 = performance.now();
console.log("Call to doSomething took " + (t1 - t0) + " milliseconds.");
