// Code generated by "stringer -type=Code -linecomment"; DO NOT EDIT.

package multicodec

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Identity-0]
	_ = x[Cidv1-1]
	_ = x[Cidv2-2]
	_ = x[Cidv3-3]
	_ = x[Ip4-4]
	_ = x[Tcp-6]
	_ = x[Sha1-17]
	_ = x[Sha2_256-18]
	_ = x[Sha2_512-19]
	_ = x[Sha3_512-20]
	_ = x[Sha3_384-21]
	_ = x[Sha3_256-22]
	_ = x[Sha3_224-23]
	_ = x[Shake128-24]
	_ = x[Shake256-25]
	_ = x[Keccak224-26]
	_ = x[Keccak256-27]
	_ = x[Keccak384-28]
	_ = x[Keccak512-29]
	_ = x[Blake3-30]
	_ = x[Dccp-33]
	_ = x[Murmur3_128-34]
	_ = x[Murmur3_32-35]
	_ = x[Ip6-41]
	_ = x[Ip6zone-42]
	_ = x[Path-47]
	_ = x[Multicodec-48]
	_ = x[Multihash-49]
	_ = x[Multiaddr-50]
	_ = x[Multibase-51]
	_ = x[Dns-53]
	_ = x[Dns4-54]
	_ = x[Dns6-55]
	_ = x[Dnsaddr-56]
	_ = x[Protobuf-80]
	_ = x[Cbor-81]
	_ = x[Raw-85]
	_ = x[DblSha2_256-86]
	_ = x[Rlp-96]
	_ = x[Bencode-99]
	_ = x[DagPb-112]
	_ = x[DagCbor-113]
	_ = x[Libp2pKey-114]
	_ = x[GitRaw-120]
	_ = x[TorrentInfo-123]
	_ = x[TorrentFile-124]
	_ = x[LeofcoinBlock-129]
	_ = x[LeofcoinTx-130]
	_ = x[LeofcoinPr-131]
	_ = x[Sctp-132]
	_ = x[DagJose-133]
	_ = x[DagCose-134]
	_ = x[EthBlock-144]
	_ = x[EthBlockList-145]
	_ = x[EthTxTrie-146]
	_ = x[EthTx-147]
	_ = x[EthTxReceiptTrie-148]
	_ = x[EthTxReceipt-149]
	_ = x[EthStateTrie-150]
	_ = x[EthAccountSnapshot-151]
	_ = x[EthStorageTrie-152]
	_ = x[BitcoinBlock-176]
	_ = x[BitcoinTx-177]
	_ = x[BitcoinWitnessCommitment-178]
	_ = x[ZcashBlock-192]
	_ = x[ZcashTx-193]
	_ = x[Caip50-202]
	_ = x[Streamid-206]
	_ = x[StellarBlock-208]
	_ = x[StellarTx-209]
	_ = x[Md4-212]
	_ = x[Md5-213]
	_ = x[Bmt-214]
	_ = x[DecredBlock-224]
	_ = x[DecredTx-225]
	_ = x[IpldNs-226]
	_ = x[IpfsNs-227]
	_ = x[SwarmNs-228]
	_ = x[IpnsNs-229]
	_ = x[Zeronet-230]
	_ = x[Secp256k1Pub-231]
	_ = x[Bls12_381G1Pub-234]
	_ = x[Bls12_381G2Pub-235]
	_ = x[X25519Pub-236]
	_ = x[Ed25519Pub-237]
	_ = x[Bls12_381G1g2Pub-238]
	_ = x[DashBlock-240]
	_ = x[DashTx-241]
	_ = x[SwarmManifest-250]
	_ = x[SwarmFeed-251]
	_ = x[Udp-273]
	_ = x[P2pWebrtcStar-275]
	_ = x[P2pWebrtcDirect-276]
	_ = x[P2pStardust-277]
	_ = x[P2pCircuit-290]
	_ = x[DagJson-297]
	_ = x[Udt-301]
	_ = x[Utp-302]
	_ = x[Unix-400]
	_ = x[Thread-406]
	_ = x[P2p-421]
	_ = x[Ipfs-421]
	_ = x[Https-443]
	_ = x[Onion-444]
	_ = x[Onion3-445]
	_ = x[Garlic64-446]
	_ = x[Garlic32-447]
	_ = x[Tls-448]
	_ = x[Noise-454]
	_ = x[Quic-460]
	_ = x[Ws-477]
	_ = x[Wss-478]
	_ = x[P2pWebsocketStar-479]
	_ = x[Http-480]
	_ = x[Json-512]
	_ = x[Messagepack-513]
	_ = x[Libp2pPeerRecord-769]
	_ = x[Libp2pRelayRsvp-770]
	_ = x[CarIndexSorted-1024]
	_ = x[Sha2_256Trunc254Padded-4114]
	_ = x[Ripemd128-4178]
	_ = x[Ripemd160-4179]
	_ = x[Ripemd256-4180]
	_ = x[Ripemd320-4181]
	_ = x[X11-4352]
	_ = x[P256Pub-4608]
	_ = x[P384Pub-4609]
	_ = x[P521Pub-4610]
	_ = x[Ed448Pub-4611]
	_ = x[X448Pub-4612]
	_ = x[Ed25519Priv-4864]
	_ = x[Secp256k1Priv-4865]
	_ = x[X25519Priv-4866]
	_ = x[Kangarootwelve-7425]
	_ = x[Sm3_256-21325]
	_ = x[Blake2b8-45569]
	_ = x[Blake2b16-45570]
	_ = x[Blake2b24-45571]
	_ = x[Blake2b32-45572]
	_ = x[Blake2b40-45573]
	_ = x[Blake2b48-45574]
	_ = x[Blake2b56-45575]
	_ = x[Blake2b64-45576]
	_ = x[Blake2b72-45577]
	_ = x[Blake2b80-45578]
	_ = x[Blake2b88-45579]
	_ = x[Blake2b96-45580]
	_ = x[Blake2b104-45581]
	_ = x[Blake2b112-45582]
	_ = x[Blake2b120-45583]
	_ = x[Blake2b128-45584]
	_ = x[Blake2b136-45585]
	_ = x[Blake2b144-45586]
	_ = x[Blake2b152-45587]
	_ = x[Blake2b160-45588]
	_ = x[Blake2b168-45589]
	_ = x[Blake2b176-45590]
	_ = x[Blake2b184-45591]
	_ = x[Blake2b192-45592]
	_ = x[Blake2b200-45593]
	_ = x[Blake2b208-45594]
	_ = x[Blake2b216-45595]
	_ = x[Blake2b224-45596]
	_ = x[Blake2b232-45597]
	_ = x[Blake2b240-45598]
	_ = x[Blake2b248-45599]
	_ = x[Blake2b256-45600]
	_ = x[Blake2b264-45601]
	_ = x[Blake2b272-45602]
	_ = x[Blake2b280-45603]
	_ = x[Blake2b288-45604]
	_ = x[Blake2b296-45605]
	_ = x[Blake2b304-45606]
	_ = x[Blake2b312-45607]
	_ = x[Blake2b320-45608]
	_ = x[Blake2b328-45609]
	_ = x[Blake2b336-45610]
	_ = x[Blake2b344-45611]
	_ = x[Blake2b352-45612]
	_ = x[Blake2b360-45613]
	_ = x[Blake2b368-45614]
	_ = x[Blake2b376-45615]
	_ = x[Blake2b384-45616]
	_ = x[Blake2b392-45617]
	_ = x[Blake2b400-45618]
	_ = x[Blake2b408-45619]
	_ = x[Blake2b416-45620]
	_ = x[Blake2b424-45621]
	_ = x[Blake2b432-45622]
	_ = x[Blake2b440-45623]
	_ = x[Blake2b448-45624]
	_ = x[Blake2b456-45625]
	_ = x[Blake2b464-45626]
	_ = x[Blake2b472-45627]
	_ = x[Blake2b480-45628]
	_ = x[Blake2b488-45629]
	_ = x[Blake2b496-45630]
	_ = x[Blake2b504-45631]
	_ = x[Blake2b512-45632]
	_ = x[Blake2s8-45633]
	_ = x[Blake2s16-45634]
	_ = x[Blake2s24-45635]
	_ = x[Blake2s32-45636]
	_ = x[Blake2s40-45637]
	_ = x[Blake2s48-45638]
	_ = x[Blake2s56-45639]
	_ = x[Blake2s64-45640]
	_ = x[Blake2s72-45641]
	_ = x[Blake2s80-45642]
	_ = x[Blake2s88-45643]
	_ = x[Blake2s96-45644]
	_ = x[Blake2s104-45645]
	_ = x[Blake2s112-45646]
	_ = x[Blake2s120-45647]
	_ = x[Blake2s128-45648]
	_ = x[Blake2s136-45649]
	_ = x[Blake2s144-45650]
	_ = x[Blake2s152-45651]
	_ = x[Blake2s160-45652]
	_ = x[Blake2s168-45653]
	_ = x[Blake2s176-45654]
	_ = x[Blake2s184-45655]
	_ = x[Blake2s192-45656]
	_ = x[Blake2s200-45657]
	_ = x[Blake2s208-45658]
	_ = x[Blake2s216-45659]
	_ = x[Blake2s224-45660]
	_ = x[Blake2s232-45661]
	_ = x[Blake2s240-45662]
	_ = x[Blake2s248-45663]
	_ = x[Blake2s256-45664]
	_ = x[Skein256_8-45825]
	_ = x[Skein256_16-45826]
	_ = x[Skein256_24-45827]
	_ = x[Skein256_32-45828]
	_ = x[Skein256_40-45829]
	_ = x[Skein256_48-45830]
	_ = x[Skein256_56-45831]
	_ = x[Skein256_64-45832]
	_ = x[Skein256_72-45833]
	_ = x[Skein256_80-45834]
	_ = x[Skein256_88-45835]
	_ = x[Skein256_96-45836]
	_ = x[Skein256_104-45837]
	_ = x[Skein256_112-45838]
	_ = x[Skein256_120-45839]
	_ = x[Skein256_128-45840]
	_ = x[Skein256_136-45841]
	_ = x[Skein256_144-45842]
	_ = x[Skein256_152-45843]
	_ = x[Skein256_160-45844]
	_ = x[Skein256_168-45845]
	_ = x[Skein256_176-45846]
	_ = x[Skein256_184-45847]
	_ = x[Skein256_192-45848]
	_ = x[Skein256_200-45849]
	_ = x[Skein256_208-45850]
	_ = x[Skein256_216-45851]
	_ = x[Skein256_224-45852]
	_ = x[Skein256_232-45853]
	_ = x[Skein256_240-45854]
	_ = x[Skein256_248-45855]
	_ = x[Skein256_256-45856]
	_ = x[Skein512_8-45857]
	_ = x[Skein512_16-45858]
	_ = x[Skein512_24-45859]
	_ = x[Skein512_32-45860]
	_ = x[Skein512_40-45861]
	_ = x[Skein512_48-45862]
	_ = x[Skein512_56-45863]
	_ = x[Skein512_64-45864]
	_ = x[Skein512_72-45865]
	_ = x[Skein512_80-45866]
	_ = x[Skein512_88-45867]
	_ = x[Skein512_96-45868]
	_ = x[Skein512_104-45869]
	_ = x[Skein512_112-45870]
	_ = x[Skein512_120-45871]
	_ = x[Skein512_128-45872]
	_ = x[Skein512_136-45873]
	_ = x[Skein512_144-45874]
	_ = x[Skein512_152-45875]
	_ = x[Skein512_160-45876]
	_ = x[Skein512_168-45877]
	_ = x[Skein512_176-45878]
	_ = x[Skein512_184-45879]
	_ = x[Skein512_192-45880]
	_ = x[Skein512_200-45881]
	_ = x[Skein512_208-45882]
	_ = x[Skein512_216-45883]
	_ = x[Skein512_224-45884]
	_ = x[Skein512_232-45885]
	_ = x[Skein512_240-45886]
	_ = x[Skein512_248-45887]
	_ = x[Skein512_256-45888]
	_ = x[Skein512_264-45889]
	_ = x[Skein512_272-45890]
	_ = x[Skein512_280-45891]
	_ = x[Skein512_288-45892]
	_ = x[Skein512_296-45893]
	_ = x[Skein512_304-45894]
	_ = x[Skein512_312-45895]
	_ = x[Skein512_320-45896]
	_ = x[Skein512_328-45897]
	_ = x[Skein512_336-45898]
	_ = x[Skein512_344-45899]
	_ = x[Skein512_352-45900]
	_ = x[Skein512_360-45901]
	_ = x[Skein512_368-45902]
	_ = x[Skein512_376-45903]
	_ = x[Skein512_384-45904]
	_ = x[Skein512_392-45905]
	_ = x[Skein512_400-45906]
	_ = x[Skein512_408-45907]
	_ = x[Skein512_416-45908]
	_ = x[Skein512_424-45909]
	_ = x[Skein512_432-45910]
	_ = x[Skein512_440-45911]
	_ = x[Skein512_448-45912]
	_ = x[Skein512_456-45913]
	_ = x[Skein512_464-45914]
	_ = x[Skein512_472-45915]
	_ = x[Skein512_480-45916]
	_ = x[Skein512_488-45917]
	_ = x[Skein512_496-45918]
	_ = x[Skein512_504-45919]
	_ = x[Skein512_512-45920]
	_ = x[Skein1024_8-45921]
	_ = x[Skein1024_16-45922]
	_ = x[Skein1024_24-45923]
	_ = x[Skein1024_32-45924]
	_ = x[Skein1024_40-45925]
	_ = x[Skein1024_48-45926]
	_ = x[Skein1024_56-45927]
	_ = x[Skein1024_64-45928]
	_ = x[Skein1024_72-45929]
	_ = x[Skein1024_80-45930]
	_ = x[Skein1024_88-45931]
	_ = x[Skein1024_96-45932]
	_ = x[Skein1024_104-45933]
	_ = x[Skein1024_112-45934]
	_ = x[Skein1024_120-45935]
	_ = x[Skein1024_128-45936]
	_ = x[Skein1024_136-45937]
	_ = x[Skein1024_144-45938]
	_ = x[Skein1024_152-45939]
	_ = x[Skein1024_160-45940]
	_ = x[Skein1024_168-45941]
	_ = x[Skein1024_176-45942]
	_ = x[Skein1024_184-45943]
	_ = x[Skein1024_192-45944]
	_ = x[Skein1024_200-45945]
	_ = x[Skein1024_208-45946]
	_ = x[Skein1024_216-45947]
	_ = x[Skein1024_224-45948]
	_ = x[Skein1024_232-45949]
	_ = x[Skein1024_240-45950]
	_ = x[Skein1024_248-45951]
	_ = x[Skein1024_256-45952]
	_ = x[Skein1024_264-45953]
	_ = x[Skein1024_272-45954]
	_ = x[Skein1024_280-45955]
	_ = x[Skein1024_288-45956]
	_ = x[Skein1024_296-45957]
	_ = x[Skein1024_304-45958]
	_ = x[Skein1024_312-45959]
	_ = x[Skein1024_320-45960]
	_ = x[Skein1024_328-45961]
	_ = x[Skein1024_336-45962]
	_ = x[Skein1024_344-45963]
	_ = x[Skein1024_352-45964]
	_ = x[Skein1024_360-45965]
	_ = x[Skein1024_368-45966]
	_ = x[Skein1024_376-45967]
	_ = x[Skein1024_384-45968]
	_ = x[Skein1024_392-45969]
	_ = x[Skein1024_400-45970]
	_ = x[Skein1024_408-45971]
	_ = x[Skein1024_416-45972]
	_ = x[Skein1024_424-45973]
	_ = x[Skein1024_432-45974]
	_ = x[Skein1024_440-45975]
	_ = x[Skein1024_448-45976]
	_ = x[Skein1024_456-45977]
	_ = x[Skein1024_464-45978]
	_ = x[Skein1024_472-45979]
	_ = x[Skein1024_480-45980]
	_ = x[Skein1024_488-45981]
	_ = x[Skein1024_496-45982]
	_ = x[Skein1024_504-45983]
	_ = x[Skein1024_512-45984]
	_ = x[Skein1024_520-45985]
	_ = x[Skein1024_528-45986]
	_ = x[Skein1024_536-45987]
	_ = x[Skein1024_544-45988]
	_ = x[Skein1024_552-45989]
	_ = x[Skein1024_560-45990]
	_ = x[Skein1024_568-45991]
	_ = x[Skein1024_576-45992]
	_ = x[Skein1024_584-45993]
	_ = x[Skein1024_592-45994]
	_ = x[Skein1024_600-45995]
	_ = x[Skein1024_608-45996]
	_ = x[Skein1024_616-45997]
	_ = x[Skein1024_624-45998]
	_ = x[Skein1024_632-45999]
	_ = x[Skein1024_640-46000]
	_ = x[Skein1024_648-46001]
	_ = x[Skein1024_656-46002]
	_ = x[Skein1024_664-46003]
	_ = x[Skein1024_672-46004]
	_ = x[Skein1024_680-46005]
	_ = x[Skein1024_688-46006]
	_ = x[Skein1024_696-46007]
	_ = x[Skein1024_704-46008]
	_ = x[Skein1024_712-46009]
	_ = x[Skein1024_720-46010]
	_ = x[Skein1024_728-46011]
	_ = x[Skein1024_736-46012]
	_ = x[Skein1024_744-46013]
	_ = x[Skein1024_752-46014]
	_ = x[Skein1024_760-46015]
	_ = x[Skein1024_768-46016]
	_ = x[Skein1024_776-46017]
	_ = x[Skein1024_784-46018]
	_ = x[Skein1024_792-46019]
	_ = x[Skein1024_800-46020]
	_ = x[Skein1024_808-46021]
	_ = x[Skein1024_816-46022]
	_ = x[Skein1024_824-46023]
	_ = x[Skein1024_832-46024]
	_ = x[Skein1024_840-46025]
	_ = x[Skein1024_848-46026]
	_ = x[Skein1024_856-46027]
	_ = x[Skein1024_864-46028]
	_ = x[Skein1024_872-46029]
	_ = x[Skein1024_880-46030]
	_ = x[Skein1024_888-46031]
	_ = x[Skein1024_896-46032]
	_ = x[Skein1024_904-46033]
	_ = x[Skein1024_912-46034]
	_ = x[Skein1024_920-46035]
	_ = x[Skein1024_928-46036]
	_ = x[Skein1024_936-46037]
	_ = x[Skein1024_944-46038]
	_ = x[Skein1024_952-46039]
	_ = x[Skein1024_960-46040]
	_ = x[Skein1024_968-46041]
	_ = x[Skein1024_976-46042]
	_ = x[Skein1024_984-46043]
	_ = x[Skein1024_992-46044]
	_ = x[Skein1024_1000-46045]
	_ = x[Skein1024_1008-46046]
	_ = x[Skein1024_1016-46047]
	_ = x[Skein1024_1024-46048]
	_ = x[PoseidonBls12_381A2Fc1-46081]
	_ = x[PoseidonBls12_381A2Fc1Sc-46082]
	_ = x[ZeroxcertImprint256-52753]
	_ = x[FilCommitmentUnsealed-61697]
	_ = x[FilCommitmentSealed-61698]
	_ = x[HolochainAdrV0-8417572]
	_ = x[HolochainAdrV1-8483108]
	_ = x[HolochainKeyV0-9728292]
	_ = x[HolochainKeyV1-9793828]
	_ = x[HolochainSigV0-10645796]
	_ = x[HolochainSigV1-10711332]
	_ = x[SkynetNs-11639056]
}

