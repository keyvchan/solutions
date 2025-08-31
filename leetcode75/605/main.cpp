#include <cassert>
#include <vector>
using namespace std;

class Solution {
  public:
    bool canPlaceFlowers(vector<int> &flowerbed, int n) {
        if (n == 0) {
            return true;
        }

        bool last = false;

        for (int i = 0; i < flowerbed.size(); i++) {
            // current position is not 0, can't place here, set last position to
            // 1 and proceed
            if (flowerbed[i] != 0) {
                last = true;
                continue;
            }

            // current is 0, check last
            if (last) {
                // already placed, set last to 0
                last = false;
                continue;
            }

            // current is 0, and last is 0, check next
            if (i == flowerbed.size() - 1 || flowerbed[i + 1] == 0) {
                // next is 0, can place here
                n--;
                last = true;
            }

            if (n == 0) {
                break;
            }
        }

        return n == 0;
    }
};

int main() {
    vector<int> flowerbed = {1, 0, 0, 0, 1};
    assert(Solution().canPlaceFlowers(flowerbed, 1) == true);

    vector<int> flowerbed2 = {1, 0, 0, 0, 1};
    assert(Solution().canPlaceFlowers(flowerbed2, 2) == false);

    vector<int> flowerbed3 = {0, 0, 1, 0, 1};
    assert(Solution().canPlaceFlowers(flowerbed3, 1) == true);

    vector<int> flowerbed4 = {0, 0, 1, 0, 0};
    assert(Solution().canPlaceFlowers(flowerbed4, 1) == true);

    return 0;
}
