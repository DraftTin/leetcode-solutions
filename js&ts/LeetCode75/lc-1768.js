/**
 * @param {string} word1
 * @param {string} word2
 * @return {string}
 */
var mergeAlternately = function (word1, word2) {
  let i = 0;
  let j = 0;
  let res = "";
  while (i < word1.length && j < word2.length) {
    res += word1[i++];
    res += word2[j++];
  }
  res = res + word1.substring(i, word1.length);
  res = res + word2.substring(j, word2.length);

  return res;
};
