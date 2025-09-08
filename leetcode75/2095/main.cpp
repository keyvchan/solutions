#include <iostream>
#include <ostream>
#include <vector>

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
  public:
    ListNode *deleteMiddle(ListNode *head) {
        // create a vector to store the nodes
        std::vector<ListNode *> nodes{};

        ListNode *iter = head;
        while (iter != nullptr) {
            nodes.push_back(iter);
            iter = iter->next;
        }

        if (nodes.size() == 1) {
            return nullptr;
        }

        if (nodes.size() == 2) {
            head->next = nullptr;
            return head;
        }

        int middle = nodes.size() / 2;

        // remove middle node

        int left = middle - 1;
        int right = middle + 1;

        ListNode *leftNode = nodes[left];
        ListNode *rightNode = nodes[right];

        leftNode->next = rightNode;

        return head;
    }
};

int main() {
    ListNode *head = new ListNode(
        1, new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5)))));

    Solution s;
    head = s.deleteMiddle(head);

    while (head != nullptr) {
        std::cout << head->val << std::endl;
        head = head->next;
    }

    ListNode *head2 = new ListNode(1, nullptr);

    head2 = s.deleteMiddle(head2);

    while (head2 != nullptr) {
        std::cout << head2->val << std::endl;
        head2 = head2->next;
    }

    ListNode *head3 = new ListNode(1, new ListNode(2, nullptr));

    head3 = s.deleteMiddle(head3);

    while (head3 != nullptr) {
        std::cout << head3->val << std::endl;
        head3 = head3->next;
    }

    return 0;
}
