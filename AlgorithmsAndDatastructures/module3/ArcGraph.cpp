#include "ArcGraph.h"

std::vector<int> ArcGraph::GetPrevVertices(int vertex) const{
    std::vector<int> ans;
    for (int i=0; i<massive.size(); ++i){
        std::pair<int, int> current = massive[i];
        if (current.second == vertex){
            ans.push_back(current.first);
        }
    }
    return ans;
};

std::vector<int> ArcGraph::GetNextVertices(int vertex) const{
    std::vector<int> ans;
    for (int i=0; i<massive.size(); ++i){
        std::pair<int, int> current = massive[i];
        if (current.first == vertex){
            ans.push_back(current.second);
        }
    }
    return ans;
};

int ArcGraph::VerticesCount() const{
    return count;
};


void ArcGraph::AddEdge(int from, int to) {
    std::pair<int, int> myPair(from, to);
    massive.push_back(myPair);
};

ArcGraph::ArcGraph(const IGraph& graph){
    count = graph.VerticesCount();
    for (int i=0; i<count; ++i){
        std::vector<int> vertexes = graph.GetNextVertices(i);
        while (!vertexes.empty()){
            AddEdge(i, vertexes.back());
            vertexes.pop_back();
        }
    }
};

ArcGraph::ArcGraph(int vertexcount){
    int maxpairs = vertexcount*(vertexcount-1);
    massive.resize(maxpairs);
    count = vertexcount;
};
