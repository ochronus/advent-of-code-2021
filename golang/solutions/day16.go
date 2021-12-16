package solutions

import (
	"math"
	"ochronus/aoc2021/utils"
)

func parseDay16Input() (binarystr string) {
	binarystr = utils.Hex2Bin(utils.ReadFileToString("../inputs/16.txt"))
	for {
		if len(binarystr)%4 == 0 {
			break
		}
		binarystr = "0" + binarystr
	}
	return
}

func Day16P01() int64 {
	bin := parseDay16Input()
	_, sum, _ := decode(bin, 0)

	return sum
}

func Day16P02() int64 {
	bin := parseDay16Input()
	_, _, val := decode(bin, 0)

	return val
}

func decode(binaryStr string, ic int) (int, int64, int64) {
	packetVersion := utils.Bin2Dec(binaryStr[ic+0 : ic+0+3])
	packetType := utils.Bin2Dec(binaryStr[ic+3 : ic+3+3])

	if packetType == 4 { // it's a literal
		ic += 6
		binNum := ""
		for {
			isLastChunk := binaryStr[ic] == '0'
			binNum += binaryStr[ic+1 : ic+5]
			ic += 5
			if isLastChunk {
				break
			}
		}
		return ic, packetVersion, utils.Bin2Dec(binNum)

	} else { // decode all the subpackets
		lengthType := string(binaryStr[ic+6])
		ic += 7
		var subVals []int64
		if lengthType == "0" {
			subPacketLength := int(utils.Bin2Dec(binaryStr[ic : ic+15]))
			ic += 15                       // start of the first subpacket
			lastIc := ic + subPacketLength // end of subpackets in the bit stream
			for ok := true; ok; ok = ic < lastIc {
				nextIc, subVersion, subVal := decode(binaryStr, ic)
				packetVersion += subVersion
				ic = nextIc
				subVals = append(subVals, subVal)
			}
		} else {
			subPacketCount := utils.Bin2Dec(binaryStr[ic : ic+11])
			ic += 11 // start of the first subpacket
			for i := int64(0); i < subPacketCount; i++ {
				nextIc, subVersion, subVal := decode(binaryStr, ic)
				packetVersion += subVersion
				ic = nextIc
				subVals = append(subVals, subVal)
			}
		}

		var result int64

		switch packetType {
		case 0:
			for _, v := range subVals {
				result += v
			}
			return ic, packetVersion, result

		case 1:
			result = 1
			for _, v := range subVals {
				result *= v
			}
			return ic, packetVersion, result

		case 2:
			result = math.MaxInt64
			for _, v := range subVals {
				result = utils.Min64(result, v)
			}
			return ic, packetVersion, result

		case 3:
			for _, v := range subVals {
				result = utils.Max64(result, v)
			}
			return ic, packetVersion, result

		case 5:
			if subVals[0] > subVals[1] {
				result = 1
			}
			return ic, packetVersion, result

		case 6:
			if subVals[0] < subVals[1] {
				result = 1
			}
			return ic, packetVersion, result
		case 7:
			if subVals[0] == subVals[1] {
				result = 1
			}
			return ic, packetVersion, result
		}

	}
	return 0, 0, 0
}
