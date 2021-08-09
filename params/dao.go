// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// DAOForkBlockExtra is the block header extra-data field to set for the DAO fork
// point and a number of consecutive blocks to allow fast/light syncers to correctly
// pick the side they want  ("dao-hard-fork").
var DAOForkBlockExtra = common.FromHex("0x64616f2d686172642d666f726b")

// DAOForkExtraRange is the number of consecutive blocks from the DAO fork point
// to override the extra-data in to prevent no-fork attacks.
var DAOForkExtraRange = big.NewInt(10)

// DAORefundContract is the address of the refund contract to send DAO balances to.
var DAORefundContract = common.HexToAddress("Exbf4ed7b27f1d666546e30d74d50d173d20bca754")

// DAODrainList is the list of accounts whose full balances will be moved into a
// refund contract at the beginning of the dao-fork block.
func DAODrainList() []common.Address {
	return []common.Address{
		common.HexToAddress("Exd4fe7bc31cedb7bfb8a345f31e668033056b2728"),
		common.HexToAddress("Exb3fb0e5aba0e20e5c49d252dfd30e102b171a425"),
		common.HexToAddress("Ex2c19c7f9ae8b751e37aeb2d93a699722395ae18f"),
		common.HexToAddress("Execd135fa4f61a655311e86238c92adcd779555d2"),
		common.HexToAddress("Ex1975bd06d486162d5dc297798dfc41edd5d160a7"),
		common.HexToAddress("Exa3acf3a1e16b1d7c315e23510fdd7847b48234f6"),
		common.HexToAddress("Ex319f70bab6845585f412ec7724b744fec6095c85"),
		common.HexToAddress("Ex06706dd3f2c9abf0a21ddcc6941d9b86f0596936"),
		common.HexToAddress("Ex5c8536898fbb74fc7445814902fd08422eac56d0"),
		common.HexToAddress("Ex6966ab0d485353095148a2155858910e0965b6f9"),
		common.HexToAddress("Ex779543a0491a837ca36ce8c635d6154e3c4911a6"),
		common.HexToAddress("Ex2a5ed960395e2a49b1c758cef4aa15213cfd874c"),
		common.HexToAddress("Ex5c6e67ccd5849c0d29219c4f95f1a7a93b3f5dc5"),
		common.HexToAddress("Ex9c50426be05db97f5d64fc54bf89eff947f0a321"),
		common.HexToAddress("Ex200450f06520bdd6c527622a273333384d870efb"),
		common.HexToAddress("Exbe8539bfe837b67d1282b2b1d61c3f723966f049"),
		common.HexToAddress("Ex6b0c4d41ba9ab8d8cfb5d379c69a612f2ced8ecb"),
		common.HexToAddress("Exf1385fb24aad0cd7432824085e42aff90886fef5"),
		common.HexToAddress("Exd1ac8b1ef1b69ff51d1d401a476e7e612414f091"),
		common.HexToAddress("Ex8163e7fb499e90f8544ea62bbf80d21cd26d9efd"),
		common.HexToAddress("Ex51e0ddd9998364a2eb38588679f0d2c42653e4a6"),
		common.HexToAddress("Ex627a0a960c079c21c34f7612d5d230e01b4ad4c7"),
		common.HexToAddress("Exf0b1aa0eb660754448a7937c022e30aa692fe0c5"),
		common.HexToAddress("Ex24c4d950dfd4dd1902bbed3508144a54542bba94"),
		common.HexToAddress("Ex9f27daea7aca0aa0446220b98d028715e3bc803d"),
		common.HexToAddress("Exa5dc5acd6a7968a4554d89d65e59b7fd3bff0f90"),
		common.HexToAddress("Exd9aef3a1e38a39c16b31d1ace71bca8ef58d315b"),
		common.HexToAddress("Ex63ed5a272de2f6d968408b4acb9024f4cc208ebf"),
		common.HexToAddress("Ex6f6704e5a10332af6672e50b3d9754dc460dfa4d"),
		common.HexToAddress("Ex77ca7b50b6cd7e2f3fa008e24ab793fd56cb15f6"),
		common.HexToAddress("Ex492ea3bb0f3315521c31f273e565b868fc090f17"),
		common.HexToAddress("Ex0ff30d6de14a8224aa97b78aea5388d1c51c1f00"),
		common.HexToAddress("Ex9ea779f907f0b315b364b0cfc39a0fde5b02a416"),
		common.HexToAddress("Exceaeb481747ca6c540a000c1f3641f8cef161fa7"),
		common.HexToAddress("Excc34673c6c40e791051898567a1222daf90be287"),
		common.HexToAddress("Ex579a80d909f346fbfb1189493f521d7f48d52238"),
		common.HexToAddress("Exe308bd1ac5fda103967359b2712dd89deffb7973"),
		common.HexToAddress("Ex4cb31628079fb14e4bc3cd5e30c2f7489b00960c"),
		common.HexToAddress("Exac1ecab32727358dba8962a0f3b261731aad9723"),
		common.HexToAddress("Ex4fd6ace747f06ece9c49699c7cabc62d02211f75"),
		common.HexToAddress("Ex440c59b325d2997a134c2c7c60a8c61611212bad"),
		common.HexToAddress("Ex4486a3d68fac6967006d7a517b889fd3f98c102b"),
		common.HexToAddress("Ex9c15b54878ba618f494b38f0ae7443db6af648ba"),
		common.HexToAddress("Ex27b137a85656544b1ccb5a0f2e561a5703c6a68f"),
		common.HexToAddress("Ex21c7fdb9ed8d291d79ffd82eb2c4356ec0d81241"),
		common.HexToAddress("Ex23b75c2f6791eef49c69684db4c6c1f93bf49a50"),
		common.HexToAddress("Ex1ca6abd14d30affe533b24d7a21bff4c2d5e1f3b"),
		common.HexToAddress("Exb9637156d330c0d605a791f1c31ba5890582fe1c"),
		common.HexToAddress("Ex6131c42fa982e56929107413a9d526fd99405560"),
		common.HexToAddress("Ex1591fc0f688c81fbeb17f5426a162a7024d430c2"),
		common.HexToAddress("Ex542a9515200d14b68e934e9830d91645a980dd7a"),
		common.HexToAddress("Exc4bbd073882dd2add2424cf47d35213405b01324"),
		common.HexToAddress("Ex782495b7b3355efb2833d56ecb34dc22ad7dfcc4"),
		common.HexToAddress("Ex58b95c9a9d5d26825e70a82b6adb139d3fd829eb"),
		common.HexToAddress("Ex3ba4d81db016dc2890c81f3acec2454bff5aada5"),
		common.HexToAddress("Exb52042c8ca3f8aa246fa79c3feaa3d959347c0ab"),
		common.HexToAddress("Exe4ae1efdfc53b73893af49113d8694a057b9c0d1"),
		common.HexToAddress("Ex3c02a7bc0391e86d91b7d144e61c2c01a25a79c5"),
		common.HexToAddress("Ex0737a6b837f97f46ebade41b9bc3e1c509c85c53"),
		common.HexToAddress("Ex97f43a37f595ab5dd318fb46e7a155eae057317a"),
		common.HexToAddress("Ex52c5317c848ba20c7504cb2c8052abd1fde29d03"),
		common.HexToAddress("Ex4863226780fe7c0356454236d3b1c8792785748d"),
		common.HexToAddress("Ex5d2b2e6fcbe3b11d26b525e085ff818dae332479"),
		common.HexToAddress("Ex5f9f3392e9f62f63b8eac0beb55541fc8627f42c"),
		common.HexToAddress("Ex057b56736d32b86616a10f619859c6cd6f59092a"),
		common.HexToAddress("Ex9aa008f65de0b923a2a4f02012ad034a5e2e2192"),
		common.HexToAddress("Ex304a554a310c7e546dfe434669c62820b7d83490"),
		common.HexToAddress("Ex914d1b8b43e92723e64fd0a06f5bdb8dd9b10c79"),
		common.HexToAddress("Ex4deb0033bb26bc534b197e61d19e0733e5679784"),
		common.HexToAddress("Ex07f5c1e1bc2c93e0402f23341973a0e043f7bf8a"),
		common.HexToAddress("Ex35a051a0010aba705c9008d7a7eff6fb88f6ea7b"),
		common.HexToAddress("Ex4fa802324e929786dbda3b8820dc7834e9134a2a"),
		common.HexToAddress("Ex9da397b9e80755301a3b32173283a91c0ef6c87e"),
		common.HexToAddress("Ex8d9edb3054ce5c5774a420ac37ebae0ac02343c6"),
		common.HexToAddress("Ex0101f3be8ebb4bbd39a2e3b9a3639d4259832fd9"),
		common.HexToAddress("Ex5dc28b15dffed94048d73806ce4b7a4612a1d48f"),
		common.HexToAddress("Exbcf899e6c7d9d5a215ab1e3444c86806fa854c76"),
		common.HexToAddress("Ex12e626b0eebfe86a56d633b9864e389b45dcb260"),
		common.HexToAddress("Exa2f1ccba9395d7fcb155bba8bc92db9bafaeade7"),
		common.HexToAddress("Exec8e57756626fdc07c63ad2eafbd28d08e7b0ca5"),
		common.HexToAddress("Exd164b088bd9108b60d0ca3751da4bceb207b0782"),
		common.HexToAddress("Ex6231b6d0d5e77fe001c2a460bd9584fee60d409b"),
		common.HexToAddress("Ex1cba23d343a983e9b5cfd19496b9a9701ada385f"),
		common.HexToAddress("Exa82f360a8d3455c5c41366975bde739c37bfeb8a"),
		common.HexToAddress("Ex9fcd2deaff372a39cc679d5c5e4de7bafb0b1339"),
		common.HexToAddress("Ex005f5cee7a43331d5a3d3eec71305925a62f34b6"),
		common.HexToAddress("Ex0e0da70933f4c7849fc0d203f5d1d43b9ae4532d"),
		common.HexToAddress("Exd131637d5275fd1a68a3200f4ad25c71a2a9522e"),
		common.HexToAddress("Exbc07118b9ac290e4622f5e77a0853539789effbe"),
		common.HexToAddress("Ex47e7aa56d6bdf3f36be34619660de61275420af8"),
		common.HexToAddress("Exacd87e28b0c9d1254e868b81cba4cc20d9a32225"),
		common.HexToAddress("Exadf80daec7ba8dcf15392f1ac611fff65d94f880"),
		common.HexToAddress("Ex5524c55fb03cf21f549444ccbecb664d0acad706"),
		common.HexToAddress("Ex40b803a9abce16f50f36a77ba41180eb90023925"),
		common.HexToAddress("Ex17802f43a0137c506ba92291391a8a8f207f487d"),
		common.HexToAddress("Ex253488078a4edf4d6f42f113d1e62836a942cf1a"),
		common.HexToAddress("Ex86af3e9626fce1957c82e88cbf04ddf3a2ed7915"),
		common.HexToAddress("Exb136707642a4ea12fb4bae820f03d2562ebff487"),
		common.HexToAddress("Exdbe9b615a3ae8709af8b93336ce9b477e4ac0940"),
		common.HexToAddress("Exf14c14075d6c4ed84b86798af0956deef67365b5"),
		common.HexToAddress("Exca544e5c4687d109611d0f8f928b53a25af72448"),
		common.HexToAddress("Exaeeb8ff27288bdabc0fa5ebb731b6f409507516c"),
		common.HexToAddress("Excbb9d3703e651b0d496cdefb8b92c25aeb2171f7"),
		common.HexToAddress("Ex6d87578288b6cb5549d5076a207456a1f6a63dc0"),
		common.HexToAddress("Exb2c6f0dfbb716ac562e2d85d6cb2f8d5ee87603e"),
		common.HexToAddress("Exaccc230e8a6e5be9160b8cdf2864dd2a001c28b6"),
		common.HexToAddress("Ex2b3455ec7fedf16e646268bf88846bd7a2319bb2"),
		common.HexToAddress("Ex4613f3bca5c44ea06337a9e439fbc6d42e501d0a"),
		common.HexToAddress("Exd343b217de44030afaa275f54d31a9317c7f441e"),
		common.HexToAddress("Ex84ef4b2357079cd7a7c69fd7a37cd0609a679106"),
		common.HexToAddress("Exda2fef9e4a3230988ff17df2165440f37e8b1708"),
		common.HexToAddress("Exf4c64518ea10f995918a454158c6b61407ea345c"),
		common.HexToAddress("Ex7602b46df5390e432ef1c307d4f2c9ff6d65cc97"),
		common.HexToAddress("Exbb9bc244d798123fde783fcc1c72d3bb8c189413"),
		common.HexToAddress("Ex807640a13483f8ac783c557fcdf27be11ea4ac7a"),
	}
}
