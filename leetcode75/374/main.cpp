#include <iostream>
using namespace std;

int guess(int n) {
    if (n == 6) {
        return 0;
    } else if (n < 6) {
        return 1;
    } else {
        return -1;
    }
}

class Solution {
  public:
    int guessNumber(int n) {
        int input = n;

        while (true) {
            int res = guess(input);
            switch (res) {
            case 1: {
                input = input + 1;
                break;
            }
            case -1: {
                input = input - 1;
                break;
            }
            case 0: {
                return input;
                break;
            }
            default:
                return -1;
            }
        }
    }
};

int main() {
    Solution s;
    int res = s.guessNumber(10);
    cout << res << endl;
    return 0;
}
