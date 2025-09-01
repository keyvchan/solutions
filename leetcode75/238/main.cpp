#include <cassert>
#include <vector>
using namespace std;

class Solution {
  public:
    vector<int> productExceptSelf(vector<int> &nums) {
        vector<int> prefix(nums.size(), 1);
        vector<int> suffix(nums.size(), 1);
        for (int i = 1; i < nums.size(); i++) {
            prefix[i] = prefix[i - 1] * nums[i - 1];
        }
        for (int i = nums.size() - 2; i >= 0; i--) {
            suffix[i] = suffix[i + 1] * nums[i + 1];
        }
        vector<int> res(nums.size());
        for (int i = 0; i < nums.size(); i++) {
            res[i] = prefix[i] * suffix[i];
        }
        return res;
    }
};

int main() {
    vector<int> nums1 = {1, 2, 3, 4};
    assert(Solution().productExceptSelf(nums1) == vector<int>({24, 12, 8, 6}));
    vector<int> nums2 = {-1, 1, 0, -3, 3};
    assert(Solution().productExceptSelf(nums2) == vector<int>({0, 0, 9, 0, 0}));
    return 0;
}
