#include <iostream>
#include <vector>
using namespace std;

class Solution {
  public:
    void moveZeroes(vector<int> &nums) {
        int p = 0;
        int q = 0;
        while (p < nums.size()) {

            // p tracker the index, q tracking the zero index
            if (nums[p] != 0) {
                // search zero from q to p
                while (q < p && nums[q] != 0) {
                    q++;
                }
                if (q < p) {
                    nums[q] = nums[p];
                    nums[p] = 0;
                }
            }
            p++;
        }
    }
};

int main() {
    vector<int> nums = {0, 1, 0, 3, 12};
    Solution solution;
    solution.moveZeroes(nums);
    for (int i = 0; i < nums.size(); i++) {
        cout << nums[i] << " ";
    }
    cout << endl;

    vector<int> nums2 = {0, 0, 1};
    solution.moveZeroes(nums2);
    for (int i = 0; i < nums2.size(); i++) {
        cout << nums2[i] << " ";
    }
    cout << endl;

    vector<int> nums3 = {0, 0, 0};
    solution.moveZeroes(nums3);
    for (int i = 0; i < nums3.size(); i++) {
        cout << nums3[i] << " ";
    }
    cout << endl;
}
