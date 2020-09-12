def solution(numbers):
    s = set()
    size = len(numbers)
    
    for i in range(size-1):
        for j in range(i + 1, size):
            s.add(numbers[i] + numbers[j])
    
    
    answer = sorted(list(s))
    return answer