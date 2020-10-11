import java.util.*;
import java.util.stream.Collectors;

public class Solution04 {
    public int solution(String depar, String hub, String dest, String[][] roads) {
        Map<String, Set<String>> map = new HashMap<>();

        for (String[] road : roads) {
            String from = road[0];
            String to = road[1];
            map.putIfAbsent(from, new HashSet<>());
            map.get(from).add(to);
        }

        Set<String> isVisit = new HashSet<>();
        Set<String> routes = new HashSet<>();
        dfs(depar, dest, depar, map, isVisit, routes);

        int answer = 0;

        for (String route : routes) {
            if (route.contains(depar) && route.contains(hub) && route.contains(dest)) {
                answer += 1;
            }
        }

        return (answer <= 0) ? answer : (answer % 10007);
    }

    public void dfs(String from, String to, String route, Map<String, Set<String>> map, Set<String> isVisit, Set<String> routes) {
        if (from.equals(to)) {
            routes.add(route);
            return;
        }

        if (isVisit.contains(from)) {
            return;
        }

        isVisit.add(from);
        Set<String> nextRoutes = map.get(from);

        for (String next : nextRoutes) {
            if (!isVisit.contains(next)) {
                dfs(next, to, route + "," + next, map, new HashSet<>(isVisit), routes);
            }
        }
    }
}
