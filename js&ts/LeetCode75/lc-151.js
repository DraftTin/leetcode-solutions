/**
 * @param {string} s
 * @return {string}
 */
var reverseWords = function (s) {
  let words = s.trim().split(" ");
  let reversed = [];
  for (let i = words.length - 1; i >= 0; i--) {
    if (words[i] === "") {
      continue;
    }
    reversed.push(words[i]);
  }
  return reversed.join(" ");
};