const _Code_name = "identitycidv1cidv2cidv3ip4tcpsha1sha2-256sha2-512sha3-512sha3-384sha3-256sha3-224shake-128shake-256keccak-224keccak-256keccak-384keccak-512blake3dccpmurmur3-128murmur3-32ip6ip6zonepathmulticodecmultihashmultiaddrmultibasednsdns4dns6dnsaddrprotobufcborrawdbl-sha2-256rlpbencodedag-pbdag-cborlibp2p-keygit-rawtorrent-infotorrent-fileleofcoin-blockleofcoin-txleofcoin-prsctpdag-josedag-coseeth-blocketh-block-listeth-tx-trieeth-txeth-tx-receipt-trieeth-tx-receipteth-state-trieeth-account-snapshoteth-storage-triebitcoin-blockbitcoin-txbitcoin-witness-commitmentzcash-blockzcash-txcaip-50streamidstellar-blockstellar-txmd4md5bmtdecred-blockdecred-txipld-nsipfs-nsswarm-nsipns-nszeronetsecp256k1-pubbls12_381-g1-pubbls12_381-g2-pubx25519-pubed25519-pubbls12_381-g1g2-pubdash-blockdash-txswarm-manifestswarm-feedudpp2p-webrtc-starp2p-webrtc-directp2p-stardustp2p-circuitdag-jsonudtutpunixthreadp2phttpsoniononion3garlic64garlic32tlsnoisequicwswssp2p-websocket-starhttpjsonmessagepacklibp2p-peer-recordlibp2p-relay-rsvpcar-index-sortedsha2-256-trunc254-paddedripemd-128ripemd-160ripemd-256ripemd-320x11p256-pubp384-pubp521-pubed448-pubx448-pubed25519-privsecp256k1-privx25519-privkangarootwelvesm3-256blake2b-8blake2b-16blake2b-24blake2b-32blake2b-40blake2b-48blake2b-56blake2b-64blake2b-72blake2b-80blake2b-88blake2b-96blake2b-104blake2b-112blake2b-120blake2b-128blake2b-136blake2b-144blake2b-152blake2b-160blake2b-168blake2b-176blake2b-184blake2b-192blake2b-200blake2b-208blake2b-216blake2b-224blake2b-232blake2b-240blake2b-248blake2b-256blake2b-264blake2b-272blake2b-280blake2b-288blake2b-296blake2b-304blake2b-312blake2b-320blake2b-328blake2b-336blake2b-344blake2b-352blake2b-360blake2b-368blake2b-376blake2b-384blake2b-392blake2b-400blake2b-408blake2b-416blake2b-424blake2b-432blake2b-440blake2b-448blake2b-456blake2b-464blake2b-472blake2b-480blake2b-488blake2b-496blake2b-504blake2b-512blake2s-8blake2s-16blake2s-24blake2s-32blake2s-40blake2s-48blake2s-56blake2s-64blake2s-72blake2s-80blake2s-88blake2s-96blake2s-104blake2s-112blake2s-120blake2s-128blake2s-136blake2s-144blake2s-152blake2s-160blake2s-168blake2s-176blake2s-184blake2s-192blake2s-200blake2s-208blake2s-216blake2s-224blake2s-232blake2s-240blake2s-248blake2s-256skein256-8skein256-16skein256-24skein256-32skein256-40skein256-48skein256-56skein256-64skein256-72skein256-80skein256-88skein256-96skein256-104skein256-112skein256-120skein256-128skein256-136skein256-144skein256-152skein256-160skein256-168skein256-176skein256-184skein256-192skein256-200skein256-208skein256-216skein256-224skein256-232skein256-240skein256-248skein256-256skein512-8skein512-16skein512-24skein512-32skein512-40skein512-48skein512-56skein512-64skein512-72skein512-80skein512-88skein512-96skein512-104skein512-112skein512-120skein512-128skein512-136skein512-144skein512-152skein512-160skein512-168skein512-176skein512-184skein512-192skein512-200skein512-208skein512-216skein512-224skein512-232skein512-240skein512-248skein512-256skein512-264skein512-272skein512-280skein512-288skein512-296skein512-304skein512-312skein512-320skein512-328skein512-336skein512-344skein512-352skein512-360skein512-368skein512-376skein512-384skein512-392skein512-400skein512-408skein512-416skein512-424skein512-432skein512-440skein512-448skein512-456skein512-464skein512-472skein512-480skein512-488skein512-496skein512-504skein512-512skein1024-8skein1024-16skein1024-24skein1024-32skein1024-40skein1024-48skein1024-56skein1024-64skein1024-72skein1024-80skein1024-88skein1024-96skein1024-104skein1024-112skein1024-120skein1024-128skein1024-136skein1024-144skein1024-152skein1024-160skein1024-168skein1024-176skein1024-184skein1024-192skein1024-200skein1024-208skein1024-216skein1024-224skein1024-232skein1024-240skein1024-248skein1024-256skein1024-264skein1024-272skein1024-280skein1024-288skein1024-296skein1024-304skein1024-312skein1024-320skein1024-328skein1024-336skein1024-344skein1024-352skein1024-360skein1024-368skein1024-376skein1024-384skein1024-392skein1024-400skein1024-408skein1024-416skein1024-424skein1024-432skein1024-440skein1024-448skein1024-456skein1024-464skein1024-472skein1024-480skein1024-488skein1024-496skein1024-504skein1024-512skein1024-520skein1024-528skein1024-536skein1024-544skein1024-552skein1024-560skein1024-568skein1024-576skein1024-584skein1024-592skein1024-600skein1024-608skein1024-616skein1024-624skein1024-632skein1024-640skein1024-648skein1024-656skein1024-664skein1024-672skein1024-680skein1024-688skein1024-696skein1024-704skein1024-712skein1024-720skein1024-728skein1024-736skein1024-744skein1024-752skein1024-760skein1024-768skein1024-776skein1024-784skein1024-792skein1024-800skein1024-808skein1024-816skein1024-824skein1024-832skein1024-840skein1024-848skein1024-856skein1024-864skein1024-872skein1024-880skein1024-888skein1024-896skein1024-904skein1024-912skein1024-920skein1024-928skein1024-936skein1024-944skein1024-952skein1024-960skein1024-968skein1024-976skein1024-984skein1024-992skein1024-1000skein1024-1008skein1024-1016skein1024-1024poseidon-bls12_381-a2-fc1poseidon-bls12_381-a2-fc1-sczeroxcert-imprint-256fil-commitment-unsealedfil-commitment-sealedholochain-adr-v0holochain-adr-v1holochain-key-v0holochain-key-v1holochain-sig-v0holochain-sig-v1skynet-ns"

