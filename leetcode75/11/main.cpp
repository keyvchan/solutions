#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
  public:
    int maxArea(vector<int> &height) {
        int left = 0;
        int right = height.size() - 1;

        int maxArea = 0;

        maxArea = min(height[left], height[right]) * (right - left);

        while (left < right) {

            if (left == right) {
                break;
            }

            left++;
            // check if move left will make area bigger
            if (height[left + 1] < height[left]) {
                left--;
            }

            right--;
            // check if move right will make area bigger
            if (height[right - 1] < height[right]) {
                right++;
            }

            return maxArea;
        }

        return maxArea;
    }
};

int main() {
    vector<int> height = {1, 8, 6, 2, 5, 4, 8, 3, 7};

    Solution s;
    assert(49 == s.maxArea(height));

    vector<int> height2 = {3, 6, 1};

    assert(3 == s.maxArea(height2));
}
