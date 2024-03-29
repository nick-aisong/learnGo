// 6-4 listing04.go
// 这个示例程序展示 goroutine 调度器是如何在单个线程上
// 切分时间片的
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg 用来等待程序完成
var wg sync.WaitGroup

// main 是所有 Go 程序的入口
func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)
	// 计数加 2，表示要等待两个 goroutine
	wg.Add(2)

	// 创建两个 goroutine
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// 等待 goroutine 结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime 显示 5000 以内的素数值
func printPrime(prefix string) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

// G:\GitHub\learnGo\Go语言实战\ch06_并发>go run listing04.go
// Create Goroutines
// Waiting To Finish
// B:2
// B:3
// B:5
// B:7
// B:11
// B:13
// B:17
// B:19
// B:23
// B:29
// B:31
// B:37
// B:41
// B:43
// B:47
// B:53
// B:59
// B:61
// B:67
// B:71
// B:73
// B:79
// B:83
// B:89
// B:97
// B:101
// B:103
// B:107
// B:109
// B:113
// B:127
// B:131
// B:137
// B:139
// B:149
// B:151
// B:157
// B:163
// B:167
// B:173
// B:179
// A:2
// A:3
// A:5
// A:7
// A:11
// A:13
// A:17
// A:19
// A:23
// A:29
// A:31
// A:37
// A:41
// A:43
// A:47
// A:53
// A:59
// A:61
// A:67
// A:71
// A:73
// A:79
// A:83
// A:89
// A:97
// A:101
// A:103
// A:107
// A:109
// A:113
// A:127
// A:131
// A:137
// A:139
// A:149
// A:151
// A:157
// A:163
// A:167
// A:173
// A:179
// A:181
// A:191
// A:193
// A:197
// A:199
// A:211
// A:223
// A:227
// A:229
// A:233
// A:239
// A:241
// A:251
// A:257
// A:263
// A:269
// A:271
// A:277
// A:281
// A:283
// A:293
// A:307
// A:311
// A:313
// A:317
// A:331
// A:337
// A:347
// A:349
// A:353
// A:359
// A:367
// A:373
// A:379
// A:383
// A:389
// A:397
// A:401
// A:409
// A:419
// A:421
// A:431
// A:433
// A:439
// A:443
// B:181
// B:191
// B:193
// B:197
// B:199
// B:211
// B:223
// B:227
// B:229
// B:233
// B:239
// B:241
// B:251
// B:257
// B:263
// B:269
// B:271
// B:277
// B:281
// B:283
// B:293
// B:307
// B:311
// B:313
// B:317
// B:331
// B:337
// B:347
// B:349
// B:353
// B:359
// B:367
// B:373
// B:379
// B:383
// B:389
// B:397
// B:401
// B:409
// B:419
// B:421
// B:431
// B:433
// B:439
// B:443
// B:449
// B:457
// A:449
// A:457
// A:461
// A:463
// A:467
// A:479
// A:487
// A:491
// A:499
// A:503
// A:509
// A:521
// A:523
// A:541
// A:547
// A:557
// A:563
// A:569
// A:571
// A:577
// A:587
// A:593
// A:599
// A:601
// A:607
// B:461
// B:463
// B:467
// B:479
// B:487
// B:491
// B:499
// B:503
// B:509
// B:521
// B:523
// B:541
// B:547
// B:557
// B:563
// B:569
// B:571
// B:577
// B:587
// B:593
// B:599
// B:601
// B:607
// B:613
// B:617
// B:619
// B:631
// B:641
// B:643
// B:647
// B:653
// A:613
// A:617
// A:619
// A:631
// A:641
// A:643
// A:647
// A:653
// A:659
// A:661
// A:673
// A:677
// A:683
// A:691
// B:659
// B:661
// B:673
// B:677
// B:683
// B:691
// B:701
// B:709
// B:719
// B:727
// B:733
// B:739
// B:743
// B:751
// B:757
// B:761
// B:769
// B:773
// B:787
// B:797
// B:809
// B:811
// B:821
// B:823
// B:827
// B:829
// B:839
// B:853
// B:857
// B:859
// B:863
// B:877
// A:701
// A:709
// A:719
// A:727
// A:733
// A:739
// A:743
// A:751
// A:757
// A:761
// A:769
// A:773
// A:787
// A:797
// A:809
// A:811
// A:821
// A:823
// A:827
// A:829
// A:839
// A:853
// A:857
// A:859
// A:863
// A:877
// A:881
// A:883
// B:881
// B:883
// B:887
// B:907
// B:911
// B:919
// B:929
// B:937
// B:941
// B:947
// B:953
// B:967
// B:971
// B:977
// B:983
// B:991
// A:887
// A:907
// A:911
// A:919
// A:929
// A:937
// A:941
// A:947
// A:953
// A:967
// A:971
// A:977
// A:983
// A:991
// A:997
// B:997
// B:1009
// B:1013
// B:1019
// B:1021
// B:1031
// B:1033
// B:1039
// B:1049
// B:1051
// B:1061
// A:1009
// A:1013
// A:1019
// A:1021
// A:1031
// A:1033
// A:1039
// A:1049
// A:1051
// A:1061
// A:1063
// A:1069
// A:1087
// A:1091
// A:1093
// A:1097
// A:1103
// A:1109
// A:1117
// A:1123
// A:1129
// A:1151
// A:1153
// A:1163
// A:1171
// A:1181
// A:1187
// A:1193
// A:1201
// A:1213
// A:1217
// A:1223
// A:1229
// A:1231
// A:1237
// A:1249
// A:1259
// A:1277
// B:1063
// B:1069
// B:1087
// B:1091
// B:1093
// B:1097
// B:1103
// B:1109
// B:1117
// B:1123
// B:1129
// B:1151
// B:1153
// B:1163
// B:1171
// B:1181
// B:1187
// B:1193
// B:1201
// B:1213
// B:1217
// B:1223
// B:1229
// B:1231
// B:1237
// B:1249
// B:1259
// B:1277
// B:1279
// B:1283
// B:1289
// A:1279
// A:1283
// A:1289
// A:1291
// A:1297
// A:1301
// A:1303
// A:1307
// A:1319
// A:1321
// A:1327
// A:1361
// A:1367
// A:1373
// A:1381
// A:1399
// A:1409
// A:1423
// A:1427
// A:1429
// A:1433
// A:1439
// B:1291
// B:1297
// B:1301
// B:1303
// B:1307
// B:1319
// B:1321
// B:1327
// B:1361
// B:1367
// B:1373
// B:1381
// B:1399
// B:1409
// B:1423
// B:1427
// B:1429
// B:1433
// B:1439
// B:1447
// B:1451
// B:1453
// B:1459
// B:1471
// B:1481
// B:1483
// B:1487
// A:1447
// A:1451
// A:1453
// A:1459
// A:1471
// A:1481
// A:1483
// A:1487
// A:1489
// A:1493
// A:1499
// A:1511
// A:1523
// A:1531
// A:1543
// A:1549
// A:1553
// A:1559
// A:1567
// A:1571
// A:1579
// A:1583
// A:1597
// A:1601
// A:1607
// A:1609
// A:1613
// A:1619
// B:1489
// B:1493
// B:1499
// B:1511
// B:1523
// B:1531
// B:1543
// B:1549
// B:1553
// B:1559
// B:1567
// B:1571
// B:1579
// B:1583
// B:1597
// B:1601
// B:1607
// B:1609
// B:1613
// B:1619
// B:1621
// B:1627
// B:1637
// A:1621
// A:1627
// A:1637
// A:1657
// A:1663
// A:1667
// A:1669
// A:1693
// A:1697
// A:1699
// A:1709
// A:1721
// A:1723
// A:1733
// A:1741
// A:1747
// A:1753
// A:1759
// A:1777
// B:1657
// B:1663
// B:1667
// B:1669
// B:1693
// B:1697
// B:1699
// B:1709
// B:1721
// B:1723
// B:1733
// B:1741
// A:1783
// A:1787
// A:1789
// B:1747
// B:1753
// B:1759
// B:1777
// B:1783
// B:1787
// B:1789
// B:1801
// B:1811
// B:1823
// B:1831
// B:1847
// B:1861
// B:1867
// B:1871
// B:1873
// B:1877
// B:1879
// B:1889
// B:1901
// B:1907
// B:1913
// B:1931
// B:1933
// B:1949
// B:1951
// B:1973
// B:1979
// A:1801
// A:1811
// A:1823
// A:1831
// A:1847
// A:1861
// A:1867
// A:1871
// A:1873
// A:1877
// A:1879
// A:1889
// A:1901
// A:1907
// A:1913
// A:1931
// A:1933
// A:1949
// A:1951
// A:1973
// A:1979
// A:1987
// A:1993
// A:1997
// A:1999
// A:2003
// A:2011
// A:2017
// A:2027
// A:2029
// A:2039
// A:2053
// A:2063
// A:2069
// A:2081
// B:1987
// B:1993
// B:1997
// B:1999
// B:2003
// B:2011
// B:2017
// B:2027
// B:2029
// B:2039
// B:2053
// B:2063
// B:2069
// B:2081
// B:2083
// B:2087
// B:2089
// B:2099
// B:2111
// B:2113
// B:2129
// B:2131
// B:2137
// B:2141
// B:2143
// B:2153
// B:2161
// B:2179
// B:2203
// B:2207
// B:2213
// B:2221
// B:2237
// B:2239
// B:2243
// B:2251
// B:2267
// B:2269
// B:2273
// B:2281
// B:2287
// B:2293
// B:2297
// B:2309
// B:2311
// B:2333
// B:2339
// B:2341
// B:2347
// B:2351
// B:2357
// B:2371
// B:2377
// B:2381
// B:2383
// B:2389
// B:2393
// B:2399
// B:2411
// B:2417
// B:2423
// B:2437
// B:2441
// B:2447
// B:2459
// B:2467
// B:2473
// A:2083
// A:2087
// A:2089
// A:2099
// A:2111
// A:2113
// A:2129
// A:2131
// A:2137
// A:2141
// A:2143
// A:2153
// A:2161
// A:2179
// A:2203
// B:2477
// B:2503
// B:2521
// B:2531
// B:2539
// B:2543
// B:2549
// B:2551
// B:2557
// B:2579
// B:2591
// B:2593
// B:2609
// B:2617
// B:2621
// B:2633
// B:2647
// B:2657
// B:2659
// B:2663
// B:2671
// A:2207
// A:2213
// A:2221
// A:2237
// A:2239
// A:2243
// A:2251
// A:2267
// A:2269
// A:2273
// A:2281
// A:2287
// A:2293
// A:2297
// A:2309
// A:2311
// A:2333
// A:2339
// A:2341
// A:2347
// A:2351
// A:2357
// B:2677
// B:2683
// B:2687
// B:2689
// B:2693
// B:2699
// B:2707
// B:2711
// B:2713
// B:2719
// B:2729
// B:2731
// B:2741
// B:2749
// B:2753
// B:2767
// B:2777
// B:2789
// B:2791
// A:2371
// A:2377
// A:2381
// A:2383
// A:2389
// A:2393
// A:2399
// A:2411
// A:2417
// A:2423
// A:2437
// A:2441
// A:2447
// A:2459
// B:2797
// A:2467
// A:2473
// A:2477
// A:2503
// A:2521
// A:2531
// A:2539
// A:2543
// A:2549
// A:2551
// A:2557
// A:2579
// A:2591
// A:2593
// A:2609
// A:2617
// A:2621
// A:2633
// A:2647
// A:2657
// A:2659
// A:2663
// A:2671
// A:2677
// A:2683
// A:2687
// A:2689
// B:2801
// B:2803
// B:2819
// B:2833
// B:2837
// B:2843
// B:2851
// B:2857
// B:2861
// B:2879
// B:2887
// B:2897
// B:2903
// B:2909
// A:2693
// A:2699
// A:2707
// A:2711
// A:2713
// A:2719
// A:2729
// A:2731
// A:2741
// A:2749
// A:2753
// B:2917
// B:2927
// B:2939
// B:2953
// B:2957
// B:2963
// B:2969
// B:2971
// B:2999
// B:3001
// B:3011
// B:3019
// B:3023
// B:3037
// A:2767
// A:2777
// A:2789
// A:2791
// A:2797
// A:2801
// A:2803
// A:2819
// A:2833
// A:2837
// A:2843
// A:2851
// A:2857
// A:2861
// A:2879
// A:2887
// A:2897
// A:2903
// B:3041
// B:3049
// B:3061
// B:3067
// B:3079
// B:3083
// B:3089
// B:3109
// B:3119
// B:3121
// B:3137
// B:3163
// B:3167
// B:3169
// B:3181
// B:3187
// B:3191
// B:3203
// A:2909
// A:2917
// A:2927
// A:2939
// A:2953
// A:2957
// A:2963
// A:2969
// A:2971
// A:2999
// A:3001
// A:3011
// A:3019
// A:3023
// A:3037
// A:3041
// A:3049
// A:3061
// A:3067
// A:3079
// A:3083
// A:3089
// A:3109
// B:3209
// B:3217
// B:3221
// B:3229
// B:3251
// B:3253
// B:3257
// B:3259
// B:3271
// B:3299
// B:3301
// B:3307
// B:3313
// B:3319
// B:3323
// B:3329
// B:3331
// B:3343
// B:3347
// B:3359
// B:3361
// B:3371
// B:3373
// B:3389
// B:3391
// A:3119
// A:3121
// A:3137
// A:3163
// A:3167
// A:3169
// A:3181
// A:3187
// A:3191
// A:3203
// A:3209
// A:3217
// A:3221
// A:3229
// A:3251
// A:3253
// B:3407
// B:3413
// B:3433
// B:3449
// B:3457
// B:3461
// B:3463
// B:3467
// B:3469
// B:3491
// B:3499
// B:3511
// A:3257
// A:3259
// A:3271
// A:3299
// A:3301
// A:3307
// A:3313
// A:3319
// A:3323
// A:3329
// A:3331
// A:3343
// A:3347
// A:3359
// A:3361
// A:3371
// A:3373
// A:3389
// B:3517
// B:3527
// B:3529
// B:3533
// B:3539
// B:3541
// B:3547
// B:3557
// B:3559
// B:3571
// B:3581
// B:3583
// B:3593
// B:3607
// B:3613
// B:3617
// B:3623
// B:3631
// A:3391
// A:3407
// A:3413
// A:3433
// A:3449
// A:3457
// A:3461
// A:3463
// A:3467
// A:3469
// A:3491
// A:3499
// A:3511
// A:3517
// A:3527
// A:3529
// A:3533
// A:3539
// A:3541
// A:3547
// B:3637
// B:3643
// A:3557
// A:3559
// A:3571
// A:3581
// A:3583
// A:3593
// A:3607
// A:3613
// A:3617
// A:3623
// A:3631
// A:3637
// A:3643
// A:3659
// A:3671
// A:3673
// A:3677
// A:3691
// B:3659
// B:3671
// B:3673
// B:3677
// B:3691
// B:3697
// B:3701
// B:3709
// B:3719
// B:3727
// B:3733
// B:3739
// B:3761
// B:3767
// B:3769
// B:3779
// B:3793
// A:3697
// B:3797
// B:3803
// B:3821
// B:3823
// B:3833
// B:3847
// B:3851
// B:3853
// B:3863
// B:3877
// B:3881
// B:3889
// B:3907
// B:3911
// B:3917
// B:3919
// B:3923
// A:3701
// A:3709
// A:3719
// A:3727
// A:3733
// A:3739
// A:3761
// A:3767
// A:3769
// A:3779
// A:3793
// A:3797
// A:3803
// B:3929
// B:3931
// B:3943
// B:3947
// B:3967
// B:3989
// B:4001
// B:4003
// B:4007
// B:4013
// B:4019
// B:4021
// B:4027
// B:4049
// B:4051
// B:4057
// B:4073
// B:4079
// B:4091
// B:4093
// B:4099
// B:4111
// B:4127
// A:3821
// A:3823
// A:3833
// B:4129
// B:4133
// B:4139
// B:4153
// B:4157
// B:4159
// B:4177
// B:4201
// B:4211
// B:4217
// B:4219
// B:4229
// B:4231
// B:4241
// B:4243
// B:4253
// A:3847
// B:4259
// B:4261
// B:4271
// B:4273
// B:4283
// B:4289
// B:4297
// B:4327
// B:4337
// B:4339
// B:4349
// B:4357
// B:4363
// B:4373
// B:4391
// B:4397
// B:4409
// B:4421
// B:4423
// A:3851
// A:3853
// A:3863
// A:3877
// A:3881
// A:3889
// A:3907
// A:3911
// A:3917
// A:3919
// A:3923
// A:3929
// A:3931
// A:3943
// A:3947
// A:3967
// A:3989
// A:4001
// A:4003
// A:4007
// A:4013
// A:4019
// A:4021
// A:4027
// A:4049
// A:4051
// A:4057
// A:4073
// A:4079
// A:4091
// A:4093
// A:4099
// A:4111
// B:4441
// B:4447
// B:4451
// B:4457
// B:4463
// B:4481
// B:4483
// B:4493
// B:4507
// B:4513
// B:4517
// B:4519
// B:4523
// B:4547
// B:4549
// B:4561
// B:4567
// B:4583
// B:4591
// B:4597
// B:4603
// B:4621
// B:4637
// B:4639
// B:4643
// B:4649
// A:4127
// A:4129
// A:4133
// A:4139
// A:4153
// A:4157
// A:4159
// A:4177
// A:4201
// A:4211
// B:4651
// B:4657
// B:4663
// B:4673
// B:4679
// B:4691
// B:4703
// B:4721
// B:4723
// B:4729
// B:4733
// B:4751
// B:4759
// B:4783
// B:4787
// B:4789
// B:4793
// B:4799
// B:4801
// B:4813
// A:4217
// A:4219
// A:4229
// A:4231
// A:4241
// A:4243
// A:4253
// A:4259
// A:4261
// A:4271
// A:4273
// A:4283
// A:4289
// A:4297
// A:4327
// A:4337
// A:4339
// A:4349
// A:4357
// A:4363
// A:4373
// A:4391
// B:4817
// B:4831
// B:4861
// B:4871
// B:4877
// B:4889
// B:4903
// B:4909
// B:4919
// B:4931
// B:4933
// B:4937
// B:4943
// B:4951
// B:4957
// B:4967
// B:4969
// B:4973
// B:4987
// B:4993
// B:4999
// Completed B
// A:4397
// A:4409
// A:4421
// A:4423
// A:4441
// A:4447
// A:4451
// A:4457
// A:4463
// A:4481
// A:4483
// A:4493
// A:4507
// A:4513
// A:4517
// A:4519
// A:4523
// A:4547
// A:4549
// A:4561
// A:4567
// A:4583
// A:4591
// A:4597
// A:4603
// A:4621
// A:4637
// A:4639
// A:4643
// A:4649
// A:4651
// A:4657
// A:4663
// A:4673
// A:4679
// A:4691
// A:4703
// A:4721
// A:4723
// A:4729
// A:4733
// A:4751
// A:4759
// A:4783
// A:4787
// A:4789
// A:4793
// A:4799
// A:4801
// A:4813
// A:4817
// A:4831
// A:4861
// A:4871
// A:4877
// A:4889
// A:4903
// A:4909
// A:4919
// A:4931
// A:4933
// A:4937
// A:4943
// A:4951
// A:4957
// A:4967
// A:4969
// A:4973
// A:4987
// A:4993
// A:4999
// Completed A
// Terminating Program
