/*
 *
 *    This file is part of go-palletone.
 *    go-palletone is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *    go-palletone is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *    You should have received a copy of the GNU General Public License
 *    along with go-palletone.  If not, see <http://www.gnu.org/licenses/>.
 * /
 *
 *  * @author PalletOne core developer <dev@pallet.one>
 *  * @date 2018
 *
 */

package modules

import (
	"encoding/json"
	"github.com/martinlindhe/base36"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsset_MaxSymbol(t *testing.T) {
	s := base36.DecodeToBytes("ZZZZZ")
	t.Logf("Data:%08b", s)
	a, _ := NewAsset("ZZZZZ", AssetType_NonFungibleToken, 18, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, UniqueIdType_Sequence, IDType16{0xff, 0xff, 0xff})
	t.Logf("asset:%x,\r\nstr:%s", a.Bytes(), a.String())
}

func TestAsset_String(t *testing.T) {
	s := base36.DecodeToBytes("DEVIN")
	t.Logf("Data:%08b", s)
	t.Logf("Data:%08b", (byte(5)<<5)|s[0])
	//t.Logf("Data:%08b", base36.DecodeToBytes("00112"))
	//t.Logf("Data:%08b", base36.DecodeToBytes("ZZZZ"))
	//t.Logf("Data:%08b", base36.DecodeToBytes("ZZZ"))
	//t.Logf("Data:%08b", base36.DecodeToBytes("ZZ"))
	//t.Logf("Data:%08b", base36.DecodeToBytes("Z"))
	//symbol := base36.DecodeToBytes("Z")
	//id := IDType16{}
	//copy(id[4-len(symbol):4], symbol)
	//t.Logf("Data:%08b", id)
	asset, err := NewAsset("DEVIN", AssetType_FungibleToken, 4, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16}, UniqueIdType_Null, IDType16{})
	assert.Nil(t, err)
	t.Log(asset.String())
	t.Logf("AssetId:%08b", asset.AssetId)
	asset2, err := NewAsset("ABC", AssetType_FungibleToken, 18, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16}, UniqueIdType_Null, IDType16{})
	assetStr := asset2.String()
	t.Log("Asset2:" + assetStr)
	t.Logf("AssetId:%08b", asset2.AssetId)
	a := Asset{}
	a.SetString(assetStr)
	t.Logf("Asset:%08b,String:%s", a.AssetId, a.String())
	assert.Equal(t, asset2.Bytes(), a.Bytes())

	decimal := byte(8)
	dStr := base36.EncodeBytes([]byte{decimal})
	t.Log(dStr)
}
func TestAsset_SetString(t *testing.T) {
	asset := &Asset{}
	asset.SetString("PTN")
	t.Logf("%08b", asset.AssetId)
	t.Logf("ptn string:%s", asset.String())
	assert.Equal(t, asset.String(), "PTN")
}
func TestPTNAsset(t *testing.T) {
	asset, err := NewAssetId("PTN", AssetType_FungibleToken, 8, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, UniqueIdType_Null)
	assert.Nil(t, err)
	t.Logf("PTN hex:%X,String:%s", asset.Bytes(), asset.ToAssetId())
	assert.Equal(t, asset, PTNCOIN)
}
func TestAssetToString(t *testing.T) {
	t1, _ := NewAssetId("T1", AssetType_FungibleToken, 8, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, UniqueIdType_Null)
	t.Logf("Hex: %x,Str: %s", t1.Bytes(), t1.ToAssetId())
	t2, _ := NewAssetId("T1", AssetType_FungibleToken, 8, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, UniqueIdType_Null)
	t.Logf("Hex: %x,Str: %s", t2.Bytes(), t2.ToAssetId())
}
func TestAsset_MarshalJSON(t *testing.T) {
	ptn := NewPTNAsset()
	js, _ := json.Marshal(ptn)
	t.Logf("%s", string(js))
	str := ptn.String()
	js, _ = json.Marshal(str)
	t.Logf("%s", string(js))
}
func TestAsset721(t *testing.T) {
	t1, _ := NewAsset("CAT0", AssetType_NonFungibleToken, 0, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, UniqueIdType_Sequence, IDType16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 88})
	t.Logf("PRC721 string:%s", t1.String())
	t2, _ := StringToAsset(t1.String())
	assert.Equal(t, t1, t2)

	t11, _ := NewAsset("CAT1", AssetType_NonFungibleToken, 0, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, UniqueIdType_Uuid, IDType16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 88})
	t.Logf("PRC721 string:%s", t11.String())
	t22, _ := StringToAsset(t11.String())
	assert.Equal(t, t11, t22)

	t111, _ := NewAsset("CAT1", AssetType_NonFungibleToken, 0, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, UniqueIdType_UserDefine, IDType16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 88})
	t.Logf("PRC721 string:%s", t111.String())
	t222, _ := StringToAsset(t111.String())
	assert.Equal(t, t111, t222)

}
