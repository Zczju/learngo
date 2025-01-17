package main

import (
	"fmt"
	"time"
)

var sampleData = []int64{
	1789, 337, 2299, 2069, 175, 1539, 2513, 994, 2203, 1353, 2852, 544, 1511, 2723, 500, 795, 664, 136, 156, 1744, 1300, 576, 365, 1836, 617, 1522, 2952, 1632, 164, 1677, 1360, 2014, 1655, 406, 1043, 2982, 2643, 2324, 751, 1543, 1140, 2276, 2943, 1927, 1461, 137, 475, 2238, 798, 1382, 222, 1505, 2915, 617, 2095, 493, 1352, 1407, 1306, 2238, 518, 384, 2510, 987, 2152, 896, 2613, 1441, 2610, 2575, 2603, 764, 1952, 715, 1895, 2383, 2243, 215, 2351, 579, 1266, 1351, 1693, 1687, 1343, 1372, 1806, 689, 1715, 1354, 2652, 377, 2233, 1068, 732, 1417, 819, 2955, 2720, 2195, 96, 1868, 633, 2250, 1804, 282, 199, 859, 2514, 265, 1619, 2967, 2974, 1270, 1041, 2313, 2238, 120, 1341, 277, 1031, 2620, 2224, 277, 77, 189, 1849, 336, 752, 291, 2258, 644, 980, 2298, 2331, 2177, 2252, 737, 1991, 1112, 858, 2439, 2610, 1624, 493, 1434, 2589, 2097, 2536, 2522, 2943, 1925, 1857, 2246, 1428, 2992, 2163, 896, 2533, 188, 2147, 2813, 1576, 2927, 278, 1106, 648, 1491, 1529, 2247, 2849, 2557, 1278, 2370, 1795, 512, 611, 2210, 1849, 682, 263, 580, 1981, 1546, 1060, 1682, 1511, 881, 2112, 1808, 952, 184, 1890, 922, 2479, 2433, 2477, 2903, 1030, 1687, 114, 862, 1640, 2539, 2712, 1259, 174, 1469, 2262, 1977, 1588, 1665, 2514, 1596, 1886, 2771, 2909, 2586, 1795, 192, 2934, 2614, 2277, 668, 2828, 1305, 1127, 875, 1591, 964, 37, 2898, 40, 126, 2066, 13, 1865, 2, 2686, 437, 2175, 58, 1717, 2338, 918, 2594, 1860, 1256, 1708, 282, 2749, 1624, 2956, 497, 1433, 2861, 911, 1025, 2466, 415, 1251, 2432, 2021, 1446, 2461, 2798, 1118, 1205, 1588, 213, 2955, 2782, 2985, 1334, 1327, 455, 63, 1387, 2326, 654, 1663, 2727, 2877, 1002, 190, 1536, 1588, 1412, 2652, 1608, 2483, 286, 2654, 2455, 2141, 1471, 998, 959, 932, 896, 1467, 2298, 608, 2735, 1375, 2411, 1826, 787, 837, 396, 1205, 997, 480, 90, 2887, 2376, 515, 2035, 1625, 1389, 691, 2406, 384, 162, 2744, 1921, 893, 2727, 2758, 300, 1144, 2764, 2361, 2495, 1180, 1296, 820, 596, 1354, 543, 100, 1910, 1562, 730, 2751, 1952, 828, 2405, 2205, 1247, 2998, 1750, 2202, 1546, 462, 2738, 1688, 1949, 1707, 2725, 1836, 1214, 2953, 2076, 2861, 1802, 811, 1222, 895, 3, 456, 2072, 2493, 34, 2783, 715, 1888, 743, 2792, 2420, 1128, 2025, 1740, 2185, 1900, 2772, 1844, 1333, 1086, 2298, 1765, 2560, 1936, 2588, 1844, 2743, 2689, 690, 345, 233, 2820, 2346, 107, 59, 248, 871, 189, 2994, 1546, 540, 745, 771, 309, 651, 2657, 2394, 793, 2597, 2511, 720, 1755, 2809, 2744, 2417, 31, 365, 1763, 1582, 2012, 2967, 221, 2015, 339, 1919, 2210, 567, 1080, 81, 841, 1969, 2316, 425, 2908, 2125, 141, 1079, 2816, 1717, 1893, 2214, 2492, 1794, 2032, 101, 1392, 2959, 2546, 1546, 507, 1070, 2251, 1750, 1129, 1304, 404, 667, 2224, 2446, 2212, 274, 321, 1468, 2220, 2068, 82, 2656, 1136, 127, 184, 1060, 1864, 1121, 1364, 897, 551, 979, 482, 1718, 892, 33, 1061, 2448, 1363, 2756, 2306, 2884, 22, 1664, 1248, 1952, 2228, 1371, 34, 1967, 944, 546, 1606, 1050, 1928, 355, 1896, 1724, 803, 91, 2331, 2859, 730, 2396, 114, 1818, 2226, 1418, 589, 1945, 389, 11, 2737, 1190, 1105, 2070, 1074, 1792, 35, 46, 67, 1702, 808, 2559, 11, 2307, 614, 2155, 1729, 1067, 536, 1322, 892, 1445, 593, 1567, 1103, 905, 1820, 1725, 2368, 445, 908, 2337, 2729, 2790, 1206, 887, 1005, 1967, 653, 2345, 2786, 546, 815, 2409, 1438, 902, 114, 850, 543, 1007, 2905, 473, 957, 1391, 2004, 450, 1954, 725, 775, 2793, 2366, 110, 849, 2243, 253, 546, 2715, 2734, 1104, 2512, 987, 1582, 268, 303, 979, 1114, 2936, 575, 474, 2280, 822, 1537, 310, 785, 1690, 1688, 456, 665, 2203, 431, 512, 1936, 1422, 1528, 756, 1537, 1793, 477, 1030, 2671, 1816, 2628, 2840, 2963, 540, 1452, 581, 1323, 2590, 542, 534, 2124, 2054, 77, 1072, 2888, 1792, 745, 120, 197, 993, 2670, 1069, 2181, 1206, 2010, 1298, 510, 518, 2700, 1725, 1477, 2267, 2749, 1646, 680, 2073, 2433, 2006, 2570, 2355, 1705, 697, 287, 2997, 2934, 985, 495, 2545, 966, 521, 2933, 1222, 1138, 1382, 2378, 138, 2388, 1764, 774, 1345, 2052, 2976, 2849, 1969, 1497, 2646, 931, 2967, 2206, 152, 2771, 2776, 2112, 2816, 271, 2894, 994, 625, 622, 546, 1387, 2688, 2489, 1172, 1345, 1299, 1787, 120, 1639, 1533, 2340, 1031, 1257, 2970, 1932, 2347, 2708, 2149, 2251, 308, 836, 1008, 1321, 1058, 2953, 2766, 1378, 270, 1806, 1557, 2601, 2502, 2212, 27, 2501, 987, 1910, 1708, 1265, 1221, 2461, 1678, 15, 1603, 40, 856, 2743, 665, 1083, 2036, 2722, 2887, 938, 1235, 701, 745, 1907, 675, 442, 495, 432, 1181, 151, 1676, 2871, 2862, 1472, 144, 761, 2511, 1581, 1920, 2377, 1365, 451, 2855, 1980, 1902, 2812, 2900, 2944, 2141, 2574, 2417, 1912, 682, 1888, 888, 2975, 1607, 2019, 971, 2904, 2022, 319, 2990, 1557, 1423, 558, 2973, 2456, 2565, 924, 2615, 2850, 1111, 1730, 1021, 216, 619, 1042, 2223, 963, 1033, 1445, 1341, 492, 2490, 4, 665, 21, 1380, 1362, 1571, 2877, 2923, 844, 562, 335, 161, 1064, 1479, 614, 1224, 1207, 144, 946, 2962, 916, 2773, 2662, 1330, 2427, 1437, 944, 2055, 780, 263, 1690, 606, 1186, 570, 1368, 2935, 2432, 1684, 2937, 1829, 2210, 2094, 1809, 1606, 70, 1417, 891, 1391, 2489, 1174, 1755, 2316, 1682, 964, 498, 1193, 2027, 1375, 2019, 2164, 1926, 1515, 1529, 2584, 674, 2259, 1829, 2648, 2452, 202, 2761, 2563, 2161, 736, 1136, 552, 2088, 2428, 2306, 1888, 1222, 2987, 2910, 766, 2908, 1054, 532, 889, 1756, 2267, 1346, 2039, 2319, 473, 2912, 110, 2679, 2014, 1607, 1929, 1006, 1247, 2258, 856, 1429, 92, 451, 1220, 1936, 1908, 650, 2855, 1702, 2970, 2508, 2793, 1692, 1099, 2047, 858, 1343, 646, 0, 2812, 1774, 2021, 2637, 2422, 2484, 1620, 240, 517, 2269, 679, 1273, 861, 2253, 901, 1283, 2943, 776, 1872, 2674, 1302, 638, 485, 2320, 806, 2321, 1026, 2884, 2641, 751, 1977, 596, 1637, 2254, 2770, 2593, 2414, 2982, 1386, 713, 1757, 2194, 1671, 2231, 170, 309, 696, 1676, 684, 1645, 919, 538,
}

