//Требуется отыскать самый выгодный маршрут между городами.
//Требования: время работы O((N+M)logN), где N-количество городов, M-известных дорог между ними
// Алгоримт Дейкстры
#include "vector"
#include "iostream"
#include "set"

struct Node{
    Node(int To, int Weight){
        to = To;
        weight = Weight;
    }
    int to;
    int weight;
};


class Graph{
private:
    std::vector<std::vector<Node>> adjlist;
public:
    Graph(int countvertex){
        adjlist.resize(countvertex);
    }

    ~Graph() {}

    Graph(const Graph& graph){
        adjlist.resize(graph.VerticesCount());
        adjlist = graph.adjlist;
    }

    int VerticesCount() const {
        return adjlist.size();
    }

    std::vector<Node> GetNextVertices(int vertex) const {
        return adjlist[vertex];
    }

    void AddEdge(int from, int to, int weight){
        adjlist[from].push_back(Node(to, weight));
        adjlist[to].push_back(Node(from, weight));
    }
};

int SearchWay(const Graph& graph, int startvertex, int endvertex){
    std::vector<int> distances (graph.VerticesCount(), 99999999);
    std::vector<bool> visited (graph.VerticesCount(), false);
    std::set<std::pair<int, int>> queue; // (d[v], v) - пары такого формата при ум приоритете удаляем пару и добавляем новую
    queue.insert(std::make_pair(0, startvertex));
    distances[startvertex] = 0;
    while(!queue.empty()){
        int from = queue.begin()->second;
        visited[from] = true;
        queue.erase(queue.begin());
        std::vector<Node> NextVertexes = graph.GetNextVertices(from);
        for (int i = 0; i < NextVertexes.size(); i++){
            Node node = NextVertexes[i];
            if (!visited[node.to]){
                int new_distance = distances[from] + node.weight;
                if(distances[node.to] > new_distance){
                    distances[node.to] = new_distance;
                    if (queue.find(std::make_pair(node.weight, node.to)) != queue.end()) {
                        queue.erase(queue.find(std::make_pair(node.weight, node.to)));
                    }
                    if (!visited[node.to]){
                        queue.insert(std::make_pair(new_distance, node.to));
                    }
                }
            }
        }

    }
    return distances[endvertex];
}

int main(){
    int N, M; // количество вершин и ребер
    std::cin >> N >> M;
    Graph g(N);
    for (int i=0; i<M; ++i){
        int from, to, weight;
        std::cin >> from >> to >> weight;
        g.AddEdge(from, to, weight);
    }
    int start, end;
    std::cin >> start >> end;
    std::cout << SearchWay(g, start, end);
    return 0;
}
