package validPath

func validPath(n int, edges [][]int, source int, destination int) bool {
    p := make([]int,n)
    for i:= range p {
        p[i] = i;
    }
    var find func(x int) int
    find = func(x int) int {
        if x != p[x]{
            p[x] = find(p[x])
        }
        return p[x]
    }

    for _,edge := range edges {
        p[find(edge[0])] = find(edge[1])
    }
    return find(source) == find(destination)
}