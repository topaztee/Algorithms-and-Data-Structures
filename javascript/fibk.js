// k - numbers to add
// n - numbers in array

//n - index to return

// k = 3
// n = 5;
// [1,1,1,3,5,9] =
function fibK(sk, n) {
	// return fib(n-1) + (n-2) ...
	let x = [];
	for(let i = 0; i < k; i++){
		x.push(1);
	}
	// x = [1,1,1, ..]
	for(let i = k; i < n; i++){
		// eahc loop sum the previous k
		let sum = 0;
		for(let ii = 1; ii <= k; ii++) {
			sum+= x[i-ii];
		}
//    x.push(b);
	}

	return x[n];
}

// fibK(3, 5, [1,1,1,3,5,9])
