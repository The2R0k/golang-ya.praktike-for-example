def Solution(node):
    if node is not None:
        print(node.value)
        Solution(node.next_item)