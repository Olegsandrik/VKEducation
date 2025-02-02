#include "vector"

struct IGraph {
    virtual ~IGraph() {}
    virtual void AddEdge(int from, int to) = 0;

    virtual int VerticesCount() const  = 0;

    virtual std::vector<int> GetNextVertices(int vertex) const = 0;
    virtual std::vector<int> GetPrevVertices(int vertex) const = 0;
};

#ifndef IGraph_H
#define IGraph_H

#endif //UNTITLED_IGRAPH_H
