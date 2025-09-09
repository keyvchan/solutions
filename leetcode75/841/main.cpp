#include <cassert>
#include <iostream>
#include <unordered_set>
#include <vector>
using namespace std;

class Solution {
  public:
    int remainRooms_;
    unordered_set<int> visitedRooms;

    void dfs(vector<vector<int>> &rooms, vector<int> &needsToVisited) {

        for (int i = 0; i < needsToVisited.size(); i++) {
            // check if we visited rooms before
            if (visitedRooms.find(needsToVisited[i]) != visitedRooms.end()) {
                // skip this room
                continue;
            }
            // visit this room
            visitedRooms.insert(needsToVisited[i]);
            dfs(rooms, rooms[needsToVisited[i]]);
        }
    }

    bool canVisitAllRooms(vector<vector<int>> &rooms) {

        visitedRooms.insert(0);
        dfs(rooms, rooms[0]);

        for (int i = 0; i < rooms.size(); i++) {
            cout << *visitedRooms.find(i) << endl;
        }
        return visitedRooms.size() == rooms.size();
    }
};

int main() {
    vector<vector<int>> rooms = {{1}, {2}, {3}, {}};
    assert(Solution().canVisitAllRooms(rooms));

    // [[1,3],[3,0,1],[2],[0]]
    vector<vector<int>> rooms2 = {{1, 3}, {3, 0, 1}, {2}, {0}};
    assert(!Solution().canVisitAllRooms(rooms2));
    return 0;
}
