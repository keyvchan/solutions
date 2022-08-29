# Definition for a Node.
class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def treeToDoublyList(self, root: Node) -> Node:
        self.nodes: list[Node] = []
        # put all node into a list
        def dfs(root):
            if not root:
                return
            dfs(root.left)
            self.nodes.append(root)
            dfs(root.right)

        dfs(root)

        last = None
        for node in self.nodes:
            if last == None:
                last = node
                continue
            last.right = node
            node.left = last
            node = node.right
            last = last.right

        self.nodes[0].left = last
        self.nodes[len(self.nodes) - 1].right = self.nodes[0]

        return self.nodes[0]
