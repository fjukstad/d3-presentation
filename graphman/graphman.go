package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Graph struct {
	Nodes []Node
	Edges []Edge
}

type Node struct {
	Id int `json:"id"`
}

type Edge struct {
	Source int `json:"source"`
	Target int `json:"target"`
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./graph.html")
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	// "upgrade" http connection to websocket connection
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		time.Sleep(1 * time.Second)

		// Turn graph into json
		b, err := json.Marshal(g)
		if err != nil {
			fmt.Println("error:", err)
		}

		// send to client
		err = conn.WriteMessage(websocket.TextMessage, []byte(b))
		if err != nil { // this happens when connection has been closed
			return
		}
	}
}

var g *Graph

func (g Graph) hasNode(a Node) bool {
	for _, n := range g.Nodes {
		if a.Id == n.Id {
			return true
		}
	}
	return false
}

func (g Graph) getIndex(a Node) int {
	for i, n := range g.Nodes {
		if a.Id == n.Id {
			return i
		}
	}
	return -1
}

func updateGraph() {
	max := 100
	for {

		// Create a new node with a random id
		a_id := rand.Intn(max)
		a := Node{a_id}
		if !g.hasNode(a) {
			g.Nodes = append(g.Nodes, a)
		}

		// Create a second node also with a random id
		b_id := rand.Intn(max)
		b := Node{b_id}
		if !g.hasNode(b) {
			g.Nodes = append(g.Nodes, b)
		}

		// If we're lucky, create a new edge as well
		if rand.Intn(max) < max/2 {
			source := g.getIndex(a)
			target := g.getIndex(b)

			e := Edge{source, target}
			g.Edges = append(g.Edges, e)
		}

		time.Sleep(1 * time.Second)

	}
}

func main() {

	g = new(Graph)

	go updateGraph()

	// set up and run http server
	routes := mux.NewRouter()
	routes.HandleFunc("/", htmlHandler)
	routes.HandleFunc("/graph", websocketHandler)
	http.ListenAndServe("localhost:4040", routes)

}
