#include <iostream>
#include <vector>
using namespace std;

class Solution {
  public:
    int findPeakElement(vector<int> &nums) {
        for (int i = 0; i < nums.size(); i++) {
            int left = i - 1 >= 0 ? nums[i - 1] : INT_MIN;
            int right = i + 1 < nums.size() ? nums[i + 1] : INT_MIN;
            if (nums[i] > left && nums[i] > right) {
                return i;
            }
        }
        return 0;
    }
};

int main() {
    vector<int> nums = {1, 2, 1, 3, 5, 6, 4};
    Solution s;
    int index = s.findPeakElement(nums);
    cout << index << endl;
    return 0;
}
