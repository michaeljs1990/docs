Spark
=====

## API

Check out the Spark API here for digging into much of what will be
mentioned in the section below.

https://spark.apache.org/docs/2.1.0/api/scala/index.html

## Config

You can configure spark to run locally when developing a script using the following
configuration.

```
import org.apache.spark.SparkConf
import org.apache.spark.SparkContext
import org.apache.spark.SparkContext._

val conf: SparkConf = new SparkConf().setMaster("local").setAppName("Wikipedia Ranking")
val sc: SparkContext = new SparkContext(conf)
```

After running the above you will also get an output that tells you the sparkUI has started
and you can visit is on port 4040 in your local browser for troubleshooting.

## RDDs 

Resilient Distributed Dataset are immutable collections of data that are designed to be operated
on in parallel. Below is an example of how to create a dataset from a file and then parse is into
an object using map.

```
import org.apache.spark.SparkConf
import org.apache.spark.SparkContext
import org.apache.spark.SparkContext._

import org.apache.spark.rdd.RDD

val wikiRdd: RDD[WikipediaArticle] = sc.textFile(WikipediaData.filePath).map(WikipediaData.parse _)
```

For testing purposes you can also create an RDD using the following.

```
val rdd: RDD[Tuple2[String, String]] = sc.parallelize(List(
  ("I", "India"),
  ("U", "USA"),
  ("W", "West")))
```

I have included the type of `RDD[Tuple2[String, String]]` to show you a few things. First that
the `parallelize` function takes a `List` but returns an RDD and second that scala has a `Tuple*`
type. This goes from `Tuple1` to `Tuple22`. You can read more about this here.

https://underscore.io/blog/posts/2016/10/11/twenty-two.html

Note that you can also use `sample` or `takeSample` to return a sub-section of the RDD when you
are trying to write some code that utilizes it.

## Pair RDDs

Pair RDDs are very similar to normal RDDs except they offer a few additional functions
that let you take advantage of the key more easily.

Here is an example of how to create one.

```
val rdd: RDD[Tuple2[String, String]] = sc.parallelize(List(
  ("I", "India"),
  ("U", "USA"),
  ("I", "Inga"),
  ("W", "West")))
```

You may notice the RDD in the `RDDs` section is infact a pair RDD as well. You can take
advantage of this with the following methods `reduceByKey`, `groupByKey`, and `join`.

`groupByKey`

Collect is only called below because without it nothing actually would happen.

```
val rdd: RDD[Tuple2[String, String]] = sc.parallelize(List(
  ("I", "India"),
  ("U", "USA"),
  ("I", "Inga"),
  ("W", "West")))

rdd.groupByKey().collect()
# res17: Array[(String, Iterable[String])] = Array((I,CompactBuffer(India, Inga)), (U,CompactBuffer(USA)), (W,CompactBuffer(West)))
```

`reduceByKey`

This is another contrived example but just serves to show you how to call it.

```
val rdd: RDD[Tuple2[String, Int]] = sc.parallelize(List(
  ("I", 10),
  ("U", 23),
  ("I", 42),
  ("W", 24)))

rdd.reduceByKey(_ + _).collect()
# res0: Array[(String, Int)] = Array((I,52), (U,23), (W,24))
```

`mapValue`

This allows you to apply a function or whatever to each value in a pair RDD.

```
val rdd: RDD[Tuple2[String, Int]] = sc.parallelize(List(
  ("I", 10),
  ("U", 23),
  ("I", 42),
  ("W", 24)))

rdd.mapValues(_ * 2).collect()
# res6: Array[(String, Int)] = Array((I,20), (U,46), (I,84), (W,48))
```

`countByKey`

Count number of elements that each key has.

```
val rdd: RDD[Tuple2[String, Int]] = sc.parallelize(List(
  ("I", 10),
  ("U", 23),
  ("I", 42),
  ("W", 24)))

rdd.countByKey()
# res13: scala.collection.Map[String,Long] = Map(I -> 2, U -> 1, W -> 1) 
```

## RDDs and Joins

Here is an example of the `join` function in spark which is actually an inner join.

```
val rdd1 = sc.parallelize(List((1, "Mike"), (2, "Abby"), (1, "Jeff"), (4, "Harry"), (5, "Potter")))
val rdd2 = sc.parallelize(List((1, ("A", "B")), (5, ("B", "B")), (6, ("C", "B")), (1, ("Z", "T"))))

rdd1.join(rdd2).collect()
# res19: Array[(Int, (String, (String, String)))] = Array((1,(Mike,(A,B))), (1,(Mike,(Z,T))), (1,(Jeff,(A,B))), (1,(Jeff,(Z,T))), (5,(Potter,(B,B))))
```

The important thing to note is that the keys with 4 and 6 in them are no longer in the list. This is because
it had nothing in common with the other list so it was dropped.

Outer joins allow you to pick what data is most important to you see the following example.

