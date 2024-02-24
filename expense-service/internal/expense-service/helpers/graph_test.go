package helpers

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/google/uuid"
)

func Test_Graph(t *testing.T) {
	var adjacencyList = map[string][]Connection{
		"ABC": {{Name: "DEF", Value: 3.0}, {Name: "DEF", Value: 4.0}},
		"DEF": {{Name: "GHI", Value: 3.4}, {Name: "ABC", Value: 99.24}},
		"GHI": {{Name: "DEF", Value: 11.3}},
		"JKL": {{Name: "ABC", Value: 978}, {Name: "DEF", Value: 6477}, {Name: "GHI", Value: 46578}},
	}

	g := NewGraph(WithAdjacencyList(adjacencyList))

	for key := range adjacencyList {
		totalAmount := g.UserLentTotalAmount(key)
		fmt.Println("cost for " + key + ": " + strconv.FormatFloat(totalAmount, 'f', -1, 64))
	}

}

func NewGraph(opts ...GraphOption) *GroupGraph {
	g := &GroupGraph{GroupMembers: map[string]*Member{}}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

func WithAdjacencyList(list map[string][]Connection) GraphOption {
	return func(this *GroupGraph) {
		for member, edges := range list {
			if _, ok := this.GroupMembers[member]; !ok {
				this.AddMember(member)
			}

			for _, edge := range edges {
				if _, ok := this.GroupMembers[edge.Name]; !ok {
					this.AddMember(edge.Name)
				}

				this.AddEdge(member, edge.Name, edge.Value)
			}
		}
	}
}

type GraphOption func(this *GroupGraph)

type GroupGraph struct {
	GroupID      string
	GroupMembers map[string]*Member
}

type Member struct {
	UserID string
	Edges  map[string]*Edge
}

type Edge struct {
	EdgeID string
	Amount float64
	Member *Member
}

type Connection struct {
	Name  string
	Value float64
}

func (this *GroupGraph) AddMember(UserID string) {
	this.GroupMembers[UserID] = &Member{UserID: UserID, Edges: map[string]*Edge{}}
}

func (this *GroupGraph) AddEdge(sourceMember, destMember string, amount float64) {
	if _, ok := this.GroupMembers[sourceMember]; !ok {
		return
	}
	if _, ok := this.GroupMembers[destMember]; !ok {
		return
	}

	this.GroupMembers[sourceMember].Edges[destMember] = &Edge{EdgeID: uuid.New().String(),
		Amount: amount,
		Member: this.GroupMembers[destMember]}
}

func (this *GroupGraph) UserLentTotalAmount(sourceMember string) float64 {
	result := float64(0)

	if _, ok := this.GroupMembers[sourceMember]; !ok {
		return result
	}

	for _, edge := range this.GroupMembers[sourceMember].Edges {
		result += edge.Amount
	}

	return result
}

// func (graph *GroupGraph) String() string {

// }