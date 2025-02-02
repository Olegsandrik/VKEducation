//Дан невзвешенный неориентированный граф.
// В графе может быть несколько кратчайших путей между какими-то вершинами.
// Найдите количество различных кратчайших путей между заданными вершинами. Требуемая сложность O(V+E).

#include "vector"
#include "queue"
#include "iostream"

struct myIGraph {
    virtual ~myIGraph() {}
    virtual void AddEdge(int from, int to) = 0;

    virtual int VerticesCount() const  = 0;

    virtual std::vector<int> GetNextVertices(int vertex) const = 0;
    virtual std::vector<int> GetPrevVertices(int vertex) const = 0;
};

class myListGraph: public myIGraph {
public:
    myListGraph( int vertexcount );

    myListGraph(const myIGraph& graph);

    ~myListGraph() override = default;

    virtual void AddEdge(int from, int to) override;


    virtual int VerticesCount() const  override;

    virtual std::vector<int> GetNextVertices(int vertex) const override;

    virtual std::vector<int> GetPrevVertices(int vertex) const override;
private:
    std::vector<std::vector<int>> adjlist;
    std::vector<std::vector<int>> prevadjlist;
};

myListGraph::myListGraph( int vertexcount ){
    adjlist.resize(vertexcount);
    prevadjlist.resize(vertexcount);
}

myListGraph::myListGraph(const myIGraph& graph){
    adjlist.resize(graph.VerticesCount());
    prevadjlist.resize(graph.VerticesCount());
    for (int i=0; i<graph.VerticesCount(); ++i){
        adjlist[i] = graph.GetNextVertices( i );
        prevadjlist[i] = graph.GetPrevVertices( i );
    }
}

void myListGraph::AddEdge(int from, int to){
    adjlist[from].push_back(to);
    prevadjlist[to].push_back(from);
}
int myListGraph::VerticesCount() const {
    return adjlist.size();
}


std::vector<int> myListGraph::GetNextVertices(int vertex) const {
    return adjlist[vertex];
}

std::vector<int> myListGraph::GetPrevVertices(int vertex) const{
    return prevadjlist[vertex];
}

int BFS(const myIGraph& graph, int startvertex, int endvertex){
    std::vector<int> countways(graph.VerticesCount(), 0);
    std::vector<int> distences(graph.VerticesCount(), -1);
    std::queue<int> bfsQ;
    distences[startvertex] = 0;
    countways[startvertex] = 1;
    bfsQ.push(startvertex);
    while (!bfsQ.empty()){
        int current = bfsQ.front();
        bfsQ.pop();
        std::vector<int> nextvertexes = graph.GetNextVertices(current);
        for (int i=0; i<nextvertexes.size(); ++i){
            if (distences[nextvertexes[i]]==-1){
                bfsQ.push(nextvertexes[i]);
                distences[nextvertexes[i]] = distences[current]+1;
                countways[nextvertexes[i]] = countways[current];
            } else if(distences[nextvertexes[i]]-distences[current] == 1){
                countways[nextvertexes[i]]+=countways[current];
            }
        }
    }
    return countways[endvertex];
}





int main(){
    int N, M; // количество вершин и ребер
    std::cin >> N >> M;
    myListGraph g(N);
    for (int i=0; i<M; ++i){
        int from, to;
        std::cin >> from >> to;
        g.AddEdge(from, to);
        g.AddEdge(to, from);
    }
    int start, end;
    std::cin >> start >> end;
    std::cout << BFS(g, start, end);
    return 0;
}
