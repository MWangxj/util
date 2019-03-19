## httpEncry 使用fangfa

    1、go get -u -v github.com/MWangxj/util

    2、import(
        "github.com/MWangxj/util/httpEncrypt"
    )

    3、使用

        httpEncrypt.NewApp (source, secret, salt string) *app

        httpEncrypt.Do (ap *app, method HttpMethod ,urlStr string, param interface{}, ver version) (resp []byte,err error)

