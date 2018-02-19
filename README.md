# golang_tour
In order to work through the exercises in the tour, I created this repo. 

## lessons learned
### concurrency model limitations
Go's go function and channels are kinda sad. The blocking channel default behavior opens up lots of race condition issues. I chased one for quite some time in the web crawler exercise. eventually, fixing the problem was easy, making the buffers big enough.

Also, the blend of (wierd) queues and traditional mutexes is an odd paradigm. Its trapped between C/Java and Erlang. 

tear down is interesting... See the web-crawler ex... but to detect everything was done, then tear everything down without leaking processes... interesting. Thanks to the bredthfirst tree traversal, detecting that its time to stop is "easy" but I think that there are cases where the approach used isn't safe. (eg a really long running single page with unique children.)

### pointers to structs and maps
there bifucation of types into "Base" and "referenced" (i think) takes some getting used to. sometimes a syntactic pass by value is a semantic pass by reference. at least so it seems. 

