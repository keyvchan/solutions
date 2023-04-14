#include <utility>
#include <vector>
using std::make_pair;
using std::pair;
using std::vector;

class Solution {
public:
  void setZeroes(vector<vector<int>> &matrix) {
    // iterate through the matrix, find all 0

    auto zeros = vector<pair<int, int>>();

    for (int i = 0; i < matrix.size(); i++) {
      for (int j = 0; j < matrix[i].size(); j++) {
        if (matrix[i][j] == 0) {
          zeros.push_back(make_pair(i, j));
        }
      }
    }

    // set all rows and columns to 0
    for (auto zero : zeros) {
      for (int i = 0; i < matrix.size(); i++) {
        matrix[i][zero.second] = 0;
      }
      for (int j = 0; j < matrix[zero.first].size(); j++) {
        matrix[zero.first][j] = 0;
      }
    }
  }
};
