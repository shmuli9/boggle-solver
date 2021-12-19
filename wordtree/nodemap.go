package wordtree

type NodeMap map[string]*WTNode

func (nodeMap NodeMap) String() string {
	output := ""
	for _, element := range nodeMap {
		output += element.String()
	}
	return output
}
