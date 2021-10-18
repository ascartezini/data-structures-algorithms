function validAnagram(arr1, arr2) {
    // if the length is different then it's not a valid anagram
    if (arr1.length !== arr2.length) {
        return false;
    }

    let frequencyCounter1 = {};
    let frequencyCounter2 = {};

    // creates an object containing the count for each character for array1
    for (const val of arr1) {
        frequencyCounter1[val] = (frequencyCounter1[val] || 0) + 1;
    }

    // creates an object containing the count for each character for array2
    for (const val of arr2) {
        frequencyCounter2[val] = (frequencyCounter2[val] || 0) + 1;
    }

    // loop through the keys of array1
    for (const key in frequencyCounter1) {
        // checks if there is a letter key of arr1 in arr2
        if (!(key in frequencyCounter2)) {
            return false;
        }

        // checks if the count of the letter key of arr1 in arr2 is the same
        if (frequencyCounter2[key] !== frequencyCounter1[key]) {
            return false;            
        }
    }

    return true;
}

console.log(validAnagram("god", "dog"));
console.log(validAnagram("race", "care"));
console.log(validAnagram("earth", "heart"));