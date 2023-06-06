#include <algorithm>
#include <vector>
using std::sort;
using std::vector;

class Solution {
public:
  bool canMakeArithmeticProgression(vector<int> &arr) {
    // sort the vector
    sort(arr.begin(), arr.end());

    // check if the difference between each element is the same
    int diff = arr[1] - arr[0];
    for (int i = 2; i < arr.size(); i++) {
      if (arr[i] - arr[i - 1] != diff) {
        return false;
      }
    }
    return true;
  }
};
