__url__ = 'https://www.hackerrank.com/challenges/is-binary-search-tree/problem'

""" Node is defined as
class node:
  def __init__(self, data):
      self.data = data
      self.left = None
      self.right = None
"""
def in_order_traverse(node, output):
    if node.left:
        in_order_traverse(node.left, output)
    if node.data:
        if output and node.data < output[-1]:
            raise Exception('error')
        output.append(node.data)
    if node.right:
        in_order_traverse(node.right, output)



def check_binary_search_tree_(root):
    output = []
    try:
        in_order_traverse(root, output)
    except Exception:
        return False
    return True