var _Code_map = map[Code]string{
	0:        _Code_name[0:8],
	1:        _Code_name[8:13],
	2:        _Code_name[13:18],
	3:        _Code_name[18:23],
	4:        _Code_name[23:26],
	6:        _Code_name[26:29],
	17:       _Code_name[29:33],
	18:       _Code_name[33:41],
	19:       _Code_name[41:49],
	20:       _Code_name[49:57],
	21:       _Code_name[57:65],
	22:       _Code_name[65:73],
	23:       _Code_name[73:81],
	24:       _Code_name[81:90],
	25:       _Code_name[90:99],
	26:       _Code_name[99:109],
	27:       _Code_name[109:119],
	28:       _Code_name[119:129],
	29:       _Code_name[129:139],
	30:       _Code_name[139:145],
	33:       _Code_name[145:149],
	34:       _Code_name[149:160],
	35:       _Code_name[160:170],
	41:       _Code_name[170:173],
	42:       _Code_name[173:180],
	47:       _Code_name[180:184],
	48:       _Code_name[184:194],
	49:       _Code_name[194:203],
	50:       _Code_name[203:212],
	51:       _Code_name[212:221],
	53:       _Code_name[221:224],
	54:       _Code_name[224:228],
	55:       _Code_name[228:232],
	56:       _Code_name[232:239],
	80:       _Code_name[239:247],
	81:       _Code_name[247:251],
	85:       _Code_name[251:254],
	86:       _Code_name[254:266],
	96:       _Code_name[266:269],
	99:       _Code_name[269:276],
	112:      _Code_name[276:282],
	113:      _Code_name[282:290],
	114:      _Code_name[290:300],
	120:      _Code_name[300:307],
	123:      _Code_name[307:319],
	124:      _Code_name[319:331],
	129:      _Code_name[331:345],
	130:      _Code_name[345:356],
	131:      _Code_name[356:367],
	132:      _Code_name[367:371],
	133:      _Code_name[371:379],
	134:      _Code_name[379:387],
	144:      _Code_name[387:396],
	145:      _Code_name[396:410],
	146:      _Code_name[410:421],
	147:      _Code_name[421:427],
	148:      _Code_name[427:446],
	149:      _Code_name[446:460],
	150:      _Code_name[460:474],
	151:      _Code_name[474:494],
	152:      _Code_name[494:510],
	176:      _Code_name[510:523],
	177:      _Code_name[523:533],
	178:      _Code_name[533:559],
	192:      _Code_name[559:570],
	193:      _Code_name[570:578],
	202:      _Code_name[578:585],
	206:      _Code_name[585:593],
	208:      _Code_name[593:606],
	209:      _Code_name[606:616],
	212:      _Code_name[616:619],
	213:      _Code_name[619:622],
	214:      _Code_name[622:625],
	224:      _Code_name[625:637],
	225:      _Code_name[637:646],
	226:      _Code_name[646:653],
	227:      _Code_name[653:660],
	228:      _Code_name[660:668],
	229:      _Code_name[668:675],
	230:      _Code_name[675:682],
	231:      _Code_name[682:695],
	234:      _Code_name[695:711],
	235:      _Code_name[711:727],
	236:      _Code_name[727:737],
	237:      _Code_name[737:748],
	238:      _Code_name[748:766],
	240:      _Code_name[766:776],
	241:      _Code_name[776:783],
	250:      _Code_name[783:797],
	251:      _Code_name[797:807],
	273:      _Code_name[807:810],
	275:      _Code_name[810:825],
	276:      _Code_name[825:842],
	277:      _Code_name[842:854],
	290:      _Code_name[854:865],
	297:      _Code_name[865:873],
	301:      _Code_name[873:876],
	302:      _Code_name[876:879],
	400:      _Code_name[879:883],
	406:      _Code_name[883:889],
	421:      _Code_name[889:892],
	443:      _Code_name[892:897],
	444:      _Code_name[897:902],
	445:      _Code_name[902:908],
	446:      _Code_name[908:916],
	447:      _Code_name[916:924],
	448:      _Code_name[924:927],
	454:      _Code_name[927:932],
	460:      _Code_name[932:936],
	477:      _Code_name[936:938],
	478:      _Code_name[938:941],
	479:      _Code_name[941:959],
	480:      _Code_name[959:963],
	512:      _Code_name[963:967],
	513:      _Code_name[967:978],
	769:      _Code_name[978:996],
	770:      _Code_name[996:1013],
	1024:     _Code_name[1013:1029],
	4114:     _Code_name[1029:1053],
	4178:     _Code_name[1053:1063],
	4179:     _Code_name[1063:1073],
	4180:     _Code_name[1073:1083],
	4181:     _Code_name[1083:1093],
	4352:     _Code_name[1093:1096],
	4608:     _Code_name[1096:1104],
	4609:     _Code_name[1104:1112],
	4610:     _Code_name[1112:1120],
	4611:     _Code_name[1120:1129],
	4612:     _Code_name[1129:1137],
	4864:     _Code_name[1137:1149],
	4865:     _Code_name[1149:1163],
	4866:     _Code_name[1163:1174],
	7425:     _Code_name[1174:1188],
	21325:    _Code_name[1188:1195],
	45569:    _Code_name[1195:1204],
	45570:    _Code_name[1204:1214],
	45571:    _Code_name[1214:1224],
	45572:    _Code_name[1224:1234],
	45573:    _Code_name[1234:1244],
	45574:    _Code_name[1244:1254],
	45575:    _Code_name[1254:1264],
	45576:    _Code_name[1264:1274],
	45577:    _Code_name[1274:1284],
	45578:    _Code_name[1284:1294],
	45579:    _Code_name[1294:1304],
	45580:    _Code_name[1304:1314],
	45581:    _Code_name[1314:1325],
	45582:    _Code_name[1325:1336],
	45583:    _Code_name[1336:1347],
	45584:    _Code_name[1347:1358],
	45585:    _Code_name[1358:1369],
	45586:    _Code_name[1369:1380],
	45587:    _Code_name[1380:1391],
	45588:    _Code_name[1391:1402],
	45589:    _Code_name[1402:1413],
	45590:    _Code_name[1413:1424],
	45591:    _Code_name[1424:1435],
	45592:    _Code_name[1435:1446],
	45593:    _Code_name[1446:1457],
	45594:    _Code_name[1457:1468],
	45595:    _Code_name[1468:1479],
	45596:    _Code_name[1479:1490],
	45597:    _Code_name[1490:1501],
	45598:    _Code_name[1501:1512],
	45599:    _Code_name[1512:1523],
	45600:    _Code_name[1523:1534],
	45601:    _Code_name[1534:1545],
	45602:    _Code_name[1545:1556],
	45603:    _Code_name[1556:1567],
	45604:    _Code_name[1567:1578],
	45605:    _Code_name[1578:1589],
	45606:    _Code_name[1589:1600],
	45607:    _Code_name[1600:1611],
	45608:    _Code_name[1611:1622],
	45609:    _Code_name[1622:1633],
	45610:    _Code_name[1633:1644],
	45611:    _Code_name[1644:1655],
	45612:    _Code_name[1655:1666],
	45613:    _Code_name[1666:1677],
	45614:    _Code_name[1677:1688],
	45615:    _Code_name[1688:1699],
	45616:    _Code_name[1699:1710],
	45617:    _Code_name[1710:1721],
	45618:    _Code_name[1721:1732],
	45619:    _Code_name[1732:1743],
	45620:    _Code_name[1743:1754],
	45621:    _Code_name[1754:1765],
	45622:    _Code_name[1765:1776],
	45623:    _Code_name[1776:1787],
	45624:    _Code_name[1787:1798],
	45625:    _Code_name[1798:1809],
	45626:    _Code_name[1809:1820],
	45627:    _Code_name[1820:1831],
	45628:    _Code_name[1831:1842],
	45629:    _Code_name[1842:1853],
	45630:    _Code_name[1853:1864],
	45631:    _Code_name[1864:1875],
	45632:    _Code_name[1875:1886],
	45633:    _Code_name[1886:1895],
	45634:    _Code_name[1895:1905],
	45635:    _Code_name[1905:1915],
	45636:    _Code_name[1915:1925],
	45637:    _Code_name[1925:1935],
	45638:    _Code_name[1935:1945],
	45639:    _Code_name[1945:1955],
	45640:    _Code_name[1955:1965],
	45641:    _Code_name[1965:1975],
	45642:    _Code_name[1975:1985],
	45643:    _Code_name[1985:1995],
	45644:    _Code_name[1995:2005],
	45645:    _Code_name[2005:2016],
	45646:    _Code_name[2016:2027],
	45647:    _Code_name[2027:2038],
	45648:    _Code_name[2038:2049],
	45649:    _Code_name[2049:2060],
	45650:    _Code_name[2060:2071],
	45651:    _Code_name[2071:2082],
	45652:    _Code_name[2082:2093],
	45653:    _Code_name[2093:2104],
	45654:    _Code_name[2104:2115],
	45655:    _Code_name[2115:2126],
	45656:    _Code_name[2126:2137],
	45657:    _Code_name[2137:2148],
	45658:    _Code_name[2148:2159],
	45659:    _Code_name[2159:2170],
	45660:    _Code_name[2170:2181],
	45661:    _Code_name[2181:2192],
	45662:    _Code_name[2192:2203],
	45663:    _Code_name[2203:2214],
	45664:    _Code_name[2214:2225],
	45825:    _Code_name[2225:2235],
	45826:    _Code_name[2235:2246],
	45827:    _Code_name[2246:2257],
	45828:    _Code_name[2257:2268],
	45829:    _Code_name[2268:2279],
	45830:    _Code_name[2279:2290],
	45831:    _Code_name[2290:2301],
	45832:    _Code_name[2301:2312],
	45833:    _Code_name[2312:2323],
	45834:    _Code_name[2323:2334],
	45835:    _Code_name[2334:2345],
	45836:    _Code_name[2345:2356],
	45837:    _Code_name[2356:2368],
	45838:    _Code_name[2368:2380],
	45839:    _Code_name[2380:2392],
	45840:    _Code_name[2392:2404],
	45841:    _Code_name[2404:2416],
	45842:    _Code_name[2416:2428],
	45843:    _Code_name[2428:2440],
	45844:    _Code_name[2440:2452],
	45845:    _Code_name[2452:2464],
	45846:    _Code_name[2464:2476],
	45847:    _Code_name[2476:2488],
	45848:    _Code_name[2488:2500],
	45849:    _Code_name[2500:2512],
	45850:    _Code_name[2512:2524],
	45851:    _Code_name[2524:2536],
	45852:    _Code_name[2536:2548],
	45853:    _Code_name[2548:2560],
	45854:    _Code_name[2560:2572],
	45855:    _Code_name[2572:2584],
	45856:    _Code_name[2584:2596],
	45857:    _Code_name[2596:2606],
	45858:    _Code_name[2606:2617],
	45859:    _Code_name[2617:2628],
	45860:    _Code_name[2628:2639],
	45861:    _Code_name[2639:2650],
	45862:    _Code_name[2650:2661],
	45863:    _Code_name[2661:2672],
	45864:    _Code_name[2672:2683],
	45865:    _Code_name[2683:2694],
	45866:    _Code_name[2694:2705],
	45867:    _Code_name[2705:2716],
	45868:    _Code_name[2716:2727],
	45869:    _Code_name[2727:2739],
	45870:    _Code_name[2739:2751],
	45871:    _Code_name[2751:2763],
	45872:    _Code_name[2763:2775],
	45873:    _Code_name[2775:2787],
	45874:    _Code_name[2787:2799],
	45875:    _Code_name[2799:2811],
	45876:    _Code_name[2811:2823],
	45877:    _Code_name[2823:2835],
	45878:    _Code_name[2835:2847],
	45879:    _Code_name[2847:2859],
	45880:    _Code_name[2859:2871],
	45881:    _Code_name[2871:2883],
	45882:    _Code_name[2883:2895],
	45883:    _Code_name[2895:2907],
	45884:    _Code_name[2907:2919],
	45885:    _Code_name[2919:2931],
	45886:    _Code_name[2931:2943],
	45887:    _Code_name[2943:2955],
	45888:    _Code_name[2955:2967],
	45889:    _Code_name[2967:2979],
	45890:    _Code_name[2979:2991],
	45891:    _Code_name[2991:3003],
	45892:    _Code_name[3003:3015],
	45893:    _Code_name[3015:3027],
	45894:    _Code_name[3027:3039],
	45895:    _Code_name[3039:3051],
	45896:    _Code_name[3051:3063],
	45897:    _Code_name[3063:3075],
	45898:    _Code_name[3075:3087],
	45899:    _Code_name[3087:3099],
	45900:    _Code_name[3099:3111],
	45901:    _Code_name[3111:3123],
	45902:    _Code_name[3123:3135],
	45903:    _Code_name[3135:3147],
	45904:    _Code_name[3147:3159],
	45905:    _Code_name[3159:3171],
	45906:    _Code_name[3171:3183],
	45907:    _Code_name[3183:3195],
	45908:    _Code_name[3195:3207],
	45909:    _Code_name[3207:3219],
	45910:    _Code_name[3219:3231],
	45911:    _Code_name[3231:3243],
	45912:    _Code_name[3243:3255],
	45913:    _Code_name[3255:3267],
	45914:    _Code_name[3267:3279],
	45915:    _Code_name[3279:3291],
	45916:    _Code_name[3291:3303],
	45917:    _Code_name[3303:3315],
	45918:    _Code_name[3315:3327],
	45919:    _Code_name[3327:3339],
	45920:    _Code_name[3339:3351],
	45921:    _Code_name[3351:3362],
	45922:    _Code_name[3362:3374],
	45923:    _Code_name[3374:3386],
	45924:    _Code_name[3386:3398],
	45925:    _Code_name[3398:3410],
	45926:    _Code_name[3410:3422],
	45927:    _Code_name[3422:3434],
	45928:    _Code_name[3434:3446],
	45929:    _Code_name[3446:3458],
	45930:    _Code_name[3458:3470],
	45931:    _Code_name[3470:3482],
	45932:    _Code_name[3482:3494],
	45933:    _Code_name[3494:3507],
	45934:    _Code_name[3507:3520],
	45935:    _Code_name[3520:3533],
	45936:    _Code_name[3533:3546],
	45937:    _Code_name[3546:3559],
	45938:    _Code_name[3559:3572],
	45939:    _Code_name[3572:3585],
	45940:    _Code_name[3585:3598],
	45941:    _Code_name[3598:3611],
	45942:    _Code_name[3611:3624],
	45943:    _Code_name[3624:3637],
	45944:    _Code_name[3637:3650],
	45945:    _Code_name[3650:3663],
	45946:    _Code_name[3663:3676],
	45947:    _Code_name[3676:3689],
	45948:    _Code_name[3689:3702],
	45949:    _Code_name[3702:3715],
	45950:    _Code_name[3715:3728],
	45951:    _Code_name[3728:3741],
	45952:    _Code_name[3741:3754],
	45953:    _Code_name[3754:3767],
	45954:    _Code_name[3767:3780],
	45955:    _Code_name[3780:3793],
	45956:    _Code_name[3793:3806],
	45957:    _Code_name[3806:3819],
	45958:    _Code_name[3819:3832],
	45959:    _Code_name[3832:3845],
	45960:    _Code_name[3845:3858],
	45961:    _Code_name[3858:3871],
	45962:    _Code_name[3871:3884],
	45963:    _Code_name[3884:3897],
	45964:    _Code_name[3897:3910],
	45965:    _Code_name[3910:3923],
	45966:    _Code_name[3923:3936],
	45967:    _Code_name[3936:3949],
	45968:    _Code_name[3949:3962],
	45969:    _Code_name[3962:3975],
	45970:    _Code_name[3975:3988],
	45971:    _Code_name[3988:4001],
	45972:    _Code_name[4001:4014],
	45973:    _Code_name[4014:4027],
	45974:    _Code_name[4027:4040],
	45975:    _Code_name[4040:4053],
	45976:    _Code_name[4053:4066],
	45977:    _Code_name[4066:4079],
	45978:    _Code_name[4079:4092],
	45979:    _Code_name[4092:4105],
	45980:    _Code_name[4105:4118],
	45981:    _Code_name[4118:4131],
	45982:    _Code_name[4131:4144],
	45983:    _Code_name[4144:4157],
	45984:    _Code_name[4157:4170],
	45985:    _Code_name[4170:4183],
	45986:    _Code_name[4183:4196],
	45987:    _Code_name[4196:4209],
	45988:    _Code_name[4209:4222],
	45989:    _Code_name[4222:4235],
	45990:    _Code_name[4235:4248],
	45991:    _Code_name[4248:4261],
	45992:    _Code_name[4261:4274],
	45993:    _Code_name[4274:4287],
	45994:    _Code_name[4287:4300],
	45995:    _Code_name[4300:4313],
	45996:    _Code_name[4313:4326],
	45997:    _Code_name[4326:4339],
	45998:    _Code_name[4339:4352],
	45999:    _Code_name[4352:4365],
	46000:    _Code_name[4365:4378],
	46001:    _Code_name[4378:4391],
	46002:    _Code_name[4391:4404],
	46003:    _Code_name[4404:4417],
	46004:    _Code_name[4417:4430],
	46005:    _Code_name[4430:4443],
	46006:    _Code_name[4443:4456],
	46007:    _Code_name[4456:4469],
	46008:    _Code_name[4469:4482],
	46009:    _Code_name[4482:4495],
	46010:    _Code_name[4495:4508],
	46011:    _Code_name[4508:4521],
	46012:    _Code_name[4521:4534],
	46013:    _Code_name[4534:4547],
	46014:    _Code_name[4547:4560],
	46015:    _Code_name[4560:4573],
	46016:    _Code_name[4573:4586],
	46017:    _Code_name[4586:4599],
	46018:    _Code_name[4599:4612],
	46019:    _Code_name[4612:4625],
	46020:    _Code_name[4625:4638],
	46021:    _Code_name[4638:4651],
	46022:    _Code_name[4651:4664],
	46023:    _Code_name[4664:4677],
	46024:    _Code_name[4677:4690],
	46025:    _Code_name[4690:4703],
	46026:    _Code_name[4703:4716],
	46027:    _Code_name[4716:4729],
	46028:    _Code_name[4729:4742],
	46029:    _Code_name[4742:4755],
	46030:    _Code_name[4755:4768],
	46031:    _Code_name[4768:4781],
	46032:    _Code_name[4781:4794],
	46033:    _Code_name[4794:4807],
	46034:    _Code_name[4807:4820],
	46035:    _Code_name[4820:4833],
	46036:    _Code_name[4833:4846],
	46037:    _Code_name[4846:4859],
	46038:    _Code_name[4859:4872],
	46039:    _Code_name[4872:4885],
	46040:    _Code_name[4885:4898],
	46041:    _Code_name[4898:4911],
	46042:    _Code_name[4911:4924],
	46043:    _Code_name[4924:4937],
	46044:    _Code_name[4937:4950],
	46045:    _Code_name[4950:4964],
	46046:    _Code_name[4964:4978],
	46047:    _Code_name[4978:4992],
	46048:    _Code_name[4992:5006],
	46081:    _Code_name[5006:5031],
	46082:    _Code_name[5031:5059],
	52753:    _Code_name[5059:5080],
	61697:    _Code_name[5080:5103],
	61698:    _Code_name[5103:5124],
	8417572:  _Code_name[5124:5140],
	8483108:  _Code_name[5140:5156],
	9728292:  _Code_name[5156:5172],
	9793828:  _Code_name[5172:5188],
	10645796: _Code_name[5188:5204],
	10711332: _Code_name[5204:5220],
	11639056: _Code_name[5220:5229],
}

func (i Code) String() string {
	if str, ok := _Code_map[i]; ok {
		return str
	}
	return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
}