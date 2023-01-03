package com.github.lzj09.datastructure.linked;

import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;

public class TestSinglyLinkedList {
	private static SinglyLinkedList<Integer> linked;
	
	@BeforeClass
    public static void beforeClass() {
		linked = new SinglyLinkedList<Integer>();
	}
	
	@Before
	public void add() {
		linked.clear();
		
		Assert.assertEquals(true, linked.isEmpty());
		
		Assert.assertEquals(0, linked.add(1));
		Assert.assertEquals(1, linked.add(2));
		Assert.assertEquals(2, linked.add(3));
		Assert.assertEquals(3, linked.add(4));
		Assert.assertEquals(4, linked.add(5));
		
		Assert.assertEquals(5, linked.size());
	}
	
	@Test
	public void indexOf() {
		Assert.assertEquals(1, linked.indexOf(2));
		Assert.assertEquals(3, linked.indexOf(4));
		Assert.assertEquals(4, linked.indexOf(5));
	}
	
	@Test
	public void get() {
		Assert.assertEquals(1, (int) linked.get(0));
		Assert.assertEquals(3, (int) linked.get(2));
	}
	
	@Test
	public void remove() {
		Assert.assertEquals(5, linked.size());
		Assert.assertEquals(4, (int) linked.remove(3));
		Assert.assertEquals(4, linked.size());
	}
	
	@Test
	public void update() {
		Assert.assertEquals(4, (int) linked.get(3));
		linked.update(3, 99);
		Assert.assertEquals(99, (int) linked.get(3));
	}
}