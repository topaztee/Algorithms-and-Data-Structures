let str = "hello world!";

function reverse(str) {
	str = str.split("");
	let strr = [];
	for(let i = 0; i < str.length; i++){
		strr.unshift(str[i]);
	}
	return strr.join("");
}

console.log(reverse(str));

// A much shorter way

str.split("").reverse().join('');
