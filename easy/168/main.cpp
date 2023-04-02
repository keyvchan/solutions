#include <bits/stdc++.h>

class Solution {
public:
  std::string convertToTitle(int columnNumber) {

    // convert to base 26
    std::string result;
    while (columnNumber > 0) {
      int remainder = columnNumber % 26;
      if (remainder == 0) {
        result += 'Z';
        columnNumber = columnNumber / 26 - 1;
      } else {
        result += 'A' + remainder - 1;
        columnNumber = columnNumber / 26;
      }
    }
    return std::string(result.rbegin(), result.rend());
  }
};
