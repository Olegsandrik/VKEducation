#include "vector"
#include "SetGraph.h"
class MatrixGraph: public IGraph{
private:
    std::vector<std::vector<int>> matrix;
public:
    MatrixGraph( int vertexcount );

    MatrixGraph(const IGraph& graph);

    virtual ~MatrixGraph() override = default;

    virtual void AddEdge(int from, int to) override;

    virtual int VerticesCount() const override;

    virtual std::vector<int> GetNextVertices(int vertex) const override;

    virtual std::vector<int> GetPrevVertices(int vertex) const override;
};




#ifndef UNTITLED_MATRIX_H
#define UNTITLED_MATRIX_H

#endif //UNTITLED_MATRIX_H
