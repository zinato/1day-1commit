
def solution(n):
    DOWN = (0, 1)
    RIGHT = (1, 0)
    LEFT_TOP = (-1, -1)
    direction = [DOWN, RIGHT, LEFT_TOP]
    
    board = [[ 0 for _ in range(n) ] for _ in range(n)]
    d_idx = 0
    k = 1
    x, y = (0, 0)    
    board[y][x] = k
    
    move_cnt = [ n-i for i in range(n) ]
    move_cnt[0] -= 1
    
    for c in move_cnt:
        for _ in range(c):
            _x, _y = direction[d_idx]
            x += _x
            y += _y
            k += 1
            board[y][x] = k
            
        d_idx = (d_idx + 1) % 3
    
    answer = []
    
    for i in range(n):
        for j in range(n):
            if board[i][j] != 0:
                answer.append(board[i][j])
    
    return answer