type Node struct {
	data  int
	root  *Node
	left  *Node
	right *Node
}

func buildTree() *Node {
	n1 := &Node{data: 51}
	n2 := &Node{data: 35}
	n3 := &Node{data: 65}

	n1.left = n2
	n2.root = n1
	n1.right = n3
	n3.root = n1

	return n1
}

func insertNode(root *Node, newNode *Node) *Node {
	if root == nil {
		return newNode
	}
	if newNode.data == root.data {
		return root
	}
	if newNode.data < root.data {
		if root.left == nil {
			//找到位置，插入数据
			root.left = newNode
			newNode.root = root
		} else {
			insertNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			//找到位置，插入数据
			root.right = newNode
			newNode.root = root
		} else {
			insertNode(root.right, newNode)
		}
	}
	return root
}

func deleteNode(root *Node, v int) *Node {
	if v < root.data {
		deleteNode(root.left, v)
	} else if v > root.data {
		deleteNode(root.right, v)

	} else {
		//现在root指向要删除的节点
		leftNextGen := findNextGenFromLeft(root.left)
		rightNextGen := findNextGenFromRight(root.right)
		if leftNextGen == nil && rightNextGen == nil {
			// 现在要删除的是叶子节点，即最底部的节点
			top := root.root
			down := root
			if top.left == down {
				// 表示要删除的在左侧
				top.left = nil
				down.root = nil
				return nil //为什么是return nil???
			} else {
				//表示要删除的在右侧
				top.right = nil
				down.root = nil
				return nil
			}

		} else if leftNextGen != nil {
			root.data = leftNextGen.data
			deleteNode(leftNextGen, leftNextGen.data)
		} else if rightNextGen != nil {
			root.data = rightNextGen.data
			deleteNode(rightNextGen, rightNextGen.data)
		}
	}
	return root
}

