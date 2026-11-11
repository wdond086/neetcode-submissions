class HistoryNode:

    def __init__(self, url: str, next: Optional["HistoryNode"] = None, prev: Optional["HistoryNode"] = None) -> None:
        self.url = url
        self.next = next
        self.prev = prev

class BrowserHistory:

    def __init__(self, homepage: str):
        newNode = HistoryNode(homepage)
        self._head = newNode
        self._tail = newNode
        

    def visit(self, url: str) -> None:
        newNode = HistoryNode(url, prev=self._tail)
        self._tail.next = newNode
        self._tail = newNode
        

    def back(self, steps: int) -> str:
        stepsTaken = 0
        while stepsTaken < steps and self._tail.prev != None:
            self._tail = self._tail.prev
            stepsTaken += 1
        return self._tail.url

    def forward(self, steps: int) -> str:
        stepsTaken = 0
        while stepsTaken < steps and self._tail.next != None:
            self._tail = self._tail.next
            stepsTaken += 1
        return self._tail.url


# Your BrowserHistory object will be instantiated and called as such:
# obj = BrowserHistory(homepage)
# obj.visit(url)
# param_2 = obj.back(steps)
# param_3 = obj.forward(steps)