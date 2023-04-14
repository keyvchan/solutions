#include <vector>

using std::vector;

class Solution {
public:
  vector<int> getConcatenation(vector<int> &nums) {
    // new vector with double the size of nums
    auto newVec = vector<int>(nums.size() * 2);

    for (int i = 0; i < nums.size(); i++) {
      // copy the first half of nums into newVec
      newVec[i] = nums[i];
      // copy the second half of nums into newVec
      newVec[i + nums.size()] = nums[i];
    }

    return newVec;
  }
};
