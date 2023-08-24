#include <vector>
using std::vector;

class Solution {
public:
  int diagonalSum(vector<vector<int>> &mat) {
    // start from 0,0
    auto result = 0;

    for (auto i = 0; i < mat.size(); ++i) {
      result += mat[i][i];
    }

    // start from 0, n-1
    for (auto i = 0; i < mat.size(); ++i) {
      result += mat[i][mat.size() - 1 - i];
    }

    // if mat size is odd, we need to remove the middle element
    if (mat.size() % 2 == 1) {
      result -= mat[mat.size() / 2][mat.size() / 2];
    }

    return result;
  }
};
