#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  bool isValidSudoku(vector<vector<char>> &board) {
    // valid line
    // create a temp vector to make sure don't have duplicates
    auto row_valid = unordered_set<char>();
    auto column_valid = unordered_set<char>();
    auto square = unordered_set<char>();
    for (int i = 0; i < board.size(); i++) {
      // valid line
      for (int j = 0; j < board.size(); j++) {
        // i line
        if (row_valid.find(board[i][j]) != row_valid.end()) {
          if (board[i][j] != '.') {
            return false;
          }
        } else {
          row_valid.insert(board[i][j]);
        }
        // i column
        if (column_valid.find(board[j][i]) != column_valid.end()) {
          if (board[j][i] != '.') {
            return false;
          }
        } else {
          column_valid.insert(board[j][i]);
        }
        // i square i/3*3+j/3, i%3*3+j%3
        if (square.find(board[i / 3 * 3 + j / 3][i % 3 * 3 + j % 3]) !=
            square.end()) {
          if (board[i / 3 * 3 + j / 3][i % 3 * 3 + j % 3] != '.') {
            return false;
          }
        } else {
          square.insert(board[i / 3 * 3 + j / 3][i % 3 * 3 + j % 3]);
        }
      }
      row_valid.clear();
      column_valid.clear();
      square.clear();
    }
    return true;
  }
};
