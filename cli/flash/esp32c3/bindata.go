// Code generated for package esp32c3 by go-bindata DO NOT EDIT. (@generated)
// sources:
// stub/stub.json
package esp32c3

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _stubStubJson = []byte(`{"params_start": 1077411840, "code": "37070c603c4f411106c693f707c03ccf130580029700c8ffe780c05673d0007e73d0107e9700c8ffe780006fb240814541011703c8ff6700a3129967aa97c207dc4b4111416722c47d1713944700d183798cbd8b3e94b305b4021375f50f06c6b3d5c5029700c8ffe780c0012285b240224441018280b787c93f9387070011673e971457056693058600b386b602b697d44363fbc60013861600d0c3b6972384a7005c5785075cd782807971b707006022d456ca3784c93f83aa870026d24ece52ccb7c9c93f130a0400916406d64ad05ac85ec693890903d294b707006003a9c7011379f93f630c0908814b37cbc93fb707006088438326040091471375f50f63e0d70293972600ce979c4382879307000c6317f50085472320f40023a60402850be31779fd55bf9307b00d6316f5008d472320f400edb79305000c9c541307040013060b006311b506850791466387d7049cd41456856793858700b386b6023697ba97238407005c562322070095eb91472320f400b707006023a60700b2502254b707006023a8570192540259f249624ad24a424bb24b4561828023a4040255bf23200400adbf0547639de600056793068700b387d702d297ba972384a700894785bf4d3da9bf9307c00d6309f5009307d00de31ef5f81305b00d19a01305000c513df1bf35714ac97d73056922cd26cb4ec756c35ac1dede7d7406cf52c5930709071a91930984fa8a97be99aa8b4e85ae8a328b9700c8ffe780603f130709070a9733098700130484f981443a9463ed54030567930707077d74930584fa8a97130484f93e94930707078a97be9522859700c8ffe780203c2285c1459700c8ffe78020e40145b1a8338a9a4063734b015a8a5286ca85338574019700c8ffe780c0ea15edd2854a859700c8ffe78060e191452285232c09f8ef0020789147631ff502032689f9e3f4c4f8058eca854e859700c8ffe780e035832489f98dbf1305300505631a91fa406a44da444a49ba492a4a9a4a0a4bf65b0d61828013054005cdb70111b707ce3f22cc2a8403a507ff2ec606ce9700c8ffe780e0e0b7270060984fb70600e0fd16758f98cfd84fb7060004fd16758fb706005c558fb245d8cf3707004098c398437dffc0c3f24062448cc30145056182802971232a9112aa8488082328211323206113232e1112232c8112232631132324411323225113232e7111232c8111232a91112328a1112326b1112e8b32c29700c8ffe78040298567fd1733f7f400130920036312072cb377fb0013093003639c072a9700c8ffe780c0d92a84130940036313052a3789c93f1166130606038145130509009700c8ffe780e0f8370ac93f2166814513050a009700c8ffe780a0f7b705384001469385250b15459700c8ffe7806021b7070060d853130500023ac03707340013074706d8d313071010d8c79700c8ffe780a01f13064002814568109700c8ffe78020f36146814548089700c8ffe78040f2e1454808c523f327207e85693edec16b930a0900130c0900d24763e367038c0888009700c8ffe780201cf327207e7257930540026810998f3ede01496523f9a2f32c207ec167fd17b3fdf400924789c7d247ce976364f4067327207ee257b3879741ba973edc7326207eb797c93f93870700dc4703a74a008d66d6963e978967d697dc4bd44eb7c8c93f3e9736973acc1167569783270900032d470237172707938c08009145130707e083a68c026395a605639bb7021309800389aab3078b4063ec7701639a0d00b70580002685dd33631c05125e949dbfb70500012685d93b631d05164e949db7f326207e918ee37ed7fa1309700319aa1147e38ee7fa7327207eb257918fba973ed6732e207e93878900b307fd0213838700e29783ad4700ce9783c787006293858b9dc3ee8605471a86a16513050a0072c49700c8ffe780e0b8fd57aa8d630df510224e13030a007327207ec2576e869a85b387c741ba9788083ed81ac69700c8ffe780a0077327207eb707ce3f03a507ff3ac49700c8ffe78080b5f327207ee2562247b697998f3edcf326207e32436e8626859a8536c49700c8ffe78040b55ded7327207ed257a246958fba973eda938789003308fd0283a74c0211478507629823220800638ae70223a2fc02d246130680058c08ee96281136ca9700c8ffe780c0d32c1168089700c8ffe78040ffe1454808ee94dd2621bd23a20c02c1bf130950030247b707006013050002d8d39700c8ffe78080fb8320c113032481134a8583244113032901138329c112032a8112832a4112032b0112832bc111032c8111832c4111032d0111832dc11031618280130960037db71309900365b7130900044db7856763e3c70091b61305200582805171cecfd6cbdec786d7a2d5a6d3cad1d2cddac9e2c57d731a91aa8bae89b28a0dc6856713052006328b63f1c70205631a91be502e549e540e59fe496e4ade4a4e4bbe4b2e4c6d618280056b05697d749307090c930404f58a97be9426859700c8ffe78040f01307090c130a84fa0a973a9a1307090c0a97330c8700130404f43a94639d090205679307070c7d74930504f58a97130404f43e949307070c8a97be9522859700c8ffe78060ec2285c1459700c8ffe7806094014595bf52859700c8ffe78040ea4e8963733b015a894a86e2855e859700c8ffe780a09a01c513053006b1b74a86e28526859700c8ffe780c0e763850a024a86e28552859700c8ffe780a0e6d28522859700c8ffe78020e6c14522859700c8ffe780208eca9bb389294185b7011106ce02c6b72700603707001098c3984393163700e3cd06febc4f370700017d176800f98f91453ec69700c8ffe780808af240014505618280797126d24ad093075005817437c9c93f22d406d62305f10037040060fd14130949041c50370702008545d98f1cd01c501305a100e58f1cd0f12a85476311f504fd57a305f1008347a1002547fd1793f7f70f6362f7028a07ca979c438287c14508084d2ab1476316f5026246d2454245c934a305a10085451305b100412a0345a10005479307a5ff93f7f70fe36bf7f88da093071004a305f100f1bfc1450808952ab1476317f5006246d2454245dd33c9b793071005c5b7c1450808a12ab1476317f5006246d2454245e13b5db793071006d1b7cd3d71bf9700c8ffe780808349bf1305b1008545a305010005220345a100b25022549254025945618280c1450808092aaa8508081122a3050100a5b7c1450808012291476318f500c247914568009c433ec6cdb7930710f9adb7c1450808cd20a1476316f500c247524798c3e9b7930710fa81bfb7073840011193870700014506ce22cc26cac043844397f0c7ffe780e074eff02fe9c167fd171307001085664166b7050001014597f0c7ffe780807a02c611cc09651305057197f0c7ffe780206b2286a6850145eff06fe92ac631651305053597f0c7ffe780806991456800a128a9352a8409651305057197f0c7ffe780006899476306f40097f0c7ffe780206bf2406244d244056182803705c93f37c6c93f93070500130606031d8e4111098681451305050006c69700c8ffe7806094b24041013dbf17f3c7ff6700c368797122d44ece52cc06d626d24ad056ca5ac85ec662c42a8aae891304000c97f0c7ffe78020632a89e31b85fe0144930a000c130bb00d930bc00d130cd00d631734039304000c97f0c7ffe780a060e31c95fe2285b250225492540259f249624ad24a424bb24b224c4561828097f0c7ffe780405eaa84e30e55fd631d650197f0c7ffe780205d630675016305850101444dbfca84b3078a0023809700050445b7", "code_start": 1077411848, "entry": 1077414384, "data": "160138403001384030013840d0013840860138406e083840ac083840c8083840e4083840e8083840f2083840f20838400e093840200938403e093840", "data_start": 1070186544, "num_params": 2, "code_size": 2748, "data_size": 60}`)

func stubStubJsonBytes() ([]byte, error) {
	return _stubStubJson, nil
}

func stubStubJson() (*asset, error) {
	bytes, err := stubStubJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stub/stub.json", size: 5794, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"stub/stub.json": stubStubJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"stub": &bintree{nil, map[string]*bintree{
		"stub.json": &bintree{stubStubJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}