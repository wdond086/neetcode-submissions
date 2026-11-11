type BrowserHistory struct {
    head *HistoryNode
	tail *HistoryNode
}

type HistoryNode struct {
	url	string
	next *HistoryNode
	prev *HistoryNode
}


func Constructor(homepage string) BrowserHistory {
    newNode := &HistoryNode{url: homepage}
	return BrowserHistory{head: newNode, tail: newNode}
}


func (this *BrowserHistory) Visit(url string)  {
    newNode := &HistoryNode{
		url: url,
		prev: this.tail,
	}
	this.tail.next = newNode
	this.tail = newNode
}


func (this *BrowserHistory) Back(steps int) string {
    stepsTaken := 0
	for stepsTaken < steps && this.tail.prev != nil {
		this.tail = this.tail.prev
		stepsTaken++
	}
	return this.tail.url
}


func (this *BrowserHistory) Forward(steps int) string {
    stepsTaken := 0
	for stepsTaken < steps && this.tail.next != nil {
		this.tail = this.tail.next
		stepsTaken++
	}
	return this.tail.url
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */