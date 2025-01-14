/**
 * @param {string} s
 * @return {string}
 */
var reverseVowels = function (s) {
  const vowels = {
    a: true,
    e: true,
    i: true,
    o: true,
    u: true,
    A: true,
    E: true,
    I: true,
    O: true,
    U: true,
  };
  let pos = [];
  for (let i = 0; i < length(s); i++) {
    if (vowels[s[i]] === true) {
      pos.push(i);
    }
  }
  let i = 0;
  let j = pos.length - 1;
  while (i < j) {
    tmp = s[pos[i]];
    s[pos[i]] = s[pos[j]];
    s[pos[j]] = tmp;
  }
  return s;
};
