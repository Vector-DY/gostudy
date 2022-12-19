package main

import "fmt"

const N = 510

var n,m int 
var g [N][N] int
var dist [N] int
var st [N] bool

func dijkstra() int{
	for i := 1;i <= n;i++{
		dist[i] = 0x3f3f3f3f
	}
	dist[1] = 0

	for i := 0;i < n;i++{
		t := -1
		for j := 1;j <= n;j++{
			if !st[j] && (t == -1 || dist[t] > dist[j]){
				t = j
			}
		}
		st[t] = true

		for j := 1;j <= n;j++{
			dist[j] = min(dist[j],dist[t] + g[t][j])
		}
	}

	if dist[n] == 0x3f3f3f3f {
		return 1
	}
	return dist[n]
}

func min(a,b int) int{
	if a > b{
		return b
	}
	return a
}

func main(){
	fmt.Scanf("%d%d",&n,&m)

	for i := 1;i <= n;i++{
		for j := 1;j <= n;j++{
			g[i][j] = 0x3f3f3f3f
		}
	}

	for ;m != 0;m--{
		var a,b,c int
		fmt.Scanf("%d%d%d",&a,&b,&c)

		g[a][b] = min(g[a][b],c)
	}

	fmt.Println(dijkstra())
}