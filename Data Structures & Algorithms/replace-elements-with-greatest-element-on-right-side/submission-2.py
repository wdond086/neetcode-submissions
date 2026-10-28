class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        for i in range(len(arr) - 1):
            arr[i] = arr[i + 1]
            for j in range(i + 2, len(arr)):
                if arr[j] > arr[i]:
                    arr[i] = arr[j]
        arr[len(arr) - 1] = -1
        return arr