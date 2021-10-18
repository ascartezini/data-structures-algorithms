
//  Given two non-empty arrays of integers, write a function that determines whether the second array is a subsequence of the first one.

//  A subsequence of an array is a set of numbers that aren't necessarily adjacent
//  in the array but that are in the same order as they appear in the array. For
//  instance, the numbers <span>[1, 3, 4]</span> form a subsequence of the array
//  <span>[1, 2, 3, 4]</span>, and so do the numbers <span>[2, 4]</span>. Note
//  that a single number in an array and the array itself are both valid
//  subsequences of the array.

// Sample Input
//  = [5, 1, 22, 25, 6, -1, 8, 10]
//  = [1, 6, -1, 10]
// Sample Output
// true

// O(n) time | O(1) space
function isValidSubsequence(array, sequence) {

    let sequenceIndex = 0;
    for (let i = 0; i < array.length; i++) {
        if (array[i] === sequence[sequenceIndex]) sequenceIndex++;
    }

    return sequenceIndex === sequence.length;
}

let array = [5, 1, 22, 25, 6, -1, 8, 10];
let sequence = [1, 6, -1, 10];
console.log(isValidSubsequence(array, sequence));



