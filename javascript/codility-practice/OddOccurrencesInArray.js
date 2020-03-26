// Find value that occurs in odd number of elements.

// Complexity:
// 		expected worst-case time complexity is O(N);
//
// expected worst-case space complexity is O(1)

function OddOccurrencesInArray(A) {
		let result = 0;
		// using in instead of will return the index
		for (let element of A) {
			// Apply Bitwise XOR to the current and next element
			result ^= element
			console.log(result)
		}

		return result
}


	console.log(OddOccurrencesInArray([5,1,1]))
