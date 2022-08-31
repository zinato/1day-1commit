# 4장. 블록, 섀도, 제어 구조

## 블록

p101. "섀도잉 변수는 포함된 블록 내에 이름이 같은 변수가 있는 것을 의미한다. 섀도우 변수가 존재하는 한 섀도잉 대상이 된 변수는 접근할 수 없다."

```go
x := 10
fmt.Println(x, &x) // 10 0xc0000b2008

if x > 5 {
    fmt.Println(x, &x) // 10 0xc0000b2008
	// x는 섀도잉된다.
    x, y := 5, 20
    fmt.Println(x, &x) //5 0xc0000b2018
    fmt.Println(y, &y) // 20 0xc0000b2030
}

fmt.Println(x, &x) // 10 0xc0000b2008
```

`shadow`도구를 설치하면 섀도잉된 변수를 잡아낼 수 있다. (버그의 원인이 될 수 있으므로 Makefile에 해당 검출 단계를 추가하는 것이 좋다.)

```bash
$ go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
```

## if

p105. "Go에서 사용되는 if 문이 다른 언어와 가장 큰 차이점을 보이는 것은 조건을 감싸는 괄호가 없다는 것이다."

p105. "Go에서 추가된 것은 조건과 if 혹은 else 블록의 범위내에서만 사용가능한 변수를 선언하는 것이다."

```go
if n := rand.Intn(10); n == 0 {
    fmt.Println("zero", n)
} else if n > 5 {
    fmt.Println("range 5 < x", n)
} else {
    fmt.Println("range 0 < x <= 5", n)
}

fmt.Println(n) // <- compile error
```

## for

p111. "for-range 루프는 두 개의 변수를 얻는다는 부분이 재미있다. 첫 번째 변수는 현재 순회중인 자료구조에 있는 값의 위치이고 두 번쨰는 해당 위치의 값이다."

```go
arr := []int{1, 2, 3, 4, 5, 6}
for i, v := range arr {
    fmt.Println(i, v)
}
```


```go
arr := []int{1, 2, 3, 4, 5, 6}
for _, elem := range arr {
    fmt.Println(elem)
}
```

p111. "맵을 순회하는 경우에는 i 대신에 k를 사용한다."

```go
m := make(map[string]string)
m["hello"] = "world"
m["go"] = "lang"

for k, v := range m {
    fmt.Println(k, v)
}
```

p114. "언급된 이 두문제를 막기 위해, Go 팀은 맵 구현에 두 가지 변경을 했다. 첫 번째는 맵을 위해 해시 알고리즘을 수정하여 맵 변수가 생성될 때마다 무작위의 숫자를 포함하도록 했다. 두 번째는 맵을 for-range로 순회의 순서를 루프가 반복될 때 마다 조금씩 달라지게 했다."

p116. "for-range 값은 복사본, 고루틴을 for-range 루프에서 실행하게 된다면 고루틴으로 인덱스와 값을 전달하는 방법에 매우 주의해야 할 것이다."

p118. "for-range 루프 내에서 if, continue, break를 적절히 조합하여 구헝할 수 있지만 표준 for 루프는 순회의 처음과 끝을 나타내는 조금 더 명확한 방법이다."

for-range + if + continue + break
```go
arr := []int{1, 2, 3, 4, 5, 6}

for i, v := range arr {
    if i == 0 {
        continue
    }

	if i == len(arr)-1 {
        break
    }
    fmt.Println(i, v)
}
```

표준 for
```go
arr := []int{1, 2, 3, 4, 5, 6}

for i := 1; i < len(arr)-1; i++ {
    fmt.Println(i, arr[i])
}
```

## switch

p121. "비어 있는 case는 아무 일도 일어나지 않는다."

```go
words := make([]string, 0)
words = append(words, "a", "cow", "smile", "gopher", "octopus", "anthropologist")

for _, word := range words {
	switch size := len(word); size {
	case 1, 2, 3, 4:
		fmt.Println(word, "is a short word")
    case 5:
		fmt.Println(word, "is length ", size)
    case 6, 7, 8, 9:
    default:
        fmt.Println(word, "is long word")
    }
}
```

p122. "만약 fallthrough 키워드를 사용할 필요를 발견한다면, 로직을 재구성하거나 case 문 간에 의존성을 제거해보도록 하자."

p124. "공백 switch는 각 case 문에 불리언 결과를 내는 비교도 모두 가능하다."

```go
words := make([]string, 0)
words = append(words, "a", "cow", "smile", "gopher", "octopus", "anthropologist")
for _, word := range words {
	switch size := len(word); {
	case size < 5:
		fmt.Println(word, "is a short word")
	case size == 5:
		fmt.Println(word, "is length ", size)
	case 5 < size && size < 10:
	default:
		fmt.Println(word, "is long word")
	}
}
```

## goto

p126. "Go의 네 번째 제어문이 있는데 절대 사용하지 않을 가능성이 있다."

p128. "goto문 사용을 하지 않도록 노력해야 한다. 하지만 가독성을 위해서 드문 상황에서는 선택적으로 사용하도록 하자." -> 가독성을 좋게 짜서 goto문을 사용하지 말자.