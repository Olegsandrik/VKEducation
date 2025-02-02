#include "ListGraph.h"

ListGraph::ListGraph( int vertexcount ){
    adjlist.resize(vertexcount);
    prevadjlist.resize(vertexcount);
}

ListGraph::ListGraph(const IGraph& graph){
    adjlist.resize(graph.VerticesCount());
    prevadjlist.resize(graph.VerticesCount());
    for (int i=0; i<graph.VerticesCount(); ++i){
        adjlist[i] = graph.GetNextVertices( i );
        prevadjlist[i] = graph.GetPrevVertices( i );
    }
}

void ListGraph::AddEdge(int from, int to){
    adjlist[from].push_back(to);
    prevadjlist[to].push_back(from);
}
int ListGraph::VerticesCount() const {
    return adjlist.size();
}


std::vector<int> ListGraph::GetNextVertices(int vertex) const {
    return adjlist[vertex];
}

std::vector<int> ListGraph::GetPrevVertices(int vertex) const{
    return prevadjlist[vertex];
}
