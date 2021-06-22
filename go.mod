module github.com/qri-io/qri

go 1.13

replace github.com/qri-io/starlib => ../starlib

require (
	github.com/beme/abide v0.0.0-20190723115211-635a09831760
	github.com/cube2222/octosql v0.2.1-0.20200319150444-e5a71fa20dbe
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dustin/go-humanize v1.0.0
	github.com/fatih/color v1.9.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/ghodss/yaml v1.0.0
	github.com/gofrs/flock v0.7.1
	github.com/google/flatbuffers v1.12.1-0.20200706154056-969d0f7a6317
	github.com/google/go-cmp v0.5.3
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/schema v1.2.0
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-datastore v0.4.4
	github.com/ipfs/go-ipfs v0.6.0
	github.com/ipfs/go-ipfs-config v0.8.0
	github.com/ipfs/go-ipld-format v0.2.0
	github.com/ipfs/go-log v1.0.4
	github.com/ipfs/interface-go-ipfs-core v0.3.0
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/libp2p/go-libp2p v0.11.0
	github.com/libp2p/go-libp2p-circuit v0.3.1
	github.com/libp2p/go-libp2p-connmgr v0.2.4
	github.com/libp2p/go-libp2p-core v0.6.1
	github.com/libp2p/go-libp2p-crypto v0.1.0
	github.com/libp2p/go-libp2p-peerstore v0.2.6
	github.com/libp2p/go-libp2p-swarm v0.2.8
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-multiaddr v0.3.2-0.20210122024440-7274874c78df
	github.com/multiformats/go-multicodec v0.1.6
	github.com/multiformats/go-multihash v0.0.14
	github.com/olekukonko/tablewriter v0.0.4
	github.com/pkg/errors v0.9.1
	github.com/qri-io/dag v0.2.2
	github.com/qri-io/dataset v0.3.1-0.20210611163416-f6059e0f298a
	github.com/qri-io/deepdiff v0.2.1
	github.com/qri-io/didmod v0.1.0
	github.com/qri-io/doggos v0.1.0
	github.com/qri-io/ioes v0.1.1
	github.com/qri-io/jsonschema v0.2.0
	github.com/qri-io/qfs v0.6.1-0.20210610133154-676ac303918e
	github.com/qri-io/starlib v0.4.2
	github.com/russross/blackfriday/v2 v2.0.2-0.20190629151518-3e56bb68c887
	github.com/sergi/go-diff v1.1.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/ugorji/go/codec v1.1.7
	github.com/vbauerster/mpb/v5 v5.3.0
	go.starlark.net v0.0.0-20210406145628-7a1108eaa012
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f
	golang.org/x/text v0.3.3
	gopkg.in/yaml.v2 v2.3.0
	nhooyr.io/websocket v1.8.6
)
