function rotateAr(A, K) {
	if(A.length < 1) return;
	for(let i = 0; i < K; i++){
		let last = A[A.length-1]
		A.pop()
		A.unshift(last);
	}
	return A;
}

console.log(rotateArr([1,2,3,4], 2)); // 3,4,1,2


// pop - add to end o farray
// push - remove from end of array
//
// shift add to start of array
// unshift remove from stary of array
