class DynamicArray:
    
    def __init__(self, capacity: int):
        self.array = [''] * capacity
        self.count = 0
        self.capacity = capacity

    def get(self, i: int) -> int:
        return self.array[i]


    def set(self, i: int, n: int) -> None:
        self.array[i] = n


    def pushback(self, n: int) -> None:
        if self.count == self.capacity:
            self.resize()
        self.set(self.count, n)
        self.count += 1


    def popback(self) -> int:
        self.count -= 1
        last_item = self.get(self.count)
        self.set(self.count, '')
        return last_item

    def resize(self) -> None:
        self.capacity *= 2
        new_array = [''] * self.capacity
        for index, item in enumerate(self.array):
            new_array[index] = item
        self.array = new_array


    def getSize(self) -> int:
        return self.count
        
    
    def getCapacity(self) -> int:
        return self.capacity
