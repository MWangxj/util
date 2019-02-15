package comment

import (
	"bufio"
	"git.dian.so/leto/util/byte2str"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 
// getGoFileListByDir
// param pathname string
// return fileNames []string
// return  err error
func GetGoFileListByDir(pathname string) (fileNames []string, err error) {
	pathname, err = filepath.Abs(pathname)
	if err != nil {
		return nil, err
	}
	dir_err := filepath.Walk(pathname,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				//fileNames = append(fileNames, path)
				return nil
			} else {
				if strings.HasSuffix(path, ".go") && !strings.Contains(path, "vendor") {
					fileNames = append(fileNames, path)
				}
			}

			return nil
		})
	return fileNames, dir_err
}

// 
// GetFileLine
// param filename string
// return err error
func GetFileLine(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	regFunc := regexp.MustCompile(`^func.*{\n`)
	regComm := regexp.MustCompile(`^[//]+[\w,\W]*`)
	rd := bufio.NewReader(f)
	fc := ""
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		if regFunc.Match([]byte(line)) {
			fi := parseFuncInfo(line)
			/*fc += "// struct " + fi.StructName + "\r\n"
			if len(fi.Params) == 0 || fi.Params[0] == "" {
				fc += "// param null \r\n"
			} else {
				for _, param := range fi.Params {
					fc += "// param " + param + "\r\n"
				}
			}
			if len(fi.Returns) == 0 {
				fc += "// return null \r\n"
			} else {
				for _, rt := range fi.Returns {
					fc += "// return " + rt + "\r\n"
				}
			}*/
			fc += "// " + fi.FuncName + "\r\n"
		}
		if regComm.Match([]byte(line)) {
			line = ""
		}
		fc += line
	}
	if err := ioutil.WriteFile(filename, byte2str.StringToBytes(fc), 0); err != nil {
		return err
	}
	return nil
}

type funcInfo struct {
	StructName string
	FuncName   string
	Params     []string
	Returns    []string
}

// 
// parseFuncInfo
// param funcStr string
// return fi *funcInfo
// 
// parseFuncInfo
// param funcStr string
// return fi *funcInfo
func parseFuncInfo(funcStr string) (fi *funcInfo) {
	fi = &funcInfo{}
	regStruct := regexp.MustCompile(`^func[ ]*[(]`)
	if regStruct.Match(byte2str.StringToBytes(funcStr)) {
		regStruct = regexp.MustCompile(`[(]+[\w,\W]*[)]`)
		if rs := regStruct.FindAllString(funcStr, -1); len(rs) > 0 {
			//funcStr = funcStr[:]
			funcStr = funcStr[:5] + funcStr[strings.Index(funcStr, ")")+1:]
			fi.StructName = rs[0][1 : len(rs[0])-1]
		}
	}
	regFuncName := regexp.MustCompile(`[\w]+`)
	fi.FuncName = regFuncName.FindAllString(funcStr, -1)[1]
	regParams := regexp.MustCompile(`[(]+[a-z,A-Z, ,*,[,\]]*[)]`)
	paramsAndReturns := regParams.FindAllString(funcStr, -1)
	if len(paramsAndReturns) > 0 {
		fi.Params = strings.Split(paramsAndReturns[0][1:len(paramsAndReturns[0])-1], ",")
	}
	if len(paramsAndReturns) > 1 {
		fi.Returns = strings.Split(paramsAndReturns[1][1:len(paramsAndReturns[1])-1], ",")
	}
	regReturns := regexp.MustCompile(`[)][\w]*[{]`)
	rt := regReturns.FindAllString(funcStr, -1)
	if len(rt) > 0 {
		fi.Returns = rt
	}
	return fi
}