func findNextGenFromLeft(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmpNode := root

	for {
		if tmpNode.right != nil {
			tmpNode = tmpNode.right
		} else {
			break
		}
	}
	return tmpNode
}

func findNextGenFromRight(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmpNode := root
	for {
		if tmpNode.left != nil {
			tmpNode = tmpNode.left
		} else {
			break
		}
	}
	return tmpNode
}

var totalCompare int = 0

func main() {
	//mainForPreSearch()
	//mainForMidSearch()
	mainForPostSearch()
}

/*
func mainForPreSearch() {
	var root *Node
	startTime := time.Now()
	for _, v := range sampleData {
		root = insertNode(root, &Node{data: int(v)})
	}
	buildFinishTime := time.Now()
	fmt.Println("构建结束", buildFinishTime.Sub(startTime))

	for i := 0; i < 1000000; i++ {
		preSearch(root, 501)
		preSearch(root, 888)
		preSearch(root, 900)
		preSearch(root, 3)
	}
	finishTime := time.Now()
	fmt.Println("总比较次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))
}

*/

/*
func mainForMidSearch() {
	var root *Node
	startTime := time.Now()
	for _, v := range sampleData {
		root = insertNode(root, &Node{data: int(v)})
	}
	buildFinishTime := time.Now()
	fmt.Println("构建结束：", buildFinishTime.Sub(startTime))

	for i := 0; i < 1000000; i++ {
		midSearch(root, 501)
		midSearch(root, 888)
		midSearch(root, 900)
		midSearch(root, 3)
	}
	finishTime := time.Now()
	fmt.Println("总比较次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))
}
*/

func mainForPostSearch() {
	var root *Node

	startTime := time.Now()
	for _, v := range sampleData {
		root = insertNode(root, &Node{data: int(v)})
	}
	buildFinishTime := time.Now()
	fmt.Println("构建结束：", buildFinishTime.Sub(startTime))

	for i := 0; i < 1000000; i++ {
		postSearch(root, 501)
		postSearch(root, 888)
		postSearch(root, 900)
		postSearch(root, 3)
	}
	finishTime := time.Now()

	fmt.Println("总比较次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))
}
func postSearch(root *Node, targetNum int) bool {
	if root == nil {
		return false
	}
	totalCompare++
	if root.data > targetNum {
		if postSearch(root.left, targetNum) {
			return true
		}
	}
	if root.data < targetNum {
		if postSearch(root.right, targetNum) {
			return true
		}
	}
	if root.data == targetNum {
		return true
	}
	return false
}
