// https://leetcode.com/problems/min-stack/description/
package main

import (
	"fmt"
)

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println("getMin():", minStack.GetMin()) // возвращает -3
	minStack.Pop()
	fmt.Println("top():", minStack.Top())       // возвращает 0
	fmt.Println("getMin():", minStack.GetMin()) // возвращает -2
}

// Time complexity O(1), Space complexity O(n)
// MinStack implements a stack that supports push, pop, top, and retrieving the minimum element in constant time.
// Реализация стека MinStack поддерживает операции push, pop, top и получение минимального элемента за константное время.

// Time complexity O(1) for all operations, Space complexity O(n)
// Временная сложность O(1) для всех операций, пространственная сложность O(n)
type MinStack struct {
	stack    []int // Main stack to store all elements / Основной стек для хранения всех элементов
	minStack []int // Auxiliary stack to track minimum values / Вспомогательный стек для отслеживания минимальных значений
}

// Constructor initializes a new MinStack
// Constructor инициализирует новый MinStack
func Constructor() MinStack {
	ms := MinStack{
		stack:    make([]int, 0), // Initialize empty main stack / Инициализация пустого основного стека
		minStack: make([]int, 0), // Initialize empty minimum stack / Инициализация пустого стека минимумов
	}

	return ms
}

// Push adds a value to the stack and updates minimum stack if necessary
// Push добавляет значение в стек и при необходимости обновляет стек минимумов
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val) // Add value to main stack / Добавление значения в основной стек

	// Update minimum stack if current value is less than or equal to current minimum
	// Обновление стека минимумов, если текущее значение меньше или равно текущему минимуму
	if len(this.minStack) == 0 || this.minStack[len(this.minStack)-1] >= val {
		this.minStack = append(this.minStack, val)
	}
}

// Pop removes the top element from the stack and updates minimum stack if necessary
// Pop удаляет верхний элемент из стека и при необходимости обновляет стек минимумов
func (this *MinStack) Pop() {
	// If the top element is the current minimum, remove it from minimum stack too
	// Если верхний элемент является текущим минимумом, удаляем его также из стека минимумов
	if this.GetMin() == this.stack[len(this.stack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}

	// Remove top element from main stack
	// Удаление верхнего элемента из основного стека
	this.stack = this.stack[:len(this.stack)-1]
}

// Top returns the top element of the stack without removing it
// Top возвращает верхний элемент стека без его удаления
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// GetMin returns the minimum element in the stack in O(1) time
// GetMin возвращает минимальный элемент в стеке за время O(1)
func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

//
//// Более эффективное решение | Efficient solution
//// Time complexity O(1), Space complexity O(n)
//// MinStack is a stack that keeps track of the minimum value
//// MinStack - это стек, который отслеживает минимальное значение
//type MinStack struct {
//	min   int   // Current minimum value / Текущее минимальное значение
//	stack []int // Stack storage with encoded values / Хранилище стека с закодированными значениями
//}
//
//// Constructor initializes a new MinStack
//// Constructor инициализирует новый MinStack
//func Constructor() MinStack {
//	return MinStack{
//		min:   math.MaxInt64, // Initialize with maximum possible integer / Инициализация максимально возможным целым числом
//		stack: []int{},
//	}
//}
//
//// Push adds a value to the stack and updates minimum if necessary
//// Push добавляет значение в стек и обновляет минимум при необходимости
//func (this *MinStack) Push(val int) {
//	if len(this.stack) == 0 {
//		// For first element, store 0 and set minimum to the value
//		// Для первого элемента сохраняем 0 и устанавливаем минимум равным значению
//		this.stack = append(this.stack, 0)
//		this.min = val
//	} else {
//		// Store the difference between value and current minimum
//		// Сохраняем разницу между значением и текущим минимумом
//		this.stack = append(this.stack, val-this.min)
//		if val < this.min {
//			// Update minimum if new value is smaller
//			// Обновляем минимум, если новое значение меньше
//			this.min = val
//		}
//	}
//}
//
//// Pop removes the top element from the stack and updates minimum if necessary
//// Pop удаляет верхний элемент из стека и обновляет минимум при необходимости
//func (this *MinStack) Pop() {
//	if len(this.stack) == 0 {
//		return
//	}
//	pop := this.stack[len(this.stack)-1]
//	this.stack = this.stack[:len(this.stack)-1]
//	if pop < 0 {
//		// If popped value was negative, it means the popped element was a new minimum
//		// Recalculate the previous minimum
//		// Если извлеченное значение отрицательное, это означает, что извлеченный элемент был новым минимумом
//		// Пересчитываем предыдущий минимум
//		this.min = this.min - pop
//	}
//}
//
//// Top returns the top element of the stack
//// Top возвращает верхний элемент стека
//func (this *MinStack) Top() int {
//	top := this.stack[len(this.stack)-1]
//	if top > 0 {
//		// If stored value is positive, add it to the minimum to get the actual value
//		// Если сохраненное значение положительное, добавляем его к минимуму, чтобы получить фактическое значение
//		return top + this.min
//	}
//	// If stored value is 0 or negative, the actual value is the current minimum
//	// Если сохраненное значение равно 0 или отрицательное, фактическое значение - текущий минимум
//	return this.min
//}
//
//// GetMin returns the minimum element in the stack
//// GetMin возвращает минимальный элемент в стеке
//func (this *MinStack) GetMin() int {
//	return this.min
//}
