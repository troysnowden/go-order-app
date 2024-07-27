package main

type requiredPackInfo struct {
	PackSize int `json:"packSize"`
	Amount   int `json:"amount"`
}

func getPackResponse(itemsOrdered int) []requiredPackInfo {

	var minItemsToSend = minimumItemsToSend(itemsOrdered)

	return requiredPacks(minItemsToSend)
}

func minimumItemsToSend(itemsOrdered int) int {
	var itemsRemaining = itemsOrdered
	var minItemsToSend = 0

	for _, packSize := range packSizes() {
		if itemsRemaining <= 0 {
			break
		}

		// number of packs required, accounting for any left over.
		var packsOfCurrentSizeRequired = itemsRemaining / packSize

		if packsOfCurrentSizeRequired > 0 {

			var itemsInPack = packsOfCurrentSizeRequired * packSize

			itemsRemaining -= itemsInPack
			minItemsToSend += itemsInPack
		}
	}

	if itemsRemaining > 0 {
		minItemsToSend += packSizes()[len(packSizes())-1]
	}
	return minItemsToSend
}

func requiredPacks(itemsRemaining int) []requiredPackInfo {

	var requiredPacks []requiredPackInfo

	// loop through again to get minimum number of packages
	for _, packSize := range packSizes() {
		if itemsRemaining <= 0 {
			break
		}

		var packsOfCurrentSizeRequired = itemsRemaining / packSize

		if packsOfCurrentSizeRequired > 0 {
			requiredPacks = append(requiredPacks, requiredPackInfo{PackSize: packSize, Amount: packsOfCurrentSizeRequired})

			var itemsInPack = packsOfCurrentSizeRequired * packSize

			itemsRemaining -= itemsInPack
		}
	}

	if itemsRemaining > 0 {
		var smallestPackage = packSizes()[len(packSizes())-1]

		// number of packs required, accounting for any remaining.
		// var numberOfPackages = (itemsRemaining + smallestPackage - 1) / smallestPackage

		requiredPacks = addPack(smallestPackage, requiredPacks)
	}
	return requiredPacks
}

// add pack, checking that pack does not already exist in the list. Only possible with leftover items
func addPack(packSize int, packList []requiredPackInfo) []requiredPackInfo {
	for i, pack := range packList {
		if pack.PackSize == packSize {
			packList[i].Amount += 1
			return packList
		}
	}
	return append(packList, requiredPackInfo{PackSize: packSize, Amount: 1})
}
