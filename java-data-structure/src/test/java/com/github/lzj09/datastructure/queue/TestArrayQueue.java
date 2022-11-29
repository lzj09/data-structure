package com.github.lzj09.datastructure.queue;

import org.junit.Assert;
import org.junit.BeforeClass;
import org.junit.Test;


public class TestArrayQueue {
	private static ArrayQueue<Integer> queue;

	@BeforeClass
    public static void beforeClass() {
		queue = new ArrayQueue<Integer>(4);
	}

	@Test
	public void add() {
		Assert.assertEquals(true, queue.isEmpty());

		queue.add(1);
		queue.add(2);
		queue.add(3);
		queue.add(4);
		
		Assert.assertEquals(true, queue.isFull());
	}
	
	@Test
	public void get() {
		Assert.assertEquals(1, (int) queue.get());
		Assert.assertEquals(2, (int) queue.get());
		Assert.assertEquals(3, (int) queue.get());
		Assert.assertEquals(4, (int) queue.get());
		
		Assert.assertEquals(true, queue.isEmpty());
	}
}