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