```
val rdd1 = sc.parallelize(List((1, "Mike"), (2, "Abby"), (1, "Jeff"), (4, "Harry"), (5, "Potter")))
val rdd2 = sc.parallelize(List((1, ("A", "B")), (5, ("B", "B")), (6, ("C", "B")), (1, ("Z", "T"))))

rdd1.leftOuterJoin(rdd2).collect()
# res1: Array[(Int, (String, Option[(String, String)]))] = Array((4,(Harry,None)), (1,(Mike,Some((A,B)))), (1,(Mike,Some((Z,T)))), (1,(Jeff,Some((A,B)))), (1,(Jeff,Some((Z,T)))), (5,(Potter,Some((B,B)))), (2,(Abby,None)))
```

With left outter join we have told it that we care most about the data on "left" or rdd1 in this case.
As you can see the key of 4 has been included in the computed list with an extra value of None added 
since it had nothing to join with.

```
val rdd1 = sc.parallelize(List((1, "Mike"), (2, "Abby"), (1, "Jeff"), (4, "Harry"), (5, "Potter")))
val rdd2 = sc.parallelize(List((1, ("A", "B")), (5, ("B", "B")), (6, ("C", "B")), (1, ("Z", "T"))))

rdd1.leftOuterJoin(rdd2).collect()
# res2: Array[(Int, (Option[String], (String, String)))] = Array((1,(Some(Mike),(A,B))), (1,(Some(Mike),(Z,T))), (1,(Some(Jeff),(A,B))), (1,(Some(Jeff),(Z,T))), (6,(None,(C,B))), (5,(Some(Potter),(B,B))))
```

This is doing the same as left join except we have decided that rdd2 is the data we care about so
the key of 6 has been included in the output and 4 has been dropped.

## Shuffling

Shuffling is important to consider when performing operations on RDDs. Take for example the case that you
have 10 nodes and an RDD in the form of `(Key, Value)` that has been partitioned across the cluster. If you
wanted to `groupByKey` each node in the cluster has to talk to each and every other node in the cluster to
find out what Values belong to what Key. This is a lot of network traffic if you have large sets of data. If
you instead reducedByKey it will reduce locally before talking over the network so instead of every `(Key, Value)`
pair needing to make a network call you are only doing it once per node in the form of `(Key, (Values))` after
your reduce has finished.

The following functions might cause shuffles.

```
cogroup
groupWith
join
leftOuterJoin
rightOuterJoin
groupByKey
reduceByKey
combineByKey
distinct
intersection
repartition
coalesce
```

## Partitioning

You can partition your data in a few different ways in Spark depending on what makes sense. The most
common are range partitioning and hash partitioning.

Here is an example of range partitioning.

```
import org.apache.spark.RangePartitioner

val rdd: RDD[Tuple2[String, String]] = sc.parallelize(List(
  ("I", "India"),
  ("U", "USA"),
  ("W", "West")))

val rp = new RangePartitioner(3, rdd)
val parts = rdd.partitionBy(rp).cache()

// This will let you peek inside each partition to see what is going on
// try changing the partition size and see what happens
parts.mapPartitionsWithIndex( (x,y) => { println(x); y.foreach(println); y } ).collect()
```

The following is how you would use a Hash partitioner that has 3 partitions.

```
import org.apache.spark.HashPartitioner

val rp = new HashPartitioner(3)
val parts = rdd.partitionBy(rp).cache()
```

When you have a partitioned set of data you can perform most operations except `map` or `flatMap`.
This is because you can alter the keys of the partition when running these which is extreamly
expensive.

## Caching

By default an RDD is recomputed everytime you run an action on it. If your RDD is not going to change
you likely will want to cache it in memory so multiple operations can be run without the performance
hit of loading it into memory.

```
val rdd: RDD[Tuple2[String, String]] = sc.parallelize(List(
  ("I", "India"),
  ("U", "USA"),
  ("W", "West"))).cache()
```

You can also use `persist` instead of `cache` if you so fancy.

A benifit of using cache can be seen in the following example of code.

```
val rdd: RDD[Tuple2[String, String]] = sc.parallelize(List(
  ("I", "India"),
  ("U", "USA"),
  ("W", "West")))

val xx = rdd.flatMap( x => x._1 + x._2).cache()

xx.collect()
xx.count()
```

If we had not called `cache` when making the rdd xx would have been computes twice.
Once for when we called collect and once when we called count. However since it was
cached the rdd is only loaded once.

Read more about it here https://stackoverflow.com/questions/28981359/why-do-we-need-to-call-cache-or-persist-on-a-rdd 

## Debug Workflow

After reading all of the above you are likely now concerned that you have written some super shitty 
scala code that is constantly shuffling data. The bad news is you 100% did but the good news is that
you have a few functions to debug it.

On any RDD you can call the function `dependencies` which will output a high level list of operations
that are about to be performed. You can all use `toDebugString` which will give you a little more
insight into what is happening.
