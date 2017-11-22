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

