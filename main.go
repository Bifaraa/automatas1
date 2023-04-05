package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Graph struct {
	adjacencyList     map[string][]string
	termina           map[string][]bool
	adjacencyListPeso map[string]map[string][]int
	mensaje           map[string][]string
}

func (g *Graph) AddNode(node string, termina bool, mensaje string) {
	if _, ok := g.adjacencyList[node]; !ok {
		g.adjacencyList[node] = []string{}
		/* 		g.termina = make([]bool, 3) */
		g.termina[node] = append(g.termina[node], termina)
		/* g.mensaje = make(map[string][]string) */
		g.adjacencyListPeso[node] = make(map[string][]int)
		g.mensaje[node] = append(g.mensaje[node], mensaje)
	}
}

func (g *Graph) AddEdge(src string, dest string, peso []int) {
	g.adjacencyList[src] = append(g.adjacencyList[src], dest)
	for _, valor := range peso {
		g.AddEdgePeso(valor, src, dest)
	}
}

func (g *Graph) AddEdgePeso(peso int, src string, dest string) {
	if _, ok := g.adjacencyListPeso[src]; !ok {
		g.adjacencyListPeso[src] = make(map[string][]int)
	}
	if _, ok := g.adjacencyListPeso[src][dest]; !ok {
		g.adjacencyListPeso[src][dest] = []int{}
	}
	g.adjacencyListPeso[src][dest] = append(g.adjacencyListPeso[src][dest], peso)
}

