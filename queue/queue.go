package queue

import (
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
)

// Server 对外服务接口
type Server interface {
	Len() int
	Name() string
	Front() *Node
	Get(index int) *Data
	Push(*Data) *Node
	Unshift(*Data) *Node
	Shift() *Node
	Pop() *Node
	InsertBefore(*Node, *Node) *Node
	InsertAfter(*Node, *Node) *Node
	Header() *Node
	Tailed() *Node
	WriteBuffer(interface{}) chan interface{}
	Buffer() <-chan interface{}
	Remove(*Node) interface{}
	Listen(Call) error
	List() []Data
	Loop(Call) error
}

// Node ...
type Node struct {
	queue      *queue
	prev, next *Node
	Data       *Data
}

// Call ...
type Call func(*Node) error

// NewNode ...
func NewNode(data *Data) *Node {
	node := new(Node).init()
	node.Data = data
	return node
}

// Next returns the next list Node or nil.
func (e *Node) init() *Node {
	e.Data = nil
	e.prev = nil
	e.next = nil
	e.queue = nil
	return e
}

// Next returns the next list Node or nil.
func (e *Node) Next() *Node {
	if p := e.next; e.queue != nil && p != &e.queue.head {
		return p
	}
	return nil
}

// Prev returns the previous list Node or nil.
func (e *Node) Prev() *Node {
	if p := e.prev; e.queue != nil && p != &e.queue.head {
		return p
	}
	return nil
}

// Queue ...
type Queue queue

// queue ...
type queue struct {
	opts       options
	len        int
	head, tail Node
}

// NewQueue new queue
func NewQueue(opts ...Option) Server {
	queue := new(queue).Init()
	queue.opts = newOptions(opts...)
	return queue
}

// Init 初始化队列
func (h *queue) Init() *queue {
	h.head.next = &h.head
	h.tail.prev = &h.tail
	h.len = 0
	return h
}

// Len 长度
func (h *queue) Len() int { return h.len }

// Name 名称
func (h *queue) Name() string {
	return h.opts.name
}

// lazyInit lazily initializes a zero List value.
func (h *queue) lazyInit() {
	if h.head.next == nil {
		h.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (h *queue) insert(e, at *Node) *Node {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.queue = h
	h.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (h *queue) insertValue(v *Node, at *Node) *Node {
	return h.insert(v, at)
}

// remove removes e from its list, decrements l.len, and returns e.
func (h *queue) remove(node *Node) *Node {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil // avoid memory leaks
	node.prev = nil // avoid memory leaks
	node.queue = nil
	h.len--
	return node
}

// move moves e to next to at and returns e.
func (h *queue) move(e, at *Node) *Node {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e

	return e
}

// Front 获取下一个
func (h *queue) Front() *Node {
	if h.len == 0 {
		return nil
	}
	return h.head.next
}

// Back returns the last element of list l or nil if the list is empty.
func (h *queue) Back() *Node {
	if h.len == 0 {
		return nil
	}
	return h.head.prev
}

// Header 头部
func (h *queue) Header() *Node {
	return &h.head
}

// Tailed 尾部
func (h *queue) Tailed() *Node {
	return &h.tail
}

// WriteBuffer 写入数据到缓冲
func (h *queue) WriteBuffer(data interface{}) chan interface{} {
	h.opts.buffer <- data
	return h.opts.buffer
}

// Buffer 获取缓冲
func (h *queue) Buffer() <-chan interface{} {
	return h.opts.buffer
}

// InsertBefore 之前插入
func (h *queue) InsertBefore(v *Node, node *Node) *Node {
	if node.queue != h {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return h.insertValue(v, node.prev)
}

// InsertAfter 之后插入
func (h *queue) InsertAfter(v *Node, node *Node) *Node {
	if node.queue != h {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return h.insertValue(v, node)
}

// Remove 删除
func (h *queue) Remove(node *Node) interface{} {
	h.opts.mutex.Lock()
	defer h.opts.mutex.Unlock()
	if node.queue == h {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		h.remove(node)
	}
	return node.Data.Content
}

func (h *queue) Loop(call Call) error {
	for node := h.Front(); node != nil; node = node.Next() {
		if err := call(node); err != nil {
			break
		}
	}
	return nil
}

// Listen 持续监听
func (h *queue) Listen(call Call) error {
	go func() {
		for {
			for node := h.Front(); node != nil; node = node.Next() {
				if err := call(node); err != nil {
					break
				}
			}
		}
	}()

	ch := make(chan os.Signal, 1)
	if h.opts.signal {
		signal.Notify(ch, os.Kill)
	}

	// wait on kill signal
	<-ch
	return nil
}

// Get ...
func (h *queue) Get(index int) *Data {
	if h.len == 0 || h.len < int(index) {
		return nil
	}
	i := 0
	for node := h.Front(); node != nil; node = node.Next() {
		if i == index {
			return node.Data
		}
		i++
	}
	return nil
}

// Unshift 开头插入
func (h *queue) Unshift(v *Data) *Node {
	h.lazyInit()
	return h.insertValue(NewNode(v), &h.head)
}

// Push 压入末尾
func (h *queue) Push(v *Data) *Node {
	h.lazyInit()
	return h.insertValue(NewNode(v), h.head.prev)
}

// Shift 移出开始第一个
func (h *queue) Shift() *Node {
	if h.len <= 0 {
		return nil
	}
	h.opts.mutex.Lock()
	defer h.opts.mutex.Unlock()
	return h.remove(&h.head)
}

// Pop 弹出结尾最后一个
func (h *queue) Pop() *Node {
	if h.len <= 0 {
		return nil
	}
	h.opts.mutex.Lock()
	defer h.opts.mutex.Unlock()

	return h.remove(h.Back())
}

// PushFrontList inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (h *queue) PushFrontList(other *queue) {
	h.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		h.insertValue(e, &h.head)
	}
}

// Unique 去重
func (h *queue) Unique() error {
	return nil
}

// Replace 替换
func (h *queue) Replace(old *Node, new *Node) error {
	return nil
}

func (h *queue) List() []Data {
	list := make([]Data, 0, h.len)
	for node := h.Front(); node != nil; node = node.Next() {
		list = append(list, *node.Data)
	}
	return list
}

// Data ...
type Data struct {
	UUID     string
	CreateAt time.Time
	ExpireAt *time.Time
	Content  interface{}
}

// NewData ...
func NewData(content interface{}) *Data {
	uuid, _ := uuid.NewUUID()
	return &Data{
		UUID:     uuid.String(),
		CreateAt: time.Now(),
		Content:  content,
	}
}

// NewExpireData ...
func NewExpireData(content interface{}, expire time.Time) *Data {
	uuid, _ := uuid.NewUUID()
	return &Data{
		UUID:     uuid.String(),
		CreateAt: time.Now(),
		Content:  content,
		ExpireAt: &expire,
	}
}

// AddData 插入数据到队列
func (h *queue) AddData(content interface{}) {
	h.opts.mutex.Lock()
	defer h.opts.mutex.Unlock()
}
