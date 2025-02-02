#include "vector"
#include "Matrix.h"

class ArcGraph: public IGraph{
private:
    std::vector<std::pair<int, int>> massive;
    int count = 0;
public:
    virtual ~ArcGraph() override = default;

    ArcGraph(int vertexcount);

    ArcGraph(const IGraph& graph);

    virtual void AddEdge(int from, int to) override;

    virtual int VerticesCount() const  override;

    virtual std::vector<int> GetNextVertices(int vertex) const override;

    virtual std::vector<int> GetPrevVertices(int vertex) const override;
};


#ifndef UNTITLED_ARCGRAPH_H
#define UNTITLED_ARCGRAPH_H

#endif //UNTITLED_ARCGRAPH_H
