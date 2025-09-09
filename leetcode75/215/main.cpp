#include <iostream>
#include <vector>
using namespace std;

class Solution {
  public:
    int findKthLargest(vector<int> &nums, int k) {
        priority_queue<int, vector<int>, less<int>> pq;
        for (int num : nums) {
            pq.push(num);
        }
        for (int i = 0; i < k - 1; i++) {
            pq.pop();
        }
        return pq.top();
    }
};

int main() {
    Solution s;
    vector<int> nums{3, 2, 1, 5, 6, 4};
    auto elem = s.findKthLargest(nums, 2);
    cout << elem << endl;
    return 0;
}
