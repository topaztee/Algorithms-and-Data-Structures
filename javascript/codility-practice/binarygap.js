// A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded
// by ones at both ends in the binary representation of N.

function solution(N) {
	// write your code in JavaScript (Node.js 8.9.4)
	const bin = Math.abs(N).toString("2")
	let finalMax = 0, currentMax=0;

	for(let i=0; i < bin.length; i++ ) {
		if(bin[i] == "0") {
			currentMax++
		}

		if(bin[i] == "1" && i != 0) {
			finalMax = Math.max(finalMax, currentMax)
			currentMax = 0
		}

	}

	return parseInt(finalMax)
}

solution(1) // 0
solution(5) // 1
solution(19) // 2
solution(1041) // 5
solution(6291457) // 20

