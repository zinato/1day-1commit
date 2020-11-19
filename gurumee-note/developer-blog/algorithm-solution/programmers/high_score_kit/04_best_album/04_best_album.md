# 알고리즘 문제 풀이: 베스트 앨범

![logo](../../logo.png)

> 프로그래머스 > 코딩테스트 연습 > 코딩테스트 고득점 Kit > 해시 > "베스트 앨범" 문제에 대한 풀이입니다.
> 
> 문제 링크는 [이 곳](https://programmers.co.kr/learn/courses/30/lessons/42579?language=java)에서 확인할 수 있습니다.


## 문제 분석

문제의 입력과 함께 문제를 풀어보면서 어떻게 풀어야 할 지 분석을 해보자.

문제 입력:
```
genres = ["classic", "pop", "classic", "classic", "pop"]
plays = [500, 600, 150, 800, 2500]
```

여기서, 각 노래는 id를 부여 받는다. id 별로 노래를 보면 다음과 같다.

```
id : 0
genre : classic
play : 500
--------------
id : 1
genre : pop
play : 600
--------------
id : 2
genre : classic
play : 150
--------------
id : 3
genre : classic
play : 800
--------------
id : 4
genre : pop
play : 2500
--------------
```

그럼 문제에서처럼, 장르 별로 플레이 횟수가 높은 순으로 모아보자.

```
genre : "pop"
total_play : 3100
songs: [
    {
        id : 4,
        genre : pop,
        play : 2500
    },
    {
        id : 1,
        genre : pop,
        play : 600
    },
]
--------------
genre : "classic"
total_play : 1450
songs: [
    {
        id : 3,
        genre : classic,
        play : 800
    },
    {
        id : 0,
        genre : classic,
        play : 500
    },
    {
        id : 2,
        genre : pop,
        play : 150
    },
]
```

이제 많이 플레이된 장르 순으로, 장르 별 제일 많이 플레이된 노래 2곡씩 뽑아서, id의 리스트 형태로 만들어주면 된다.

```
result = [4, 1, 3, 0]
```

이 문제는 해시와 우선순위 큐를 적절하게 조합하여 사용하면 쉽게 풀 수 있다.


## 문제 풀이

"우선 순위 큐"는 정렬된 상태로 객체가 저장된다. 자바 클래스는 `Comparable<T>` 인터페이스를 구현해야 적절하게 원하는 순서로 만들 수 있다. 여기서는 2개의 우선순위 큐가 필요하다. 

각각의 우선순위 큐는 다음의 정보를 가진 객체를 저장한다.

1. {id, 플레이 횟수} 형태의 객체를 플레이 횟수 별 내림차순, 만약에 같을 땐 고유번호 별 오름차순으로 정렬하는 우선순위 큐
2. {(1) 형태의 우선순위 큐와, 우선 순위큐에 저장된 총 플레이 횟수} 형태의 객체를 총 플레이 횟수 내림차순으로 정렬하는 우선 순위 큐

자바에서는 우선순위 큐에 커스텀한 정렬을 적용하려면, 객체를 정의하고 `Comparable<T>` 인터페이스를 구현해야 한다. 먼저, {id, 플레이 횟} `Song`이란 클래스로 정의하겠다. 

```java
static class Song implements Comparable<Song>  {
    public final int id;
    public final int plays;

    public Song(int id, int plays) {
        this.id = id;
        this.plays = plays;
    }

    @Override
    public int compareTo(Song other) {
        int result = Integer.compare(other.plays, plays);

        if (result == 0) {
            return Integer.compare(id, other.id);
        }

        return result;
    }
}
```

그리고, 이제 장르별로 저장된 노래들의 우선순위 큐와 그 노래들의 총 플레이 횟수 정보를 담은 것을 `Track`이란 클래스로 정의한다. 역시 우선 순위 큐에 "총 플레이 횟수" 별로 내림차순 해야 하기 때문에 `Comparable<T>` 인터페이스를 구현해야 한다.

```java
static class Track implements Comparable<Track> {
    public int totalPlay;
    public PriorityQueue<Song> queue;

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
```

이제 위에서 했던 것을 고대로 코드로 구현하면 된다. 정리하면 순서는 다음과 같다.

1. "genre"를 키로, `Track`을 값으로 갖는 해시를 만든다. // 장르 별로, 정렬된 노래 목록을 만든다.
2. 입력 "genres", "plays" 길이만큼 순회한다.(둘은 같은 길이를 가닌다.)
   1. 해시에 "genre" 키가 없다면, "genre"를 키로 빈 `Track`을 값으로 해시에 저장한다.
   2. "genre"를 키로 해시에서 `Track`을 가져온다.
   3. `Track`의 총 플레이 횟수에서 "play"를 더해준다.
   4. `Track`의 우선순위 큐에, "play" 그리고 순회하는 인덱스를 "id"로 `Song`을 만든 후 저장한다.
3. 이제 `Track`을 저장하는 우선순위 큐 "pq"를 만든다.
4. 해시를 순회하여, pq에 `Track`을 저장한다.
5. `Song`을 저장하는 리스트를 만든다. 이를 "albums"라고 한다.
6. 이제 우선순위 큐 pq가 저장된 `Track`이 없을 때까지 순회한다.
   1. pq에서 `Track`을 가져온다.
   2. 이 `Track`의 `Song` 객체를 저장하고 있는 우선순위 큐에서 2개의 `Song`을 가져와서 "albums"에 저장한다.
7. 이제 "albums"를 `Song` 클래스에서 "id"로 변환한 후 정수형 배열로 만든 후 반환한다.

이를 자바 코드로 옮기면 다음과 같다.

```java
import java.util.*;

public class Solution {
    static class Song implements Comparable<Song>  {
        public final int id;
        public final int plays;

        public Song(int id, int plays) {
            this.id = id;
            this.plays = plays;
        }

        @Override
        public int compareTo(c other) {
            int result = Integer.compare(other.plays, plays);

            if (result == 0) {
                return Integer.compare(id, other.id);
            }

            return result;
        }
    }

    static class Track implements Comparable<Track> {
        public int totalPlay;
        public PriorityQueue<Song> queue;

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
        Map<String, Track> q = new HashMap<>();

        for (int i=0; i<genres.length; i++) {
            String genre = genres[i];
            int play = plays[i];

            if (!q.containsKey(genre)) {
                q.put(genre, new Track());
            }

            Track t = q.get(genre);
            t.totalPlay += play;
            t.queue.add(new Song(i, play));
        }

        PriorityQueue<Track> pq = new PriorityQueue<>();

        for (Map.Entry<String, Track> entry : q.entrySet()) {
            pq.add(entry.getValue());
        }

        List<Song> albums = new ArrayList<>();

        while (!pq.isEmpty()) {
            Track t = pq.remove();
            int cnt = 0;

            while (!t.queue.isEmpty() && cnt < 2) {
                Song sing = t.queue.remove();
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