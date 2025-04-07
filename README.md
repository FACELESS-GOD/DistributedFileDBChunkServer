# DistributedFileDBChunkServer
This is a DistributedFileDatabase Chunk Server


This Project Implements a CHUNK Server stated in this [White Paper](https://research.google.com/archive/gfs-sosp2003.pdf)


- Here Every time a chunk is added a seperate GRPC call is send to master to update it's meta data.

