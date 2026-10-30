class MinStack:

    def __init__(self):
        self.stack = []
        self.orderStack = []

    def push(self, val: int) -> None:
        self.stack.append(val)
        stack_length = len(self.orderStack)
        if stack_length == 0:
            self.orderStack.append(val)
        else:
            currentMin = self.orderStack[-1]
            if val < currentMin:
                self.orderStack.append(val)
            else:
                self.orderStack.append(currentMin)

    def pop(self) -> None:
        self.stack.pop()
        self.orderStack.pop()
        

    def top(self) -> int:
        return self.stack[-1]
        

    def getMin(self) -> int:
        return self.orderStack[-1]
