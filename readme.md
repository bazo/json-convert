first run ./build script - this will build the app in docker container named convert

to generate json files:

./run <fileName> -l <countOfLines>

to convert:

./run convert3 <fileName> - this one has lowest memory consumption

there is also convert and convert2 using different libraries but they use too much memory

to verify the parquet files written:

./run read <fileName> or ./run read2 <fileName> - each command uses different library