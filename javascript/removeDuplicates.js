
// function takes a string and removes the duplicates

// if you see something that says unique values. Then use a set
function removeDuplicates(str) {
	let arr = str.split('');
	let set = new Set(arr);
	let newStr = [...set].join('')
	return newStr;
}

removeDuplicates("aabbbbccc")


let arr = str.split('')
let set = new Set(arr);
let newStr = [...set].join
