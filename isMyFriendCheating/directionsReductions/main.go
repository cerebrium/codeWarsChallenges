// function dirReduc(arr, index = 0){
// 	if (index+1 <= arr.length) {
// 		switch (arr[index+1]) {
// 			case "NORTH":
// 				if (arr[index] === "SOUTH") {
// 					arr.splice(index, 2, "")
// 					let FinalArray = arr.filter(ele => ele.length > 0)
// 					return dirReduc(FinalArray, index > 0 ? index - 1 : 0)
// 				} else {
// 					return dirReduc(arr, index+1)
// 				}
// 			case "SOUTH":
// 				if (arr[index] === "NORTH") {
// 					arr.splice(index, 2, "")
// 					let FinalArray = arr.filter(ele => ele.length > 0)
// 					return dirReduc(FinalArray, index > 0 ? index - 1 : 0)
// 				}
// 				else {
// 					return dirReduc(arr, index+1)
// 				}
// 			case "EAST":
// 				if (arr[index] === "WEST") {
// 					arr.splice(index, 2, "")
// 					let FinalArray = arr.filter(ele => ele.length > 0)
// 					return dirReduc(FinalArray, index > 0 ? index - 1 : 0)
// 				}
// 				else {
// 					return dirReduc(arr, index+1)
// 				}
// 			case "WEST":
// 				if (arr[index] === "EAST") {
// 					arr.splice(index, 2, "")
// 					let FinalArray = arr.filter(ele => ele.length > 0)
// 					return dirReduc(FinalArray, index > 0 ? index - 1 : 0)
// 				}
// 				else {
// 					return dirReduc(arr, index+1)
// 				}
// 		default:
// 			return arr
// 		}
// 	} else {
// 		return arr
// 	}
// }

