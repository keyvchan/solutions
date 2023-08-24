#include <algorithm>
#include <cstdio>
#include <vector>
using std::max;
using std::vector;

class Solution {
public:
  vector<int> replaceElements(vector<int> &arr) {
    // from back
    auto maxNum = -1;

    for (int i = arr.size() - 1; i >= 0; --i) {
      auto tmp = arr[i];
      arr[i] = maxNum;
      maxNum = max(maxNum, tmp);
    }

    return arr;
  }
};