func (g *Graph) PrintGraph() {
	for node, neighbors := range g.adjacencyList {
		fmt.Printf("%s -> ", node)
		for i, neighbor := range neighbors {
			pesos := g.adjacencyListPeso[node][neighbor]
			if len(pesos) > 1 {
				fmt.Printf("%s(%v)", neighbor, pesos)
			} else {
				fmt.Printf("%s(%v)", neighbor, pesos[0])
			}
			if i != len(neighbors)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}

// Estructura numero del JSON

type Data struct {
	List []struct {
		Valor string `json:"valor"`
	} `json:"list"`
}

func (g *Graph) BFS(start string) {
	visited := make(map[string]bool)
	queue := []string{start}

	for len(queue) > 0 {
		// Pop the first node from the queue
		node := queue[0]
		queue = queue[1:]

		// Skip if the node is already visited
		if visited[node] {
			continue
		}

		// Mark the node as visited
		visited[node] = true

		// Print the visited node
		fmt.Printf("%s ", node)

		// Add the neighbors of the node to the queue
		for _, neighbor := range g.adjacencyList[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
}

func (g *Graph) RecorrerAutomata(nums []int) string {
	// Se empieza en el primer nodo del grafo
	currentNode := "a"
	// Se recorre el arreglo de números
	for i, num := range nums {

		// Se busca si el número es un peso de alguna de las aristas del nodo actual
		adjacentNodes, ok := g.adjacencyList[currentNode]
		if !ok {
			return "" //
		}

		nextNode := ""
		for _, adjNode := range adjacentNodes {
			weights, ok := g.adjacencyListPeso[currentNode][adjNode]
			if !ok {
				continue // No hay pesos definidos para esta arista
			}
			for _, weight := range weights {
				if weight == num {
					nextNode = adjNode
					break
				}
			}
			if nextNode != "" {
				break
			}
		}

		if nextNode == "" {
			return "No hay una arista que coincida con el número"
		}
		currentNode = nextNode

		// Si este es el último número del arreglo, se muestra el mensaje del nodo actual (si existe)
		if (i == len(nums)-1 && g.termina[currentNode][0]) || (i == len(nums)-1 && currentNode == "a") {

			msg, ok := g.mensaje[currentNode]
			if ok {
				return msg[0]
			}
		}
	}
	return "Terminó el recorrido sin encontrar un nodo con mensaje asociado"
}

/*
func transfornJSON() [][]int {
	var datos Data
	err = json.Unmarshal(data, &datos)
	if err != nil {
		fmt.Println("Error al parsear el JSON:", err)
		return nil
	}

	fmt.Println(datos)
	var arreglo [][]int

	// Transformar cada valor de "valor" en un arreglo de int
	for _, item := range datos.List {
		num, err := strconv.Atoi(item.Valor)
		if err != nil {
			panic(err)
		}

		// divide el número por digitos y los guarda en un arreglo
		var digits []int
		for num > 0 {
			digits = append(digits, num%10)
			num /= 10
		}

		// Invertir el arreglo para que quede en el orden correcto
		for i := 0; i < len(digits)/2; i++ {
			j := len(digits) - i - 1
			digits[i], digits[j] = digits[j], digits[i]
		}

		arreglo = append(arreglo, digits)
		fmt.Println(digits)
	}
	return arreglo
}
*/

// Leer el archivo JSON
func cargarJSON() [][]int {
	data, err := ioutil.ReadFile("entradas.json")
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return nil
	}
	// Parsear el JSON
	var datos Data
	err = json.Unmarshal(data, &datos)
	if err != nil {
		fmt.Println("Error al parsear el JSON:", err)
		return nil
	}

	fmt.Println(datos)
	var arreglo [][]int

	// Transformar cada valor de "valor" en un arreglo de int
	for _, item := range datos.List {
		num, err := strconv.Atoi(item.Valor)
		if err != nil {
			panic(err)
		}

		// divide el número por digitos y los guarda en un arreglo
		var digits []int
		for num > 0 {
			digits = append(digits, num%10)
			num /= 10
		}

		// Invertir el arreglo para que quede en el orden correcto
		for i := 0; i < len(digits)/2; i++ {
			j := len(digits) - i - 1
			digits[i], digits[j] = digits[j], digits[i]
		}

		arreglo = append(arreglo, digits)
		fmt.Println(digits)
	}
	return arreglo
}

func crearGrafo() Graph {
	g := Graph{
		adjacencyList:     make(map[string][]string),
		termina:           make(map[string][]bool),
		adjacencyListPeso: make(map[string]map[string][]int),
		mensaje:           make(map[string][]string),
	}

	// agregar nodos
	g.AddNode("a", false, "se trabo en el inicio")
	g.AddNode("b", true, "Es par termina en 0")
	g.AddNode("c", true, "Es par termina en 5")
	//g.AddNode("d", true, "activo lambda, es menor a cero")

	pesosCero := []int{0}
	pesosN := []int{1, 2, 3, 4, 6, 7, 8, 9}
	/* 	pesosAD := []int{-1} */
	pesosCinco := []int{5}

	// agregar aristas
	g.AddEdge("a", "a", pesosN)
	g.AddEdge("a", "b", pesosCero)
	g.AddEdge("a", "c", pesosCinco)
	/* g.AddEdge("a", "d", pesosAD) */
	g.AddEdge("b", "a", pesosN)
	g.AddEdge("b", "c", pesosCinco)
	g.AddEdge("c", "a", pesosN)
	g.AddEdge("c", "b", pesosCero)
	g.AddEdge("c", "c", pesosCinco)
	g.AddEdge("b", "b", pesosCero)
	return g
}

func correrAutomata(g Graph) []string {
	//g := crearGrafo()
	//g.PrintGraph()
	//dotCode := g.GenerateDotCode()
	result := []string{""}
	arreglo := cargarJSON()
	for _, arr := range arreglo {
		res := g.RecorrerAutomata(arr)
		result = append(result, res)
		//fmt.Println(result)
	}
	return result
}

func (g *Graph) GenerateDotCode() string {
	nodeSet := make(map[string]bool)
	edgeSet := make(map[string]bool)
	dotCode := "digraph {"

	for src, dests := range g.adjacencyListPeso {
		if !nodeSet[src] {
			dotCode += fmt.Sprintf("%s [label=\"%s\"];", src, src)
			nodeSet[src] = true
		}
		for dest, pesos := range dests {
			edge := fmt.Sprintf("%s -> %s [label=\"%v\"]", src, dest, pesos)
			if !edgeSet[edge] {
				dotCode += fmt.Sprintf("%s;", edge)
				edgeSet[edge] = true
			}
			if !nodeSet[dest] {
				dotCode += fmt.Sprintf("%s [label=\"%s\"];", dest, dest)
				nodeSet[dest] = true
			}
		}
	}

	dotCode += "}"
	return dotCode
}

type Num struct {
	numero string `json:"num"`
}

func main() {

	/*recorrido por anchura
	fmt.Println("BFS traversal starting from node a:") */
	/* g.BFS("a") */

	//g.PrintGraph()
	/* arregloPrueba := []int{2, 5, 5} */

	g := crearGrafo()

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/run", func(c *fiber.Ctx) error {
		result := correrAutomata(g)
		return c.JSON(&fiber.Map{
			"data": result,
		})
	})
	app.Get("/dot", func(c *fiber.Ctx) error {
		dotCode := g.GenerateDotCode()
		return c.JSON(&fiber.Map{
			"data": dotCode,
		})
	})
	app.Get("/data", func(c *fiber.Ctx) error {
		data, err := ioutil.ReadFile("entradas.json")
		if err != nil {
			fmt.Println("Error al leer el archivo:", err)
		}
		// Parsear el JSON
		var datos Data
		err = json.Unmarshal(data, &datos)
		if err != nil {
			fmt.Println("Error al parsear el JSON:", err)
			return nil
		}
		return c.JSON(&fiber.Map{
			"data": datos,
		})
	})
	app.Post("/ejecutar", func(c *fiber.Ctx) error {
		p := new(Num)
		if err := c.BodyParser(p); err != nil {
			return err
		}
		fmt.Println(p.numero)
		return c.JSON(&fiber.Map{
			"data": "ok",
		})
	})
	app.Listen(":3000")

}
