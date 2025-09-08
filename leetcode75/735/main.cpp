#include <stack>
#include <vector>
using namespace std;

class Solution {
  public:
    vector<int> asteroidCollision(vector<int> &asteroids) {
        stack<int> stack = {};

        for (auto asteroid : asteroids) {
            // push to stack and check top
            if (stack.empty()) {
                stack.push(asteroid);
                continue;
            }

            if (stack.top() * asteroid > 0) {
                // is same direction

                if (abs(stack.top()) < abs(asteroid)) {
                    // left < right, lefr will exploid
                    stack.pop();
                    stack.push(asteroid);
                } else if (abs(stack.top()) == abs(asteroid)) {
                    // left = right, both will exploid
                    stack.pop();
                }
            } else {
                // is different direction
                // left > right, must exploid one
                if (abs(stack.top()) < abs(asteroid)) {
                    stack.pop();
                    stack.push(asteroid);
                    // nothing chageds
                } else if (abs(stack.top()) == asteroid) {
                    stack.pop();
                }
            }
        }
    }
};
