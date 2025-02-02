#include "vector"
#include "ArcGraph.h"

class ListGraph: public IGraph {
public:
    ListGraph( int vertexcount );

    ListGraph(const IGraph& graph);

    ~ListGraph() override = default;

    virtual void AddEdge(int from, int to) override;


    virtual int VerticesCount() const  override;

    virtual std::vector<int> GetNextVertices(int vertex) const override;

    virtual std::vector<int> GetPrevVertices(int vertex) const override;
private:
    std::vector<std::vector<int>> adjlist;
    std::vector<std::vector<int>> prevadjlist;
};



#ifndef UNTITLED_LISTGRAPH_H
#define UNTITLED_LISTGRAPH_H

#endif //UNTITLED_LISTGRAPH_H
