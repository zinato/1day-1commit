def solution(info, query):
    answer = []
    
    d = dict()
    languages = ["python", "cpp", "java", "-"]
    parts = ["backend", "frontend", "-"]
    careers = ["junior", "senior", "-"]
    foods = ["chicken", "pizza", "-"]
    
    for l in languages:
        d[l] = dict()
        for p in parts:
            d[l][p] = dict()
            for c in careers:
                d[l][p][c] = dict() 
                for f in foods:
                    d[l][p][c][f] = []
    
    for i in info:
        language, part, career, food, code = i.split(" ")
        point = int(code)
        
        for l in ["-", language]:        
            for p in ["-", part]:
                for c in ["-", career]:
                    for f in ["-", food]:                                
                        d[l][p][c][f].append(point)
        
          
    for q in query:
        (language, part, career, f_c) = [e.strip() for e in q.split("and")]
        food, code = f_c.split(" ")
        point = int(code)
        result = d[language][part][career][food]
        cnt = 0
        
        for elem in result:
            if elem >= point:
                cnt += 1
                
        answer.append(cnt)
            
    return answer