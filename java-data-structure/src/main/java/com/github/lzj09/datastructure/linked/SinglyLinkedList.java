package com.github.lzj09.datastructure.linked;

/**
 * 单向链表
 * 
 * @author lzj
 * @date 2022-12-11
 */
public class SinglyLinkedList<T> {
	// 链表的表头，表头不存放数据
	private Node<T> head = new Node<T>();
	// 用于记录节点个数
	private int num;

	/**
	 * 往链表中加入数据，直接将数据挂在链表尾部
	 * 
	 * @param data
	 */
	public int add(T data) {
		// 当next为null时，则定位为尾部
		Node<T> tmp = head;
		while (tmp.next != null) {
			tmp = tmp.next;
		}
		
		tmp.next = new Node<T>(data);
		return num++;
	}

	/**
	 * 获取数据的索引，未找到则返回-1
	 * 
	 * @param data
	 * @return
	 */
	public int indexOf(T data) {
		Node<T> tmp = head;
		for (int i = 0; tmp.next != null; i++) {
			tmp = tmp.next;
			if (tmp.data.equals(data)) {
				return i;
			}
		}
		
		return -1;
	}
	
	/**
	 * 获取索引获取数据
	 * 
	 * @param i
	 * @return
	 */
	public T get(int i) {
		int n = i + 1;
		if (n > num || i < 0) {
			throw new RuntimeException("超出索引范围");
		}
		
		Node<T> tmp = head;
		for (int j = 0; j < n; j++) {
			tmp = tmp.next;
		}
		
		return tmp.data;
	}
	
	/**
	 * 根据索引位置删除数据
	 * 
	 * @param i
	 * @return
	 */
	public T remove(int i) {
		int n = i + 1;
		if (n > num || i < 0) {
			throw new RuntimeException("超出索引范围");
		}
		
		// 先找到该索引前一位节点
		Node<T> tmp = head;
		for (int j = 0; j < i; j++) {
			tmp = tmp.next;
		}
		// 需要删除的节点
		Node<T> delNode = tmp.next;
		tmp.next = delNode.next;
		num--;
		
		return delNode.data;
	}
	
	/**
	 * 根据索引更新节点数据
	 * 
	 * @param i
	 * @param data
	 * @return
	 */
	public void update(int i, T data) {
		int n = i + 1;
		if (n > num || i < 0) {
			throw new RuntimeException("超出索引范围");
		}
		
		// 先找到该索引前一位节点
		Node<T> tmp = head;
		for (int j = 0; j < i; j++) {
			tmp = tmp.next;
		}
		// 需要更新的节点
		Node<T> updateNode = tmp.next;
		updateNode.data = data;
	}
	
	/**
	 * 清空链表
	 */
	public void clear() {
		head.next = null;
		num = 0;
	}
	
	/**
	 * 判断链表是否为空
	 * 
	 * @return
	 */
	public boolean isEmpty() {
		return head.next == null;
	}
	
	/**
	 * 链表节点数
	 * 
	 * @return
	 */
	public int size() {
		return num;
	}
	
	/**
	 * 链表中的节点
	 *
	 * @param <E>
	 */
	class Node<E> {
		E data;
		Node<E> next;
		
		public Node() {
		}
		
		public Node(E data) {
			this.data = data;
		}
	}
}