# Fan-Out Fan-In Pattern
When you need to split a task into multiple sub-tasks and then combine the results, the Fan-Out, Fan-In pattern is your go-to solution.
This pattern is particularly useful for parallel processing of data or when you need to distribute work across multiple goroutines and then aggregate the results.
The Fan-Out part involves distributing the work to multiple goroutines or functions that can run concurrently.
The Fan-In part involves collecting the results from all these concurrent operations and combining them into a single stream.


Sample fan out fan in pattern in Go
In this example, the fanOut function distributes the work to multiple worker goroutines.
Each worker processes items from the input channel concurrently.
The fanIn function then collects the results from all these workers and multiplexes them into a single channel.
This pattern shines in scenarios like parallel data processing or distributed computing tasks.
For instance, you could use it to process large datasets by splitting them into chunks, processing each chunk concurrently, and then aggregating the results.
It’s also useful in scenarios where you need to make multiple API calls in parallel and combine their responses.
The Fan-Out, Fan-In pattern allows you to leverage parallelism effectively, potentially leading to significant performance improvements in your Go applications.

However, it’s important to consider the overhead of creating and managing multiple goroutines and channels.

In some cases, if the processing of each item is very quick, the overhead might outweigh the benefits of parallelism.
