#include "Matrix.h"

std::vector<int> MatrixGraph::GetPrevVertices(int vertex) const{
    std::vector<int> ans;
    for (int i=0; i<matrix.size(); ++i){
        if (matrix[i][vertex]==1){
            ans.push_back(i);
        }
    }
    return ans;
};

std::vector<int> MatrixGraph::GetNextVertices(int vertex) const{
    std::vector<int> ans;
    for (int i=0; i<matrix.size(); ++i){
        if (matrix[vertex][i]==1){
            ans.push_back(i);
        }
    }
    return ans;
};

int MatrixGraph::VerticesCount() const{
    return matrix.size();
};

void MatrixGraph::AddEdge(int from, int to) {
    matrix[from][to] = 1;
};

MatrixGraph::MatrixGraph( int vertexcount ){
    matrix.resize(vertexcount);
    for (int i=0; i<vertexcount; ++i){
        matrix[i].resize(vertexcount);
    }
}

MatrixGraph::MatrixGraph(const IGraph& graph){
    matrix.resize(graph.VerticesCount());
    for (int i=0; i<graph.VerticesCount(); ++i){
        matrix[i].resize(graph.VerticesCount());
    }

    for(int i=0; i<matrix.size(); ++i){
        std::vector<int> vertexes = graph.GetNextVertices(i);
        while (!vertexes.empty()){
            int current = vertexes.back();
            vertexes.pop_back();
            matrix[i][current] = 1;
        }
    }
}

