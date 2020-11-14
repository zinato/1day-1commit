# 알고리즘 문제 풀이: 베스트 앨범

![logo](../../logo.png)

> 프로그래머스 > 코딩테스트 연습 > 코딩테스트 고득점 Kit > 해시 > "베스트 앨범" 문제에 대한 풀이입니다.
> 
> 문제 링크는 [이 곳](https://programmers.co.kr/learn/courses/30/lessons/42579?language=java)에서 확인할 수 있습니다.


## 문제 분석

## 문제 풀이

이를 자바 코드로 옮기면 다음과 같다.

```java
import java.util.*;

public class Solution {
    static class Sing implements Comparable<Sing>  {
        public final int id;
        public final int plays;

        public Sing(int id, int plays) {
            this.id = id;
            this.plays = plays;
        }

        @Override
        public int compareTo(Sing other) {
            int result = Integer.compare(other.plays, plays);

            if (result == 0) {
                return Integer.compare(id, other.id);
            }

            return result;
        }
    }

    static class Track implements Comparable<Track> {
        public int totalPlay;
        public PriorityQueue<Sing> queue;

        public Track() {
            totalPlay = 0;
            queue = new PriorityQueue<>();
        }

        @Override
        public int compareTo(Track other) {
            int result = Integer.compare(other.totalPlay, totalPlay);
            return result;
        }
    }

    public int[] solution(String[] genres, int[] plays) {
        List<Sing> albums = new ArrayList<>();
        Map<String, Track> q = new HashMap<>();

        for (int i=0; i<genres.length; i++) {
            String genre = genres[i];
            int play = plays[i];

            if (!q.containsKey(genre)) {
                q.put(genre, new Track());
            }

            Track t = q.get(genre);
            t.totalPlay += play;
            t.queue.add(new Sing(i, play));
        }

        PriorityQueue<Track> pq = new PriorityQueue<>();

        for (Map.Entry<String, Track> entry : q.entrySet()) {
            pq.add(entry.getValue());
        }

        while (!pq.isEmpty()) {
            Track t = pq.remove();
            int cnt = 0;

            while (!t.queue.isEmpty() && cnt < 2) {
                Sing sing = t.queue.remove();
                albums.add(sing);
                cnt += 1;
            }
        }

        int [] answer = albums.stream()
                .mapToInt(s -> s.id)
                .toArray();
        return answer;
    }
}
```