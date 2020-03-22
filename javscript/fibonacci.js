// First naive solution
// Recursion

function fibRec(n) {
	return n <= 1 ? n : fib(n - 1) + fib(n - 2)
}

// dynamic programming bottom-up approach
// Improvement over recursive solution. which is (2^N) time complexity exponential
// That’s because the function makes too many subcalls. The same values are re-evaluated again and again.

// First implementation - loop
// calculate the next number by adding the current number to the old number.

function fibLoop(n) {
	let a = 1;
	let b = 2;
	for (let i = 3; i < n; i++) {
		let c = a + b;
		a = b;
		b = c;
	}
	return b;
}

// The loop starts with i=3, because the first and the second sequence values are hard-coded into variables a=1, b=1.

// memoization
// Is an optimization technique used primarily to speed up computer programs by storing the results of expensive function calls.

function fib(n, memo = {}) {
	if (memo[n]) return memo[n];
	if (n <= 1) return n;

	return memo[n] = fib(n - 1, memo) + fib(n - 2, memo)

}


// |             | Time complexity | Space complexity |   |   |
// |-------------|-----------------|------------------|---|---|
// | loop        | O(n)            | O(1)             |   |   |
// | recursion   | O(2^n)          | O(n)             |   |   |
// | memoization | O(2n)           | O(n)             |   |   |
//
//
// The memoization version will take O(n) time on first run,
// 		since each number is only computed once. However, in exchange,
// 		it also take O(n) memory for your current implementation
// (the n comes from storing the computed value, and also for the stack on the first run).
// If you run it many times, the time complexity will become O(M + q)
// where M is the max of all input n and q is the number of queries.
// 		The memory complexity will become O(M), which comes from the array
// which holds all the computed values.
//
//
// The loop(iterative) implementation is the best if you consider one run,
// 		as it also runs in O(n), but uses constant amount of memory O(1) to compute.
// 		For a large number of runs, it will recompute everything,
// 		so its performance may not be as good as memoization version.
