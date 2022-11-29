package com.github.lzj09.datastructure.queue;

/**
 * 基于数组实现队列
 * 
 * @author lzj
 * @date 2022-11-29
 */
public class ArrayQueue<T> {
	// 指向队头的下标
	private int front;
	
	// 指向队尾的下标下一个下标
	private int rear;
	
	// 队列的最大容量
	private int size;
	
	// 存放数据的数组容器
	private T[] arrs;

	@SuppressWarnings("unchecked")
	public ArrayQueue(int size) {
		// 由于rear指向队尾的下一个下标，为此需要多1个空间
		this.size = size + 1;
		arrs = (T[]) new Object[this.size];
	}

	/**
	 * 队列是否满
	 * 
	 * @return
	 */
	public boolean isFull() {
		return (rear - front + 1) % size == 0;
	}
	
	/**
	 * 队列是否为空
	 * 
	 * @return
	 */
	public boolean isEmpty() {
		return rear == front;
	}

	/**
	 * 往队列中加入数据
	 * 
	 * @param data
	 */
	public void add(T data) {
		if (isFull()) {
			throw new RuntimeException("队列已满");
		}
		
		arrs[rear] = data;
		rear = (rear + 1) % size;
	}

	/**
	 * 从队列中获取数据
	 * 
	 * @return
	 */
	public T get() {
		if (isEmpty()) {
			throw new RuntimeException("队列为空");
		}
		
		int tmp = front;
		front = (front + 1) % size;
		
		return arrs[tmp];
	}
}