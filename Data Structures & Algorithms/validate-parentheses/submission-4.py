class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        # If every open paran is closed, len should be even
        if len(s) % 2 != 0:
            return False

        for char in s:
            if char in '{([':
                stack.append(char)
            else:
                expected_char = ''
                # If this is true, then no open parantheses was added
                if len(stack) == 0:
                    return False
                match stack[-1]:
                    case '{':
                        expected_char = '}'
                    case '[':
                        expected_char = ']'
                    case '(':
                        expected_char = ')'
                # Here a closing parantheses appears before it was opened
                if char != expected_char:
                    return False
                stack.pop()
        return len(stack) == 0